---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_system_version_minor_versions"
sidebar_current: "docs-oci-datasource-database-system_version_minor_versions"
description: |-
  Provides the list of System Version Minor Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_system_version_minor_versions
This data source provides the list of System Version Minor Versions in Oracle Cloud Infrastructure Database service.

Retrieves a list of supported minor versions for the specified Exadata System Software major version. You must provide either a `shape` or `resourceId` value.

## Example Usage

```hcl
data "oci_database_system_version_minor_versions" "test_system_version_minor_versions" {
	#Required
	compartment_id = var.compartment_id
	gi_version = var.system_version_minor_version_gi_version
	major_version = var.system_version_minor_version_major_version

	#Optional
	is_latest = var.system_version_minor_version_is_latest
	resource_id = oci_cloud_guard_resource.test_resource.id
	shape = var.system_version_minor_version_shape
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `gi_version` - (Required) Specifies gi version query parameter.
* `is_latest` - (Optional) If provided, return highest versions from each major version family.
* `major_version` - (Required) The System major version.
* `resource_id` - (Optional) If provided, filters the results for the specified resource Id.
* `shape` - (Optional) If provided, filters the results for the given shape.


## Attributes Reference

The following attributes are exported:

* `system_version_minor_version_collection` - The list of system_version_minor_version_collection.

### SystemVersionMinorVersion Reference

The following attributes are exported:

* `items` - List of System minor versions.
	* `version` - A valid system minor version.

