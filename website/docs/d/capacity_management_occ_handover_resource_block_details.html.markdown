---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_handover_resource_block_details"
sidebar_current: "docs-oci-datasource-capacity_management-occ_handover_resource_block_details"
description: |-
  Provides the list of Occ Handover Resource Block Details in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_handover_resource_block_details
This data source provides the list of Occ Handover Resource Block Details in Oracle Cloud Infrastructure Capacity Management service.

List details about a given occHandoverResourceBlock.


## Example Usage

```hcl
data "oci_capacity_management_occ_handover_resource_block_details" "test_occ_handover_resource_block_details" {
	#Required
	occ_handover_resource_block_id = oci_capacity_management_occ_handover_resource_block.test_occ_handover_resource_block.id

	#Optional
	host_id = oci_capacity_management_host.test_host.id
}
```

## Argument Reference

The following arguments are supported:

* `host_id` - (Optional) This fiter is applicable only for COMPUTE namespace. It helps in fetching of all resource block details for which the hostId is equal to the one provided in this query param. 
* `occ_handover_resource_block_id` - (Required) The OCID of the OccHandoverResource which is a required query parameter for listing OccHandoverResourceDetails. 


## Attributes Reference

The following attributes are exported:

* `occ_handover_resource_block_detail_collection` - The list of occ_handover_resource_block_detail_collection.

### OccHandoverResourceBlockDetail Reference

The following attributes are exported:

* `items` - An array of details about an occ handover resource block. 
	* `details` - A map that contains additional details for a given handover resource. For example for compute namespace this includes host ocid, host serial etc. 
	* `occ_resource_handover_block_id` - The OCID of the occResourceHandoverBlock. 

