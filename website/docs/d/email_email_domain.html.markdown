---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_domain"
sidebar_current: "docs-oci-datasource-email-email_domain"
description: |-
  Provides details about a specific Email Domain in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_email_domain
This data source provides details about a specific Email Domain resource in Oracle Cloud Infrastructure Email service.

Retrieves the specified email domain.

## Example Usage

```hcl
data "oci_email_email_domain" "test_email_domain" {
	#Required
	email_domain_id = oci_email_email_domain.test_email_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `email_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this email domain.


## Attributes Reference

The following attributes are exported:

* `active_dkim_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM key that will be used to sign mail sent from this email domain. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email domain. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of a email domain.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain. 
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components). 
* `name` - The name of the email domain in the Internet Domain Name System (DNS).  Example: `example.net` 
* `state` - The current state of the email domain.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the email domain was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z` 

