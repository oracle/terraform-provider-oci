---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_virtual_circuit_associated_tunnels"
sidebar_current: "docs-oci-datasource-core-virtual_circuit_associated_tunnels"
description: |-
  Provides the list of Virtual Circuit Associated Tunnels in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_virtual_circuit_associated_tunnels
This data source provides the list of Virtual Circuit Associated Tunnels in Oracle Cloud Infrastructure Core service.

Gets the specified virtual circuit's associatedTunnelsInfo.

## Example Usage

```hcl
data "oci_core_virtual_circuit_associated_tunnels" "test_virtual_circuit_associated_tunnels" {
	#Required
	virtual_circuit_id = oci_core_virtual_circuit.test_virtual_circuit.id
}
```

## Argument Reference

The following arguments are supported:

* `virtual_circuit_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual circuit.


## Attributes Reference

The following attributes are exported:

* `virtual_circuit_associated_tunnel_details` - The list of virtual_circuit_associated_tunnel_details.

### VirtualCircuitAssociatedTunnel Reference

The following attributes are exported:

* `ipsec_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of IPSec connection associated with the virtual circuit.
* `tunnel_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec tunnel associated with the virtual circuit.
* `tunnel_type` - The type of the tunnel associated with the virtual circuit.

