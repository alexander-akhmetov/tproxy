# tproxy

Simple SOCKS5 proxy with user-password authentication.

## Start docker container

To start on your own computer build a Docker container:

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

Do not forget to allow this port in your providers firewall: `1080`.
