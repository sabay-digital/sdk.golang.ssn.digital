FROM docker-registry.sabay.com/docker/golang1.12/build:latest

WORKDIR $GOPATH/src/git.sabay.com/payment-network/sdk/sdk.go-cashier.ssn.digital

COPY . .

RUN go get -d -v ./... \
    && go install -v ./...