---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_rules_summary"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_rules_summary"
description: |-
  Provides details about a specific Namespace Rules Summary in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_rules_summary
This data source provides details about a specific Namespace Rules Summary resource in Oracle Cloud Infrastructure Log Analytics service.

Returns the count of detection rules in a compartment.


## Example Usage

```hcl
data "oci_log_analytics_namespace_rules_summary" "test_namespace_rules_summary" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.namespace_rules_summary_namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `ingest_time_rules_count` - The count of ingest time rules.
* `saved_search_rules_count` - The count of saved search rules.
* `total_count` - The total count of detection rules.

