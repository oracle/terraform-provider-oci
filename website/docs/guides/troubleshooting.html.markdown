---
layout: "oci"
page_title: "Troubleshooting"
sidebar_current: "docs-oci-guide-troubleshooting"
description: |-
  The Oracle Cloud Infrastructure provider. Troubleshooting
---

## Troubleshooting

When troubleshooting or getting support for the OCI Terraform Provider, it is often useful to first check the status of the OCI services and to collect verbose logging.


### Checking OCI service status and outages

To check on the latest status and whether there are any outages in OCI, see [OCI Status](https://ocistatus.oraclecloud.com/)


### Verbose logging for OCI Terraform Provider

To get verbose console output when the provider is running, precede your Terraform command with the `TF_LOG` and `OCI_GO_SDK_DEBUG` flags:

```sh
TF_LOG=DEBUG OCI_GO_SDK_DEBUG=v terraform plan
```

The [tf_log](https://www.terraform.io/docs/internals/debugging.html) level and `OCI_GO_SDK_DEBUG` flags can also be set as environment variables.


## Common Issues

### Not Authenticated Error when configuring Terraform

_If the Terraform CLI gives an error message like:_

```
* oci_core_vcn.resource1: Service error:NotAuthenticated. The required information to complete authentication was not provided or was incorrect.. http status code: 401
```

* Verify you have properly set `user_ocid`, `tenancy_ocid`, `fingerprint` and `private_key_path` 
* Verify your `private_key_path` is pointing to your private key and not the corresponding public key
* Verify you have added the corresponding public key to the user account you have specified with `user_ocid`
* Verify the public/private key pairs you are using are of the correct format
  * see: [Required Keys](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm) for details on the correct format and how to generate keys
* Verify the user account is part of a group with the appropriate permissions to perform the actions in the plan you are executing
* Verify your Tenancy has been subscribed to the Region you are targeting in your plan


### Error message after upgrading OCI Terraform Provider

_If the Terraform CLI gives an error message like:_

```
Error asking for user input: 1 error(s) occurred:
 
* provider.oci: dial unix /var/folders/6r/8fk5dmbj4_z3sl0mc_y_fhjw0000gn/T/plugin811254328|netrpc: connect: no such file or directory
```

You are likely using a version of the OCI Terraform Provider that is not compatible with the Terraform binary you have 
installed. For OCI Provider versions v3.x.x and above, a minimum Terraform version of v.0.10.1 is required. 

### Error message when field cannot be set

_If the Terraform CLI gives an error message like:_

``` 
* Error: "field_name": this field cannot be set
```

You are likely using an older version of the OCI Terraform Provider and the field you are trying to set was released in later version. The OCI Terraform Provider documentation reflects the [latest version](https://github.com/terraform-providers/terraform-provider-oci/releases).


### Dial tcp i/o timeout when connecting via proxy

_If the Terraform CLI gives an error message like:_

```
* provider.oci: ... dial tcp 134.70.16.0:443: i/o timeout
```

Then you may not have properly configured your proxy settings. The OCI terraform provider does support `http_proxy`, `https_proxy` and `no_proxy` variables where the inclusion or exclusion lists can be defined as follows:

```
export http_proxy=http://www.your-proxy.com:80/
export https_proxy=http://www.your-proxy.com:80/
export no_proxy=localhost,127.0.0.1
```


### Errors about service limits

While using Terraform, you may encounter errors indicating that you have reached or exceeded the service limits for a resource.

To understand more about your OCI service limits and how to request a limit increase, see [Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm)
