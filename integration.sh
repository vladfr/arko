#!/bin/bash

log () {
    echo "test >> $1"
}

MPORT=$(awk 'BEGIN{srand();print int(rand()*(33000-2000))+10000 }')
SPORT=$(awk 'BEGIN{srand();print int(rand()*(33000-2000))+10001 }')
DBNAME="arkotest.$MPORT.db"
if [ $MPORT -eq $SPORT ]; then
    log 'Failed to generate 2 random numbers'
    exit 1
fi

# this function will kill all bg jobs, but will not reap their children
# e.g. `go run` spawns a child, but does not reap it, so be careful!
function cleanup()
{
    rm $DBNAME
    log "Killing bg jobs and exiting..."
    jobs
    [[ -z "$(jobs -p)" ]] || kill -9 $(jobs -p)
}

trap cleanup EXIT

### Build
make build

## Run with fresh builds from bin (`make build` should put them there)
cd bin

### Master
log "Starting Master on port $MPORT"

bash -c "./master --port $MPORT --db $DBNAME" &
sleep 3

nc -w 1 127.0.0.1 $MPORT
retVal=$?

if [ $retVal -ne 0 ]; then
    log "Master failed to bind on port $MPORT"
    exit 1
fi
log "Master is up and running"

### Slave
log "Starting Slave on port $SPORT"
bash -c "./slave --port ${SPORT} --master \"127.0.0.1:${MPORT}\"" &
sleep 3

nc -w 1 127.0.0.1 $SPORT
retVal=$?

if [ $retVal -ne 0 ]; then
    log "Slave failed to bind on port $SPORT"
    exit 1
fi
log "Slave is up and running"

exit 0
