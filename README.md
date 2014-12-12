lipservice
==========

A simple TCP server that returns a preset response for a given port.

Takes a json from STDIN.

~~~shell
% echo '{ "51234": "HELLO", "55556": "HI" }' | json | go run lipservice.go
2014/12/11 18:06:48 starting service at 51234...
response: HELLO...
2014/12/11 18:06:48 starting service at 55556...
response: HI...
~~~

Then on the client side,
~~~bash
% telnet localhost 51234
Trying ::1...
Connected to localhost.
Escape character is '^]'.
fsdfd
HELLOfsdfsf
HELLOfsdfds
HELLOjljlkj
HELLO
HELLOfsdfdsfsd
~~~

~~~bash
% telnet localhost 55556
Trying ::1...
Connected to localhost.
Escape character is '^]'.
fdfsdfds
HI
~~~
