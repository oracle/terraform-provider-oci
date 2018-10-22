---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_fast_connect_provider_services"
sidebar_current: "docs-oci-datasource-core-fast_connect_provider_services"
description: |-
  Provides the list of Fast Connect Provider Services in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_fast_connect_provider_services
This data source provides the list of Fast Connect Provider Services in Oracle Cloud Infrastructure Core service.

Lists the service offerings from supported providers. You need this
information so you can specify your desired provider and service
offering when you create a virtual circuit.

For the compartment ID, provide the OCID of your tenancy (the root compartment).

For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).


## Example Usage

```hcl
data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


## Attributes Reference

The following attributes are exported:

* `fast_connect_provider_services` - The list of fast_connect_provider_services.

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

