---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entities_summary"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entities_summary"
description: |-
  Provides details about a specific Log Analytics Entities Summary in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entities_summary
This data source provides details about a specific Log Analytics Entities Summary resource in Oracle Cloud Infrastructure Log Analytics service.

Returns log analytics entities count summary report.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entities_summary" "test_log_analytics_entities_summary" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.log_analytics_entities_summary_namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `active_entities_count` - Total number of ACTIVE entities 
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `entities_with_has_logs_collected_count` - Entities with log collection enabled 
* `entities_with_management_agent_count` - Entities with management agent 

