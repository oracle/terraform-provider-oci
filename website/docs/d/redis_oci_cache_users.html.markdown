---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_users"
sidebar_current: "docs-oci-datasource-redis-oci_cache_users"
description: |-
  Provides the list of Oci Cache Users in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_users
This data source provides the list of Oci Cache Users in Oracle Cloud Infrastructure Redis service.

Lists the Oracle Cloud Infrastructure Cache users based on the supplied parameters.

## Example Usage

```hcl
data "oci_redis_oci_cache_users" "test_oci_cache_users" {

	#Optional
	compartment_id = var.compartment_id
	name = var.oci_cache_user_name
	state = var.oci_cache_user_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `name` - (Optional) A filter to return the resources that match with the given Oracle Cloud Infrastructure cache user name.
* `state` - (Optional) A filter to return the resources, whose lifecycleState matches with the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `oci_cache_user_collection` - The list of oci_cache_user_collection.

### OciCacheUser Reference

The following attributes are exported:

* `acl_string` - ACL string of Oracle Cloud Infrastructure cache user.
* `authentication_mode` - These are the Authentication details of an Oracle Cloud Infrastructure cache user.
	* `authentication_type` - This is Authentication Type of Oracle Cloud Infrastructure cache user
	* `hashed_passwords` - SHA-256 hashed passwords for Oracle Cloud Infrastructure Cache user,required if authenticationType is set to PASSWORD.
* `compartment_id` - Oracle Cloud Infrastructure Cache user compartment ID.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of Oracle Cloud Infrastructure cache user.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Oracle Cloud Infrastructure Cache user unique ID.
* `name` - Oracle Cloud Infrastructure Cache user name.
* `state` - Oracle Cloud Infrastructure Cache user lifecycle state.
* `status` - Oracle Cloud Infrastructure Cache user status. ON enables and OFF disables the Oracle Cloud Infrastructure cache user to login to the cluster.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time, when the Oracle Cloud Infrastructure cache user was created.
* `time_updated` - The date and time, when the Oracle Cloud Infrastructure cache user was updated.

