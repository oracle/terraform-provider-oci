---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint"
sidebar_current: "docs-oci-datasource-oda-oda_private_endpoint"
description: |-
  Provides details about a specific Oda Private Endpoint in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_private_endpoint
This data source provides details about a specific Oda Private Endpoint resource in Oracle Cloud Infrastructure Digital Assistant service.

Gets the specified ODA Private Endpoint.

## Example Usage

```hcl
data "oci_oda_oda_private_endpoint" "test_oda_private_endpoint" {
	#Required
	oda_private_endpoint_id = oci_oda_oda_private_endpoint.test_oda_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `oda_private_endpoint_id` - (Required) Unique ODA Private Endpoint identifier which is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

