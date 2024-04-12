---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ip_inventory_subnet"
sidebar_current: "docs-oci-datasource-core-ip_inventory_subnet"
description: |-
  Provides details about a specific Ip Inventory Subnet in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ip_inventory_subnet
This data source provides details about a specific Ip Inventory Subnet resource in Oracle Cloud Infrastructure Core service.

Gets the IP Inventory data of the specified subnet. Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_core_ip_inventory_subnet" "test_ip_inventory_subnet" {
	#Required
	subnet_id = oci_core_subnet.test_subnet.id
}
```

## Argument Reference

The following arguments are supported:

* `subnet_id` - (Required) Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment of the subnet. 
* `ip_inventory_subnet_resource_summary` - Lists `SubnetResourceSummary` objects. 
	* `address_type` - Address type of the allocated private IP address.
	* `assigned_resource_name` - Name of the created resource.
	* `assigned_resource_type` - Type of the resource.
	* `assigned_time` - Assigned time of the private IP address.
	* `associated_public_ip` - Associated public IP address for the private IP address.
	* `associated_public_ip_pool` - Public IP address Pool the IP address is allocated from.
	* `dns_host_name` - DNS hostname of the IP address.
	* `ip_address` - Lists the allocated private IP address.
	* `ip_address_lifetime` - Lifetime of the allocated private IP address.
	* `ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IP address.
	* `parent_cidr` - The address range the IP address is assigned from.
	* `public_ip_lifetime` - Lifetime of the assigned public IP address.
* `ip_inventory_subnet_count` - Specifies the count for the number of results for the response.
* `last_updated_timestamp` - The Timestamp of the latest update from the database in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `message` - Indicates the status of the data. 

