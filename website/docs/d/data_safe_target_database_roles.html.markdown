---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database_roles"
sidebar_current: "docs-oci-datasource-data_safe-target_database_roles"
description: |-
  Provides the list of Target Database Roles in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_database_roles
This data source provides the list of Target Database Roles in Oracle Cloud Infrastructure Data Safe service.

Returns a list of role metadata objects.


## Example Usage

```hcl
data "oci_data_safe_target_database_roles" "test_target_database_roles" {
	#Required
	target_database_id = oci_data_safe_target_database.test_target_database.id

	#Optional
	authentication_type = var.target_database_role_authentication_type
	is_oracle_maintained = var.target_database_role_is_oracle_maintained
	role_name = var.target_database_role_role_name
	role_name_contains = var.target_database_role_role_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `authentication_type` - (Optional) A filter to return roles based on authentication type.
* `is_oracle_maintained` - (Optional) A filter to return roles based on whether they are maintained by oracle or not.
* `role_name` - (Optional) A filter to return only a specific role based on role name.
* `role_name_contains` - (Optional) A filter to return only items if role name contains a specific string.
* `target_database_id` - (Required) The OCID of the Data Safe target database.


## Attributes Reference

The following attributes are exported:

* `roles` - The list of roles.

### TargetDatabaseRole Reference

The following attributes are exported:

* `authentication_type` - Type of authentication.
* `is_common` - Is the role common.
* `is_implicit` - Is the role implicit.
* `is_inherited` - Is the role inherited.
* `is_oracle_maintained` - Is the role oracle maintained.
* `is_password_required` - Is password required.
* `role_name` - Name of the role.

