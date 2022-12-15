---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_versions"
sidebar_current: "docs-oci-datasource-database-autonomous_container_database_versions"
description: |-
  Provides the list of Autonomous Container Database Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_database_versions
This data source provides the list of Autonomous Container Database Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Autonomous Container Database versions.

## Example Usage

```hcl
data "oci_database_autonomous_container_database_versions" "test_autonomous_container_database_versions" {
	#Required
	compartment_id = var.compartment_id
	service_component = var.autonomous_container_database_version_service_component
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `service_component` - (Required) The service component to use, either ADBD or EXACC.


## Attributes Reference

The following attributes are exported:

* `autonomous_container_database_versions` - The list of autonomous_container_database_versions.

### AutonomousContainerDatabaseVersion Reference

The following attributes are exported:

* `details` - A URL that points to a detailed description of the Autonomous Container Database version.
* `supported_apps` - The list of applications supported for the given version.
	* `end_of_support` - The Autonomous Container Database version end of support date.
	* `is_certified` - Indicates if the image is certified.
	* `release_date` - The Autonomous Container Database version release date.
	* `supported_app_name` - The name of the supported application.
* `version` - A valid Oracle Database version for provisioning an Autonomous Container Database.

