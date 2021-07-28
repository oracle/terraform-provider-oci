---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_dkims"
sidebar_current: "docs-oci-datasource-email-dkims"
description: |-
  Provides the list of Dkims in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_dkims
This data source provides the list of Dkims in Oracle Cloud Infrastructure Email service.

Lists DKIMs for a email domain.

## Example Usage

```hcl
data "oci_email_dkims" "test_dkims" {
	#Required
	email_domain_id = oci_email_email_domain.test_email_domain.id

	#Optional
	id = var.dkim_id
	name = var.dkim_name
	state = var.dkim_state
}
```

## Argument Reference

The following arguments are supported:

* `email_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain to which this DKIM belongs. 
* `id` - (Optional) A filter to only return resources that match the given id exactly. 
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) Filter returned list by specified lifecycle state. This parameter is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `dkim_collection` - The list of dkim_collection.

### Dkim Reference

The following attributes are exported:

* `cname_record_value` - The DNS CNAME record value to provision to the DKIM DNS subdomain, when using the CNAME method for DKIM setup (preferred). 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this DKIM. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the DKIM. Avoid entering confidential information.
* `dns_subdomain_name` - The name of the DNS subdomain that must be provisioned to enable email recipients to verify DKIM signatures. It is usually created with a CNAME record set to the cnameRecordValue 
* `email_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain that this DKIM belongs to. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource. 
* `name` - The DKIM selector. If the same domain is managed in more than one region, each region must use different selectors. 
* `state` - The current state of the DKIM.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the DKIM was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z` 
* `time_updated` - The time of the last change to the DKIM configuration, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ". 
* `txt_record_value` - The DNS TXT record value to provision to the DKIM DNS subdomain in place of using a CNAME record. This is used in cases where a CNAME can not be used, such as when the cnameRecordValue would exceed the maximum length for a DNS entry. This can also be used by customers who have an existing procedure to directly provision TXT records for DKIM. Be aware that many DNS APIs will require you to break this string into segments of less than 255 characters. 

