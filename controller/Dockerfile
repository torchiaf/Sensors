FROM golang:alpine AS build
ADD . /sensors
RUN cd /sensors && go build -o controller

FROM alpine
WORKDIR /controller
COPY --from=build /app /controller

ENTRYPOINT ["./controller"]
