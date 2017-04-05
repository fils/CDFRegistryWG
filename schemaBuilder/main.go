package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kazarena/json-gold/ld"
)

type SchemaOrgMetadata struct {
	Context         Context      `json:"@context"` // was []interface{}  should be Context struct (which has 3 items in it for each voc)
	Type            string       `json:"@type"`
	ID              string       `json:"@id"` // need an ID for each of these subsections
	URL             string       `json:"url"`
	ContactPoint    ContactPoint `json:"contactPoint"`
	PotentialAction SearchAction `json:"potentialAction"`
}

type SearchAction struct {
	Type   string `json:"@type"`
	Target Target `json:"target"`
}

type Target struct {
	Type        string `json:"@type"`
	URLTemplate string `json:"urlTemplate"`
	Description string `json:"description"`
	HTTPMethod  string `json:"httpMethod"`
}

type Context struct {
	Schema   string `json:"@vocab"`
	GeoLink  string `json:"geolink"`         // namespace prefix in the rest of the struct
	OpenCore string `json:"opencore"`        // namespace prefix in the rest of the struct
	Base     string `json:"@base,omitempty"` // used in CSVW to prevent relative IRI from becoming an absolute IRI
}

type ContactPoint struct {
	Type        string `json:"@type"`
	ID          string `json:"@id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Email       string `json:"email"`
	ContactType string `json:"contactType"`
}

func main() {

	result := SchemaOrgMetadata{}
	result.URL = "http://example.org"
	result.Type = "Organization"
	result.ID = fmt.Sprintf("%s", result.URL) //  %s.json ??

	// context setting
	result.Context.Schema = "http://schema.org/" // this "schema" is the @voc in the struct..  confusing when not using schema.org
	result.Context.OpenCore = "http://opencore.org/voc/1/"
	result.Context.GeoLink = "http://glview.org/voc/1/"
	result.Context.Base = fmt.Sprintf("%s", result.URL) //  %s.json

	// contactPoint setting
	result.ContactPoint.Type = "ContactPoint"
	result.ContactPoint.ID = "ID"
	result.ContactPoint.Name = "John Doe"
	result.ContactPoint.URL = "http://url.to/person"
	result.ContactPoint.Email = "joe@example.org"
	result.ContactPoint.ContactType = "technical support"

	// potentialAction
	result.PotentialAction.Type = "SearchAction"
	result.PotentialAction.Target.Description = "Swagger 1.2 description document"
	result.PotentialAction.Target.HTTPMethod = "GET"
	result.PotentialAction.Target.Type = "EntryPoint"
	result.PotentialAction.Target.URLTemplate = "http://opencoredata.org/apidocs.json"

	jsonldtext, _ := json.MarshalIndent(result, "", " ") // results as embeddale JSON-LD

	fmt.Println("jsonld text--------------------------------")
	fmt.Println(string(jsonldtext))

	// fmt.Println("jsonLDToRDF--------------------------------")
	// fmt.Println(jsonLDToRDF(string(jsonldtext)))

}

// Trys to take a simple JSON-LD string and process it to RDF.
func jsonLDToRDF(jsonld string) string {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	triples, err := proc.ToRDF(myInterface, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err.Error()
	}

	return triples.(string)
}
