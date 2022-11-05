# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine AS build

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./

# compile application, disables CGO for scratch image
RUN go build -o ./go-api
# RUN CGO_ENABLED=0 go build -o ./go-api

# FROM scratch
FROM alpine:3.16.2
WORKDIR /
COPY --from=build /app/go-api /go-api
EXPOSE 8080
ENTRYPOINT ["/go-api"]