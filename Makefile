GPC_PROJECT_ID=hkraft-prod-com-089eacdc
SERVICE_NAME=lora-payload-parsers
CONTAINER_NAME=eu.gcr.io/hkraft-ci/$(SERVICE_NAME)
REGION=europe-west1
URL=lora-payload-parsers.hkraft.dev

run: build
	docker run -p 8080:8080 $(CONTAINER_NAME)
build:
	docker build -t $(CONTAINER_NAME) .
deploy: build
	docker push $(CONTAINER_NAME)
	gcloud run deploy $(SERVICE_NAME) --project $(GPC_PROJECT_ID) --allow-unauthenticated -q --region $(REGION) --image $(CONTAINER_NAME)
test:
	go test ./...
domain-mapping:
	gcloud beta run domain-mappings create --platform managed  --region $(REGION) --service $(SERVICE_NAME)  --domain $(URL) --project $(GPC_PROJECT_ID)
