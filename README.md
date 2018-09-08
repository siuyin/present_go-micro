# Trying out Go Micro
I present my go-micro experiments,
deployed the "go modules"(aka vgo) way.

## If running in a docker container

```
expose port 3999 in your container.

present -http 0.0.0.0:3999 -orighost {exposed ip address}

eg.
present -http 0.0.0.0:3999 -orighost 192.168.99.100
```
