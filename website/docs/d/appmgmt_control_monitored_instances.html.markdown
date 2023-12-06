---
subcategory: "Appmgmt Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_appmgmt_control_monitored_instances"
sidebar_current: "docs-oci-datasource-appmgmt_control-monitored_instances"
description: |-
  Provides the list of Monitored Instances in Oracle Cloud Infrastructure Appmgmt Control service
---

# Data Source: oci_appmgmt_control_monitored_instances
This data source provides the list of Monitored Instances in Oracle Cloud Infrastructure Appmgmt Control service.

Returns a list of monitored instances.


## Example Usage

```hcl
data "oci_appmgmt_control_monitored_instances" "test_monitored_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.monitored_instance_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.


## Attributes Reference

The following attributes are exported:

* `monitored_instance_collection` - The list of monitored_instance_collection.

### MonitoredInstance Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `display_name` - A user-friendly name of the monitored instance. It is binded to [Compute Instance](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm). DisplayName is fetched from [Core Service API](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance/). 
* `instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored instance.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `management_agent_id` - Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Used to invoke manage operations on Management Agent Cloud Service. 
* `monitoring_state` - Monitoring status. Can be either enabled or disabled.
* `state` - The current state of the monitored instance.
* `time_created` - The time the MonitoredInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the MonitoredInstance was updated. An RFC3339 formatted datetime string

