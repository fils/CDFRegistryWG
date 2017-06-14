package search

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
)

type SearchResults struct {
	Place     int
	Index     string
	Score     float64
	ID        string
	Fragments []Fragment
}

type Fragment struct {
	Key   string
	Value []string
}

// DoSearch is there to do searching..  (famous documentation style intact!)
func DoSearch(w http.ResponseWriter, r *http.Request) {
	log.Printf("r path: %s\n", r.URL.Query())
	queryterm := r.URL.Query().Get("q")

	// Make a var in case I want other templates I switch to later...
	templateFile := "./templates/rwg.html"

	// var queryResults DocumentMatchCollection{}
	queryResults := indexCall(queryterm)
	fmt.Println(queryResults)

	len := len(queryResults)
	fmt.Println(len)

	// get SPARQL results
	// need a for each loop on the search results
	sparqlCall("url")

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	// If we have a term.. search the index
	topResult := queryResults[0] // pass this as a new template section TR!

	fmt.Printf("Top results %v \n", topResult)

	err = ht.ExecuteTemplate(w, "T", queryResults) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

func sparqlCall(uri string) {
	log.Println("SPARQL call..   results to struct pointer")

	/*
	   Something like this but also with logo

	   SELECT DISTINCT *
	   WHERE {
	     ?s <http://schema.org/url> <http://www.bco-dmo.org> .
	     optional {?s <http://schema.org/description> ?desc } .
	     optional {?s rdf:type ?type } .
	     ?s ?pred ?obj
	   }

	*/

}

// return JSON string..  enables use of func for REST call too
func indexCall(phrase string) []SearchResults {
	indexPath := "/Users/dfils/src/go/src/oceanleadership.org/CDFRegistryWG/server/webui/index/rwg.bleve"

	index, err := bleve.OpenUsing(indexPath, map[string]interface{}{
		"read_only": true,
	})
	if err != nil {
		log.Printf("error opening index %s: %v", indexPath, err)
	} else {
		log.Printf("registered index: at %s", indexPath)
	}

	// query := bleve.NewMatchQuery(phrase)
	query := bleve.NewQueryStringQuery(phrase)
	search := bleve.NewSearchRequestOptions(query, 10, 0, false) // no explanation
	// search. = bleve.NewTextFieldMapping  ["potentialAction.target.description"]
	search.Highlight = bleve.NewHighlight() // need Stored and IncludeTermVectors in index
	searchResults, err := index.Search(search)

	hits := searchResults.Hits // array of struct DocumentMatch

	var results []SearchResults

	for k, item := range hits {
		fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
		// fmt.Printf("%v\n", item.Fields["potentialAction.target.description"])
		var frags []Fragment
		for key, frag := range item.Fragments {
			fmt.Printf("%s   %s\n", key, frag)
			frags = append(frags, Fragment{key, frag})
		}
		results = append(results, SearchResults{k, item.Index, item.Score, item.ID, frags})
	}

	// TODO..  just return the documentmatch struct collection (hits) and parse it in the template...
	return results

	// jsonResults, _ := json.MarshalIndent(hits, " ", " ")
	// return string(jsonResults)
}
