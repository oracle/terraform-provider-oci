# Overview
This is a Terraform configuration that creates the Logging service on Oracle Cloud Infrastructure.

The Terraform code is used to create a Resource Manager stack, that creates the required resources and configures the application on the created resources.

Build preview provider and put it in
`~/.terraform.d/plugins`
or
Put the terraform config in the tf file
```terraform
terraform {
  # specify provider
  required_providers {
    # put preview provider in ~/.terraform.d/plugins/terraform.local/local/oci/4.80.1/darwin_amd64
    oci = {
      source = "terraform.local/local/oci"
      version = "4.80.1"
    }
  }
  # specify terraform version
  # required_version = ">= 0.12.31"
  required_version = ">= 1.2.3"
}

```
Then go to `examples/logging` and run the terraform commands to test the examples.

## Magic Button 
[![Deploy to Oracle Cloud](https://oci-resourcemanager-plugin.plugins.oci.oraclecloud.com/latest/deploy-to-oracle-cloud.svg)](https://cloud.oracle.com/resourcemanager/stacks/create?zipUrl=https://github.com/oracle/terraform-provider-oci/raw/master/examples/zips/logging.zip)