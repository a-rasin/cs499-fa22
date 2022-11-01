# Part 1

## Report
90%    4.56ms
99%    4.89ms
Requests/sec:  21388.43

90%    4.82ms
99%    5.15ms
Requests/sec:  32924.07

90%    3.34ms
99%    4.56ms
Requests/sec:  46348.97

90%    2.81ms
99%    5.05ms
Requests/sec:  59684.51

## Collected data
### 1 webserver
Running 5s test @ http://node0/
  10 threads and 99 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.18ms  326.71us   9.88ms   74.21%
    Req/Sec     2.16k    72.26     2.66k    96.06%
  Latency Distribution
     50%    4.16ms
     75%    4.39ms
     90%    4.56ms
     99%    4.89ms
  109083 requests in 5.10s, 50.24MB read
Requests/sec:  21388.43
Transfer/sec:      9.85MB

### 2 webserver
Running 5s test @ http://node0/
  10 threads and 99 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.72ms    1.67ms   8.54ms   49.28%
    Req/Sec     3.31k   118.09     3.67k    66.86%
  Latency Distribution
     50%    2.82ms
     75%    4.39ms
     90%    4.82ms
     99%    5.15ms
  167932 requests in 5.10s, 77.35MB read
Requests/sec:  32924.07
Transfer/sec:     15.16MB

### 3 webserver
Running 5s test @ http://node0/
  10 threads and 99 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.94ms    1.05ms   8.81ms   65.39%
    Req/Sec     4.68k   267.38     6.14k    80.71%
  Latency Distribution
     50%    1.86ms
     75%    2.65ms
     90%    3.34ms
     99%    4.56ms
  236342 requests in 5.10s, 108.85MB read
Requests/sec:  46348.97
Transfer/sec:     21.35MB

### 4 webserver
Running 5s test @ http://node0/
  10 threads and 99 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.56ms    1.01ms   8.65ms   73.90%
    Req/Sec     6.01k   312.09     6.83k    66.01%
  Latency Distribution
     50%    1.36ms
     75%    2.07ms
     90%    2.81ms
     99%    5.05ms
  304417 requests in 5.10s, 140.21MB read
Requests/sec:  59684.51
Transfer/sec:     27.49MB

