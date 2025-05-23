FROM golang:alpine as build-env
WORKDIR /app
COPY .git/ .git/
COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod go.mod
COPY go.sum go.sum
RUN apk update && apk add git
COPY build.sh build.sh
RUN /bin/sh build.sh

FROM alpine as production
WORKDIR /app
COPY --from=build-env /app/web-server /app/web-server
USER 1000:1000
EXPOSE 8080
CMD [ "/app/web-server" ]
