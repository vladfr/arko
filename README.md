# arko

ARKO is an SDK for API-driven CI/CD pipelines.

Read up on the what and why: https://writetogether.space/posts/vladfr/my-new-software-project


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
* Slaves should report Master connection errors (and retry connecting?)
* Master should save a list of slaves to file, and try to find them at start
* Slave should try to reconnect to master whenever Master goes away
* slaveList on master needs to be deduped, OR when a new slave registers, we need to check to see if slave is already registered (based on host:port)

#### Executing Jobs on slaves
* ~~Master needs a grpc method to call a slave (execution Service) - DONE~~
* the Execution service needs to call a Job Scheduler
* The Job Scheduler needs to receive the list of slaves with all their methods and it needs to schedule a job on one of the slaves
* Scheduler opens a connection to the slave, runs the job and waits for a reply
* Execution service receives the response from the Scheduler and prints it

#### Authentication
* all gRPC comms should be done over TLS (is this enough? do we need an extra token for slaves to auth?)
* Do slaves need AuthZ at all?
* slaves should not accept any connections from anywhere else than master

### The UI
* build a simple UI with React/Vue
* it should connect to the master
* list all slaves with their methods
* tell master to run a method
* wait for and print the result
