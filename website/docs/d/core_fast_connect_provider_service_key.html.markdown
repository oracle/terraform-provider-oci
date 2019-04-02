---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_fast_connect_provider_service_key"
sidebar_current: "docs-oci-datasource-core-fast_connect_provider_service_key"
description: |-
  Provides details about a specific Fast Connect Provider Service Key in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_fast_connect_provider_service_key
This data source provides details about a specific Fast Connect Provider Service Key resource in Oracle Cloud Infrastructure Core service.

Gets the specified provider service key's information.


## Example Usage

```hcl
data "oci_core_fast_connect_provider_service_key" "test_fast_connect_provider_service_key" {
	#Required
	provider_service_id = "${oci_core_provider_service.test_provider_service.id}"
	provider_service_key_name = "${var.fast_connect_provider_service_key_provider_service_key_name}"
}
```

## Argument Reference

The following arguments are supported:

* `provider_service_id` - (Required) The OCID of the provider service.
* `provider_service_key_name` - (Required) The provider service key name.


## Attributes Reference

The following attributes are exported:

* `bandwidth_shape_name` - The provisioned data rate of the connection.  To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/FastConnectProviderService/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps` 
* `name` - The name of the service key offered by the provider. 
* `peering_location` - The provider's peering location.

