ARG BUILD_IMAGE=no-image

FROM ${BUILD_IMAGE} AS build

FROM python:3.9.21-alpine

WORKDIR /app

COPY --from=build /dist/init /app

CMD ["./init"]
