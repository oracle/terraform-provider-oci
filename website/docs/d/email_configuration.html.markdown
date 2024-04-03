---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_configuration"
sidebar_current: "docs-oci-datasource-email-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Email service.

Returns  email configuration associated with the specified compartment.


## Example Usage

```hcl
data "oci_email_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The root compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) (same as the tenancy OCID)
* `http_submit_endpoint` - Endpoint used to submit emails via the HTTP email submission API
* `smtp_submit_endpoint` - Endpoint used to submit emails via the standard SMTP submission protocol. Note that TLS 1.2 and standard SMTP authentication is required for submission.

