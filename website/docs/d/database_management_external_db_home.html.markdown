---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_home"
sidebar_current: "docs-oci-datasource-database_management-external_db_home"
description: |-
  Provides details about a specific External Db Home in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_home
This data source provides details about a specific External Db Home resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the external DB home specified by `externalDbHomeId`.


## Example Usage

```hcl
data "oci_database_management_external_db_home" "test_external_db_home" {
	#Required
	external_db_home_id = oci_database_management_external_db_home.test_external_db_home.id
}
```

## Argument Reference

The following arguments are supported:

* `external_db_home_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database home.


## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the DB home defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external DB home.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the external DB home. The name does not have to be unique.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the DB home is a part of.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_directory` - The location of the DB home.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB home.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external DB home.
* `time_created` - The date and time the external DB home was created.
* `time_updated` - The date and time the external DB home was last updated.

