# CDFRegistryWG


## registryC*.json
Some test files to explore the use of schema.org objects and properties to 
describe data repositories with.

* registryC5.json : A test of multiple context entries..  to allow us to use
other context for things like controlled voc. 
* registryC3.json : A test to see how use of subOrganization would work for 
groups that scope several projects 
* registryC2.json : Uses the SearchAction object.  Best candidate so far.
* registryC1.json : Early test... used DigitalDocument which is likely not a proper 
use of this type.  However, it might apply to the VoID document in some ways.   
* opencore.json : Example of encoding Open Core services into this schema.org
 
## Notes
The registryC2.json is the best candidate so far.  It is encoding the Open Core Data
information right now.  A copy of this is in opencore.json as well.  A reference work flow would go something like this.

* Sites place this JSON-LD into their index page or some other page they designate.
* The white list of site/URLs is feed through something like [https://github.com/fils/contextBuilder](https://github.com/fils/contextBuilder) or by DateOne tools.  This example code will look for schema.org JSON-LD packages defined 
by a line like
```
    <script type="application/ld+json" class="cdfregistry">
```
* After reading in the JSON-LD it could be converted to RDF for a triple store 
or other store approaches.


## Tests
The following links show results of tests on the opencore.json file.  Note that the 
URL error at the Google tool is related to hosting the file at a domain different than what it referees to.

#### Results from Google Structure Data Testing Tool
[View results](https://search.google.com/structured-data/testing-tool#url=https%3A%2F%2Fraw.githubusercontent.com%2Ffils%2FCDFRegistryWG%2Fmaster%2Fopencore.json)

#### Structured Data Linter
[View results](http://linter.structured-data.org/?url=https:%2F%2Fraw.githubusercontent.com%2Ffils%2FCDFRegistryWG%2Fmaster%2Fopencore.json)

