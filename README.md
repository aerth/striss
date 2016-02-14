# icecast-stress

Stress Test Real Icecast server

Slightly modified [netTestShoutcast](https://github.com/afriza/netTestShoutcast)

## Works great. Here is how to go get it.

```
go get -v -u github.com/hazrd/icecast-stress

```

## Usage

```
icecast-stress

Usage: icecast-stress <host.server.com:port> <mountpoint> <num_conn> [interval]

Example: # icecast-stress example.com:8000 mount 5 600 # five connections, mountpoint /mount
Example: # icecast-stress example.com:8080 stream3 40 600 # forty connections, mountpoint /stream3


```

```
icecast-stress yourserver.com:8000 5 # start small.
icecast-stress yourserver.com:8000 9000 # go hard!

```


## Bugs

* When going high numbers like 600 it seems to have a tough time responding to Ctrl+C

* icecast-stress will panic if you dont do it correctly. "future: add error handling"

* Report any new ones you find
