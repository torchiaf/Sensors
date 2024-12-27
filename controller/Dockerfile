FROM golang:alpine AS build
ADD . /sensors
RUN cd /sensors && go build -o controller

FROM alpine
WORKDIR /controller
COPY --from=build /sensors /controller

CMD ["./controller"]
