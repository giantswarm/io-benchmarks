# I/O Benchmarks

A project to perform I/O benchmarks.

## Usage

```
docker run --rm -it \
  -v /:/host \
  --workdir=/host/$(pwd) \
  giantswarm.io/io-benchmarks /tests/<test_file>.fio
```
