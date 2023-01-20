FROM golang:1.19.4
WORKDIR /go/src/Vitae/
RUN go mod init
RUN go mod tidy
COPY . /go/src/Vitae/
RUN CGO_ENABLED=0 go build -a -o server .

FROM alpine:latest

ENV GIN_MODE=debug
ENV PORT=
ENV LOG_LEVEL=
ENV MONGO_USERNAME=
ENV MONGO_PASSWORD=
ENV MONGO_CLUSTER_ID=

RUN apk add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/Vitae/server ./
CMD ["./server"]
