VERSION ?= latest
IMAGE := harbor.mammoth.com/mammoth:$(VERSION)

build:
	@docker build -t $(IMAGE) .

push:
	@docker push $(IMAGE)

start:
	@docker run -d -p 8080:8080 --name mammoth $(IMAGE)

stop:
	@docker rm -f mammoth

logs:
	@docker logs -f mammoth --tail=200