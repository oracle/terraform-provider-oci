---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resources_associate_monitored_resource"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resources_associate_monitored_resource"
description: |-
  Provides the Monitored Resources Associate Monitored Resource resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resources_associate_monitored_resource
This resource provides the Monitored Resources Associate Monitored Resource resource in Oracle Cloud Infrastructure Stack Monitoring service.

Create an association between two monitored resources.

## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resources_associate_monitored_resource" "test_monitored_resources_associate_monitored_resource" {
	#Required
	association_type = var.monitored_resources_associate_monitored_resource_association_type
	compartment_id = var.compartment_id
	destination_resource_id = oci_stack_monitoring_destination_resource.test_destination_resource.id
	source_resource_id = oci_stack_monitoring_source_resource.test_source_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `association_type` - (Required) Association type to be created between source and destination resources
* `compartment_id` - (Required) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `destination_resource_id` - (Required) Destination Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `source_resource_id` - (Required) Source Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `association_type` - Association Type
* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `destination_resource_details` - Association Resource Details
	* `name` - Monitored Resource Name
	* `type` - Monitored Resource Type
* `destination_resource_id` - Destination Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `source_resource_details` - Association Resource Details
	* `name` - Monitored Resource Name
	* `type` - Monitored Resource Type
* `source_resource_id` - Source Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `tenant_id` - Tenancy Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `time_created` - The time when the association was created. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resources Associate Monitored Resource
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resources Associate Monitored Resource
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resources Associate Monitored Resource


## Import

MonitoredResourcesAssociateMonitoredResources can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resources_associate_monitored_resource.test_monitored_resources_associate_monitored_resource "id"
```

