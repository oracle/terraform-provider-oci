---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_maintenance_reboot"
sidebar_current: "docs-oci-datasource-core-instance_maintenance_reboot"
description: |-
  Provides details about a specific Instance Maintenance Reboot in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_maintenance_reboot
This data source provides details about a specific Instance Maintenance Reboot resource in Oracle Cloud Infrastructure Core service.

Gets the maximum possible date that a maintenance reboot can be extended.

## Example Usage

```hcl
data "oci_core_instance_maintenance_reboot" "test_instance_maintenance_reboot" {
	#Required
	instance_id = oci_core_instance.test_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.


## Attributes Reference

The following attributes are exported:

* `time_maintenance_reboot_due_max` - The maximum extension date and time for the maintenance reboot, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). The range for the maintenance extension is between 1 and 14 days from the initial scheduled maintenance date. Example: `2018-05-25T21:10:29.600Z` 

