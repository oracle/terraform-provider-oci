# Development Guide

# Required Reading
Before starting development it's important to get familiarized with
Terraform and how it's providers work.

Here's a list of recommended material:
* https://www.hashicorp.com/blog/terraform-custom-providers.html (Error
  Handling and partial state are very important concepts)
* https://www.terraform.io/docs/plugins/basics.html
* https://www.terraform.io/docs/plugins/provider.html


# Setup
To get started we recommend to have this repository located at
$GOPATH/src/github.com/mustwin/terraform-Oracle-BareMetal-Provider

# Build
## For local testing
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
[Oracle Bare Metal Iaas API Docs](https://docs.us-az-phoenix-1.oracleiaas.com/)
