package sparql

import (
	"bytes"
	"log"

	"github.com/knakk/sparql"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: orgInfo
PREFIX schemaorg: <http://schema.org/>
SELECT DISTINCT ?repository ?name ?url ?logo ?description ?contact_name ?contact_email ?contact_url ?contact_role
WHERE {
  {
     ?repository schemaorg:url <{{.URL}}> .
  }
  UNION
  {
     ?repository <http://schema.org/url> "{{.URL}}" .
  }
  ?repository rdf:type <http://schema.org/Organization>   .
  ?repository schemaorg:name ?name .
  ?repository schemaorg:url ?url .
  OPTIONAL { ?repository schemaorg:description ?description . }
  OPTIONAL { ?repository schemaorg:logo [ schemaorg:url ?logo ] . }
  OPTIONAL {
    ?repository schemaorg:contactPoint ?contact .
    ?contact schemaorg:name ?contact_name .
    ?contact schemaorg:email ?contact_email .
    ?contact schemaorg:contactType ?contact_role .
    ?contact schemaorg:url ?contact_url .
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
// ?repository ?name ?url ?logo ?description ?contact_name ?contact_email ?contact_url ?contact_role
type SPres struct {
	Repository   string
	Name         string
	URL          string
	Logo         string
	Description  string
	ContactName  string
	ContactEmail string
	ContactURL   string
	ContactRole  string
}

// SPARQLCall calls triple store and returns results
func DoCall(url string) (SPres, error) {
	data := SPres{}
	repo, err := sparql.NewRepo("http://rwgsparql:9999/blazegraph/namespace/ecrwg/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
		return data, err
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("orgInfo", struct{ URL string }{url})
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
		return data, err
	}

	res, err := repo.Query(q)
	if err != nil {
		log.Printf("query call: %v\n", err)
		return data, err
	}

	bindingsTest2 := res.Bindings() // map[string][]rdf.Term

	// This whole aspect seems verbose... there has to be a better Go way to do this check?
	data.Description = "No description provided by facility"
	if len(bindingsTest2) > 0 {
		data.Repository = bindingsTest2["repository"][0].String()
		if len(bindingsTest2["description"]) > 0 {
			data.Description = bindingsTest2["description"][0].String()
		}
		if len(bindingsTest2["name"]) > 0 {
			data.Name = bindingsTest2["name"][0].String()
		}
		if len(bindingsTest2["url"]) > 0 {
			data.URL = bindingsTest2["url"][0].String()
		}
		if len(bindingsTest2["logo"]) > 0 {
			data.Logo = bindingsTest2["logo"][0].String()
		}
		if len(bindingsTest2["contact_name"]) > 0 {
			data.ContactName = bindingsTest2["contact_name"][0].String()
		}
		if len(bindingsTest2["contact_email"]) > 0 {
			data.ContactEmail = bindingsTest2["contact_email"][0].String()
		}
		if len(bindingsTest2["contact_url"]) > 0 {
			data.ContactURL = bindingsTest2["contact_url"][0].String()
		}
		if len(bindingsTest2["contact_role"]) > 0 {
			data.ContactRole = bindingsTest2["contact_role"][0].String()
		}
	}

	return data, err
}
