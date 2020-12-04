# cf8k
----
Catfeeder 8k is a news agregator, written in Go for a few selected
news web sites that does not use a news feed, only by means of html parsing.

## Requirements
- go (1.14, tested might work with other versions)
- make

## How to build
To compile just run:
```bash
make
```

## How to run
To run it after it has been compiled it is possible to just go the build
directory and then do:
```bash
./cf8k
```

But it is also possible to run it via make by setting the variable NEWS to the
newsreader you want to run just like:
```bash
NEWS=globo make run
```

To see documentation for exported parts of the code type:
```bash
make doc
```
