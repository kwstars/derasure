apps = 'derasure'

run: wire
	for app in $(apps) ;\
	do \
		 go run ./cmd/$$app -f configs/$$app.yml  & \
	done

wire:
	wire ./...

build-in-docker:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -o dist/$$app ./cmd/$$app/; \
	done

build-docker-images: wire
	 docker build -f build/Dockerfile -t kwstars/derasure .
