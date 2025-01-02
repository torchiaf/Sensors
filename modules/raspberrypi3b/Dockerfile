FROM python:3.9.21-alpine

WORKDIR /app

COPY --from=build /dist/init /app

CMD ["./init"]
