---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_fast_launch_job_configs"
sidebar_current: "docs-oci-datasource-datascience-fast_launch_job_configs"
description: |-
  Provides the list of Fast Launch Job Configs in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_fast_launch_job_configs
This data source provides the list of Fast Launch Job Configs in Oracle Cloud Infrastructure Data Science service.

List fast launch capable job configs in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_fast_launch_job_configs" "test_fast_launch_job_configs" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `fast_launch_job_configs` - The list of fast_launch_job_configs.

### FastLaunchJobConfig Reference

The following attributes are exported:

* `core_count` - The number of cores associated with this fast launch job shape. 
* `managed_egress_support` - The managed egress support 
* `memory_in_gbs` - The number of cores associated with this fast launch job shape. 
* `name` - The name of the fast launch job config 
* `shape_name` - The name of the fast launch job shape. 
* `shape_series` - The family that the compute shape belongs to. 

