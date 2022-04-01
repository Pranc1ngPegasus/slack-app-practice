.PHONY: prepare-tools
prepare-tools:
	go get github.com/google/wire/cmd/wire
	go get github.com/cosmtrek/air

.PHONY: wiregen
wiregen: prepare-tools
	go run github.com/google/wire/cmd/wire gen ./...

.PHONY: dev
dev:
	go run github.com/cosmtrek/air
