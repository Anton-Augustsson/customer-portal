FROM golang:1.19 AS build
WORKDIR /go/src
COPY src ./src
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o robot_remote_controller_service .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/robot_remote_controller_service ./
EXPOSE 50052/tcp
ENTRYPOINT ["./robot_remote_controller_service"]
