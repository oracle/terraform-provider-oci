---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_software_update"
sidebar_current: "docs-oci-datasource-bds-bds_instance_software_update"
description: |-
  Provides details about a specific Bds Instance Software Update in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_software_update
This data source provides details about a specific Bds Instance Software Update resource in Oracle Cloud Infrastructure Big Data Service service.

Get the details of the software update of the given SoftwareUpdateId


## Example Usage

```hcl
data "oci_bds_bds_instance_software_update" "test_bds_instance_software_update" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	software_update_key = var.bds_instance_software_update_software_update_key
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `software_update_key` - (Required) The unique identifier of the software update.


## Attributes Reference

The following attributes are exported:

* `software_update_key` - Unique identifier of a given software update
* `software_update_type` - type of current software update.
	* Big Data Service's micro service. BDS version will be changed after upgrade. 
* `software_update_version` - The version of the software update.
* `state` - The lifecycle state of the software update.
* `time_due` - The due date for the software update. Big Data Service will be updated automatically after this date.
* `time_released` - The time when the software update was released.

