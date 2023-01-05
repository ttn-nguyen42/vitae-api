FROM golang:1.19.4
WORKDIR /go/src/Vitae/
RUN go get -d -v github.com/gin-gonic/gin@v1.8.2 \
    github.com/rs/zerolog@v1.28.0 \
    go.mongodb.org/mongo-driver@v1.11.1
COPY . /go/src/Vitae/
RUN CGO_ENABLED=0 go build -a -o server .

FROM alpine:latest

ENV GIN_MODE=debug
ENV PORT=8080
ENV LOG_LEVEL=DEBUG
ENV MONGO_USERNAME=
ENV MONGO_PASSWORD=
ENV MONGO_CLUSTER_ID=

RUN apk add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/Vitae/server ./
CMD ["./server"]
