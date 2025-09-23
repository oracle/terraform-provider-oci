---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_tenancy_attachments"
sidebar_current: "docs-oci-datasource-resource_analytics-tenancy_attachments"
description: |-
  Provides the list of Tenancy Attachments in Oracle Cloud Infrastructure Resource Analytics service
---

# Data Source: oci_resource_analytics_tenancy_attachments
This data source provides the list of Tenancy Attachments in Oracle Cloud Infrastructure Resource Analytics service.

Gets a list of TenancyAttachments.


## Example Usage

```hcl
data "oci_resource_analytics_tenancy_attachments" "test_tenancy_attachments" {

	#Optional
	id = var.tenancy_attachment_id
	resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
	state = var.tenancy_attachment_state
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the TenancyAttachment.
* `resource_analytics_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a ResourceAnalyticsInstance.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `tenancy_attachment_collection` - The list of tenancy_attachment_collection.

### TenancyAttachment Reference

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

