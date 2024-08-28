---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_maintenance_windows_stop"
sidebar_current: "docs-oci-resource-stack_monitoring-maintenance_windows_stop"
description: |-
  Provides the Maintenance Windows Stop resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_maintenance_windows_stop
This resource provides the Maintenance Windows Stop resource in Oracle Cloud Infrastructure Stack Monitoring service.

Stop a maintenance window before the end time is reached.


## Example Usage

```hcl
resource "oci_stack_monitoring_maintenance_windows_stop" "test_maintenance_windows_stop" {
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
	* `create` - (Defaults to 20 minutes), when creating the Maintenance Windows Stop
	* `update` - (Defaults to 20 minutes), when updating the Maintenance Windows Stop
	* `delete` - (Defaults to 20 minutes), when destroying the Maintenance Windows Stop


## Import

MaintenanceWindowsStop can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_maintenance_windows_stop.test_maintenance_windows_stop "id"
```

