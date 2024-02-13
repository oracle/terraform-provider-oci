---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_dkim"
sidebar_current: "docs-oci-resource-email-dkim"
description: |-
  Provides the Dkim resource in Oracle Cloud Infrastructure Email service
---

# oci_email_dkim
This resource provides the Dkim resource in Oracle Cloud Infrastructure Email service.

Creates a new DKIM for an email domain.
This DKIM signs all approved senders in the tenancy that are in this email domain.
Best security practices indicate to periodically rotate the DKIM that is doing the signing.
When a second DKIM is applied, all senders seamlessly pick up the new key
without interruption in signing.


## Example Usage

```hcl
resource "oci_email_dkim" "test_dkim" {
	#Required
	email_domain_id = oci_email_email_domain.test_email_domain.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.dkim_description
	freeform_tags = {"Department"= "Finance"}
	name = var.dkim_name
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A string that describes the details about the DKIM. It does not have to be unique, and you can change it. Avoid entering confidential information. 
* `email_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Optional) The DKIM selector. This selector is required to be globally unique for this email domain. If you do not provide the selector, we will generate one for you. If you do provide the selector, we suggest adding a short region indicator to differentiate from your signing of emails in other regions you might be subscribed to. Selectors limited to ASCII characters can use alphanumeric, dash ("-"), and dot (".") characters. Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).

	Avoid entering confidential information.

	Example: `mydomain-phx-20210228`


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cname_record_value` - The DNS CNAME record value to provision to the DKIM DNS subdomain, when using the CNAME method for DKIM setup (preferred).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this DKIM.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - The description of the DKIM. Avoid entering confidential information.
* `dns_subdomain_name` - The name of the DNS subdomain that must be provisioned to enable email recipients to verify DKIM signatures. It is usually created with a CNAME record set to the cnameRecordValue. 
* `email_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain that this DKIM belongs to. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource. 
* `name` - The DKIM selector. If the same domain is managed in more than one region, each region must use different selectors. 
* `state` - The current state of the DKIM.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the DKIM was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z` 
* `time_updated` - The time of the last change to the DKIM configuration, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ". 
* `txt_record_value` - The DNS TXT record value to provision to the DKIM DNS subdomain in place of using a CNAME record. This is used in cases where a CNAME cannot be used, such as when the cnameRecordValue would exceed the maximum length for a DNS entry. You can also use this if you have an existing procedure to directly provision TXT records for DKIM. Many DNS APIs require you to break this string into segments of fewer than 255 characters.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dkim
	* `update` - (Defaults to 20 minutes), when updating the Dkim
	* `delete` - (Defaults to 20 minutes), when destroying the Dkim

## Import

Dkims can be imported using the `id`, e.g.

```
$ terraform import oci_email_dkim.test_dkim "id"
```
