# FROM  java:8
FROM jetty:9.4.2-jre8-alpine

LABEL authors="dfils@oceanleadership.org"

ENV BLAZEGRAPH_VERSION 2_1_5
ENV BLAZEGRAPH_VERSION_URL https://github.com/blazegraph/database/releases/download/BLAZEGRAPH_RELEASE_CANDIDATE_2_1_5/blazegraph.jar

RUN mkdir /blazegraph
ADD . /blazegraph

# Create a non-root user to run as
# RUN groupadd -r gorunner -g 433 && \
# mkdir /home/gorunner && \
# useradd -u 431 -r -g gorunner -d /home/gorunner -s /sbin/nologin -c "User to run go apps on high ports" gorunner && \
# chown -R gorunner:gorunner  /home/gorunner && \
# chown -R gorunner:gorunner /blazegraph

# set user
# USER gorunner

# Move to a workign directory for running codices so it can see it's static files
# future version should take this as a param so static content can be anywhere
RUN apk --no-cache add wget
RUN apk --no-cache add curl

WORKDIR /blazegraph
RUN wget -O blazegraph.jar $BLAZEGRAPH_VERSION_URL
#RUN java -server -Xmx4g -jar blazegraph.jar
RUN  /blazegraph/startAndLoad.sh RDF

# Build indexes?  https://wiki.blazegraph.com/wiki/index.php/Rebuild_Text_Index_Procedure

# Exec on start
ENTRYPOINT /usr/bin/java -server -Xmx4g -Djetty.overrideWebXml=./readonly.xml  -jar blazegraph.jar

EXPOSE 9999
