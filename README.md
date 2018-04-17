# tproxy

Simple SOCKS5 proxy with user-password authentication.

[Docker Hub](https://hub.docker.com/r/akhmetov/tproxy/)

## Start

To start just run:

```sh
docker run --rm \
           --name socks5 \
           -p 0.0.0.0:1080:1080 \
           -e SOCKS5_LISTEN=0.0.0.0:1080 \
           -e SOCKS5_USERNAME=user_1 \
           -e SOCKS5_PASSWORD=pass_1 \
           akhmetov/tproxy
```

Do not forget to allow this port (`1080`) in your providers firewall.

## Build docker container

To build a container and start:

```sh
make build-docker

docker run --rm \
           --name socks5 \
           -p 0.0.0.0:1080:1080 \
           -e SOCKS5_LISTEN=0.0.0.0:1080 \
           -e SOCKS5_USERNAME=user_1 \
           -e SOCKS5_PASSWORD=pass_1 \
           socks5
```

And now you can connect to `127.0.0.1:1080` with `user_1:pass_1`.
