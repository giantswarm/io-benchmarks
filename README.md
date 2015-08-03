# I/O Benchmarks

A project to perform I/O benchmarks.

The tool makes heavy use of [fio](https://github.com/axboe/fio) it's scripts to
plot graphs via `gnuplot`.

## Tests

Shipped with the source you can find some `fio` job definitions in the the `tests/`
directory. Most of them are seperated for write and read specific I/O to test
them individually. All jobs focused on read specific I/O start with `10*_*`
while write specific jobs start with `20*_*`. Jobs starting with `00*_*` read
a small file to test if the configuration works, not to benchmark the system.

## Usage

To run benchmarks you need to have `fio`, `fio2gnplot` and `gnuplot` available
on your system. You can also use Docker to it which has all dependencies
installed:

```
docker run --rm -it \
  -v $(pwd)/.io-benchmark:/mountpoint \
  -v $(pwd)/io-benchmark-results:/results \
  registry.giantswarm.io/giantswarm/io-benchmarks \
  test <test>.fio
```
