package search

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
)

// TODO
// - place the bleve results into a struct array
// - pass said array to the template for display
// - also use results to pick a winning URI to use in a SPARQL search
// - pass results of SPARQL search to template for display ala knowledge box.

func DoSearch(w http.ResponseWriter, r *http.Request) {
	log.Printf("r path: %s\n", r.URL.Query())
	queryterm := r.URL.Query().Get("q")

	// Make a var in case I want other templates I switch to later...
	templateFile := "./templates/rwg.html"

	queryResults := indexCall(queryterm)

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	// If we have a term.. search the index

	err = ht.ExecuteTemplate(w, "T", queryResults) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}

}

func indexCall(phrase string) string {
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
	search.Highlight = bleve.NewHighlight()                      // need Stored and IncludeTermVectors in index
	searchResults, err := index.Search(search)

	hits := searchResults.Hits // array of struct DocumentMatch

	for k, item := range hits {
		fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
		for key, frag := range item.Fragments {
			fmt.Printf("%s   %s\n", key, frag)
		}
	}

	jsonResults, _ := json.MarshalIndent(hits, " ", " ")

	return string(jsonResults)
}
