---
layout: "oci"
page_title: "OCI: oci_core_instance_credentials"
sidebar_current: "docs-oci-datasource-core-instance_credentials"
description: |-
  Provides details about a specific InstanceCredential
---

# Data Source: oci_core_instance_credentials
The InstanceCredentials data source provides details about a specific InstanceCredentials

Gets the generated credentials for the instance. Only works for instances that require password to log in (E.g. Windows).
For certain OS'es, users will be forced to change the initial credentials.


## Example Usage

```hcl
data "oci_core_instance_credentials" "test_instance_credential" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The OCID of the instance.


## Attributes Reference

The following attributes are exported:

* `password` - The password for the username.
* `username` - The username.

