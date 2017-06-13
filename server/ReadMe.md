## Server

#### About
A simple package that builds into a docker image for serving 
out the results of the working group crawl.

### Notes

```
docker run -d -p 9999:9999 opencoredata/blazegraph:0.2
```

```
docker run --name my-virtuoso \
    -p 8890:8890 -p 1111:1111 \
    -e DBA_PASSWORD=myDbaPassword \
    -e SPARQL_UPDATE=true \
    -e DEFAULT_GRAPH=http://www.example.com/my-graph \
    -v /my/path/to/the/virtuoso/db:/data \
    -d tenforce/virtuoso
```