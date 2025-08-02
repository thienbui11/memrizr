FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

<<<<<<< HEAD
RUN go install github.com/cespare/reflex@latest
=======
RUN go get github.com/cespare/reflex
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]