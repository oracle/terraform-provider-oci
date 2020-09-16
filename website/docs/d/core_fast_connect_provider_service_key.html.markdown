---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_fast_connect_provider_service_key"
sidebar_current: "docs-oci-datasource-core-fast_connect_provider_service_key"
description: |-
  Provides details about a specific Fast Connect Provider Service Key in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_fast_connect_provider_service_key
This data source provides details about a specific Fast Connect Provider Service Key resource in Oracle Cloud Infrastructure Core service.

Gets the specified provider service key's information. Use this operation to validate a
provider service key. An invalid key returns a 404 error.


## Example Usage

```hcl
data "oci_core_fast_connect_provider_service_key" "test_fast_connect_provider_service_key" {
	#Required
	provider_service_id = data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id
	provider_service_key_name = var.fast_connect_provider_service_key_provider_service_key_name
}
```

## Argument Reference

The following arguments are supported:

* `provider_service_id` - (Required) The OCID of the provider service.
* `provider_service_key_name` - (Required) The provider service key that the provider gives you when you set up a virtual circuit connection from the provider to Oracle Cloud Infrastructure. You can set up that connection and get your provider service key at the provider's website or portal. For the portal location, see the `description` attribute of the [FastConnectProviderService](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/FastConnectProviderService/). 


## Attributes Reference

The following attributes are exported:

* `bandwidth_shape_name` - The provisioned data rate of the connection.  To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/FastConnectProviderService/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps` 
* `name` - The service key that the provider gives you when you set up a virtual circuit connection from the provider to Oracle Cloud Infrastructure. Use this value as the `providerServiceKeyName` query parameter for [GetFastConnectProviderServiceKey](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/FastConnectProviderServiceKey/GetFastConnectProviderServiceKey). 
* `peering_location` - The provider's peering location.

