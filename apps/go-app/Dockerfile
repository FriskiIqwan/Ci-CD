# Build the Go server
FROM golang:1.23-bookworm AS server
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . /go/src/app
RUN CGO_ENABLED=0 go build -o build/go-app main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=server /go/src/app/build/go-app .
CMD ["./go-app"]