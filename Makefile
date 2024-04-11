GEN_DIR=internal/gen
SERVICE_NAME = backend-service
SERVICE_PORT = 8888

swagger-gen:
	if ! [ -d $(GEN_DIR) ]; then \
	    mkdir $(GEN_DIR); \
	elif [ -d $(GEN_DIR) ]; then \
		rm -rf $(GEN_DIR); \
		mkdir $(GEN_DIR); \
	fi && \
	swagger generate server -t internal/gen -f ./api/swagger.yml --exclude-main -A $(SERVICE_NAME) && \
	go mod tidy && \
	git add $(GEN_DIR)

go-test:
	go test -timeout=10s -count=1 -v ./internal/...

go-run+build-dev:
	go build ./cmd/service/ &&\
	./service -config_path="$(PWD)/configs/dev.yml"

go-build-windows-dev:
	env GOOS=windows GOARCH=amd64 go build ./cmd/service/

docker-build:
	docker build -t $(SERVICE_NAME) .

docker-save:
	docker save $(SERVICE_NAME) > $(SERVICE_NAME).tar

docker-load:
	docker load < $(SERVICE_NAME).tar

docker-restart:
	sudo docker stop $(SERVICE_NAME) &&\
	sudo docker rm $(SERVICE_NAME) &&\
	make docker-run-stage

docker-build+restart:
	make docker-build &&\
	make docker-restart

docker-run-stage:
	docker run -v $(PWD)/user:/app/user -v $(PWD)/secret:/app/secret -v $(PWD)/configs/stage.yml:/app/configs/config.yml -v $(PWD)/logs:/app/logs --network host -p $(SERVICE_PORT):$(SERVICE_PORT) -d --restart=always --name $(SERVICE_NAME) $(SERVICE_NAME)

project-send-on-stage:
	make docker-build &&\
	make docker-save &&\
	tar -czvf $(SERVICE_NAME).tar.gz --files-from=./project_migration/stagelist.txt &&\
	scp $(SERVICE_NAME).tar.gz root@79.174.95.104:~/project/