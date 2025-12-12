---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_engine_options"
sidebar_current: "docs-oci-datasource-redis-oci_cache_engine_options"
description: |-
  Provides the list of Oci Cache Engine Options in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_engine_options
This data source provides the list of Oci Cache Engine Options in Oracle Cloud Infrastructure Redis service.

Lists Oracle Cloud Infrastructure Cache Engine options


## Example Usage

```hcl
data "oci_redis_oci_cache_engine_options" "test_oci_cache_engine_options" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.


## Attributes Reference

The following attributes are exported:

* `oci_cache_engine_options_collection` - The list of oci_cache_engine_options_collection.

### OciCacheEngineOption Reference

The following attributes are exported:

* `items` - List of Oracle Cloud Infrastructure Cache Engine Options
	* `engine_versions` - List of available engine versions
		* `name` - Oracle Cloud Infrastructure Cache engine version friendly name
		* `version` - Oracle Cloud Infrastructure Cache engine version

