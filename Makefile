PACKAGES =	$(notdir $(wildcard ./cmd/*))
SOURCE :=	$(shell find . -name "*.go")


all: build


$(PACKAGES): $(SOURCE)
	go install -v ./cmd/$@


.PHONY: build
build: $(PACKAGES)


.PHONY: docker
docker:
	docker build -t moul/number-to-words .


.PHONY: test
test:
	go test -i .
	go test -v .


.PHONY: cover
cover: profile.out


profile.out: $(SOURCE)
	rm -f $@
	go test -covermode=count -coverpkg=. -coverprofile=$@ .
