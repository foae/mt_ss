all: lint test install_mtss

install_mtss:
	go install github.com/foae/mtss/cmd/mtss

lint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;
	which gometalinter; if [ $$? -ne 0 ]; then go get -u github.com/alecthomas/gometalinter && gometalinter --install; fi
	gometalinter --vendor --exclude=repos --disable-all --enable=golint ./...
	go vet ./...

test:
	go test -v -short -cover ./...

run: install_mtss
	ENV="dev" \
	SUBREDDIT="romania_ss" \
	REDDIT_AGENT_CREDENTIALS_FILE_NAME="reddit_credentials" \
	$(GOPATH)/bin/mtss
