---
layout: "oci"
page_title: "OCI: oci_core_fast_connect_provider_service"
sidebar_current: "docs-oci-datasource-core-fast_connect_provider_service"
description: |-
  Provides details about a specific FastConnectProviderService
---

# Data Source: oci_core_fast_connect_provider_service
The FastConnectProviderService data source provides details about a specific FastConnectProviderService

Gets the specified provider service.
For more information, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).


## Example Usage

```hcl
data "oci_core_fast_connect_provider_service" "test_fast_connect_provider_service" {
	#Required
	provider_service_id = "${var.fast_connect_provider_service_provider_service_id}"
}
```

## Argument Reference

The following arguments are supported:

* `provider_service_id` - (Required) The OCID of the provider service.


## Attributes Reference

The following attributes are exported:

* `description` - A description of the service offered by the provider. 
* `id` - The OCID of the service offered by the provider. 
* `private_peering_bgp_management` - Private peering BGP management. 
* `provider_name` - The name of the provider. 
* `provider_service_name` - The name of the service offered by the provider. 
* `public_peering_bgp_management` - Public peering BGP management. 
* `supported_virtual_circuit_types` - An array of virtual circuit types supported by this service. 
* `type` - Provider service type. 

