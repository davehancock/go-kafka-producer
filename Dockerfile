FROM golang:1.11 as build

# Using go modules so want to avoid being inside $GOPATH
WORKDIR /src/app

# Building librdkafka from source as having issues with latest pkgs in stretch (and issues with libssl1.0.0 pkg w/ confluent dist)
# We need this low level wrapper for confluent-go-kafka
RUN git clone https://github.com/edenhill/librdkafka.git
RUN cd ./librdkafka &&  ./configure --prefix /usr && make && make install

COPY . .

RUN go version
RUN env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags static -o /app 

FROM gcr.io/distroless/base
COPY --from=build /app /app
COPY config.yml /
CMD ["/app"]
