---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_credentials"
sidebar_current: "docs-oci-datasource-core-instance_credential"
description: |-
  Provides details about a specific Instance Credential in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_credentials
This data source provides details about a specific Instance Credential resource in Oracle Cloud Infrastructure Core service.

Gets the generated credentials for the instance. Only works for instances that require a password to log in, such as Windows.
For certain operating systems, users will be forced to change the initial credentials.


## Example Usage

```hcl
data "oci_core_instance_credentials" "test_instance_credential" {
	#Required
	instance_id = oci_core_instance.test_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.


## Attributes Reference

The following attributes are exported:

* `password` - The password for the username.
* `username` - The username.

