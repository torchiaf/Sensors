ARG BUILD_IMAGE=no-image

FROM ${BUILD_IMAGE} AS build

FROM ubuntu

WORKDIR /app

COPY --from=build /dist/init /app

CMD ["./init"]
