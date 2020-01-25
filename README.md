# Port Scanner with Go
Easily scan and return open ports on your system.

## Use Cases
I found it frustrating to hunt down available ports when working in my dev environment, sometimes wasting time debugging code before realizing there is already a service using that port!  😓

Using `portscan` you can quickly confirm availability right away for a desired port, or return the next closest port.

You can also pipe the results through to a command by using the -n flag to ensure an open port returns. Useful for automated scripts to launch processes.
```
python3 -m http.server $(./portscan -pn 6666 -n)
node http-server -p $(./portscan -p 6666 -n)
```

## How to build from source
```
$ git clone https://github.com/ncolletti/portscan
$ cd portscan
$ export GO111MODULE=on
$ go build

```

## How to run binary from any location
```
$ mkdir ~/bin
$ mv portscan ~/bin
BASH: export PATH=$PATH:/home/<user>/bin
ZSH: path+=('/home/<user>/bin')
```

## How to use
```
Usage: portscan -p 8888 -n -nt tcp6

portscan -p - port - Port to check if available
portscan -n - next - Return next available closest port
portscan -nt - network - Specify a network - 'tcp'(default) 'tcp4' 'tcp6'
portscan -v - verboose - Enable more command feedback
```


