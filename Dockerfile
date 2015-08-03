FROM debian:jessie

RUN apt-get update && apt-get install -y \
  fio \
  gnuplot \
  --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

COPY tests /opt/io-benchmarks/tests
COPY io-benchmarks /opt/io-benchmarks/io-benchmarks

RUN mkdir /mountpoint
RUN mkdir /results

ENTRYPOINT [ "/opt/io-benchmarks/io-benchmarks", "--tests-directory=/opt/io-benchmarks/tests", "--working-directory=/mountpoint/.io-benchmark", "--output-directory=/results", "run" ]
