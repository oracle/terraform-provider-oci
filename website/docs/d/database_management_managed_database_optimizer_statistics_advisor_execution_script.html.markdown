---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_optimizer_statistics_advisor_execution_script"
sidebar_current: "docs-oci-datasource-database_management-managed_database_optimizer_statistics_advisor_execution_script"
description: |-
  Provides details about a specific Managed Database Optimizer Statistics Advisor Execution Script in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_optimizer_statistics_advisor_execution_script
This data source provides details about a specific Managed Database Optimizer Statistics Advisor Execution Script resource in Oracle Cloud Infrastructure Database Management service.

Gets the Oracle system-generated script for the specified Optimizer Statistics Advisor execution.

## Example Usage

```hcl
data "oci_database_management_managed_database_optimizer_statistics_advisor_execution_script" "test_managed_database_optimizer_statistics_advisor_execution_script" {
	#Required
	execution_name = var.managed_database_optimizer_statistics_advisor_execution_script_execution_name
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	task_name = var.managed_database_optimizer_statistics_advisor_execution_script_task_name
}
```

## Argument Reference

The following arguments are supported:

* `execution_name` - (Required) The name of the Optimizer Statistics Advisor execution.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `task_name` - (Required) The name of the optimizer statistics collection execution task.


## Attributes Reference

The following attributes are exported:

* `script` - The Optimizer Statistics Advisor execution script.

