FROM scratch
LABEL maintainer = "Hector Lovo <lovohh@gmail.com>"


# Copies the "production" version of the app into container.
# NOTE: This assumes that the lovohhwebapi binary exists locally.
COPY ./lovohhwebapi /

ENTRYPOINT ["/lovohhwebapi"]

EXPOSE 3000
