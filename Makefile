apps = 'derasure'

wire:
	wire ./...

generate:
	go generate ./...

run: wire generate
	for app in $(apps) ;\
	do \
		 go run ./cmd/$$app -f configs/$$app.yml  & \
	done

build-in-docker:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build  -ldflags="-s -w" -o dist/$$app ./cmd/$$app/; \
	done

build-docker-images: wire generate
	 docker build -f build/Dockerfile -t kwstars/derasure .
