# arko

ARKO is an SDK for API-driven CI/CD pipelines.

Read up on the what and why: https://writetogether.space/posts/vladfr/my-new-software-project


```
# Generate protos
make proto

# Run the services
go run master/main.go
go run slave/main.go

# Master will list all registered slaves and their services/methods
# Tell master to start a pipeline execution
grpcurl -d '{"method": "MyPipeline.Run"}' -plaintext -proto execution/execution.proto 127.0.0.1:10001 Execution/ExecuteJob
```

TODO
=====

#### Basic stuff:
* ~~DONE: use reflection on master to list gRPC methods in each slave~~
* use viper for configs and make sure each flag actually works
* ~~DONE: report and act on connection errors to slaves on Master~~

#### Master-slave discovery
* Master should save a list of slaves to file, and try to find them at start
* Slave should try to reconnect to master whenever Master goes away
* slaveList on master needs to be deduped, OR when a new slave registers, we need to check to see if slave is already registered (based on host:port)

#### Executing Jobs on slaves
* ~~Master needs a grpc method to call a slave (execution Service) - DONE~~
* the Execution service needs to call a Job Scheduler
* The Job Scheduler needs to receive the list of slaves with all their methods and it needs to schedule a job on one of the slaves
* Scheduler opens a connection to the slave, runs the job and waits for a reply
* Execution service receives the response from the Scheduler and prints it