---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook_import"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbook_import"
description: |-
  Provides details about a specific Runbook Import in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbook_import
This data source provides details about a specific Runbook Import resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Get the runbook import status for provided runbook and importId.

## Example Usage

```hcl
data "oci_fleet_apps_management_runbook_import" "test_runbook_import" {
	#Required
	import_id = var.runbook_import_id
	runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
}
```

## Argument Reference

The following arguments are supported:

* `import_id` - (Required) Unique tracking identifier to fetch runbook import status.
* `runbook_id` - (Required) Unique Runbook identifier.


## Attributes Reference

The following attributes are exported:

* `details` - Map of runbook import details.
* `runbook_id` - The OCID of the runbook.
* `runbook_name` - Runbook name.
* `runbook_version` - Runbook version.
* `status` - Status.
* `tracking_id` - Tracking/Import identifier.

