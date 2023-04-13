---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint_attachments"
sidebar_current: "docs-oci-datasource-oda-oda_private_endpoint_attachments"
description: |-
  Provides the list of Oda Private Endpoint Attachments in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_private_endpoint_attachments
This data source provides the list of Oda Private Endpoint Attachments in Oracle Cloud Infrastructure Digital Assistant service.

Returns a page of ODA Instances attached to this ODA Private Endpoint.

If the `opc-next-page` header appears in the response, then
there are more items to retrieve. To get the next page in the subsequent
GET request, include the header's value as the `page` query parameter.


## Example Usage

```hcl
data "oci_oda_oda_private_endpoint_attachments" "test_oda_private_endpoint_attachments" {
	#Required
	compartment_id = var.compartment_id
	oda_private_endpoint_id = oci_oda_oda_private_endpoint.test_oda_private_endpoint.id

	#Optional
	state = var.oda_private_endpoint_attachment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) List the ODA Private Endpoint Attachments that belong to this compartment.
* `oda_private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of ODA Private Endpoint.
* `state` - (Optional) List only the ODA Private Endpoint Attachments that are in this lifecycle state.


## Attributes Reference

The following attributes are exported:

* `oda_private_endpoint_attachment_collection` - The list of oda_private_endpoint_attachment_collection.

### OdaPrivateEndpointAttachment Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint attachment belongs to.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint Attachment.
* `oda_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached ODA Instance.
* `oda_private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint.
* `state` - The current state of the ODA Private Endpoint attachment.
* `time_created` - When the resource was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the resource was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.

