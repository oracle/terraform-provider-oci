---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_virtual_circuit_bandwidth_shapes"
sidebar_current: "docs-oci-datasource-core-virtual_circuit_bandwidth_shapes"
description: |-
  Provides the list of Virtual Circuit Bandwidth Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_virtual_circuit_bandwidth_shapes
This data source provides the list of Virtual Circuit Bandwidth Shapes in Oracle Cloud Infrastructure Core service.

Gets the list of available virtual circuit bandwidth levels for a provider.
You need this information so you can specify your desired bandwidth level (shape) when you create a virtual circuit.

For more information about virtual circuits, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).


## Example Usage

```hcl
data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes" {
	#Required
	provider_service_id = data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id
}
```

## Argument Reference

The following arguments are supported:

* `provider_service_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provider service.


## Attributes Reference

The following attributes are exported:

* `virtual_circuit_bandwidth_shapes` - The list of virtual_circuit_bandwidth_shapes.

### VirtualCircuitBandwidthShape Reference

The following attributes are exported:

* `bandwidth_in_mbps` - The bandwidth in Mbps.  Example: `10000` 
* `name` - The name of the bandwidth shape.  Example: `10 Gbps` 

