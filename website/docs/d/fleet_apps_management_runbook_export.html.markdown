---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook_export"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbook_export"
description: |-
  Provides details about a specific Runbook Export in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbook_export
This data source provides details about a specific Runbook Export resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Get the runbook export status for provided runbook and exportId.

## Example Usage

```hcl
data "oci_fleet_apps_management_runbook_export" "test_runbook_export" {
	#Required
	export_id = var.runbook_export_id
	runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
}
```

## Argument Reference

The following arguments are supported:

* `export_id` - (Required) Unique tracking identifier to fetch runbook export status.
* `runbook_id` - (Required) Unique Runbook identifier.


## Attributes Reference

The following attributes are exported:

* `details` - Map of runbook export details.
* `runbook_id` - The OCID of the runbook.
* `runbook_name` - Runbook name.
* `runbook_version` - Runbook version.
* `status` - Status.
* `tracking_id` - Tracking/Export identifier.

