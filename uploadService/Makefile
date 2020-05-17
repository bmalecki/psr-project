.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/uploadS3 uploadS3/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	npx serverless deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
