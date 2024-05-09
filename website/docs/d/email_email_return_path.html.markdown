---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_return_path"
sidebar_current: "docs-oci-datasource-email-email_return_path"
description: |-
  Provides details about a specific Email Return Path in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_email_return_path
This data source provides details about a specific Email Return Path resource in Oracle Cloud Infrastructure Email service.

Retrieves the specified email return path.

## Example Usage

```hcl
data "oci_email_email_return_path" "test_email_return_path" {
	#Required
	email_return_path_id = oci_email_email_return_path.test_email_return_path.id
}
```

## Argument Reference

The following arguments are supported:

* `email_return_path_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this email return path.


## Attributes Reference

The following attributes are exported:

* `cname_record_value` - The DNS CNAME record value to provision to the Return Patn DNS subdomain, when using the CNAME method for Email Return Path setup (preferred). 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email return path. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the email return path. Avoid entering confidential information.
* `dns_subdomain_name` - The name of the DNS subdomain that must be provisioned to enable email recipients to verify Email Return Path. It is usually created with a CNAME record set to the cnameRecordValue. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email return path. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state. 
* `name` - The email return path domain in the Internet Domain Name System (DNS).  Example: `iad1.rp.example.com` 
* `parent_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain that this email return path belongs to. 
* `state` - The current state of the email return path.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the email return path was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z` 
* `time_updated` - The time of the last change to the Email Return Path configuration, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ". 

