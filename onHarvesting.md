# Review of activity related to harvesting


### About
We are looking to conduct a simple test with a set of providers.  This test will involve
them publishing facility data using the approaches defined in the onImplementation.md file 
and elsewhere in this repo.

Then, using [https://github.com/fils/contextBuilder](https://github.com/fils/contextBuilder) a
set of RDF files (nquads format) will be feed into a triple store and used for basic queries.


### Issues
The biggest issue is the fact that default JSON-LD results in blank nodes in a RDF representation.  
While this is fine in terms of SPARQL where blank nodes are effectively variables it is not so good in a 
LOD approach.  However, through the use of @id in JSON-LD blank nodes can be removed.  This 
make the JSON-LD a bit more involved to author but only by 1 entry per type.  A bigger issue though is
that by declaring these ID (URIs) we become responsible for them.  Thus introducing a URI that requires
"attention" in a LOD context.  

### Example results
A simple run of this has already been done.  The resulting triples were queried with the following SPARQL which 
was crafted to extract contact point information in this case.  

```
prefix schema: <http://schema.org/>
select ?name ?pred ?obj
where {
  ?s schema:contactPoint ?cpoint .
  ?s schema:name ?name .
  ?cpoint ?pred ?obj
 }
 ```

 and results in:

 
| name	 | pred	 | obj |
| ------- | ------ | ------- | ------ |
| BCO-DMO	 | schema:contactType	 | technical support |
| BCO-DMO	 | schema:name	 | Adam Shepherd |
| BCO-DMO	 | schema:email	 | ashepherd@whoi.edu |
| BCO-DMO	 | schema:url	 | <http://orcid.org/0000-0003-4486-9448> |
| BCO-DMO	 | rdf:type	 | schema:ContactPoint |
| Open Core Data	 | schema:contactType	 | technical support |
| Open Core Data	 | schema:email	 | dfilsAToceanleadershipDOTorg |
| Open Core Data	 | schema:name	 | Douglas Fils |
| Open Core Data	 | schema:url	 | <http://orcid.org/0000-0002-2257-9127> |
| Open Core Data	 | rdf:type	 | schema:ContactPoint |
| UNAVCO	 | schema:contactType	 | technical support |
| UNAVCO	 | schema:email	 | jrileyATunavcoDOTorg |
| UNAVCO	 | schema:name	 | Jim Riley |
| UNAVCO	 | schema:url	 | <http://orcid.org/0000-0001-8163-5662> |
| UNAVCO	 | rdf:type	 | schema:ContactPoint |


