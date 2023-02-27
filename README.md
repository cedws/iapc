# goiap
Alternative client and Go library for Google Cloud's Identity-Aware Proxy. This is based on an implementation in [gartnera/gcloud](https://github.com/gartnera/gcloud) and the official gcloud source code. It aims to expose more functionality and have better performance. I'd like to add connection pooling at some point too.

## Example
The tool needs to acquire Application Default Credentials (ADC) to authenticate with the proxy, so make sure you're logged in.

If you want to create a tunnel to an instance, use the `instance` subcommand. If you want to create a tunnel to an IP or FQDN in a VPC, use the `host` subcommand. This requires BeyondCorp Enterprise and you will need to create a destination group first.

```sh
$ goiap instance prod-1 --project analog-figure-330721 --zone europe-west2-c --listen 127.0.0.1:1337
```
