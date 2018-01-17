FROM golang:alpine
LABEL maintainer = "Hector Lovo <lovohh@gmail.com>"


ENV app_dir /go/bin/

WORKDIR ${app_dir}
# Copies the "production" version of the app into /go/bin.
# NOTE: This assumes that the lovohhwebapi binary exists locally.
COPY ./lovohhwebapi ${app_dir}

ENTRYPOINT ${app_dir}/lovohhwebapi

EXPOSE 3000
