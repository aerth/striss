# icecast-stress
Stress Test your Icecast server in Go

Modified from [afriza's netTestShoutcast](https://github.com/afriza/netTestShoutcast)

## Works great. Here is how to go get it.

```
go get -v -u github.com/hazrd/icecast-stress
```

## Usage
```
icecast-stress

Usage: ./icecast-stress <host.server.com:port> <num_conn> [interval]
Defaults to mountpoint: /stream -- If you need, go in and rebuild!
```

```
icecast-stress yourserver.com:8000 5 # start small.
icecast-stress yourserver.com:8000 9000 # go hard!
```


## Bugs

* Cant change mountpoint via CLI, have to rebuild with it hardcoded. Ohwell.
* When going over 9000 it seems to have a tough time responding to Ctrl+C
