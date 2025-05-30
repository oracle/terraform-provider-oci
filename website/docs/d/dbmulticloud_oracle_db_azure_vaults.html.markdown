---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_vaults"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_vaults"
description: |-
  Provides the list of Oracle Db Azure Vaults in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_vaults
This data source provides the list of Oracle Db Azure Vaults in Oracle Cloud Infrastructure Dbmulticloud service.

Lists the all DB Azure Vaults based on filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_vaults" "test_oracle_db_azure_vaults" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_azure_vault_display_name
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
	oracle_db_azure_resource_group = var.oracle_db_azure_vault_oracle_db_azure_resource_group
	oracle_db_azure_vault_id = oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id
	state = var.oracle_db_azure_vault_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Azure Vaults.
* `oracle_db_azure_connector_id` - (Optional) A filter to return Oracle DB Azure Blob Mount Resources.
* `oracle_db_azure_resource_group` - (Optional) A filter to return Azure Vaults.
* `oracle_db_azure_vault_id` - (Optional) A filter to return Oracle DB Azure Vault Resources.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_azure_vault_summary_collection` - The list of oracle_db_azure_vault_summary_collection.

### OracleDbAzureVault Reference

The following attributes are exported:

* `azure_vault_id` - Azure Vault Id.
* `compartment_id` - The Compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that has this DB Azure Vault Resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Display name of DB Azure Vault.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB Azure Vault Resource.
* `last_modification` - Description of the latest modification of the DB Azure Vault Resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `location` - Vault Resource Location.
* `oracle_db_azure_resource_group` - Display name of Azure Resource Group.
* `oracle_db_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB Connector Resource.
* `properties` - Resource's properties.
* `state` - The lifecycle state of the DB Azure Vault Resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the DB Azure Vault was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z' 
* `time_updated` - Time when the DB Azure Vault was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z' 
* `type` - Vault Resource Type.

