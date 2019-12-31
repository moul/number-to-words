# build
FROM            golang:1.13-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/number-to-words
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM            alpine:3.11
COPY            --from=builder /go/bin/number-to-words /bin/
ENTRYPOINT      ["/bin/number-to-words"]
CMD             []
