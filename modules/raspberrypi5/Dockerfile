FROM ubuntu

WORKDIR /app

COPY --from=build /dist/init /app

CMD ["./init"]
