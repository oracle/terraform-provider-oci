---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_tenancy_attachment"
sidebar_current: "docs-oci-resource-resource_analytics-tenancy_attachment"
description: |-
  Provides the Tenancy Attachment resource in Oracle Cloud Infrastructure Resource Analytics service
---

# oci_resource_analytics_tenancy_attachment
This resource provides the Tenancy Attachment resource in Oracle Cloud Infrastructure Resource Analytics service.

Creates a TenancyAttachment.


## Example Usage

```hcl
resource "oci_resource_analytics_tenancy_attachment" "test_tenancy_attachment" {
	#Required
	resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id

	#Optional
	description = var.tenancy_attachment_description
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) (Updatable) A description of the tenancy.
* `resource_analytics_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this TenancyAttachment.
* `tenancy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy associated with this TenancyAttachment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - A description of the tenancy.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the TenancyAttachment.
* `is_reporting_tenancy` - Whether the tenancy is the tenancy used when creating Resource Analytics Instance.
* `lifecycle_details` - A message that describes the current state of the TenancyAttachment in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `resource_analytics_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this TenancyAttachment.
* `state` - The current state of the TenancyAttachment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy associated with this TenancyAttachment.
* `time_created` - The date and time the TenancyAttachment was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the TenancyAttachment was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Tenancy Attachment
	* `update` - (Defaults to 20 minutes), when updating the Tenancy Attachment
	* `delete` - (Defaults to 20 minutes), when destroying the Tenancy Attachment


## Import

TenancyAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_resource_analytics_tenancy_attachment.test_tenancy_attachment "id"
```

