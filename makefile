SHELL := /bin/bash

all: api

api:
	docker build \
		-t ecom/api-amd64:1.0 \
		--build-arg PACKAGE_NAME=service \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.
	docker system prune -f

push:
	$$(aws ecr get-login --no-include-email --region us-east-1 --profile aat)
	docker tag ecom/api-amd64:1.0 905848599188.dkr.ecr.us-east-1.amazonaws.com/aat-api
	docker push 905848599188.dkr.ecr.us-east-1.amazonaws.com/aat-api

deploy:
	ssh -i "~/Downloads/aat.pem" ec2-user@ec2-3-85-36-183.compute-1.amazonaws.com
	
up:
	docker-compose up

down:
	docker-compose down

test:
	cd "$$GOPATH/src/github.com/souravToptal/ecom"
	go test ./...

clean:
	docker system prune -f

stop-all:
	docker stop $(docker ps -aq)

remove-all:
	docker rm $(docker ps -aq)


#EC2 Steps
# sudo yum update -y
# sudo yum install -y docker
# sudo service docker start
# sudo usermod -a -G docker ec2-user
# sudo curl -L https://github.com/docker/compose/releases/download/1.21.0/docker-compose-`uname -s`-`uname -m` | sudo tee /usr/local/bin/docker-compose > /dev/null
# sudo chmod +x /usr/local/bin/docker-compose
# sudo service docker start

# $(aws ecr get-login --no-include-email --region us-east-1)
# docker pull 905848599188.dkr.ecr.us-east-1.amazonaws.com/aat-api
# docker-compose down
# docker-compose up -d
