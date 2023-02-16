---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system"
sidebar_current: "docs-oci-datasource-database_management-external_db_system"
description: |-
  Provides details about a specific External Db System in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_system
This data source provides details about a specific External Db System resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the external DB system specified by `externalDbSystemId`.


## Example Usage

```hcl
data "oci_database_management_external_db_system" "test_external_db_system" {
	#Required
	external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `external_db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_management_config` - The details required to enable Database Management for an external DB system.
	* `license_model` - The Oracle license model that applies to the external database. 
* `db_system_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `discovery_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `home_directory` - The Oracle Grid home directory in case of cluster-based DB system and Oracle home directory in case of single instance-based DB system. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
* `is_cluster` - Indicates whether the DB system is a cluster DB system or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external DB system resource.
* `time_created` - The date and time the external DB system was created.
* `time_updated` - The date and time the external DB system was last updated.

