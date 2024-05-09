---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_return_path"
sidebar_current: "docs-oci-resource-email-email_return_path"
description: |-
  Provides the Email Return Path resource in Oracle Cloud Infrastructure Email service
---

# oci_email_email_return_path
This resource provides the Email Return Path resource in Oracle Cloud Infrastructure Email service.

Creates a new email return path. Avoid entering confidential information.

## Example Usage

```hcl
resource "oci_email_email_return_path" "test_email_return_path" {
	#Required
	parent_resource_id = oci_usage_proxy_resource.test_resource.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.email_return_path_description
	freeform_tags = {"Department"= "Finance"}
	name = var.email_return_path_name
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A string that describes the details about the email return path. It does not have to be unique, and you can change it. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Optional) The name of the email return path domain in the Internet Domain Name System (DNS). The name must be a subdomain of the email domain used to send emails. The email return path name must be globally unique for this tenancy. If you do not provide the email return path name, we will generate one for you. If you do provide the email return path name, we suggest adding a short region indicator to allow using the same parent domain in other regions you might be subscribed to. Domain names limited to ASCII characters use alphanumeric, dash ("-"), and dot (".") characters. The dash and dot are only allowed between alphanumeric characters. Non-ASCII domain names should adopt IDNA2008 normalization (RFC 5891-5892). 
* `parent_resource_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this email return path. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Email Return Path
	* `update` - (Defaults to 20 minutes), when updating the Email Return Path
	* `delete` - (Defaults to 20 minutes), when destroying the Email Return Path


## Import

EmailReturnPaths can be imported using the `id`, e.g.

```
$ terraform import oci_email_email_return_path.test_email_return_path "id"
```

