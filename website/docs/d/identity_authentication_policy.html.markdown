---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_authentication_policy"
sidebar_current: "docs-oci-datasource-identity-authentication_policy"
description: |-
  Provides details about a specific Authentication Policy in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_authentication_policy
This data source provides details about a specific Authentication Policy resource in Oracle Cloud Infrastructure Identity service.

Gets the authentication policy for the given tenancy. You must specify your tenant’s OCID as the value for
the compartment ID (remember that the tenancy is simply the root compartment).


## Example Usage

```hcl
data "oci_identity_authentication_policy" "test_authentication_policy" {
	#Required
	compartment_id = var.tenancy_ocid
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID.
* `network_policy` - Network policy, Consists of a list of Network Source ids. 
	* `network_source_ids` - Network Source ids 
* `password_policy` - Password policy, currently set for the given compartment. 
	* `is_lowercase_characters_required` - At least one lower case character required.
	* `is_numeric_characters_required` - At least one numeric character required.
	* `is_special_characters_required` - At least one special character required.
	* `is_uppercase_characters_required` - At least one uppercase character required.
	* `is_username_containment_allowed` - User name is allowed to be part of the password.
	* `minimum_password_length` - Minimum password length required.

