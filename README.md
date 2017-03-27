# CDFRegistryWG


## About
Some test files to explore the use of schema.org objects and properties to 
describe data repositories with.

#### Facility examples
* opencore.json : Example of encoding Open Core services into this schema.org
* bcodmo.json : Example for BCO-DMO
* iris.json : Example for IRIS with a mix of machine and human focused resources


#### Candidate encodings :  Some test of various encoding ideas
* registryC5.json : A test of multiple context entries..  to allow us to use
other context for things like controlled voc. 
* registryC3.json : A test to see how use of subOrganization would work for 
groups that scope several projects 
* registryC2.json : Uses the SearchAction object.  Best candidate so far.
* registryC1.json : Early test... used DigitalDocument which is likely not a proper 
use of this type.  However, it might apply to the VoID document in some ways.   
 
## Notes
A reference work flow would go something like this.


* Make a simple SKOS or OWL out of RE3 XML schema terms
* Use schema.org/Organization plus external voc from above (valid JOSN-LD)
* Sites place this JSON-LD into their index page or some other page they designate.
* The white list of site/URLs is feed through something like [https://github.com/fils/contextBuilder](https://github.com/fils/contextBuilder) or by DateOne tools.  This example code will look for schema.org JSON-LD packages defined 
by a line like
```
    <script type="application/ld+json" class="cdfregistry">
```
* After reading in the JSON-LD it could be converted to RDF for a triple store 
or other store approaches.

## On external vocabularies
registryC5 is testing some external vocabulary uses.  It is valid JSON-LD but 
Google will always through an error since it doesn't see this as a property of some
known schema.org class.  This should be fine and I have tested this, but it is always
a worry with Google that you will not know when how they deal with this case
will be changed.   Their typical response has been, "try and get things you need 
in core schema.org".  


## Tests
The following links show results of tests on the opencore.json file.  Note that the 
URL error at the Google tool is related to hosting the file at a domain different than what it referees to.

#### Results from Google Structure Data Testing Tool
[View results](https://search.google.com/structured-data/testing-tool#url=https%3A%2F%2Fraw.githubusercontent.com%2Ffils%2FCDFRegistryWG%2Fmaster%2Fopencore.json)

#### Structured Data Linter
[View results](http://linter.structured-data.org/?url=https:%2F%2Fraw.githubusercontent.com%2Ffils%2FCDFRegistryWG%2Fmaster%2Fopencore.json)

