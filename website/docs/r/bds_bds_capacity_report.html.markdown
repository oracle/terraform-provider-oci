---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_capacity_report"
sidebar_current: "docs-oci-resource-bds-bds_capacity_report"
description: |-
  Provides the Bds Capacity Report resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_capacity_report
This resource provides the Bds Capacity Report resource in Oracle Cloud Infrastructure Big Data Service service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/bigdata/latest/BdsCapacityReport

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/big_data_service

Create a detailed capacity report for BDS service


## Example Usage

```hcl
resource "oci_bds_bds_capacity_report" "test_bds_capacity_report" {
	#Required
	compartment_id = var.compartment_id
	shape_availabilities {
		#Required
		shape = var.bds_capacity_report_shape_availabilities_shape

		#Optional
		shape_config {

			#Optional
			memory_in_gbs = var.bds_capacity_report_shape_availabilities_shape_config_memory_in_gbs
			nvmes = var.bds_capacity_report_shape_availabilities_shape_config_nvmes
			ocpus = var.bds_capacity_report_shape_availabilities_shape_config_ocpus
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment. This should always be the root compartment.
* `shape_availabilities` - (Required) Information about the shapes in the capacity report.
	* `shape` - (Required) The shape that you want to request a capacity report for.
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes.
		* `nvmes` - (Optional) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available. This parameter is used only for dense shapes.
		* `ocpus` - (Optional) The total number of OCPUs available to the node.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID for the compartment. This should always be the root compartment.
* `shape_availabilities` - Information about the capacity of each requested shape.
	* `domain_level_capacity_reports` - Information about the capacity in each domain.
		* `availability_domain` - The availability domain for the capacity report.
		* `capacity_availability` - Information about the available capacity for a shape.
			* `availability_status` - A flag denoting whether capacity is available.
			* `available_count` - The total number of new cluster nodes that can be created with the specified shape configuration.
		* `domain_type` - Type of domain level for the capacity report.
		* `fault_domain` - The fault domain for the capacity report.
	* `shape` - The shape that the capacity report was requested for.
	* `shape_config` - The shape configuration requested for the node.
		* `memory_in_gbs` - The total amount of memory available to the node, in gigabytes.
		* `nvmes` - The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
		* `ocpus` - The total number of OCPUs available to the node.
* `time_created` - The time the report was created, shown as an RFC 3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Capacity Report
	* `update` - (Defaults to 20 minutes), when updating the Bds Capacity Report
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Capacity Report


## Import

BdsCapacityReports can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_capacity_report.test_bds_capacity_report "id"
```

