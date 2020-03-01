# Build Google middleware

1. Install git and dep
    ```dockerfile
    RUN apk add --no-cache curl git
    RUN curl -fsSL -o /bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /bin/dep
    ```
2. Set ENV
    ```
    PORT=8080
    GIN_MODE=release
    ```
3. Build Golang project
    ```dockerfile
    ADD . $GOPATH/src/google-service
    WORKDIR $GOPATH/src/google-service

    RUN dep ensure
    RUN go build .
    CMD ["./google"]
    EXPOSE 8080
    ```
