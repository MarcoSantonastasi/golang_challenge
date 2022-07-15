FROM golang:1.18-bullseye

RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN apt-get update && apt-get install -y --no-install-recommends \
    protobuf-compiler \
    && apt-get clean

RUN dogo install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR $GOPATH/src
COPY . .

RUN go build -v -o $GOPATH/bin/ ./...

CMD ["server"]