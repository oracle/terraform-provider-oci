---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_config_sets"
sidebar_current: "docs-oci-datasource-redis-oci_cache_config_sets"
description: |-
  Provides the list of Oci Cache Config Sets in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_config_sets
This data source provides the list of Oci Cache Config Sets in Oracle Cloud Infrastructure Redis service.

Lists the Oracle Cloud Infrastructure Cache Config Sets in the specified compartment.


## Example Usage

```hcl
data "oci_redis_oci_cache_config_sets" "test_oci_cache_config_sets" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.oci_cache_config_set_display_name
	id = var.oci_cache_config_set_id
	software_version = var.oci_cache_config_set_software_version
	state = var.oci_cache_config_set_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique Oracle Cloud Infrastructure Cache Config Set identifier.
* `software_version` - (Optional) A filter to return the Oracle Cloud Infrastructure Cache Config Set resources, whose software version matches with the given software version.
* `state` - (Optional) A filter to return the Oracle Cloud Infrastructure Cache Config Set resources, whose lifecycle state matches with the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `oci_cache_config_set_collection` - The list of oci_cache_config_set_collection.

### OciCacheConfigSet Reference

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

