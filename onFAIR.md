## Reviewing FAIR related to this work

### Notes
It would be good to show how this approach aligns with and assists with
being FAIR compliant.  This is the beginning of some notes related to that.


### Sections
The sections below reference the sections from the FAIR GUIDING PRINCIPLES
section at the bottom. 

#### section 1
To address FAIR section 1 I am looking at @id

```
@id
Used to uniquely identify things that are being described in the document with IRIs or blank node identifiers. This keyword is described in section 3.3 Node Identifiers.
```

So in opencore.json I can put in something like
```
    "@id" : "http://opencoredata.org/id/facilityinfo",
```
However, this means I have to honor this as a thing!  It might be nice if I could  point to 
something re3 provides or someone else that lets me uniquely ID this.  (part 1 fair)


#### section 2
We should be good here, need to demonstrate 

#### section 3
We should be good here, need to demonstrate

#### section 4
We should be ok with this section if we properly address section 1.   We should try and 
demonstrate interconnection with the GeoLink graph here.  Might be a nice place to use 
web components ref [geocomponents.org](http://geocomponents.org) like

1) generate a component that reads the facility JSON-LD document
2) have it look into GeoLink for connections of interest (FAIR 4.2)
3) present a card in the component hosting page that displays the facility info along with any 
further information from GeoLink  (people, awards, etc)

### FAIR
Ref: [https://www.force11.org/fairprinciples](https://www.force11.org/fairprinciples)



```
FAIR GUIDING PRINCIPLES

1. To be Findable any Data Object should be uniquely and persistently identifiable [4]
1.1. The same Data Object should be re-findable at any point in time, thus Data Objects should be persistent, with emphasis on their metadata, [4 and JDDCP 4 and JDDCP 6]
1.2. A Data Object should minimally contain basic machine actionable metadata that allows it to be distinguished from other Data Objects [see JDDCP 5]
1.3. Identifiers for any concept used in Data Objects should therefore be Unique and Persistent [5 and JDDCP 4 and JDDCP 6].

2. Data is Accessible in that it can be always obtained by machines and humans
2.1 Upon appropriate authorization [6]
2.2 Through a well-defined protocol [7 and JDDCP 5]
2.3 Thus, machines and humans alike will be able to judge the actual accessibilty of each Data Object.

3. Data Objects can be Interoperable only if:
3.1. (Meta) data is machine-actionable [8]
3.2. (Meta) data formats utilize shared vocabularies and/or ontologies [9]
3.3  (Meta) data within the Data Object should thus be both syntactically parseable and semantically machine-accessible [10]

4. For Data Objects to be Re-usable additional criteria are:
4.1 Data Objects should be compliant with principles 1-3
4.2 (Meta) data should be sufficiently well-described and rich that it can be automatically (or with minimal human effort) linked or integrated, like-with-like, with other data sources [11 and JDDCP 7 and JDDCP 8]
4.3 Published Data Objects should refer to their sources with rich enough metadata and provenance to enable proper citation (ref to JDDCP 1-3).




```