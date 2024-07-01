---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_news_reports"
sidebar_current: "docs-oci-datasource-opsi-news_reports"
description: |-
  Provides the list of News Reports in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_news_reports
This data source provides the list of News Reports in Oracle Cloud Infrastructure Opsi service.

Gets a list of news reports based on the query parameters specified. Either compartmentId or id query parameter must be specified.


## Example Usage

```hcl
data "oci_opsi_news_reports" "test_news_reports" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.news_report_compartment_id_in_subtree
	news_report_id = oci_opsi_news_report.test_news_report.id
	state = var.news_report_state
	status = var.news_report_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments. 
* `news_report_id` - (Optional) Unique Ops Insights news report identifier
* `state` - (Optional) Lifecycle states
* `status` - (Optional) Resource Status


## Attributes Reference

The following attributes are exported:

* `news_report_collection` - The list of news_report_collection.

### NewsReport Reference

The following attributes are exported:

* `are_child_compartments_included` - A flag to consider the resources within a given compartment and all sub-compartments.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `content_types` - Content types that the news report can handle.
	* `capacity_planning_resources` - Supported resources for capacity planning content type.
	* `sql_insights_fleet_analysis_resources` - Supported resources for SQL insights - fleet analysis content type.
	* `sql_insights_performance_degradation_resources` - Supported resources for SQL insights - performance degradation content type.
	* `sql_insights_plan_changes_resources` - Supported resources for SQL insights - plan changes content type.
	* `sql_insights_top_databases_resources` - Supported resources for SQL insights - top databases content type.
	* `sql_insights_top_sql_by_insights_resources` - Supported resources for SQL insights - top SQL by insights content type.
	* `sql_insights_top_sql_resources` - Supported resources for SQL insights - top SQL content type.
* `day_of_week` - Day of the week in which the news report will be sent if the frequency is set to WEEKLY.
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
* `status` - Indicates the status of a news report in Ops Insights.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the news report was first enabled. An RFC3339 formatted datetime string.
* `time_updated` - The time the news report was updated. An RFC3339 formatted datetime string.

