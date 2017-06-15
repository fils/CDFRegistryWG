package search

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
	"oceanleadership.org/CDFRegistryWG/server/webui/sparql"
)

type FreeTextResults struct {
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

	// If we have a term.. search the triplestore
	var spres sparql.SPres
	if len > 0 {
		topResult := queryResults[0] // pass this as a new template section TR!
		fmt.Println(topResult.ID)
		spres = sparql.DoCall(topResult.ID)
		fmt.Print(spres.Desc)
	}

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", queryterm) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "T", queryResults) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "S", spres) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

// return JSON string..  enables use of func for REST call too
func indexCall(phrase string) []FreeTextResults {
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

	var results []FreeTextResults

	for k, item := range hits {
		// fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
		// fmt.Printf("%v\n", item.Fields["potentialAction.target.description"])
		var frags []Fragment
		for key, frag := range item.Fragments {
			// fmt.Printf("%s   %s\n", key, frag)
			frags = append(frags, Fragment{key, frag})
		}
		results = append(results, FreeTextResults{k, item.Index, item.Score, item.ID, frags})
	}

	// TODO..  just return the documentmatch struct collection (hits) and parse it in the template...
	return results

	// jsonResults, _ := json.MarshalIndent(hits, " ", " ")
	// return string(jsonResults)
}
