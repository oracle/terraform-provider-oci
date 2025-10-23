---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_user"
sidebar_current: "docs-oci-resource-redis-oci_cache_user"
description: |-
  Provides the Oci Cache User resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_oci_cache_user
This resource provides the Oci Cache User resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/OciCacheUser

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Creates a new Oracle Cloud Infrastructure Cache user. Oracle Cloud Infrastructure Cache user is required to authenticate to Oracle Cloud Infrastructure Cache cluster.

## Example Usage

```hcl
resource "oci_redis_oci_cache_user" "test_oci_cache_user" {
	#Required
	acl_string = var.oci_cache_user_acl_string
	authentication_mode {
		#Required
		authentication_type = var.oci_cache_user_authentication_mode_authentication_type

		#Optional
		hashed_passwords = var.oci_cache_user_authentication_mode_hashed_passwords
	}
	compartment_id = var.compartment_id
	description = var.oci_cache_user_description
	name = var.oci_cache_user_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	status = var.oci_cache_user_status
}
```

## Argument Reference

The following arguments are supported:

* `acl_string` - (Required) (Updatable) ACL string of Oracle Cloud Infrastructure cache user.
* `authentication_mode` - (Required) (Updatable) These are the Authentication details of an Oracle Cloud Infrastructure cache user.
	* `authentication_type` - (Required) (Updatable) This is Authentication Type of Oracle Cloud Infrastructure cache user
	* `hashed_passwords` - (Required when authentication_type=PASSWORD) (Updatable) SHA-256 hashed passwords for Oracle Cloud Infrastructure Cache user,required if authenticationType is set to PASSWORD.
* `compartment_id` - (Required) (Updatable) Oracle Cloud Infrastructure cache user compartment ID.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) (Updatable) Description of Oracle Cloud Infrastructure cache user.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) Oracle Cloud Infrastructure cache user name is required to connect to an Oracle Cloud Infrastructure cache cluster.
* `status` - (Optional) (Updatable) Oracle Cloud Infrastructure cache user status. ON enables and OFF disables the Oracle Cloud Infrastructure cache user to login to the associated clusters. Default value is ON.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oci Cache User
	* `update` - (Defaults to 20 minutes), when updating the Oci Cache User
	* `delete` - (Defaults to 20 minutes), when destroying the Oci Cache User


## Import

OciCacheUsers can be imported using the `id`, e.g.

```
$ terraform import oci_redis_oci_cache_user.test_oci_cache_user "id"
```

