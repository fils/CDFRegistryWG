# Review of activity related to harvesting

### About
We are looking to conduct a simple test with a set of providers.  This test will involve
a facility publishing metadata using the approaches defined in the onImplementation.md file 
and elsewhere in this repo.  This metadata will contain general information about the facility 
and links to descriptor documents for the services the facility provides.

Then, using [https://github.com/fils/contextBuilder](https://github.com/fils/contextBuilder) a
set of RDF files (nquads format) will be feed into a triple store and used for basic queries.

![havesting diagram](./Images/harvest.png)

## Thoughts
There are several approaches to extracting the triples from the JSON-LD.   In this early
testing we are justing some Go RDF libraries.   However, another group doing this might also be 
interested in using something like the Apache Any23 or other such projects to extract the 
triples.

### JSON-LD and blank nodes
An issue is the fact that it's easy to generate a JSON-LD document 
that results in blank nodes in a RDF representation.  
While this is the issue in SPARQL where blank nodes are effectively variables this is not 
a major issue.   It could be more of a factor in LOD approaches where blank nodes
are going to obviously impact resource dereferencing on the net.  
However, through the use of @id in JSON-LD blank nodes can be removed.  This 
make the JSON-LD a bit more involved to author but only by 1 entry per type.  There are a lot 
of item type uses though. 

This potential solution raises a new issue in that by declaring these ID (URIs) we 
become responsible for them.  Thus introducing a URI that requires "attention" and 
maintenance in a LOD context.  Further testing of the approach will reveal if this is a major
factor or not in implementation patterns.  

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

 and results in

![Result Table](./Images/resultTable.png)


### Visual Graph 
A simple visual of some of the relations in the graph.

![Result Table](./Images/CDFRWGgraph.png)


