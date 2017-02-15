# Development Guide

# Required Reading
Before starting development it's important to get familiarized with
Terraform and how it's providers work.

Here's a list of recommended materials:
* https://www.hashicorp.com/blog/terraform-custom-providers.html (Error
  Handling and partial state are very important concepts)
* https://www.terraform.io/docs/plugins/basics.html
* https://www.terraform.io/docs/plugins/provider.html


# Setup
We recommend cleaning up your $GOPATH/src/ before starting.

We also strongly encourage to pull this repository at:
$GOPATH/src/github.com/mustwin/terraform-Oracle-BareMetal-Provider

And the Baremetal GO SDK at:
$GOPATH/src/github.com/MustWin/baremetal-sdk-go

Once that is setup, run govendor sync to copy all the dependencies into
your $GOPATH.

# Testing
```
  $ make test
```

# Build
## For local development
Generates binary on ./terraform-Oracle-BareMetal-Provider
```
  $ make
```

## For release
Generates cross platform binaries on ./bin/
```
  $ make clean cross
```

# Vendoring
This project uses the [Go vendor folder](https://blog.gopheracademy.com/advent-2015/vendor-folder/) for dependencies.
If you need to add or update dependency, please review the [go
vendor docs](https://github.com/kardianos/govendor).

# References
[Oracle Bare Metal Iaas API Docs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apiref.htm)
