---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_summarize_resource_inventory"
sidebar_current: "docs-oci-datasource-jms-summarize_resource_inventory"
description: |-
  Provides details about a specific Summarize Resource Inventory in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_summarize_resource_inventory
This data source provides details about a specific Summarize Resource Inventory resource in Oracle Cloud Infrastructure Jms service.

Retrieve the inventory of JMS resources in the specified compartment: a list of the number of _active_ fleets, managed instances, Java Runtimes, Java installations, and applications.


## Example Usage

```hcl
data "oci_jms_summarize_resource_inventory" "test_summarize_resource_inventory" {

	#Optional
	compartment_id = var.compartment_id
	time_end = var.summarize_resource_inventory_time_end
	time_start = var.summarize_resource_inventory_time_start
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `active_fleet_count` - The number of _active_ fleets.
* `application_count` - The number of applications.
* `installation_count` - The number of Java installations.
* `jre_count` - The number of Java Runtimes.
* `managed_instance_count` - The number of managed instances.

