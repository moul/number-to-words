NAME =		number-to-words
SOURCE :=	$(shell find . -name "*.go")


all: build


$(NAME): $(SOURCE)
	go build -o ./$(NAME) ./cmd/$(NAME)/main.go


.PHONY: build
build: $(NAME)


.PHONY: docker
docker:
	docker build -t moul/number-to-words .


.PHONY: test
test:
	go test -v .


.PHONY: cover
cover: profile.out


profile.out: $(SOURCE)
	rm -f $@
	go test -covermode=count -coverpkg=. -coverprofile=$@ .
