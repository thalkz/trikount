FROM golang:1.18-alpine as builder
RUN apk add build-base
WORKDIR app/
COPY . .
RUN go mod download
RUN go build -o /main

FROM alpine as runner
COPY ./www /www
COPY ./assets /assets
COPY --from=builder /main /server/main
ENTRYPOINT ["/server/main"]