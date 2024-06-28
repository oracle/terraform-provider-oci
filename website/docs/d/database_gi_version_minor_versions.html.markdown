---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_gi_version_minor_versions"
sidebar_current: "docs-oci-datasource-database-gi_version_minor_versions"
description: |-
  Provides the list of Gi Version Minor Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_gi_version_minor_versions
This data source provides the list of Gi Version Minor Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Oracle Grid Infrastructure minor versions for the given major version and shape family.

## Example Usage

```hcl
data "oci_database_gi_version_minor_versions" "test_gi_version_minor_versions" {
	#Required
	version = var.gi_version_minor_version_version

	#Optional
	availability_domain = var.gi_version_minor_version_availability_domain
	compartment_id = var.compartment_id
	is_gi_version_for_provisioning = var.gi_version_minor_version_is_gi_version_for_provisioning
	shape = var.gi_version_minor_version_shape
	shape_family = var.gi_version_minor_version_shape_family
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The target availability domain. Only passed if the limit is AD-specific.
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `is_gi_version_for_provisioning` - (Optional) If true, returns the Grid Infrastructure versions that can be used for provisioning a cluster
* `shape` - (Optional) If provided, filters the results for the given shape.
* `shape_family` - (Optional) If provided, filters the results to the set of database versions which are supported for the given shape family.
* `version` - (Required) The Oracle Grid Infrastructure major version.


## Attributes Reference

The following attributes are exported:

* `gi_minor_versions` - The list of gi_minor_versions.

### GiVersionMinorVersion Reference

The following attributes are exported:

* `grid_image_id` - Grid Infrastructure Image Id
* `version` - A valid Oracle Grid Infrastructure (GI) software version.

