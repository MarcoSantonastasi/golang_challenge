FROM golang:1.18-bullseye

RUN apt-get update && apt-get install -y --no-install-recommends \
    protobuf-compiler \
    && apt-get clean

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && export PATH="$PATH:$(go env GOPATH)/bin"

COPY ./scripts/ /docker-entrypoint-initdb.d/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]