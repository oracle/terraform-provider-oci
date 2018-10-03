# oci_core_fast_connect_provider_service

## FastConnectProviderService Singular DataSource

### FastConnectProviderService Reference

The following attributes are exported:

* `description` - A description of the service offered by the provider. 
* `id` - The OCID of the service offered by the provider. 
* `private_peering_bgp_management` - Private peering BGP management. 
* `provider_name` - The name of the provider. 
* `provider_service_name` - The name of the service offered by the provider. 
* `public_peering_bgp_management` - Public peering BGP management. 
* `supported_virtual_circuit_types` - An array of virtual circuit types supported by this service. 
* `type` - Provider service type. 



### Get Operation
Gets the specified provider service.
For more information, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).


The following arguments are supported:

* `provider_service_id` - (Required) The OCID of the provider service.


### Example Usage

```hcl
data "oci_core_fast_connect_provider_service" "test_fast_connect_provider_service" {
	#Required
	provider_service_id = "${oci_core_fast_connect_provider_service.test_provider_service.id}"
}
```
# oci_core_fast_connect_provider_services

## FastConnectProviderService DataSource

Gets a list of fast_connect_provider_services.

### List Operation
Lists the service offerings from supported providers. You need this
information so you can specify your desired provider and service
offering when you create a virtual circuit.

For the compartment ID, provide the OCID of your tenancy (the root compartment).

For more information, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


The following attributes are exported:

* `fast_connect_provider_services` - The list of fast_connect_provider_services.

### Example Usage

```hcl
data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
### FastConnectProviderService Reference

The following attributes are exported:

* `description` - A description of the service offered by the provider. 
* `id` - The OCID of the service offered by the provider. 
* `private_peering_bgp_management` - Private peering BGP management. 
* `provider_name` - The name of the provider. 
* `provider_service_name` - The name of the service offered by the provider. 
* `public_peering_bgp_management` - Public peering BGP management. 
* `supported_virtual_circuit_types` - An array of virtual circuit types supported by this service. 
* `type` - Provider service type. 
