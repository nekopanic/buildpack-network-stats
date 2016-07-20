Dyno network stats!

Print out network statistics every 30 seconds or so:

```
2016-07-20T06:26:13.369455+00:00 app[web.1]: source=web.1 sample#iface=eth0 sample#bytes_recv=5166 sample#bytes_sent=7768 sample#packets_recv=16 sample#packets_sent=18 sample#recv_error=0 sample#sent_error=0 sample#recv_drop=0 sample#sent_drop=0
2016-07-20T06:26:13.369477+00:00 app[web.1]: source=web.1 sample#iface=lo sample#bytes_recv=0 sample#bytes_sent=0 sample#packets_recv=0 sample#packets_sent=0 sample#recv_error=0 sample#sent_error=0 sample#recv_drop=0 sample#sent_drop=0
```

The statistics are differential, printing only the amount changed since the last time it checked. You can adjust the interval like this: `heroku config:set DYNO_NETWORK_STATS_INTERVAL=60`. Bytes, packets, errors and drops are described here: http://www.onlamp.com/pub/a/linux/2000/11/16/LinuxAdmin.html

Usually it will print statistics on two interfaces:

* "eth0" - which is the dyno communicating outside.
* "lo" - the local loopback interface, used by some apps for internal dyno networking

# Hack it!

* `go get github.com/c9s/goprocinfo/linux`
* Edit [dyno-network-stats.go](dyno-network-stats.go)
* `go build -o dyno-network-stats`
