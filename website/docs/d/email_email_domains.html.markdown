---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_domains"
sidebar_current: "docs-oci-datasource-email-email_domains"
description: |-
  Provides the list of Email Domains in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_email_domains
This data source provides the list of Email Domains in Oracle Cloud Infrastructure Email service.

Lists email domains in the specified compartment.

## Example Usage

```hcl
data "oci_email_email_domains" "test_email_domains" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.email_domain_id
	name = var.email_domain_name
	state = var.email_domain_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment.
* `id` - (Optional) A filter to only return resources that match the given id exactly. 
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) Filter returned list by specified lifecycle state. This parameter is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `email_domain_collection` - The list of email_domain_collection.

### EmailDomain Reference

The following attributes are exported:

* `active_dkim_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM key that will be used to sign mail sent from this email domain. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email domain. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of an email domain.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain. 
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components). 
* `name` - The name of the email domain in the Internet Domain Name System (DNS).  Example: `mydomain.example.com` 
* `state` - The current state of the email domain.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the email domain was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z` 

