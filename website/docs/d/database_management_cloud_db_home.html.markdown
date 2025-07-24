---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_db_home"
sidebar_current: "docs-oci-datasource-database_management-cloud_db_home"
description: |-
  Provides details about a specific Cloud Db Home in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_db_home
This data source provides details about a specific Cloud Db Home resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the cloud DB home specified by `cloudDbHomeId`.


## Example Usage

```hcl
data "oci_database_management_cloud_db_home" "test_cloud_db_home" {
	#Required
	cloud_db_home_id = oci_database_management_cloud_db_home.test_cloud_db_home.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_home_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud database home.


## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the DB home defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the DB home is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud DB home.
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB home in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cloud DB home. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_directory` - The location of the DB home.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB home.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the cloud DB home.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud DB home was created.
* `time_updated` - The date and time the cloud DB home was last updated.

