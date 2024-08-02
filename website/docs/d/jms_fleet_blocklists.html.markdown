---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_blocklists"
sidebar_current: "docs-oci-datasource-jms-fleet_blocklists"
description: |-
  Provides the list of Fleet Blocklists in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_blocklists
This data source provides the list of Fleet Blocklists in Oracle Cloud Infrastructure Jms service.

Returns a list of blocklist entities contained by a fleet.


## Example Usage

```hcl
data "oci_jms_fleet_blocklists" "test_fleet_blocklists" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	managed_instance_id = var.fleet_blocklist_managed_instance_id
	operation = var.fleet_blocklist_operation
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `operation` - (Optional) The operation type.


## Attributes Reference

The following attributes are exported:

* `blocklist_collection` - The list of blocklist_collection.

### FleetBlocklist Reference

The following attributes are exported:

* `items` - The blocklist
	* `key` - The unique identifier of this blocklist record.
	* `operation` - The operation type
	* `reason` - The reason why the operation is blocklisted.
	* `target` - A resource to blocklist for certain operation.
		* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet. 
		* `installation_key` - The unique identifier for the installation of Java Runtime at a specific path on a specific operating system.
		* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 

