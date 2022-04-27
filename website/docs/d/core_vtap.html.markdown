---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vtap"
sidebar_current: "docs-oci-datasource-core-vtap"
description: |-
  Provides details about a specific Vtap in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_vtap
This data source provides details about a specific Vtap resource in Oracle Cloud Infrastructure Core service.

Gets the specified `Vtap` resource.

## Example Usage

```hcl
data "oci_core_vtap" "test_vtap" {
	#Required
	vtap_id = oci_core_vtap.test_vtap.id
}
```

## Argument Reference

The following arguments are supported:

* `vtap_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VTAP.


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

