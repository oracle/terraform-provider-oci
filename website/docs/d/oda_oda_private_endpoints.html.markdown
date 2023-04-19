---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoints"
sidebar_current: "docs-oci-datasource-oda-oda_private_endpoints"
description: |-
  Provides the list of Oda Private Endpoints in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_private_endpoints
This data source provides the list of Oda Private Endpoints in Oracle Cloud Infrastructure Digital Assistant service.

Returns a page of ODA Private Endpoints that belong to the specified
compartment.

If the `opc-next-page` header appears in the response, then
there are more items to retrieve. To get the next page in the subsequent
GET request, include the header's value as the `page` query parameter.


## Example Usage

```hcl
data "oci_oda_oda_private_endpoints" "test_oda_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oda_private_endpoint_display_name
	state = var.oda_private_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) List the ODA Private Endpoints that belong to this compartment.
* `display_name` - (Optional) List only the information for the Digital Assistant instance with this user-friendly name. These names don't have to be unique and may change.  Example: `My new resource` 
* `state` - (Optional) List only the ODA Private Endpoints that are in this lifecycle state.


## Attributes Reference

The following attributes are exported:

* `oda_private_endpoint_collection` - The list of oda_private_endpoint_collection.

### OdaPrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint belongs to.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the ODA private endpoint.
* `display_name` - User-defined name for the ODA private endpoint. Avoid entering confidential information. You can change this value. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that was assigned when the ODA private endpoint was created.
* `nsg_ids` - List of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of [network security groups](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm)
* `state` - The current state of the ODA private endpoint.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.
* `time_created` - When the resource was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the resource was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.

