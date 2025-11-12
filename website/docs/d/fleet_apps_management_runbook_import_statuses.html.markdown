---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook_import_statuses"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbook_import_statuses"
description: |-
  Provides the list of Runbook Import Statuses in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbook_import_statuses
This data source provides the list of Runbook Import Statuses in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the Runbook import status in the specified compartment.
The query parameter `compartmentId` is required.


## Example Usage

```hcl
data "oci_fleet_apps_management_runbook_import_statuses" "test_runbook_import_statuses" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `runbook_import_status_collection` - The list of runbook_import_status_collection.

### RunbookImportStatus Reference

The following attributes are exported:

* `items` - List of Runbook import status.
	* `runbook_id` - The OCID of the runbook.
	* `runbook_name` - Runbook name.
	* `runbook_version` - Runbook version.
	* `status` - Status.
	* `tracking_id` - Tracking/Import identifier.

