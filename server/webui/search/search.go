package search

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/blevesearch/bleve"
	"oceanleadership.org/CDFRegistryWG/server/webui/sparql"
)

type FreeTextResults struct {
	Place     int
	Index     string
	Score     float64
	ID        string
	Fragments []Fragment
	IconName  string
}

type Fragment struct {
	Key   string
	Value []string
}

type SearchMetaData struct {
	Term    string
	Count   int
	Message string
}

// DoSearch is there to do searching..  (famous documentation style intact!)
func DoSearch(w http.ResponseWriter, r *http.Request) {
	log.Printf("r path: %s\n", r.URL.Query()) // need to log this better so I can filter out search terms later
	queryterm := r.URL.Query().Get("q")
	queryterm = strings.TrimSpace(queryterm) // remove leading and trailing white spaces a user might put in (not internal spaces though)

	// Make a var in case I want other templates I switch to later...
	templateFile := "./templates/rwg.html"

	// var queryResults DocumentMatchCollection{}
	queryResults := indexCall(queryterm, "")
	len := len(queryResults)

	// Set up some metadata on the search results to return
	var searchmeta SearchMetaData
	searchmeta.Term = queryterm
	searchmeta.Count = len
	if len == 0 {
		if queryterm == "" {
			searchmeta.Message = "Search EarthCube CDF RWG demo index"

		} else {
			searchmeta.Message = "No results found for this search"
		}
	}

	// If we have a term.. search the triplestore
	var spres sparql.SPres
	if len > 0 {
		topResult := queryResults[0] // pass this as a new template section TR!
		fmt.Println(topResult.ID)
		spres = sparql.DoCall(topResult.ID) // turn sparql call on / off
		// fmt.Print(spres.Description)
	}

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", searchmeta) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
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

// termReWrite puts the bleve ~1 or ~2 term options on for fuzzy matching
func termReWrite(phrase string, distanceAppend string) string {
	terms := strings.Split(phrase, " ")

	for k, _ := range terms {
		var str bytes.Buffer
		str.WriteString(strings.TrimSpace(terms[k]))
		str.WriteString(distanceAppend)
		terms[k] = str.String()
	}

	fmt.Println(strings.Join(terms, " "))
	return strings.Join(terms, " ")
}

// return JSON string..  enables use of func for REST call too
func indexCall(phrase string, distance string) []FreeTextResults {
	if phrase == "" {
		return nil
	}

	// TODO ..  improve this..
	// Really need to check if it is ~1 or ~2.  If not, set to empty
	if distance == "" {
		distance = ""
	}

	// indexPath := "./index/rwg.bleve"
	// index, err := bleve.OpenUsing(indexPath, map[string]interface{}{
	// 	"read_only": true,
	// })
	// if err != nil {
	// 	log.Printf("error opening index %s: %v", indexPath, err)
	// } else {
	// 	log.Printf("registered index: at %s", indexPath)
	// }

	// Playing with index aliases
	// Open all indexes in an alias and use this in a named call
	log.Printf("Start building Codex index \n")

	index1, err := bleve.OpenUsing("./index/rwg.bleve", map[string]interface{}{
		"read_only": true,
	})
	if err != nil {
		log.Printf("Error with index alias: %v", err)
	}
	index2, err := bleve.OpenUsing("./index/rwgdata.bleve", map[string]interface{}{
		"read_only": true,
	})
	if err != nil {
		log.Printf("Error with index alias: %v", err)
	}
	index := bleve.NewIndexAlias(index1, index2)
	log.Printf("Codex index built\n")

	// parse string and add ~2 to each term/word, then rebuild as a string.
	query := bleve.NewQueryStringQuery(termReWrite(phrase, distance))
	search := bleve.NewSearchRequestOptions(query, 10, 0, false) // no explanation
	search.Highlight = bleve.NewHighlight()                      // need Stored and IncludeTermVectors in index
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

		// set up a material icon   ref:  https://material.io/icons/
		var iconName string
		if item.Index == "./index/rwgdata.bleve" {
			iconName = "file_download" // material design icon name used in template
		}
		if item.Index == "./index/rwg.bleve" {
			iconName = "http" // material design icon name used in template  alts:  web_asset or web
		}

		results = append(results, FreeTextResults{k, item.Index, item.Score, item.ID, frags, iconName})
	}

	fmt.Printf("Looping status count:%d, distance:%s\n", len(results), distance)

	// TODO..  Yet Another Ugly Section (YAUS)  (I've named the pattern..  that is just sad)
	// check here..  if results are 0 then recursive call with ~1
	// check here and if 0 then try again with ~2
	var finalResults []FreeTextResults
	if len(results) == 0 {
		if distance == "" {
			finalResults = indexCall(phrase, "~1")
		}
		if distance == "~1" {
			finalResults = indexCall(phrase, "~2")
		}
	}

	if len(results) > 0 {
		finalResults = results
	}

	return finalResults
}
