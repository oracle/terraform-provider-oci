# Terraform Provider for Oracle Cloud Infrastructure

**The OCI Terraform Provider is now available for automatic download through the Terraform Provider Registry.**

- [Documentation](https://www.terraform.io/docs/providers/oci/index.html)
- [Setup guide](https://www.terraform.io/docs/providers/oci/guides/version-3-upgrade.html)
- [Examples](https://github.com/oracle/terraform-provider-oci/tree/master/examples)
- [OCI forums](https://cloudcustomerconnect.oracle.com/resources/9c8fa8f96f/summary)
- [Github issues](https://github.com/oracle/terraform-provider-oci/issues)
- [Troubleshooting](https://www.terraform.io/docs/providers/oci/guides/troubleshooting.html)

[![wercker status](https://app.wercker.com/status/666d2ee10f45dde41189bb03248aadf9/s/master "wercker status")](https://app.wercker.com/project/byKey/666d2ee10f45dde41189bb03248aadf9)


## Requirements

- [Terraform](https://www.terraform.io/downloads.html) v0.12.31 or greater
- [Go](https://golang.org/doc/install) 1.20.7 (recommended)


## Building the Provider

Clone repository to: `$GOPATH/src/terraform-provider-oci`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-oci
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/terraform-provider-oci
$ make build
```


## Installation

If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) 
After placing it into your plugins directory,  run `terraform init` to initialize it and begin using Terraform with the Oracle Cloud Infrastructure provider.


## Troubleshooting the Provider

See [verbose logging](https://www.terraform.io/docs/providers/oci/guides/troubleshooting.html#verbose-logging-for-oci-terraform-provider) for the details.

## Developing the Provider

To add features to the provider, install [Go](http://www.golang.org) and configure your your [GOPATH](http://golang.org/doc/code.html#GOPATH)

Compile the provider by running `make build`. The provider binary will output to your `$GOPATH/bin` directory, make sure this has been added to your `$PATH`.

```sh
$ make build
```

To test the provider run `make testacc`.

```sh
$ make testacc
```

> **Note:** The tests run against live OCI service APIs, you will need to configure environment variables with valid credientials as shown in the [documentation](https://www.terraform.io/docs/providers/oci/index.html).

## Contributing

This project welcomes contributions from the community. Before submitting a pull request, please [review our contribution guide](./CONTRIBUTING.md)

## Security

Please consult the [security guide](./SECURITY.md) for our responsible security vulnerability disclosure process

## License

Copyright (c) 2017, 2023 Oracle and/or its affiliates.

Released under the Mozilla Public License 2.0
