---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_export_status"
sidebar_current: "docs-oci-datasource-jms-fleet_export_status"
description: |-
  Provides details about a specific Fleet Export Status in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_export_status
This data source provides details about a specific Fleet Export Status resource in Oracle Cloud Infrastructure Jms service.

Returns last export status for the specified Fleet.

## Example Usage

```hcl
data "oci_jms_fleet_export_status" "test_fleet_export_status" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet. 
* `latest_run_status` - The status of the latest export run. 
* `time_last_run` - The date and time of the last export run.
* `time_next_run` - The date and time of the next export run.

