---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_list_user_grants"
sidebar_current: "docs-oci-datasource-data_safe-list_user_grants"
description: |-
  Provides the list of List User Grants in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_list_user_grants
This data source provides the list of List User Grants in Oracle Cloud Infrastructure Data Safe service.

Gets a list of grants for a particular user in the specified user assessment. A user grant contains details such as the
privilege name, type, category, and depth level. The depth level indicates how deep in the hierarchy of roles granted to
roles a privilege grant is. The userKey in this operation is a system-generated identifier. Perform the operation ListUsers
to get the userKey for a particular user.


## Example Usage

```hcl
data "oci_data_safe_list_user_grants" "test_list_user_grants" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
	user_key = var.list_user_grant_user_key

	#Optional
	depth_level = var.list_user_grant_depth_level
	depth_level_greater_than_or_equal_to = var.list_user_grant_depth_level_greater_than_or_equal_to
	depth_level_less_than = var.list_user_grant_depth_level_less_than
	grant_key = var.list_user_grant_grant_key
	grant_name = var.list_user_grant_grant_name
	privilege_category = var.list_user_grant_privilege_category
	privilege_type = var.list_user_grant_privilege_type
}
```

## Argument Reference

The following arguments are supported:

* `depth_level` - (Optional) A filter to return only items that match the specified user grant depth level.
* `depth_level_greater_than_or_equal_to` - (Optional) A filter to return only items that are at a level greater than or equal to the specified user grant depth level.
* `depth_level_less_than` - (Optional) A filter to return only items that are at a level less than the specified user grant depth level.
* `grant_key` - (Optional) A filter to return only items that match the specified user grant key.
* `grant_name` - (Optional) A filter to return only items that match the specified user grant name.
* `privilege_category` - (Optional) A filter to return only items that match the specified user privilege category.
* `privilege_type` - (Optional) A filter to return only items that match the specified privilege grant type.
* `user_assessment_id` - (Required) The OCID of the user assessment.
* `user_key` - (Required) The unique user key. This is a system-generated identifier. ListUsers gets the user key for a user.


## Attributes Reference

The following attributes are exported:

* `grants` - The list of grants.

### ListUserGrant Reference

The following attributes are exported:

* `depth_level` - The grant depth level of the indirect grant. An indirectly granted role/privilege is granted to the user through another role. The depth level indicates how deep a privilege is within the grant hierarchy. 
* `grant_name` - The name of a user grant.
* `key` - The unique key of a user grant.
* `privilege_category` - The privilege category.
* `privilege_type` - The type of a user grant.

