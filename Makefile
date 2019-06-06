.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/candidates/addCandidate candidates/addCandidate.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/candidates/listCandidates candidates/listCandidates.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/candidates/getCandidate candidates/getCandidate.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/candidates/removeCandidate candidates/removeCandidate.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
