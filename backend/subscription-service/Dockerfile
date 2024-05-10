FROM golang:1.19 AS build
WORKDIR /go/src
COPY src ./src
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o subscription_service .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/subscription_service ./
EXPOSE 50051/tcp
ENTRYPOINT ["./subscription_service"]
