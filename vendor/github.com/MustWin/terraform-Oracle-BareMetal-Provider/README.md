```
__________                         _____          __         .__
\______   \_____ _______   ____   /     \   _____/  |______  |  |
 |    |  _/\__  \\_  __ \_/ __ \ /  \ /  \_/ __ \   __\__  \ |  |
 |    |   \ / __ \|  | \/\  ___//    Y    \  ___/|  |  / __ \|  |__
 |______  /(____  /__|    \___  >____|__  /\___  >__| (____  /____/
        \/      \/            \/        \/     \/          \/
```
# Oracle BareMetal Terraform Provider ![](https://circleci.com/gh/MustWin/terraform-Oracle-BareMetal-Provider.svg?style=shield&circle-token=fa06ce2af6b594812e3a756f5451a9e101d7b9f5)

This repository contains the Terraform provider for the Oracle Bare Metal Iaas

## Requirements

* Terraform v0.6.16+
* Oracle Bare Metal credentials (see below)

Credentials required:
* Tenancy OCID
* User OCID
* API Key

Also you will be needing Terraform 0.6.16+
[Terraform Install](https://www.terraform.io/intro/getting-started/install.html)

## Installation
To install the plugin, put the binary somewhere on your filesystem then configure Terraform to be able to find it.
The configuration where plugins are defined is ~/.terraformrc for Unix-like systems and %APPDATA%/terraform.rc for Windows.

```json
providers {
    oraclebaremetal = "/path/to/plugin"
}
```

## Development
[**Developer Guide**](docs/development.md)
