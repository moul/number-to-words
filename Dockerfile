FROM golang:1.12
COPY . /go/src/github.com/moul/number-to-words
WORKDIR /go/src/github.com/moul/number-to-words
RUN make
ENTRYPOINT ["/go/src/github.com/moul/number-to-words/number-to-words"]
