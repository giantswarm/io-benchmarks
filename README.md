# I/O Benchmarks

A project to perform I/O benchmarks.

## Usage

```
docker run --rm -it \
  -v $(pwd):/mountpoint \
  -v $(pwd)/io-benchmark-results:/results \
  registry.giantswarm.io/giantswarm/io-benchmarks \
  test <test>.fio
```
