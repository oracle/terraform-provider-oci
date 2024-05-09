---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_return_paths"
sidebar_current: "docs-oci-datasource-email-email_return_paths"
description: |-
  Provides the list of Email Return Paths in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_email_return_paths
This data source provides the list of Email Return Paths in Oracle Cloud Infrastructure Email service.

Lists email return paths in the specified compartment or emaildomain.

## Example Usage

```hcl
data "oci_email_email_return_paths" "test_email_return_paths" {

	#Optional
	compartment_id = var.compartment_id
	id = var.email_return_path_id
	name = var.email_return_path_name
	parent_resource_id = oci_usage_proxy_resource.test_resource.id
	state = var.email_return_path_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID for the compartment.
* `id` - (Optional) A filter to only return resources that match the given id exactly. 
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `parent_resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Email Domain to which this Email Return Path belongs. 
* `state` - (Optional) Filter returned list by specified lifecycle state. This parameter is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `email_return_path_collection` - The list of email_return_path_collection.

### EmailReturnPath Reference

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

