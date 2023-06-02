---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_report"
sidebar_current: "docs-oci-resource-core-compute_capacity_report"
description: |-
  Provides the Compute Capacity Report resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_capacity_report
This resource provides the Compute Capacity Report resource in Oracle Cloud Infrastructure Core service.

Generates a report of the host capacity within an availability domain that is available for you
to create compute instances. Host capacity is the physical infrastructure that resources such as compute
instances run on.

Use the capacity report to determine whether sufficient capacity is available for a shape before
you create an instance or change the shape of an instance.


## Example Usage

```hcl
resource "oci_core_compute_capacity_report" "test_compute_capacity_report" {
	#Required
	availability_domain = var.compute_capacity_report_availability_domain
	compartment_id = var.compartment_id
	shape_availabilities {
		#Required
		instance_shape = var.compute_capacity_report_shape_availabilities_instance_shape

		#Optional
		fault_domain = var.compute_capacity_report_shape_availabilities_fault_domain
		instance_shape_config {

			#Optional
			memory_in_gbs = var.compute_capacity_report_shape_availabilities_instance_shape_config_memory_in_gbs
			nvmes = var.compute_capacity_report_shape_availabilities_instance_shape_config_nvmes
			ocpus = var.compute_capacity_report_shape_availabilities_instance_shape_config_ocpus
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain for the capacity report.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment. This should always be the root compartment. 
* `shape_availabilities` - (Required) Information about the shapes in the capacity report. 
	* `fault_domain` - (Optional) The fault domain for the capacity report.

		If you do not specify a fault domain, the capacity report includes information about all fault domains. 
	* `instance_shape` - (Required) The shape that you want to request a capacity report for. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
	* `instance_shape_config` - (Optional) The shape configuration for a shape in a capacity report. 
		* `memory_in_gbs` - (Optional) The total amount of memory available to the instance, in gigabytes. 
		* `nvmes` - (Optional) The number of NVMe drives to be used for storage. 
		* `ocpus` - (Optional) The total number of OCPUs available to the instance. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain for the capacity report.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment. This should always be the root compartment. 
* `shape_availabilities` - Information about the available capacity for each shape in a capacity report. 
	* `availability_status` - A flag denoting whether capacity is available.
	* `available_count` - The total number of new instances that can be created with the specified shape configuration.
	* `fault_domain` - The fault domain for the capacity report.

		If you do not specify the fault domain, the capacity report includes information about all fault domains. 
	* `instance_shape` - The shape that the capacity report was requested for. 
	* `instance_shape_config` - The shape configuration for a shape in a capacity report. 
		* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
		* `nvmes` - The number of NVMe drives to be used for storage. 
		* `ocpus` - The total number of OCPUs available to the instance. 
* `time_created` - The date and time the capacity report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Capacity Report
	* `update` - (Defaults to 20 minutes), when updating the Compute Capacity Report
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Capacity Report


## Import

ComputeCapacityReports can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_capacity_report.test_compute_capacity_report "id"
```

