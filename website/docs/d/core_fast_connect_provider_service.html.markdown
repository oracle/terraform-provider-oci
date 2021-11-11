---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_fast_connect_provider_service"
sidebar_current: "docs-oci-datasource-core-fast_connect_provider_service"
description: |-
  Provides details about a specific Fast Connect Provider Service in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_fast_connect_provider_service
This data source provides details about a specific Fast Connect Provider Service resource in Oracle Cloud Infrastructure Core service.

Gets the specified provider service.
For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).


## Example Usage

```hcl
data "oci_core_fast_connect_provider_service" "test_fast_connect_provider_service" {
	#Required
	provider_service_id = data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id
}
```

## Argument Reference

The following arguments are supported:

* `provider_service_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provider service.


## Attributes Reference

The following attributes are exported:

* `bandwith_shape_management` - Who is responsible for managing the virtual circuit bandwidth. 
* `customer_asn_management` - Who is responsible for managing the ASN information for the network at the other end of the connection from Oracle. 
* `description` - The location of the provider's website or portal. This portal is where you can get information about the provider service, create a virtual circuit connection from the provider to Oracle Cloud Infrastructure, and retrieve your provider service key for that virtual circuit connection.  Example: `https://example.com` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service offered by the provider. 
* `private_peering_bgp_management` - Who is responsible for managing the private peering BGP information. 
* `provider_name` - The name of the provider. 
* `provider_service_key_management` - Who is responsible for managing the provider service key. 
* `provider_service_name` - The name of the service offered by the provider. 
* `public_peering_bgp_management` - Who is responsible for managing the public peering BGP information. 
* `required_total_cross_connects` - Total number of cross-connect or cross-connect groups required for the virtual circuit. 
* `supported_virtual_circuit_types` - An array of virtual circuit types supported by this service. 
* `type` - Provider service type. 

