# EarthCube CDF Registry Working Group

## TLDR;
The work of the registry working group can be summed up rather quickly.  Use existing 
vocabularies like schema.org and re3data terms to expose facility metadata using web architecture
patterns.   Leverage HTML5 microdata publishing, JSON-LD and standard web architecture (hypermedia) 
to both expose and collect metadata. 

## About
The EarthCube Council of Data Facilities (CDF) formed the Registry Working Group to review alignment of existing approaches to research facility description and discovery.  The involved parties include the EarthCube CDF, Coalition for Publishing Data in the Earth and Space Sciences (COPDESS) and the Registry of Research Data Repositories (re3data).   

* [EarthCube CDF](https://www.earthcube.org/group/council-data-facilities)
* [Re3data](http://www.re3data.org/) and  [RE3 schema](http://www.re3data.org/schema) 
* [COPDESS](http://www.copdess.org/)


## Documents

* [Members](members.md)
* [Presentation](./docs/ESIPSemCommJune2017/talkBody.pdf)
* [Poster](./docs/DataOneESIP_poster.pdf)
* [Harvesting (old)](onHarvesting.md)
* [Implementation (old)](onImplementation.md)


## Repository structure

* [JSON-LD Docuements](./jsonldDocuments)  A collection of JONS-LD documents being used
to test ideas and use of the schema.org and re3data types and terms.
* [Documentation](./docs)  Assorted presentations and posters.
* [Notebooks](./notebooks) A simple notebook (Jupyter) to demonstrate a potential 
where more human approachable formats like YAML allow people to more easily create
example JSON-LD documents for reference. 
* [Server code](./server)  The Go based code for hosting the test interface and triple store
This is the service available at [repograph.net](http://repograph.net/)
* [Schema Builder](./schemaBuilder) Related to the "notebooks" above this is a thought about
creating a method to allow more human approachable schema.org building.  Like what can be seen
at [Structured Markup Editor](http://www.stoumann.dk/examples/editor/) but focused on CDF needs.
  

#### Simple Scenario 

1. A facility has both metadata about the facility as well as links to service description 
documents like Swagger, OGC or Threads.  
2. These are assembled together into a JSON-LD document following schema.org patterns with possible
use of external vocabularies.  This is then placed into the facility landing page (or other designated page) via 
```
    <script type="application/ld+json">
```
3. Items that can not be defined by schema.org can be then be defined via an external vocabulary
4. The white list of site/URLs is feed through something like [https://github.com/fils/contextBuilder](https://github.com/fils/contextBuilder) or by DateOne tools.  This example code will look for schema.org JSON-LD packages defined in item 2.  More advanced crawling solutions might use tools like: https://github.com/anaskhan96/soup or https://github.com/PuerkitoBio/fetchbot 

After reading in the JSON-LD it could be converted to RDF for a triple store 
or other data storage or index approaches used by a harvesting group.   
There is no blessed harvesting or presentation site.  Any number of groups or organizations 
could harvest and provide access to this material. 

The following image gives a brief overview of how facilities might take their descriptor
documents and metadata and expose this material up through a workflow to aggregation 
and interface clients.  


![Image of Flow](./Images/bubbles.png)


## Errata 
### On ad hoc implementation 
As noted a test crawler, harvester and indexer is being developed at 
[contextBuilder](https://github.com/fils/contextBuilder).  This is a simple (and not 
production ready) application for harvesting from a whitelist and extracting the JSON-LD
package.  The next step will be to convert this JSON-LD to triples and moved into a standard 
triple store.  A focused JSON-LD crawler is also in development at 
[https://github.com/ESIPFed/snapHacks/tree/master/sh01-jsonldCrawl](https://github.com/ESIPFed/snapHacks/tree/master/sh01-jsonldCrawl)

### On external vocabularies
The registryC5 file is testing some external vocabulary uses.  It is valid JSON-LD but 
Google will always through an error since it doesn't see this as a property of some
known schema.org class.  This should be fine and I have tested this, but it is always
a worry with Google that you will not know when how they deal with this case
will be changed.   Their typical response has been, "try and get things you need 
in core schema.org".  

