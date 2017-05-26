# go_vs_sanic

Just a simple Docker image that has a Go http server and a Python sanic server

* Build the image: `docker build -t go_vs_sanic .`
* Run the Go server: `docker run --rm -it -p 8080:8080 go_vs_sanic run_go_app`
* Run the Go server: `docker run --rm -it -p 8080:8080 go_vs_sanic run_py_app`

## Workers number

The number of workers should match the logical CPUs on the host

* To set the number of workers on python http server see `app/py/go_vs_sanic.py`
* To set the number of 'workers' on go http server see `app/go/src/go_vs_sanic/main.go` -- `GOMAXPROCS`

## Some results

Server machine: AWS EC2 c4.xlarge  
Number of workers: 4

Client machine: AWS EC2 c4.xlarge  
Http client: wrk

### Python server - wrk -t4 -c1000 -d1m

```
  4 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   261.71ms  353.91ms   2.00s    85.27%
    Req/Sec     2.75k     1.04k    5.06k    73.30%
  Latency Distribution
     50%   86.07ms
     75%  287.67ms
     90%  794.54ms
     99%    1.57s 
  429882 requests in 1.00m, 52.48MB read
  Socket errors: connect 0, read 0, write 0, timeout 721
Requests/sec:   7152.93
Transfer/sec:      0.87MB
```

### Go server - wrk -t4 -c1000 -d1m

```
  4 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.60ms   10.20ms 454.96ms   94.91%
    Req/Sec    17.70k     1.23k   45.69k    96.54%
  Latency Distribution
     50%   13.23ms
     75%   14.45ms
     90%   20.20ms
     99%   34.06ms
  4226119 requests in 1.00m, 503.79MB read
Requests/sec:  70397.99
Transfer/sec:      8.39MB
```
