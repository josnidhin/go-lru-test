FROM golang:1.21.0-bookworm as builder

ENV BUILDER_APP_DIR=/app/srv/test-2

WORKDIR ${BUILDER_APP_DIR}
COPY go.mod go.sum ./

COPY . .

RUN make release-build

## Build the production image
FROM alpine:3.18.4

ENV BUILDER_APP_DIR=/app/srv/test-2
ENV APP_DIR=/app/srv/test-2
ENV APP_NAME=test-2

RUN apk update \
  && rm -rf /var/cache/apk/* \
  && adduser -D -h /home/jose jose

WORKDIR ${APP_DIR}

COPY --from=builder ${BUILDER_APP_DIR}/${APP_NAME} .

RUN chown -R jose:jose $APP_DIR

USER jose

CMD ${APP_DIR}/${APP_NAME}
