package search

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/blevesearch/bleve"
	"oceanleadership.org/CDFRegistryWG/server/webui/sparql"
)

type FreeTextResults struct {
	Place           int
	Index           string
	Score           float64
	ID              string
	Fragments       []Fragment
	IconName        string
	IconDescription string
}

type Fragment struct {
	Key   string
	Value string //[]string
}

type SearchMetaData struct {
	Term    string
	Count   int
	Message string
}

type Qstring struct {
	Query      string
	Qualifiers map[string]string
}

// DoSearch is there to do searching..  (famous documentation style intact!)
func DoSearch(w http.ResponseWriter, r *http.Request) {
	log.Printf("r path: %s\n", r.URL.Query()) // need to log this better so I can filter out search terms later
	queryterm := r.URL.Query().Get("q")
	queryterm = strings.TrimSpace(queryterm) // remove leading and trailing white spaces a user might put in (not internal spaces though)

	// Make a var in case I want other templates I switch to later...
	templateFile := "./templates/rwg.html"

	// parse the queryterm to get the colon based qualifiers
	qstring := parse(queryterm)

	// var queryResults DocumentMatchCollection{}
	distance := ""
	queryResults := indexCall(qstring, distance)
	qrl := len(queryResults)

	// moved the len test and string mod to here
	// TODO..  Yet Another Ugly Section (YAUS)  (I've named the pattern..  that is just sad)
	// check here..  if results are 0 then recursive call with ~1
	// check here and if 0 then try again with ~2
	// var finalResults []FreeTextResults
	fmt.Printf("Len: %d    distance: %s \n", qrl, distance)
	if qrl == 0 {
		if strings.Contains(distance, "") {
			fmt.Println("Call ~1")
			queryResults = indexCall(qstring, "~1")
		}
	}
	qrl = len(queryResults)

	if qrl == 0 {
		if strings.Contains(distance, "~1") {
			fmt.Println("Call ~2")
			queryResults = indexCall(qstring, "~2")
		}
	}

	// if len(results) > 0 {
	// 	finalResults = results
	// }

	// Set up some metadata on the search results to return
	var searchmeta SearchMetaData
	searchmeta.Term = queryterm // We don't use qstring.Query here since we want the full string including qualifiers, returned to the page for rendering with results
	searchmeta.Count = qrl
	if qrl == 0 {
		if qstring.Query == "" {
			searchmeta.Message = "Search EarthCube CDF RWG demo index"

		} else {
			searchmeta.Message = "No results found for this search"
		}
	}

	// If we have a term.. search the triplestore
	var spres sparql.SPres
	if qrl > 0 {
		topResult := queryResults[0] // pass this as a new template section TR!
		fmt.Println(topResult.ID)
		var err error
		spres, err = sparql.DoCall(topResult.ID) // turn sparql call on / off
		if err != nil {
			log.Printf("SPARQL call failed: %s", err)
		}
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

func parse(qstring string) Qstring {
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`) // get rid of multiple spaces
	qstring = re_inside_whtsp.ReplaceAllString(qstring, " ")
	sa := strings.Split(qstring, " ")

	var buffer bytes.Buffer
	qpairs := make(map[string]string)
	for _, item := range sa {
		if strings.ContainsAny(item, ":") {
			qualpair := strings.Split(item, ":")
			qpairs[qualpair[0]] = qualpair[1]
		} else {
			buffer.WriteString(item)
			buffer.WriteString(" ")
		}
	}

	qs := Qstring{Query: buffer.String(), Qualifiers: qpairs}
	return qs
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
func indexCall(qstruct Qstring, distance string) []FreeTextResults {
	if qstruct.Query == "" {
		return nil
	}

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

	var index bleve.IndexAlias

	if strings.Contains(qstruct.Qualifiers["type"], "organization") {
		index = bleve.NewIndexAlias(index1)
	} else if strings.Contains(qstruct.Qualifiers["type"], "data") {
		index = bleve.NewIndexAlias(index2)
	} else {
		index = bleve.NewIndexAlias(index1, index2)
	}

	log.Printf("Codex index built\n")

	// parse string and add ~2 to each term/word, then rebuild as a string.
	query := bleve.NewQueryStringQuery(termReWrite(qstruct.Query, distance))
	search := bleve.NewSearchRequestOptions(query, 100, 0, false) // no explanation
	search.Highlight = bleve.NewHighlight()                       // need Stored and IncludeTermVectors in index
	searchResults, err := index.Search(search)
	// index.Close()  // null index bug test..  didn't work

	// TODO  make var hits, check for nil and do nothing when nil....
	hits := searchResults.Hits // array of struct DocumentMatch

	var results []FreeTextResults

	for k, item := range hits {
		// fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
		// fmt.Printf("%v\n", item.Fields["potentialAction.target.description"])
		var frags []Fragment
		for key, frag := range item.Fragments {
			// fmt.Printf("%s   %s\n", key, frag)
			frags = append(frags, Fragment{key, frag[0]})
		}

		// set up a material icon   ref:  https://material.io/icons/
		var iconName string
		var iconDescription string
		if item.Index == "./index/rwgdata.bleve" {
			iconName = "file_download"                                  // material design icon name used in template
			iconDescription = "Data resource of type data landing page" // material design icon name used in template
		}
		if item.Index == "./index/rwg.bleve" {
			iconName = "http"                                                  // material design icon name used in template  alts:  web_asset or web
			iconDescription = "Organization or other related on-line resource" // material design icon name used in template  alts:  web_asset or web
		}

		results = append(results, FreeTextResults{k, item.Index, item.Score, item.ID, frags, iconName, iconDescription})
	}

	fmt.Printf("Looping status count:%d, distance:%s\n", len(results), distance)
	return results // finalResults
}
