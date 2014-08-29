HTTP server for network testing

Replies with "pong\n" in text/plain for every request

## usage

```
$ http-pong-server -h
Usage of http-pong-server:
  -listen=":8080": listen directive
```

Demo use case with [http-stress-test](https://github.com/sebcat/http-stress-test)

```
$ nohup http-pong-server &
$ http-stress-test -rate=2000 -duration=3 -url="http://127.0.0.1:8080/foo"
total 1496 (0 failed) time: 327.338967ms avg: 218.809us

```

