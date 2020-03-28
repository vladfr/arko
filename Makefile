proto:
	@echo 'Generating proto files'
	cd slave; protoc -I pipeline pipeline/pipeline.proto --go_out=plugins=grpc:pipeline
	cd master; protoc -I register register/register.proto --go_out=plugins=grpc:register
	cd master; protoc -I execution execution/execution.proto --go_out=plugins=grpc:execution