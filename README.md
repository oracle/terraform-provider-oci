## NOTICE  
**The OCI Terraform Provider is now available for automatic download through the Terraform Provider Registry. 
For more information on how to get started view the [documentation](https://www.terraform.io/docs/providers/oci/index.html) 
and [setup guide](https://www.terraform.io/docs/providers/oci/guides/version-3-upgrade.html).**


Terraform Provider for Oracle Cloud Infrastructure
==================

- [Documentation](https://www.terraform.io/docs/providers/oci/index.html)
- [OCI forums](https://cloudcustomerconnect.oracle.com/resources/9c8fa8f96f/summary)
- [Github issues](https://github.com/terraform-providers/terraform-provider-oci/issues)
- [Troubleshooting](https://www.terraform.io/docs/providers/oci/guides/guides/troubleshooting.html)

[![wercker status](https://app.wercker.com/status/666d2ee10f45dde41189bb03248aadf9/s/master "wercker status")](https://app.wercker.com/project/byKey/666d2ee10f45dde41189bb03248aadf9)


Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) v0.10.1 or greater
- [Go](https://golang.org/doc/install) 1.11.4 (recommended)

Note: You may use any version 1.11 or above to build the provider. However, the `goimports`, `go vet`, and `gofmt` code checks will only pass when using version 1.11.


Building the Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-oci`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-oci
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-oci
$ make build
```


Using the Provider
----------------------
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) 
After placing it into your plugins directory,  run `terraform init` to initialize it and begin using Terraform with the Oracle Cloud Infrastructure provider.


Troubleshooting the Provider
----------------------
To get verbose console output when the provider is running, precede your Terraform command with the `TF_LOG` and `OCI_GO_SDK_DEBUG` flags:
```sh
TF_LOG=DEBUG OCI_GO_SDK_DEBUG=1 terraform plan
```

The [tf_log](https://www.terraform.io/docs/internals/debugging.html) level and `OCI_GO_SDK_DEBUG` flags can also be set as environment variables.


Developing the Provider
---------------------------

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
