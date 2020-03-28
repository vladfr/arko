# arko

don't know what this is yet

```
# Generate protos
make proto

# Run the services
go run master/main.go
go run slave/main.go
```

TODO
=====
* DONE: use reflection on master to list gRPC methods in each slave
* use viper for configs and make sure each flag actually works
* DONE: report and act on connection errors to slaves on Master
* Master should save a list of slaves to file, and try to find them at start
* Slave should try to reconnect to master whenever Master goes away
* slaveList on master needs to be deduped, OR when a new slave registers, we need to check to see if slave is already registered (based on host:port)