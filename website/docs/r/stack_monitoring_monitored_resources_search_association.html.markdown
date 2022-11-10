---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resources_search_association"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resources_search_association"
description: |-
  Provides the Monitored Resources Search Association resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resources_search_association
This resource provides the Monitored Resources Search Association resource in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of monitored resource associations.

## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resources_search_association" "test_monitored_resources_search_association" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	association_type = var.monitored_resources_search_association_association_type
	destination_resource_id = oci_stack_monitoring_destination_resource.test_destination_resource.id
	destination_resource_name = var.monitored_resources_search_association_destination_resource_name
	destination_resource_type = var.monitored_resources_search_association_destination_resource_type
	source_resource_id = oci_stack_monitoring_source_resource.test_source_resource.id
	source_resource_name = var.monitored_resources_search_association_source_resource_name
	source_resource_type = var.monitored_resources_search_association_source_resource_type
}
```

## Argument Reference

The following arguments are supported:

* `association_type` - (Optional) Association type to be created between source and destination resources
* `compartment_id` - (Required) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `destination_resource_id` - (Optional) Destination Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `destination_resource_name` - (Optional) Source Monitored Resource Name
* `destination_resource_type` - (Optional) Source Monitored Resource Type
* `source_resource_id` - (Optional) Source Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `source_resource_name` - (Optional) Source Monitored Resource Name
* `source_resource_type` - (Optional) Source Monitored Resource Type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - List of Monitored Resource Associations.
	* `association_type` - Association type to be created between source and destination resources
	* `destination_resource_details` - Association Resource Details
		* `name` - Monitored Resource Name
		* `type` - Monitored Resource Type
	* `destination_resource_id` - Destination Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	* `source_resource_details` - Association Resource Details
		* `name` - Monitored Resource Name
		* `type` - Monitored Resource Type
	* `source_resource_id` - Source Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	* `time_created` - The time the the association was created. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resources Search Association
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resources Search Association
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resources Search Association


## Import

MonitoredResourcesSearchAssociations can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resources_search_association.test_monitored_resources_search_association "id"
```

