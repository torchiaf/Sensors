FROM golang:alpine AS build
ADD . /build-dir
WORKDIR /build-dir
RUN go build -o init

FROM alpine
WORKDIR /controller
COPY --from=build /build-dir /controller

CMD ["./init"]
