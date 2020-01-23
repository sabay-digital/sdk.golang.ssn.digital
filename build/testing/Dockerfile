FROM docker-registry.sabay.com/docker/golang/build:latest

WORKDIR $GOPATH/src/git.sabay.com/payment-network/sdk/sdk.golang.ssn.digital

COPY . .

RUN go get -d -v ./... \
    && go install -v ./...