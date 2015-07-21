FROM debian:jessie

RUN apt-get update && apt-get install -y \
  fio \
  --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

RUN mkdir /host

COPY tests /tests

ENTRYPOINT [ "fio" ]
