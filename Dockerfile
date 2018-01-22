FROM alpine:3.7
LABEL maintainer = "Hector Lovo <lovohh@gmail.com>"

RUN apk add --no-cache ca-certificates

# Copies the "production" version of the app into container.
# NOTE: This assumes that the lovohhwebapi binary exists locally.
COPY ./lovohhwebapi /home/

ENTRYPOINT ["/home/lovohhwebapi"]

EXPOSE 3000
