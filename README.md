# iapc
Alternative client and Go library for Google Cloud's Identity-Aware Proxy. This is based on an implementation in [gartnera/gcloud](https://github.com/gartnera/gcloud) and the official gcloud source code. It aims to expose more functionality and have better performance. It's worth mentioning that the IAP speaks a slightly modified version of the SSH Relay v4 protocol documented [here](https://chromium.googlesource.com/apps/libapps/+/HEAD/nassh/docs/relay-protocol.md#corp-relay-v4).

## Usage
The CLI needs to acquire Application Default Credentials (ADC) to authenticate with the proxy, so make sure you're logged in.

Here's an example of how to create a tunnel to an instance.

```sh
$ iapc to-instance prod-1 --project analog-figure-330721 --zone europe-west2-a
```

Here's an example of how to create a tunnel to a private IP or FQDN in a VPC. This requires BeyondCorp Enterprise and a TCP Destination Group.

```sh
$ iapc to-host prod-1 --project analog-figure-330721 --region europe-west2 --network prod --dest-group prod
```