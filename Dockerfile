FROM golang:1.12
WORKDIR /go/src/github.com/moul/number-to-words
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN make
ENTRYPOINT ["/go/src/github.com/moul/number-to-words/number-to-words"]
