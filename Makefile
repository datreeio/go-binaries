build:
	go build

release:
	goreleaser release --rm-dist
