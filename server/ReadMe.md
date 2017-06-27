## Server

#### About
A simple package that builds into a docker image for serving 
out the results of the working group crawl.

### Notes
Build Go binary for Linux for use in Docker

```
 env GOOS=linux go build .
 docker build -t earthcube/cdfrwgweb:0.5 .
```


```
docker-compose up -d
```

```
docker-compose down
```

Transferer command
```
docker save earthcube/cdfrwgweb:0.8 | bzip2 |  ssh fils@eccdf.cloudapp.net 'bunzip2 | docker load'
```


Since I am using docker-compose I need to reference the servers via the docker network name as defined
in the docker-compose.yml file.  Like for the SPARQL calls in sparql.go.   I should pass these in as a 
parameter at compile time to support non-docker runs.  


Blazegraph 
```
docker run -d -p 9999:9999 opencoredata/blazegraph:0.2
```

From playing with virtuoso
```
docker run --name my-virtuoso \
    -p 8890:8890 -p 1111:1111 \
    -e DBA_PASSWORD=myDbaPassword \
    -e SPARQL_UPDATE=true \
    -e DEFAULT_GRAPH=http://www.example.com/my-graph \
    -v /my/path/to/the/virtuoso/db:/data \
    -d tenforce/virtuoso
```


### SPARQL

```
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
```

```
SELECT DISTINCT *
WHERE {
  ?s <http://schema.org/url> <http://www.bco-dmo.org> .
  optional {?s <http://schema.org/description> ?desc } .
  optional {?s rdf:type ?type } .
  ?s ?pred ?obj   
}

```