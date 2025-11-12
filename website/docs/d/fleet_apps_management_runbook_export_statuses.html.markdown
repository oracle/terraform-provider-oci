---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook_export_statuses"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbook_export_statuses"
description: |-
  Provides the list of Runbook Export Statuses in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbook_export_statuses
This data source provides the list of Runbook Export Statuses in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the Runbook export status in the specified compartment.
The query parameter `compartmentId` is required.


## Example Usage

```hcl
data "oci_fleet_apps_management_runbook_export_statuses" "test_runbook_export_statuses" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `runbook_export_status_collection` - The list of runbook_export_status_collection.

### RunbookExportStatus Reference

The following attributes are exported:

* `items` - List of Runbook export status.
	* `runbook_id` - The OCID of the runbook.
	* `runbook_name` - Runbook name.
	* `runbook_version` - Runbook version.
	* `status` - Runbook export status.
	* `tracking_id` - Tracking/Export identifier.

