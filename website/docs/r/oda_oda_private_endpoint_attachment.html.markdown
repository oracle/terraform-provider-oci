---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint_attachment"
sidebar_current: "docs-oci-resource-oda-oda_private_endpoint_attachment"
description: |-
  Provides the Oda Private Endpoint Attachment resource in Oracle Cloud Infrastructure Digital Assistant service
---

# oci_oda_oda_private_endpoint_attachment
This resource provides the Oda Private Endpoint Attachment resource in Oracle Cloud Infrastructure Digital Assistant service.

Starts an asynchronous job to create an ODA Private Endpoint Attachment.

To monitor the status of the job, take the `opc-work-request-id` response
header value and use it to call `GET /workRequests/{workRequestID}`.


## Example Usage

```hcl
resource "oci_oda_oda_private_endpoint_attachment" "test_oda_private_endpoint_attachment" {
	#Required
	oda_instance_id = oci_oda_oda_instance.test_oda_instance.id
	oda_private_endpoint_id = oci_oda_oda_private_endpoint.test_oda_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `oda_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached ODA Instance.
* `oda_private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint attachment belongs to.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint Attachment.
* `oda_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached ODA Instance.
* `oda_private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint.
* `state` - The current state of the ODA Private Endpoint attachment.
* `time_created` - When the resource was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the resource was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oda Private Endpoint Attachment
	* `update` - (Defaults to 20 minutes), when updating the Oda Private Endpoint Attachment
	* `delete` - (Defaults to 20 minutes), when destroying the Oda Private Endpoint Attachment


## Import

OdaPrivateEndpointAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_oda_oda_private_endpoint_attachment.test_oda_private_endpoint_attachment "id"
```

