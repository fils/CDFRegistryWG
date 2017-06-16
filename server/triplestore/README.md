########### Get Repository info

```
PREFIX schemaorg: <http://schema.org/>
SELECT DISTINCT ?repository ?name ?url ?logo ?description ?contact_name ?contact_email ?contact_url ?contact_role
WHERE {
  {
     ?repository schemaorg:url <http://www.bco-dmo.org> .
  }
  UNION
  {
     ?repository <http://schema.org/url> "http://www.bco-dmo.org" .
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
```

```
########### Get Repository Endpoints
PREFIX schemaorg: <http://schema.org/>
SELECT DISTINCT ?repository ?endpoint_url ?endpoint_description ?endpoint_method
WHERE {
  {
     ?repository schemaorg:url <http://www.bco-dmo.org> .
  }
  UNION
  {
     ?repository <http://schema.org/url> "http://www.bco-dmo.org" .
  }
  ?repository rdf:type <http://schema.org/Organization>   .
  OPTIONAL {
    ?repository schemaorg:potentialAction [ schemaorg:target ?action ] .
    ?action schemaorg:urlTemplate ?endpoint_url .
    ?action schemaorg:description ?endpoint_description .
    ?action schemaorg:httpMethod ?endpoint_method .
  }
}
```