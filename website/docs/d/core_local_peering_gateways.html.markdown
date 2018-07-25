---
layout: "oci"
page_title: "OCI: oci_core_local_peering_gateways"
sidebar_current: "docs-oci-datasource-core-local_peering_gateways"
description: |-
  Provides a list of LocalPeeringGateways
---

# Data Source: oci_core_local_peering_gateways
The LocalPeeringGateways data source allows access to the list of OCI local_peering_gateways

Lists the local peering gateways (LPGs) for the specified VCN and compartment
(the LPG's compartment).


## Example Usage

```hcl
data "oci_core_local_peering_gateways" "test_local_peering_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.


## Attributes Reference

The following attributes are exported:

* `local_peering_gateways` - The list of local_peering_gateways.

### LocalPeeringGateway Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the Local Peering Gateway (LPG).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The LPG's Oracle ID (OCID).
* `is_cross_tenancy_peering` - Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false` 
* `peer_advertised_cidr` - The range of IP addresses available on the VCN at the other end of the peering from this LPG. The value is `null` if the LPG is not peered. You can use this as the destination CIDR for a route rule to route a subnet's traffic to this LPG.  Example: `192.168.0.0/16` 
* `peering_status` - Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the LPG at the other end of the peering has been deleted. 
* `peering_status_details` - Additional information regarding the peering status, if applicable.
* `state` - The LPG's current lifecycle state.
* `time_created` - The date and time the LPG was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the LPG belongs to.

