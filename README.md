# iapc
Alternative client and Go library for Google Cloud's Identity-Aware Proxy. This is based on an implementation in [gartnera/gcloud](https://github.com/gartnera/gcloud) and the official gcloud source code. It aims to expose more functionality and have better performance. It's worth mentioning that the IAP speaks a slightly modified version of the SSH Relay v4 protocol documented [here](https://chromium.googlesource.com/apps/libapps/+/HEAD/nassh/docs/relay-protocol.md#corp-relay-v4).

## Usage
The CLI needs to acquire Application Default Credentials (ADC) to authenticate with the proxy, so make sure you're logged in.

If you want to create a tunnel to an instance, use the `instance` subcommand.

```sh
$ iapc instance prod-1 --project analog-figure-330721 --zone europe-west2-c --listen 127.0.0.1:1337
```

If you want to create a tunnel to an IP or FQDN in a VPC, use the `host` subcommand. This requires BeyondCorp Enterprise and a TCP Destination Group.
