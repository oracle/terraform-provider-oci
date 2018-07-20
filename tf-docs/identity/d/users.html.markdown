---
layout: "oci"
page_title: "OCI: oci_identity_users"
sidebar_current: "docs-oci-datasource-users"
description: |-
Provides a list of Users
---
# Data Source: oci_identity_users
The Users data source allows access to the list of OCI users

Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_users" "test_users" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `users` - The list of users.