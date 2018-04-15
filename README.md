# Redirect Server

Implementation of a simple redirect server written in golong. The
server simply replies to all incoming HTTP requests with redirect
(status code 301) to a configured URL.

The executable is statically linked and can therefore be executed on
any linux distribution without installing any dependencies or libraries.

## Usage

When starting the server, the redirect url must be specified as
paramter. By default, the server starts on port 80 (make sure to run
the executable as root or with the according privileges).  It is also
possible to optionally specify a different port as first parameter.


```
./server [PORT] REDIRECT_URL
```

## Examples

To redirect all incoming HTTP requests on port 80 to the golang
hompage, the following command may be used:

```
./server https://golang.org/
```

To redirect all incoming HTTP requests on port 1234 to the golang
hompage, the following command may be used:

```
./server 1234 https://golang.org/
```