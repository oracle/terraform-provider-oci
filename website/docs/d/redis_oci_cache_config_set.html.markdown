---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_config_set"
sidebar_current: "docs-oci-datasource-redis-oci_cache_config_set"
description: |-
  Provides details about a specific Oci Cache Config Set in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_config_set
This data source provides details about a specific Oci Cache Config Set resource in Oracle Cloud Infrastructure Redis service.

Retrieves the specified Oracle Cloud Infrastructure Cache Config Set.

## Example Usage

```hcl
data "oci_redis_oci_cache_config_set" "test_oci_cache_config_set" {
	#Required
	oci_cache_config_set_id = oci_redis_oci_cache_config_set.test_oci_cache_config_set.id
}
```

## Argument Reference

The following arguments are supported:

* `oci_cache_config_set_id` - (Required) Unique Oracle Cloud Infrastructure Cache Config Set identifier.


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

