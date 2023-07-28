---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_news_report"
sidebar_current: "docs-oci-resource-opsi-news_report"
description: |-
  Provides the News Report resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_news_report
This resource provides the News Report resource in Oracle Cloud Infrastructure Opsi service.

Create a news report in Operations Insights. The report will be enabled in Operations Insights. Insights will be emailed as per selected frequency.


## Example Usage

```hcl
resource "oci_opsi_news_report" "test_news_report" {
	#Required
	compartment_id = var.compartment_id
	content_types {
		#Required
		capacity_planning_resources = var.news_report_content_types_capacity_planning_resources
	}
	description = var.news_report_description
	locale = var.news_report_locale
	name = var.news_report_name
	news_frequency = var.news_report_news_frequency
	ons_topic_id = oci_opsi_ons_topic.test_ons_topic.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	status = var.news_report_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier where the news report will be created.
* `content_types` - (Required) (Updatable) Content types that the news report can handle.
	* `capacity_planning_resources` - (Required) (Updatable) Supported resources for capacity planning content type.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) The description of the news report. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locale` - (Required) (Updatable) Language of the news report.
* `name` - (Required) The news report name.
* `news_frequency` - (Required) (Updatable) News report frequency.
* `ons_topic_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ONS topic.
* `status` - (Optional) (Updatable) Defines if the news report will be enabled or disabled.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `content_types` - Content types that the news report can handle.
	* `capacity_planning_resources` - Supported resources for capacity planning content type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the news report. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the news report resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `locale` - Language of the news report.
* `name` - The news report name.
* `news_frequency` - News report frequency.
* `ons_topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ONS topic.
* `state` - The current state of the news report.
* `status` - Indicates the status of a news report in Operations Insights.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the news report was first enabled. An RFC3339 formatted datetime string.
* `time_updated` - The time the news report was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the News Report
	* `update` - (Defaults to 20 minutes), when updating the News Report
	* `delete` - (Defaults to 20 minutes), when destroying the News Report


## Import

NewsReports can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_news_report.test_news_report "id"
```

