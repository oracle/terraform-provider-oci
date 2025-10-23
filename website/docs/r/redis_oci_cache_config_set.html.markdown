---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_config_set"
sidebar_current: "docs-oci-resource-redis-oci_cache_config_set"
description: |-
  Provides the Oci Cache Config Set resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_oci_cache_config_set
This resource provides the Oci Cache Config Set resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/OciCacheConfigSet

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Create a new Oracle Cloud Infrastructure Cache Config Set for the given Oracle Cloud Infrastructure cache engine version.


## Example Usage

```hcl
resource "oci_redis_oci_cache_config_set" "test_oci_cache_config_set" {
	#Required
	compartment_id = var.compartment_id
	configuration_details {
		#Required
		items {
			#Required
			config_key = var.oci_cache_config_set_configuration_details_items_config_key
			config_value = var.oci_cache_config_set_configuration_details_items_config_value
		}
	}
	display_name = var.oci_cache_config_set_display_name
	software_version = var.oci_cache_config_set_software_version

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.oci_cache_config_set_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the Oracle Cloud Infrastructure Cache Config Set.
* `configuration_details` - (Required) List of Oracle Cloud Infrastructure Cache Config Set Values.
	* `items` - (Required) List of ConfigurationInfo objects.
		* `config_key` - (Required) Key is the configuration key.
		* `config_value` - (Required) Value of the configuration as a string. Can represent a string, boolean, or number. Example: "true", "42", or "someString". 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description for the custom Oracle Cloud Infrastructure Cache Config Set.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `software_version` - (Required) The Oracle Cloud Infrastructure Cache engine version that the cluster is running.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the Oracle Cloud Infrastructure Cache Config Set.
* `configuration_details` - List of Oracle Cloud Infrastructure Cache Config Set Values.
	* `items` - List of ConfigurationInfo objects.
		* `config_key` - Key is the configuration key.
		* `config_value` - Value of the configuration as a string. Can represent a string, boolean, or number. Example: "true", "42", or "someString". 
* `default_config_set_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the default Oracle Cloud Infrastructure Cache Config Set which the custom Oracle Cloud Infrastructure Cache Config Set is based upon.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the Oracle Cloud Infrastructure Cache Config Set.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Oracle Cloud Infrastructure Cache Config Set.
* `software_version` - The Oracle Cloud Infrastructure Cache engine version that the cluster is running.
* `state` - The current state of the Oracle Cloud Infrastructure Cache Config Set.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Oracle Cloud Infrastructure Cache Config Set was created. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
* `time_updated` - The date and time the Oracle Cloud Infrastructure Cache Config Set was updated. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oci Cache Config Set
	* `update` - (Defaults to 20 minutes), when updating the Oci Cache Config Set
	* `delete` - (Defaults to 20 minutes), when destroying the Oci Cache Config Set


## Import

OciCacheConfigSets can be imported using the `id`, e.g.

```
$ terraform import oci_redis_oci_cache_config_set.test_oci_cache_config_set "id"
```

