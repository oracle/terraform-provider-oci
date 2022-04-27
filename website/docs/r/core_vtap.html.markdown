---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vtap"
sidebar_current: "docs-oci-resource-core-vtap"
description: |-
  Provides the Vtap resource in Oracle Cloud Infrastructure Core service
---

# oci_core_vtap
This resource provides the Vtap resource in Oracle Cloud Infrastructure Core service.

Creates a virtual test access point (VTAP) in the specified compartment.

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the VTAP.
For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the VTAP, otherwise a default is provided.
It does not have to be unique, and you can change it.


## Example Usage

```hcl
resource "oci_core_vtap" "test_vtap" {
	#Required
	capture_filter_id = oci_core_capture_filter.test_capture_filter.id
	compartment_id = var.compartment_id
	source_id = oci_core_source.test_source.id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.vtap_display_name
	encapsulation_protocol = var.vtap_encapsulation_protocol
	freeform_tags = {"Department"= "Finance"}
	is_vtap_enabled = var.vtap_is_vtap_enabled
	max_packet_size = var.vtap_max_packet_size
	source_private_endpoint_ip = var.vtap_source_private_endpoint_ip
	source_private_endpoint_subnet_id = oci_core_subnet.test_subnet.id
	source_type = var.vtap_source_type
	target_id = oci_cloud_guard_target.test_target.id
	target_ip = var.vtap_target_ip
	target_type = var.vtap_target_type
	traffic_mode = var.vtap_traffic_mode
	vxlan_network_identifier = var.vtap_vxlan_network_identifier
}
```

## Argument Reference

The following arguments are supported:

* `capture_filter_id` - (Required) (Updatable) The capture filter's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)). 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the `Vtap` resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `encapsulation_protocol` - (Optional) (Updatable) Defines an encapsulation header type for the VTAP's mirrored traffic. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_vtap_enabled` - (Optional) (Updatable) Used to start or stop a `Vtap` resource.
	* `TRUE` directs the VTAP to start mirroring traffic.
	* `FALSE` (Default) directs the VTAP to stop mirroring traffic. 
* `max_packet_size` - (Optional) (Updatable) The maximum size of the packets to be included in the filter.
* `source_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source point where packets are captured. 
* `source_private_endpoint_ip` - (Optional) (Updatable) The IP Address of the source private endpoint. 
* `source_private_endpoint_subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that source private endpoint belongs to. 
* `source_type` - (Optional) (Updatable) The source type for the VTAP. 
* `target_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the destination resource where mirrored packets are sent. 
* `target_ip` - (Optional) (Updatable) The IP address of the destination resource where mirrored packets are sent. 
* `target_type` - (Optional) (Updatable) The target type for the VTAP. 
* `traffic_mode` - (Optional) (Updatable) Used to control the priority of traffic. It is an optional field. If it not passed, the value is DEFAULT
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN containing the `Vtap` resource.
* `vxlan_network_identifier` - (Optional) (Updatable) The virtual extensible LAN (VXLAN) network identifier (or VXLAN segment ID) that uniquely identifies the VXLAN. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `capture_filter_id` - The capture filter's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)). 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the `Vtap` resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `encapsulation_protocol` - Defines an encapsulation header type for the VTAP's mirrored traffic. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The VTAP's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `is_vtap_enabled` - Used to start or stop a `Vtap` resource.
	* `TRUE` directs the VTAP to start mirroring traffic.
	* `FALSE` (Default) directs the VTAP to stop mirroring traffic. 
* `lifecycle_state_details` - The VTAP's current running state.
* `max_packet_size` - The maximum size of the packets to be included in the filter.
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source point where packets are captured. 
* `source_private_endpoint_ip` - The IP Address of the source private endpoint. 
* `source_private_endpoint_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that source private endpoint belongs to. 
* `source_type` - The source type for the VTAP. 
* `state` - The VTAP's administrative lifecycle state.
* `target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the destination resource where mirrored packets are sent. 
* `target_ip` - The IP address of the destination resource where mirrored packets are sent. 
* `target_type` - The target type for the VTAP. 
* `time_created` - The date and time the VTAP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2020-08-25T21:10:29.600Z` 
* `traffic_mode` - Used to control the priority of traffic. It is an optional field. If it not passed, the value is DEFAULT
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN containing the `Vtap` resource.
* `vxlan_network_identifier` - The virtual extensible LAN (VXLAN) network identifier (or VXLAN segment ID) that uniquely identifies the VXLAN. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vtap
	* `update` - (Defaults to 20 minutes), when updating the Vtap
	* `delete` - (Defaults to 20 minutes), when destroying the Vtap


## Import

Vtaps can be imported using the `id`, e.g.

```
$ terraform import oci_core_vtap.test_vtap "id"
```

