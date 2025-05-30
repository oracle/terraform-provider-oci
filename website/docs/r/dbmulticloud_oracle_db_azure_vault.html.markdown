---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_vault"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_azure_vault"
description: |-
  Provides the Oracle Db Azure Vault resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_azure_vault
This resource provides the Oracle Db Azure Vault resource in Oracle Cloud Infrastructure Dbmulticloud service.

Create DB Azure Vaults based on the provided information, this will fetch Keys related to Azure Vaults.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_azure_vault" "test_oracle_db_azure_vault" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.oracle_db_azure_vault_display_name
	oracle_db_connector_id = oci_dbmulticloud_oracle_db_connector.test_oracle_db_connector.id

	#Optional
	azure_vault_id = oci_kms_vault.test_vault.id
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	location = var.oracle_db_azure_vault_location
	oracle_db_azure_resource_group = var.oracle_db_azure_vault_oracle_db_azure_resource_group
	properties = var.oracle_db_azure_vault_properties
	type = var.oracle_db_azure_vault_type
}
```

## Argument Reference

The following arguments are supported:

* `azure_vault_id` - (Optional) (Updatable) Azure Vault Id.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains DB Azure Vault Resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Display name of DB Azure Vault.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `location` - (Optional) (Updatable) Vault Resource Location.
* `oracle_db_azure_resource_group` - (Optional) (Updatable) Display name of Azure Resource Group.
* `oracle_db_connector_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB Connector Resource.
* `properties` - (Optional) (Updatable) Resource's properties.
* `type` - (Optional) (Updatable) Vault Resource Type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Azure Vault
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Azure Vault
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Azure Vault


## Import

OracleDbAzureVaults can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault "id"
```

