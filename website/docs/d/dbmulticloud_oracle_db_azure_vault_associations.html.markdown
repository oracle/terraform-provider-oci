---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_vault_associations"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_vault_associations"
description: |-
  Provides the list of Oracle Db Azure Vault Associations in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_vault_associations
This data source provides the list of Oracle Db Azure Vault Associations in Oracle Cloud Infrastructure Dbmulticloud service.

Lists all Oracle DB Azure Associations based on the specified filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_vault_associations" "test_oracle_db_azure_vault_associations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_azure_vault_association_display_name
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
	oracle_db_azure_vault_association_id = oci_dbmulticloud_oracle_db_azure_vault_association.test_oracle_db_azure_vault_association.id
	oracle_db_azure_vault_id = oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id
	state = var.oracle_db_azure_vault_association_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB Azure Vault Association resources that match the specified display name.
* `oracle_db_azure_connector_id` - (Optional) A filter to return Oracle DB Azure Azure Identity Connector resources.
* `oracle_db_azure_vault_association_id` - (Optional) A filter to return Oracle DB Azure Vault resources that match the specified OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault Association resource.
* `oracle_db_azure_vault_id` - (Optional) A filter to return Oracle DB Azure Vault resources that match the specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault resource.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_azure_vault_association_summary_collection` - The list of oracle_db_azure_vault_association_summary_collection.

### OracleDbAzureVaultAssociation Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Vault Association resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB Azure Vault Association resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault Association resource.
* `is_resource_accessible` - The Associated resource is accessible or not.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Vault Association resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `oracle_db_azure_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector that contains Oracle DB Azure Vault Association resource.
* `oracle_db_azure_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault that contains Oracle DB Azure Vault Association resource.
* `state` - The current lifecycle state of the Oracle DB Azure Vault Association resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Vault Association resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Vault Association resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

