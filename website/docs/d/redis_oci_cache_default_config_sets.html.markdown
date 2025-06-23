---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_default_config_sets"
sidebar_current: "docs-oci-datasource-redis-oci_cache_default_config_sets"
description: |-
  Provides the list of Oci Cache Default Config Sets in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_default_config_sets
This data source provides the list of Oci Cache Default Config Sets in Oracle Cloud Infrastructure Redis service.

Lists the Oracle Cloud Infrastructure Cache Default Config Sets in the specified compartment.

## Example Usage

```hcl
data "oci_redis_oci_cache_default_config_sets" "test_oci_cache_default_config_sets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oci_cache_default_config_set_display_name
	id = var.oci_cache_default_config_set_id
	software_version = var.oci_cache_default_config_set_software_version
	state = var.oci_cache_default_config_set_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique Oracle Cloud Infrastructure Cache Default Config Set identifier.
* `software_version` - (Optional) A filter to return the Oracle Cloud Infrastructure Cache Config Set resources, whose software version matches with the given software version.
* `state` - (Optional) A filter to return the Oracle Cloud Infrastructure Cache Default Config Set resources, whose lifecycle state matches with the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `oci_cache_default_config_set_collection` - The list of oci_cache_default_config_set_collection.

### OciCacheDefaultConfigSet Reference

The following attributes are exported:

* `default_configuration_details` - List of Oracle Cloud Infrastructure Cache Default Config Set Values.
	* `items` - List of DefaultConfigurationInfo objects.
		* `allowed_values` - Allowed values for the configuration setting.
		* `config_key` - The key of the configuration setting.
		* `data_type` - The data type of the configuration setting.
		* `default_config_value` - The default value for the configuration setting.
		* `description` - Description of the configuration setting.
		* `is_modifiable` - Indicates if the configuration is modifiable.
* `description` - Description of the Oracle Cloud Infrastructure Cache Default Config Set.
* `display_name` - A user-friendly name of the Oracle Cloud Infrastructure Cache Default Config Set.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Oracle Cloud Infrastructure Cache Default Config Set.
* `software_version` - The engine version of the Oracle Cloud Infrastructure Cache Default Config Set.
* `state` - The current state of the Oracle Cloud Infrastructure Cache Default Config Set.
* `time_created` - The date and time the Oracle Cloud Infrastructure Cache Default Config Set was created. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.

