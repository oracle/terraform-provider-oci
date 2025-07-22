---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_gi_versions"
sidebar_current: "docs-oci-datasource-database-gi_versions"
description: |-
  Provides the list of Gi Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_gi_versions
This data source provides the list of Gi Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported GI versions.

## Example Usage

```hcl
data "oci_database_gi_versions" "test_gi_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.gi_version_availability_domain
	resource_id = oci_cloud_guard_resource.test_resource.id
	shape = var.gi_version_shape
	shape_attribute = var.gi_version_shape_attribute
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The target availability domain. Only passed if the limit is AD-specific.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `resource_id` - (Optional) If provided, filters the results for the specified resource Id.
* `shape` - (Optional) If provided, filters the results for the given shape.
* `shape_attribute` - (Optional) If provided and applicable, return the results based on the shapeAttribute provided


## Attributes Reference

The following attributes are exported:

* `gi_versions` - The list of gi_versions.

### GiVersion Reference

The following attributes are exported:

* `version` - A valid Oracle Grid Infrastructure (GI) software version.

