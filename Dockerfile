FROM golang:1.21-alpine as builder
RUN apk add build-base
WORKDIR app/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /main

FROM alpine as runner
RUN mkdir /home/data
WORKDIR /home
COPY ./www ./www
COPY ./assets ./assets
COPY --from=builder /main ./main
ENTRYPOINT ["./main"]