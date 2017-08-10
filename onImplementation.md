# Approaches to implementing JSON-LD microdata

## Abstract

The basic approach to including JSON-LD based HTML5 microdata in a web page is to simply
include a ``` <script type="application/ld+json"> ``` tag in the body of the page. 
Normally this is done within the ```<head>``` tag of the page but in practice can be anywhere in the 
document.

For a live example you can visit the [Open Core Data](http://opencoredata.org) page and 
do a "view source" on the page.  The contents of the ``` <script type="application/ld+json"> ```
is JSON-LD.  Examples of JSON-LD can be seen in this repository in the form of
the files:

* [BCO-DMO](https://github.com/fils/CDFRegistryWG/blob/master/bcodmo.json)  
* [OpenCore](https://github.com/fils/CDFRegistryWG/blob/master/opencore.json)  

and others.

These are JSON-LD files that would be scoped by the script tag.  Feel free to raise any questions
about these in the issues section of this repository.  

## Simple example with testing option

A complete, though simple, example of this process is fully scoped below to demonstrate the approach.  
This text can be pasted into the 
[Google testing tool](https://search.google.com/structured-data/testing-tool) to validate.  Note, it will error on the example.org URL only due to the fact this is not a valid URL and the content is not coming 
directly from that domain.   For more detailed examples of JSON-LD look at the examples in this repository.


```
<!doctype html>
<html lang="en">
<head>
    <title>Example Org</title>
    <script type="application/ld+json">
        {
            "@context": "http://schema.org/",
            "@type": "Organization",
            "name": "Example Org",
            "contactPoint": {
                "@type": "ContactPoint",
                "name": "John Doe",
                "email": "jdoe@example.org",
                "url": "http://example.org/person/jdoe",
                "contactType": "technical support"
            },
            "url": "http://www.example.org",
            "potentialAction": [{
                "@type": "SearchAction",
                "target": {
                    "@type": "EntryPoint",
                    "urlTemplate": "http://example.org/apidocs.json",
                    "description": "Swagger 1.2 description document",
                    "httpMethod": "GET"
                }
            }]
        }
    </script>

</head>
<body>
  My great facility data
</body>
</html>

```

For more information on JSON-LD as well as tools and examples 
visit the [JSON-LD Playground](http://json-ld.org/playground/).


### Some early tests / linting
The following links show results of tests on the opencore.json file.  Note that the 
URL error at the Google tool is related to hosting the file at a domain different than what it referees to.

#### Results from Google Structure Data Testing Tool
[View results](https://search.google.com/structured-data/testing-tool#url=https%3A%2F%2Fraw.githubusercontent.com%2Ffils%2FCDFRegistryWG%2Fmaster%2Fopencore.json)

#### Structured Data Linter
[View results](http://linter.structured-data.org/?url=https:%2F%2Fraw.githubusercontent.com%2Ffils%2FCDFRegistryWG%2Fmaster%2Fopencore.json)

