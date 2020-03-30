# arko

ARKO is an SDK for API-driven CI/CD pipelines.

Read up on the what and why: https://writetogether.space/posts/vladfr/my-new-software-project

Development
====

arko uses protobuf v3 with custom tags so get ready to use them:
```
# Mac
brew install protobuf

# Other OSes see https://github.com/protocolbuffers/protobuf and https://github.com/golang/protobuf

# Get go support for proto3
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# This package adds our custom tags to generated .pb.go files
go get github.com/favadi/protoc-go-inject-tag
```

Ok! Now you're ready to generate the protos and start the master and slave.

```
# Generate protos
make proto

# Run the master
go run master/main.go

# Run slave; it will register to master, which by default runs on localhost
go run slave/main.go

# Master should then list all registered slaves and their services/methods
```

For now there is no UI. To run a pipeline you need to call the master's Execution service:
```
# Tell master to start a pipeline execution
# cd master/ - master doesn't have reflection turned on so you need to give it the path to the proto file
grpcurl -d '{"method": "MyPipeline.Run"}' -plaintext -proto execution/execution.proto 127.0.0.1:10001 Execution/ExecuteJob
```

TODO
=====

#### Basic stuff:
* ~~DONE: basic gRPC setup for Master and slave - DONE~~
* ~~DONE: master should ping slaves~~
* ~~DONE: use reflection on master to list gRPC methods in each slave~~
* use viper for configs and make sure each flag actually works

#### Master/slave improvements
* ~~DONE: report and act on connection errors to slaves on Master~~
* Slaves should report Master connection errors (and retry connecting?) - currently Slaves don't do Ping, maybe they should?
* Slave should try to reconnect to master whenever Master goes away (?)
* ~~DONE slaveList on master needs to be deduped, OR when a new slave registers, we need to check to see if slave is already registered (based on host:port)~~
* In pingSlaves the Master needs to ask for the Slave status and update it. Slave should report 1 for OK and 0 for not OK. Right now, all Slaves should reply with 1; 
* in pingSlaves, if a slave doesn't reply, Master should set Status=0

#### Persistence
* ~~DONE Master should save a list of slaves to file, and try to find them at start~~
* ~~DONE use bbolt with Storm ORM~~
* add BBolt Backup endpoint for DB backups:  https://github.com/etcd-io/bbolt#database-backups
* write some tests with DB mock: https://www.alexedwards.net/blog/organising-database-access (see under "Using an interface" for a pattern)

#### Executing Jobs on slaves
* ~~DONE: Master needs to save methods on Slaves~~
* ~~Master needs a grpc method to call a slave (execution Service) - DONE~~
* the Execution service needs to call a Job Scheduler
* The Job Scheduler needs to receive the list of slaves with all their methods and it needs to schedule a job on one of the slaves
* The Scheduler needs to ask the Slave if it can accept the Job
* Scheduler opens a connection to the slave, runs the job and waits for a reply
* Execution service receives the response from the Scheduler and prints it

#### Authentication
* all gRPC comms should be done over TLS (is this enough? do we need an extra token for slaves to auth?)
* Do slaves need AuthZ at all?
* slaves should not accept any connections from anywhere else than master

#### The UI
* build a simple UI with React/Vue
* it should connect to the master
* list all slaves with their methods
* tell master to run a method
* wait for and print the result

#### Example simple pipeline
* Dogfood: create a slave that can build Docker images of Arko master and our slave
* and that can deploy master?