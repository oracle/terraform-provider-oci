---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_db_preview_versions"
sidebar_current: "docs-oci-datasource-database-autonomous_db_preview_versions"
description: |-
  Provides the list of Autonomous Db Preview Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_db_preview_versions
This data source provides the list of Autonomous Db Preview Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Autonomous Database versions. Note that preview version software is only available for
databases with [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).


## Example Usage

```hcl
data "oci_database_autonomous_db_preview_versions" "test_autonomous_db_preview_versions" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_db_preview_versions` - The list of autonomous_db_preview_versions.

### AutonomousDbPreviewVersion Reference

The following attributes are exported:

* `db_workload` - The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type. 
* `details` - A URL that points to a detailed description of the preview version.
* `time_preview_begin` - The date and time when the preview version availability begins.
* `time_preview_end` - The date and time when the preview version availability ends.
* `version` - A valid Autonomous Database preview version.

