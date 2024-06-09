.PHONY: all_protos proto_gus_client proto_ggs_client proto_gas_client build login join add status commit reset

all_protos: proto_gus_client proto_ggs_client proto_gas_client

proto_gus_client:
	rm -f services/gus/*.go
	protoc --proto_path=../shared/protos/user-service \
	       --go_out=services/gus \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=services/gus \
	       --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
	       ../shared/protos/user-service/*.proto

proto_ggs_client:
	rm -f services/ggs/*.go
	protoc --proto_path=../shared/protos/git-service \
	       --go_out=services/ggs \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=services/ggs \
	       --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
	       ../shared/protos/git-service/*.proto

proto_gas_client:
	rm -f services/gas/*.go
	protoc --proto_path=../shared/protos/assignment-service \
	       --go_out=services/gas \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=services/gas \
	       --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
	       ../shared/protos/assignment-service/*.proto

# running code..
# NUe24HV7

autoClone ?= false

build:
	go build -o gitversity .

login: build
	./gitversity login --email ${email} --password ${password}

logout: build
	./gitversity logout 

join: build
	./gitversity join --auto-clone ${autoClone} --invite-code ${inviteCode}

clone: build 
	./gitversity clone --assignment_id ${assignmentId}

add: build
	cd ${directory} && ../gitversity add ${args}

commit: build
	cd ${directory} && ../gitversity commit --message ${message}

status: build
	cd ${directory} && ../gitversity status

push: build
	cd ${directory} && ../gitversity push

reset: build 
	cd ${directory} && ../gitversity reset ${args}
	
help: build 
	./gitversity ${command} --help
