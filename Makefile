GOPKG ?= moul.io/number-to-words
GOBINS ?= ./cmd/ntw-web ./cmd/number-to-words
DOCKER_IMAGE ?= moul/number-to-words

all: test install

include rules.mk
