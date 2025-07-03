---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_vault"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_vault"
description: |-
  Provides details about a specific Oracle Db Azure Vault in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_vault
This data source provides details about a specific Oracle Db Azure Vault resource in Oracle Cloud Infrastructure Dbmulticloud service.

Get Oracle DB Azure Vault Details form a particular Container Resource ID.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_vault" "test_oracle_db_azure_vault" {
	#Required
	oracle_db_azure_vault_id = oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_azure_vault_id` - (Required) The ID of the Oracle DB Azure Vault Resource.


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

