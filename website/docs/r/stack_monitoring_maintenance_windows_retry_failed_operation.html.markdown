---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_maintenance_windows_retry_failed_operation"
sidebar_current: "docs-oci-resource-stack_monitoring-maintenance_windows_retry_failed_operation"
description: |-
  Provides the Maintenance Windows Retry Failed Operation resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_maintenance_windows_retry_failed_operation
This resource provides the Maintenance Windows Retry Failed Operation resource in Oracle Cloud Infrastructure Stack Monitoring service.

Retry the last failed operation. The operation failed will be the most recent one. It won't apply for previous failed operations.


## Example Usage

```hcl
resource "oci_stack_monitoring_maintenance_windows_retry_failed_operation" "test_maintenance_windows_retry_failed_operation" {
	#Required
	maintenance_window_id = oci_stack_monitoring_maintenance_window.test_maintenance_window.id
}
```

## Argument Reference

The following arguments are supported:

* `maintenance_window_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of maintenance window.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Maintenance Windows Retry Failed Operation
	* `update` - (Defaults to 20 minutes), when updating the Maintenance Windows Retry Failed Operation
	* `delete` - (Defaults to 20 minutes), when destroying the Maintenance Windows Retry Failed Operation


## Import

MaintenanceWindowsRetryFailedOperation can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_maintenance_windows_retry_failed_operation.test_maintenance_windows_retry_failed_operation "id"
```

