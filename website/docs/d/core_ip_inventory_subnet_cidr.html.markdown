---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ip_inventory_subnet_cidr"
sidebar_current: "docs-oci-datasource-core-ip_inventory_subnet_cidr"
description: |-
  Provides details about a specific Ip Inventory Subnet Cidr in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ip_inventory_subnet_cidr
This data source provides details about a specific Ip Inventory Subnet Cidr resource in Oracle Cloud Infrastructure Core service.

Gets the CIDR utilization data of the specified subnet. Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_core_ip_inventory_subnet_cidr" "test_ip_inventory_subnet_cidr" {
	#Required
	subnet_id = oci_core_subnet.test_subnet.id
}
```

## Argument Reference

The following arguments are supported:

* `subnet_id` - (Required) Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment of the subnet. 
* `ip_inventory_cidr_utilization_summary` - Lists 'IpInventoryCidrUtilizationSummary` object. 
	* `address_type` - Address type of the CIDR within a subnet.
	* `cidr` - The CIDR range of a subnet.
	* `utilization` - The CIDR utilisation of a subnet.
* `ip_inventory_subnet_cidr_count` - Specifies the count for the number of results for the response.
* `last_updated_timestamp` - The Timestamp of the latest update from the database in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `message` - Indicates the status of the data.

