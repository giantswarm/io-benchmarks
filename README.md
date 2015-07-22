# I/O Benchmarks

A project to perform I/O benchmarks.

## Usage

```
docker run --rm -it \
  -v $(pwd):/mountpoint \
  registry.giantswarm.io/giantswarm/io-benchmarks \
  run test <test>.fio
```
