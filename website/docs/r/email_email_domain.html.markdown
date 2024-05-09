---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_domain"
sidebar_current: "docs-oci-resource-email-email_domain"
description: |-
  Provides the Email Domain resource in Oracle Cloud Infrastructure Email service
---

# oci_email_email_domain
This resource provides the Email Domain resource in Oracle Cloud Infrastructure Email service.

Creates a new email domain. Avoid entering confidential information.

## Example Usage

```hcl
resource "oci_email_email_domain" "test_email_domain" {
	#Required
	compartment_id = var.compartment_id
	name = var.email_domain_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.email_domain_description
	domain_verification_id = oci_email_domain_verification.test_domain_verification.id
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for this email domain.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A string that describes the details about the domain. It does not have to be unique, and you can change it. Avoid entering confidential information.
* `domain_verification_id` - (Optional) (Updatable) Id for Domain in Domain Management (under governance) if DOMAINID verification method used.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `name` - (Required) The name of the email domain in the Internet Domain Name System (DNS). The email domain name must be unique in the region for this tenancy. Domain names limited to ASCII characters use alphanumeric, dash ("-"), and dot (".") characters. The dash and dot are only allowed between alphanumeric characters. For details, see [RFC 5321, section 4.1.2](https://tools.ietf.org/html/rfc5321#section-4.1.2) Non-ASCII domain names should adopt IDNA2008 normalization (RFC 5891-5892).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `active_dkim_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM key that will be used to sign mail sent from this email domain.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email domain.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - The description of an email domain.
* `domain_verification_id` - Id for Domain in Domain Management (under governance) if DOMAINID verification method used.
* `domain_verification_status` - The current domain verification status.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components). 
* `name` - The name of the email domain in the Internet Domain Name System (DNS).  Example: `mydomain.example.com`
* `state` - The current state of the email domain.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the email domain was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z`

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Email Domain
	* `update` - (Defaults to 20 minutes), when updating the Email Domain
	* `delete` - (Defaults to 20 minutes), when destroying the Email Domain

## Import

EmailDomains can be imported using the `id`, e.g.

```
$ terraform import oci_email_email_domain.test_email_domain "id"
```
