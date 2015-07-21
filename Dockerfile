FROM debian:jessie

RUN apt-get update && apt-get install -y \
  fio \
  --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

COPY tests /tests

RUN mkdir /mountpoint

WORKDIR /mountpoint

ENTRYPOINT [ "fio" ]
