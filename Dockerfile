FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

WORKDIR /app
COPY ../../go.mod go.sum ./
RUN apk add --update alpine-sdk figlet
RUN go mod download


# build
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -tags musl -o main .

# stage-2: image builder
FROM golang:1.21-alpine
WORKDIR /build
COPY --from=builder /app/main .

# run
ENTRYPOINT [ "/build/main" ,"hook-handler"]
