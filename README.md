[![Build Status](https://travis-ci.com/marek2901/golang-channels.svg?branch=master)](https://travis-ci.com/marek2901/golang-channels)

# GO csv import fun
> I was having some fun with sync.WaitGroup go feature.

> This program creates inserts records from csv electricity-consumption-by-sectors.csv to local sqlite db

## Requirements

* golang
* sqlite3 binary
* *nix os 

> you can still run on windows but there are only bash scripts so you'll end up doing everything manually

## Run

git clone, cd into project dir and run

```
make
```

poor's man debugging mode
```
CSV_ER_DEBUG=true make test
```

## License
MIT or do whatever you want :D
