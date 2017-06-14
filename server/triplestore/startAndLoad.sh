#!/bin/ash

# never lets go, so script never moves on...
# "${cmd}" &>/dev/null &disown
# & (the first one) detaches the command from stdin.
# >/dev/null detaches the shell session from stdout and stderr.
# &disown removes the command from the shell's job list.
# java -server -Xmx4g -jar blazegraph.jar &>/dev/null &disown
java -server -Xmx4g -jar blazegraph.jar &>/dev/null  &
# java -server -Xmx4g -jar blazegraph.jar  &
 
sleep 30s

FILE_OR_DIR=$1

if [ -f "/etc/default/blazegraph" ] ; then
    . "/etc/default/blazegraph" 
else
    JETTY_PORT=9999
fi

LOAD_PROP_FILE=/tmp/$$.properties

export NSS_DATALOAD_PROPERTIES=./rwg.properties

#Probably some unused properties below, but copied all to be safe.
# set namepsace to kb for default namespace

cat <<EOT >> $LOAD_PROP_FILE
quiet=false
verbose=0
#closure=false
#durableQueues=true
#Needed for quads
defaultGraph=http://opencoredata.org/ecrwg
com.bigdata.rdf.store.DataLoader.flush=false
com.bigdata.rdf.store.DataLoader.bufferCapacity=100000
com.bigdata.rdf.store.DataLoader.queueCapacity=10
#Namespace to load
namespace=ecrwg
#Files to load
fileOrDirs=$1
#Property file (if creating a new namespace)
propertyFile=$NSS_DATALOAD_PROPERTIES
EOT

echo "Loading with properties..."

cat $LOAD_PROP_FILE

curl -X POST --data-binary @${LOAD_PROP_FILE} --header 'Content-Type:text/plain' http://localhost:${JETTY_PORT}/blazegraph/dataloader

#Let the output go to STDOUT/ERR to allow script redirection

rm -f $LOAD_PROP_FILE
