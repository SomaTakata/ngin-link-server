ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd


FROM alpine:latest

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
