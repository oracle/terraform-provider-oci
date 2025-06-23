---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_default_config_set"
sidebar_current: "docs-oci-datasource-redis-oci_cache_default_config_set"
description: |-
  Provides details about a specific Oci Cache Default Config Set in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_default_config_set
This data source provides details about a specific Oci Cache Default Config Set resource in Oracle Cloud Infrastructure Redis service.

Retrieves the specified Oracle Cloud Infrastructure Cache Default Config Set.

## Example Usage

```hcl
data "oci_redis_oci_cache_default_config_set" "test_oci_cache_default_config_set" {
	#Required
	compartment_id = var.compartment_id
	oci_cache_default_config_set_id = oci_redis_oci_cache_default_config_set.test_oci_cache_default_config_set.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.
* `oci_cache_default_config_set_id` - (Required) Unique Oracle Cloud Infrastructure Cache Default Config Set identifier.


## Attributes Reference

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

