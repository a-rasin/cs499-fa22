# Part 1

## Collected data
### 1 webserver
anazir01@node5:~/wrk$ wrk -t10 -c399 -d5s --timeout 2s --latency http://node0/
Running 5s test @ http://node0/
  10 threads and 399 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   124.89ms  280.14ms   1.86s    87.12%
    Req/Sec     1.76k   602.73     5.21k    76.49%
  Latency Distribution
     50%    7.79ms
     75%    8.72ms
     90%  567.51ms
     99%    1.19s
  88128 requests in 5.10s, 40.59MB read
  Socket errors: connect 0, read 0, write 0, timeout 126
  Non-2xx or 3xx responses: 1
Requests/sec:  17282.05
Transfer/sec:      7.96MB

### 2 webservers
Running 5s test @ http://node0/
  10 threads and 399 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   145.65ms  303.04ms   1.87s    85.83%
    Req/Sec     1.78k   767.89    11.45k    83.83%
  Latency Distribution
     50%    7.85ms
     75%    8.40ms
     90%  653.25ms
     99%    1.23s
  88558 requests in 5.10s, 40.79MB read
  Socket errors: connect 0, read 0, write 0, timeout 111
  Non-2xx or 3xx responses: 4
Requests/sec:  17364.88
Transfer/sec:      8.00MB

### 3 webservers
Running 5s test @ http://node0/
  10 threads and 399 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   125.65ms  282.39ms   1.87s    87.10%
    Req/Sec     1.81k   621.90     7.79k    76.29%
  Latency Distribution
     50%    7.66ms
     75%    8.04ms
     90%  570.77ms
     99%    1.19s
  90468 requests in 5.10s, 41.67MB read
  Socket errors: connect 0, read 0, write 0, timeout 104
  Non-2xx or 3xx responses: 6
Requests/sec:  17739.91
Transfer/sec:      8.17MB

### 4 webservers
Running 5s test @ http://node0/
  10 threads and 399 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   133.95ms  288.51ms   1.87s    86.43%
    Req/Sec     1.76k   529.09     4.57k    73.16%
  Latency Distribution
     50%    7.80ms
     75%    8.80ms
     90%  606.91ms
     99%    1.19s
  88176 requests in 5.10s, 40.61MB read
  Socket errors: connect 0, read 0, write 0, timeout 92
Requests/sec:  17289.53
Transfer/sec:      7.96MB
