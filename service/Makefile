.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/uploadS3 uploadS3/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/removeS3 removeS3/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/list list/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/triggerS3 triggerS3/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/analyze analyze/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/afteranalyze afteranalyze/main.go afteranalyze/models.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	npx serverless deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
