

build:
	GOARCH=amd64 GOOS=linux go build -v -o bin/terraform-provider-finger.v0.0.1.linux


test:
	docker run -it \
		-v $(shell pwd)/test-terraform:/tf \
		-w /tf \
		--entrypoint /bin/sh \
		hashicorp/terraform:0.11.13

apply:
	terraform init && terraform plan && terraform apply