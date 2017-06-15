package sparql

import (
	"bytes"
	"log"

	"github.com/knakk/sparql"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: orgInfo
	   SELECT DISTINCT *
	   WHERE {
		{   
	     ?s <http://schema.org/url> "{{.URL}}" .
	     optional {?s <http://schema.org/description> ?desc } .
	     ?s rdf:type <http://schema.org/Organization>   .
	     ?s ?pred ?obj
		 }
		 UNION
		 		{   
	     ?s <http://schema.org/url> <{{.URL}}> .
	     optional {?s <http://schema.org/description> ?desc } .
	     ?s rdf:type <http://schema.org/Organization>   .
	     ?s ?pred ?obj
		 }
}
LIMIT 1

# tag: generalInfo
	   SELECT DISTINCT *
	   WHERE {
	     ?s <http://schema.org/url> "{{.URL}}" .
	     optional {?s <http://schema.org/description> ?desc } .
	     optional {?s rdf:type ?type } .
	     ?s ?pred ?obj
	   }
	   LIMIT 1

# tag: freeText
		prefix bds: <http://www.bigdata.com/rdf/search#>
		select DISTINCT ?name ?url ?progname ?description
		where {
		
		{ ?s <http://schema.org/name> ?name   .
		?s <http://schema.org/url> ?url}
		
		UNION
		
		{?s <http://schema.org/programName> ?progname   . 
		?s <http://schema.org/hostingOrganization> ?ho .
		?ho <http://schema.org/url> ?url
		}
		
		UNION
		
		{?s <http://schema.org/description> ?description .
		?s <http://schema.org/url> ?url
		}
}

`

// SPres SPARQL call results
type SPres struct {
	Subject string
	Desc    string
	Type    string
	Pred    string
	Obj     string
}

// SPARQLCall calls triple store and returns results
func DoCall(url string) SPres {
	repo, err := sparql.NewRepo("http://0.0.0.0:7777/blazegraph/namespace/ecrwg/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("orgInfo", struct{ URL string }{url})
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)
	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	data := SPres{}
	bindingsTest2 := res.Bindings() // map[string][]rdf.Term

	data.Desc = "No description provided by facility"
	if len(bindingsTest2) > 0 {
		data.Subject = bindingsTest2["s"][0].String()
		if len(bindingsTest2["desc"]) > 0 {
			data.Desc = bindingsTest2["desc"][0].String()
		}
		data.Pred = bindingsTest2["pred"][0].String()
		data.Obj = bindingsTest2["obj"][0].String()
	}

	return data
}
