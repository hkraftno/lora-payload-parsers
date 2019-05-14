GPC_PROJECT_ID=hkraft-iot
SERVICE_NAME=lora-payload-parsers
CONTAINER_NAME=eu.gcr.io/$(GPC_PROJECT_ID)/$(SERVICE_NAME)

run: build
	docker run -p 8080:8080 $(CONTAINER_NAME)
build:
	docker build -t $(CONTAINER_NAME) .
deploy: build
	docker push $(CONTAINER_NAME)
	gcloud beta run deploy $(SERVICE_NAME) --project $(GPC_PROJECT_ID) --allow-unauthenticated -q --region us-central1 --image $(CONTAINER_NAME)
test:
	go test ./...
