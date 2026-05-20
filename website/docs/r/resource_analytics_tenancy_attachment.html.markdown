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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/resource-analytics/latest/TenancyAttachment

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

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

* `data_population_status` - The overall status of the data population from the tenancy.
* `description` - A description of the tenancy.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the TenancyAttachment.
* `is_reporting_tenancy` - Whether the tenancy is the tenancy used when creating Resource Analytics Instance.
* `lifecycle_details` - A message that describes the current state of the TenancyAttachment in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `monitored_regions` - List of monitored regions with their data population status.
	* `data_population` - Data population status for a monitored region in the tenancy.
		* `in_progress_count` - The number of data population tasks currently in progress.
		* `status` - The overall status of the data population from the monitored region of the tenancy.
		* `succeeded_count` - The number of data population tasks that have succeeded.
		* `time_ended` - The date and time the data population task completed, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
		* `time_started` - The date and time the data population task was started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
		* `total_count` - The total number of data population tasks.
	* `region_id` - The [Region Identifier](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) of the monitored region. E.g. us-ashburn-1
* `resource_analytics_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this TenancyAttachment.
* `state` - The current state of the TenancyAttachment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy associated with this TenancyAttachment.
* `time_created` - The date and time the TenancyAttachment was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_data_population_ended` - The date and time the data population tasks completed, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_data_population_started` - The date and time the data population tasks started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
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

