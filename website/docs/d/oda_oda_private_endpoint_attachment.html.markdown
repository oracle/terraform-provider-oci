---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint_attachment"
sidebar_current: "docs-oci-datasource-oda-oda_private_endpoint_attachment"
description: |-
  Provides details about a specific Oda Private Endpoint Attachment in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_private_endpoint_attachment
This data source provides details about a specific Oda Private Endpoint Attachment resource in Oracle Cloud Infrastructure Digital Assistant service.

Gets the specified ODA Private Endpoint Attachment.

## Example Usage

```hcl
data "oci_oda_oda_private_endpoint_attachment" "test_oda_private_endpoint_attachment" {
	#Required
	oda_private_endpoint_attachment_id = oci_oda_oda_private_endpoint_attachment.test_oda_private_endpoint_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `oda_private_endpoint_attachment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of ODA Private Endpoint Attachment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint attachment belongs to.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint Attachment.
* `oda_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached ODA Instance.
* `oda_private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint.
* `state` - The current state of the ODA Private Endpoint attachment.
* `time_created` - When the resource was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the resource was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.

