---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_software_updates"
sidebar_current: "docs-oci-datasource-bds-bds_instance_software_updates"
description: |-
  Provides the list of Bds Instance Software Updates in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_software_updates
This data source provides the list of Bds Instance Software Updates in Oracle Cloud Infrastructure Big Data Service service.

List all the available software updates for current cluster.

## Example Usage

```hcl
data "oci_bds_bds_instance_software_updates" "test_bds_instance_software_updates" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `software_update_collection` - The list of software_update_collection.

### BdsInstanceSoftwareUpdate Reference

The following attributes are exported:

* `software_update_key` - Unique identifier of a given software update
* `software_update_type` - type of current software update.
	* Big Data Service's micro service. BDS version will be changed after upgrade. 
* `software_update_version` - The version of the software update.
* `state` - The lifecycle state of the software update.
* `time_due` - The due date for the software update. Big Data Service will be updated automatically after this date.
* `time_released` - The time when the software update was released.

