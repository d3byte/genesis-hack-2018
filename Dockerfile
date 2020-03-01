FROM golang:alpine

RUN apk add --no-cache curl git
RUN curl -fsSL -o /bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /bin/dep

ADD . $GOPATH/src/google
ADD token.json /tmp/token.json
ADD credentials.json /tmp/credentials.json
WORKDIR $GOPATH/src/google

RUN dep ensure
RUN go build .

CMD ["./google"]
EXPOSE 8080
