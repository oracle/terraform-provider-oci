---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_network_source"
sidebar_current: "docs-oci-datasource-identity-network_source"
description: |-
  Provides details about a specific Network Source in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_network_source
This data source provides details about a specific Network Source resource in Oracle Cloud Infrastructure Identity service.

Gets the specified network source's information.


## Example Usage

```hcl
data "oci_identity_network_source" "test_network_source" {
	#Required
	network_source_id = "${oci_identity_network_source.test_network_source.id}"
}
```

## Argument Reference

The following arguments are supported:

* `network_source_id` - (Required) The OCID of the network source.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the network source. The tenancy is the root compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the network source. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the network source.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the network source during creation. The name must be unique across the tenancy and cannot be changed. 
* `public_source_list` - A list of allowed public IP addresses and CIDR ranges. 
* `services` - A list of services allowed to make on-behalf-of requests. These requests can have different source IPs than those specified in the network source. Currently, only `all` and `none` are supported. The default is `all`. 
* `state` - The network source object's current state. After creating a network source, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `virtual_source_list` - A list of allowed VCN OCID and IP range pairs. Example:`"vcnId": "ocid1.vcn.oc1.iad.aaaaaaaaexampleuniqueID", "ipRanges": [ "129.213.39.0/24" ]` 

