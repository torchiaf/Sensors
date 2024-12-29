FROM golang:alpine AS build
ADD . /sensors
WORKDIR /sensors
RUN go build -o init

FROM alpine
WORKDIR /controller
COPY --from=build /sensors /controller

CMD ["./init"]
