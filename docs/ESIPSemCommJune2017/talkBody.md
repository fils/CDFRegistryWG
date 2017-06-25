class: center, middle, inverse

<img style="width:100px"  src="../../images/logo_earthcube_cube-only_0.png">

# Report on EarthCube CDF Registry Working Group
## Work to Date

.footnote[
  created with [remark](http://github.com/gnab/remark)
]


---

# Group Members

1. Doug Fils
2. Bob Arko
3. Adam Shephard
4. Shelley Stall
5. Danie Kinkade
6. Michael Witt
7. Lynne

Acknowledgement:  This work is support by NSF EarthCube and the ESSO office which provides
the infrastructure to support group interaction.  


---

# Talk Agenda

1. Introduction
2. Vocabulary element
3. Perspectives of involved parties
4. FAIR
5. Prototype infrastructure
6. Lessons learned
7. Future


---

# Intro and Functional goals

     
* Goal or working group  (voc alignement with re3data terms and CDF needs)
* Functional goals of the group (products we wanted to deliver)
* Structure of process (terms -> voc -> deployed arch -> search) 


---
# Early overview  
 
<img style="width:40%"  src="../../images/bubbles.png">


---
# Vocabulary elements



---
# Connections (re3data)



---
# Connections (COPDESS)


---
# Connections (Other: Google, Bing, Yandex, etc.)

The approach has benefit beyond EarthCube's goals.  It is an approach leveraging 
standards based approaches to organization description.  It's arguably both a metadata
publishing and outreach activity.


---
# FAIR



---
# Current Providers

1. Open Core Data
2. BCO-DMO
3. R2R
4. UNAVCO
5. IRIS
6. Open Topology
--

7. IEDA (in process)
8. WHOI (in process)

---
# Effort required by Providers

At the current scope effort is relatively small.  However, it will grow as we expend
exposed content.  Such expansion though can be controlled by the repository.


.small[
```html
<html lang="en">
<head>
    ...

    <script type="application/ld+json">
        {
            "@context": "http://schema.org/",
            "@type": "Organization",
            "name": "Open Core Data",
            "contactPoint": {
                "@type": "ContactPoint",
                ...
```
]


However, the JSON-LD has to be generated.  We have been working from templates, but 
that is unlikely to scale.   

Some thought needs to go into how to assist groups in generating this content if it is seen
as a way forward.  

---
# Assembling  (more on the big part)

Fortunately we have tools to help us build schema.org packages:

[https://json-ld.org/](https://json-ld.org/)

Fortunately this is very much in line with the sort of data pipelines
the data facilities are doing now.  

#### put an example link to some code here...  
---

# Harvesting

1. Build and install [aha](https://github.com/theZiz/aha):
.small[
```terminal
josh@brick ~/repos $ git clone https://github.com/theZiz/aha; cd aha
josh@brick ~/repos/aha $ make && cp aha /usr/local/bin
```
]

2. Capture output with `aha` (for dark background highlight.js styles such as
   solarized_dark, use `aha -b`):
.small[
```terminal
josh@brick ~/repos/aha $ git log -2 --color | aha -b -n | pbcopy
```
]
To capture directly to clipboard, use:
  - `pbcopy` for OS X
  - `xsel --clipboard` for Linux
  - `getclip` for Cygwin

---
# Harvesting

<img style="width:100%"  src="../../images/harvest.png">



---
# Search and visualize


---
# Lessons learned

* Ontology development elements
* issues related to JSON-LD and blank nodes
* issues about generating JSON-LD (publishing is the easy part)
* how to navigate the schema.org "graph"
* index and display elements

___
# Future

1. Extend to DataCatalog -> DataSet -> Measurements -> Variables
2. Employ emelments of Hydra?
3. How to best address blank nodes in the JSON-LD to RDF process?
4. Vocabulary gap issues identified?

---

# Thanks

1. Contact Us
2. Repo URL
3. Test site URL 




