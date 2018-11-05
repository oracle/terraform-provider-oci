---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pools"
sidebar_current: "docs-oci-datasource-core-instance_pools"
description: |-
  Provides the list of Instance Pools in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_pools
This data source provides the list of Instance Pools in Oracle Cloud Infrastructure Core service.

Lists the instance pools in the specified compartment.

## Example Usage

```hcl
data "oci_core_instance_pools" "test_instance_pools" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.instance_pool_display_name}"
	state = "${var.instance_pool_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `instance_pools` - The list of instance_pools.

### InstancePool Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the instance pool
* `display_name` - The user-friendly name.  Does not have to be unique.
* `id` - The OCID of the instance pool
* `instance_configuration_id` - The OCID of the instance configuration associated to the intance pool.
* `size` - The number of instances that should be in the instance pool.
* `state` - The current state of the instance pool.
* `time_created` - The date and time the instance pool was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

