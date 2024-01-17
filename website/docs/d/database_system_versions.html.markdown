---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_system_versions"
sidebar_current: "docs-oci-datasource-database-system_versions"
description: |-
  Provides the list of System Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_system_versions
This data source provides the list of System Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Exadata system versions for a given shape and GI version.

## Example Usage

```hcl
data "oci_database_system_versions" "test_system_versions" {
	#Required
	compartment_id = var.compartment_id
	gi_version = var.system_version_gi_version
	shape = var.system_version_shape
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `gi_version` - (Required) Specifies gi version query parameter.
* `shape` - (Required) Specifies shape query parameter.


## Attributes Reference

The following attributes are exported:

* `system_version_collection` - The list of system_version_collection.

### SystemVersion Reference

The following attributes are exported:

* `items` - List of System versions.
	* `gi_version` - Grid Infrastructure version.
	* `shape` - Exadata shape.
	* `system_versions` - Compatible Exadata system versions for a given shape and GI version.

