---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ip_inventory_vcn_overlaps"
sidebar_current: "docs-oci-datasource-core-ip_inventory_vcn_overlaps"
description: |-
  Provides the list of Ip Inventory Vcn Overlaps in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ip_inventory_vcn_overlaps
This data source provides the list of Ip Inventory Vcn Overlaps in Oracle Cloud Infrastructure Core service.

Gets the CIDR overlap information of the specified VCN in selected compartments. Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_core_ip_inventory_vcn_overlaps" "test_ip_inventory_vcn_overlaps" {
	#Required
	compartment_list = var.ip_inventory_vcn_overlap_compartment_list
	region_list = var.ip_inventory_vcn_overlap_region_list
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_list` - (Required) The list of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartments.
* `region_list` - (Required) Lists the selected regions.
* `vcn_id` - (Required) Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `ip_inventory_vcn_overlap_collection` - The list of ip_inventory_vcn_overlap_collection.

### IpInventoryVcnOverlap Reference

The following attributes are exported:

* `ip_inventory_vcn_overlap_summary` - Lists `IpInventoryVcnOverlapSummary` object. 
	* `cidr` - CIDR prefix of the VCN.
	* `overlapping_cidr` - The overlapping CIDR prefix.
	* `overlapping_vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN .
	* `overlapping_vcn_name` - Name of the overlapping VCN.
* `last_updated_timestamp` - The timestamp of the latest update from the database in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `message` - Indicates the status of the data.
* `overlap_count` - The overlap count for the given VCN and compartments.

