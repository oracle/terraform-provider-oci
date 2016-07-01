# Oracle BareMetal Terraform Provider

This repository contains the Terraform provider for the Oracle Bare Metal Iaas

## Requirements
* Tenancy OCID
* User OCID
* API Key

## Installation
To install the plugin, put the binary somewhere on your filesystem then configure Terraform to be able to find it.
The configuration where plugins are defined is ~/.terraformrc for Unix-like systems and %APPDATA%/terraform.rc for Windows.

providers {
    oraclebaremetal = "/path/to/plugin"
}

## Development
[**Developer Guide**](docs/development.md)
