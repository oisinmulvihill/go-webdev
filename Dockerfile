FROM golang:alpine as build-env
WORKDIR /app
COPY cmd/ cmd/
RUN go build -o /app/main cmd/main.go

FROM alpine as runtime
WORKDIR /app
COPY --from=build-env /app/main /app/web-server
USER 1000:1000
EXPOSE 8080
CMD [ "/app/web-server" ]
