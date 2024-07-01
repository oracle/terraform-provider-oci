---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_private_endpoint"
sidebar_current: "docs-oci-resource-opsi-operations_insights_private_endpoint"
description: |-
  Provides the Operations Insights Private Endpoint resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_operations_insights_private_endpoint
This resource provides the Operations Insights Private Endpoint resource in Oracle Cloud Infrastructure Opsi service.

Create a private endpoint resource for the tenant in Ops Insights.
This resource will be created in customer compartment.


## Example Usage

```hcl
resource "oci_opsi_operations_insights_private_endpoint" "test_operations_insights_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.operations_insights_private_endpoint_display_name
	is_used_for_rac_dbs = var.operations_insights_private_endpoint_is_used_for_rac_dbs
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.operations_insights_private_endpoint_description
	freeform_tags = {"bar-key"= "value"}
	nsg_ids = var.operations_insights_private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Private service accessed database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The description of the private endpoint.
* `display_name` - (Required) (Updatable) The display name for the private endpoint. It is changeable.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_used_for_rac_dbs` - (Required) This flag was previously used to create a private endpoint with scan proxy. Setting this to true will now create a private endpoint with a DNS proxy causing `isProxyEnabled` flag to be true; this is used exclusively for full feature support for dedicated Autonomous Databases. 
* `nsg_ids` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups that the private endpoint belongs to. 
* `subnet_id` - (Required) The Subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Private service accessed database.
* `vcn_id` - (Required) The VCN [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Private service accessed database.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID of the Private service accessed database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the private endpoint.
* `display_name` - The display name of the private endpoint.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Private service accessed database.
* `is_used_for_rac_dbs` - The flag is to identify if private endpoint is used for rac database or not. This flag is deprecated and no longer is used.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_endpoint_status_details` - A message describing the status of the private endpoint connection of this resource. For example, it can be used to provide actionable information about the validity of the private endpoint connection.
* `private_ip` - The private IP addresses assigned to the private endpoint. All IP addresses will be concatenated if it is RAC DBs. 
* `state` - The current state of the private endpoint.
* `subnet_id` - The OCID of the subnet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Operations Insights Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Operations Insights Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Operations Insights Private Endpoint


## Import

OperationsInsightsPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint "id"
```

