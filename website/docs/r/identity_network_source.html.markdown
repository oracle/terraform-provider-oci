---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_network_source"
sidebar_current: "docs-oci-resource-identity-network_source"
description: |-
  Provides the Network Source resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_network_source
This resource provides the Network Source resource in Oracle Cloud Infrastructure Identity service.

Creates a new network source in your tenancy.

You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
reside within compartments inside the tenancy. For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the network source, which must be unique across all network sources in your
tenancy, and cannot be changed.
You can use this name or the OCID when writing policies that apply to the network source. For more information
about policies, see [How Policies Work](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policies.htm).

You must also specify a *description* for the network source (although it can be an empty string). It does not
have to be unique, and you can change it anytime with [UpdateNetworkSource](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/NetworkSource/UpdateNetworkSource).
After your network resource is created, you can use it in policy to restrict access to only requests made from an allowed
IP address specified in your network source. For more information, see [Managing Network Sources](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingnetworksources.htm).


## Example Usage

```hcl
resource "oci_identity_network_source" "test_network_source" {
	#Required
	compartment_id = var.tenancy_ocid
	description = var.network_source_description
	name = var.network_source_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	public_source_list = var.network_source_public_source_list
	services = var.network_source_services
	virtual_source_list = var.network_source_virtual_source_list
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy (root compartment) containing the network source object.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the network source during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the network source during creation. The name must be unique across all groups in the tenancy and cannot be changed. 
* `public_source_list` - (Optional) (Updatable) A list of allowed public IP addresses and CIDR ranges. 
* `services` - (Optional) (Updatable) A list of services allowed to make on-behalf-of requests. These requests can have different source IP addresses than those listed in the network source. Currently, only `all` and `none` are supported. The default is `all`. 
* `virtual_source_list` - (Optional) (Updatable) A list of allowed VCN OCID and IP range pairs. Example:`"vcnId": "ocid1.vcn.oc1.iad.aaaaaaaaexampleuniqueID", "ipRanges": [ "129.213.39.0/24" ]` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Source
	* `update` - (Defaults to 20 minutes), when updating the Network Source
	* `delete` - (Defaults to 20 minutes), when destroying the Network Source


## Import

NetworkSources can be imported using the `id`, e.g.

```
$ terraform import oci_identity_network_source.test_network_source "id"
```

