
# oci_core_virtual_circuit_bandwidth_shapes

## VirtualCircuitBandwidthShape DataSource

Gets a list of virtual_circuit_bandwidth_shapes.

### List Operation
Gets the list of available virtual circuit bandwidth levels for a provider.
You need this information so you can specify your desired bandwidth level (shape) when you create a virtual circuit.

For more information about virtual circuits, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).

The following arguments are supported:

* `provider_service_id` - (Required) The OCID of the provider service.


The following attributes are exported:

* `virtual_circuit_bandwidth_shapes` - The list of virtual_circuit_bandwidth_shapes.

### Example Usage

```hcl
data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes" {
	#Required
	provider_service_id = "${oci_core_provider_service.test_provider_service.id}"
}
```
### VirtualCircuitBandwidthShape Reference

The following attributes are exported:

* `bandwidth_in_mbps` - The bandwidth in Mbps.  Example: `10000` 
* `name` - The name of the bandwidth shape.  Example: `10 Gbps` 
