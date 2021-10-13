build:
	go build

publish:
	goreleaser release --rm-dist
