---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_key"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_key"
description: |-
  Provides details about a specific Oracle Db Azure Key in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_key
This data source provides details about a specific Oracle Db Azure Key resource in Oracle Cloud Infrastructure Dbmulticloud service.

Retrieves detailed information about a Oracle DB Azure Key resource by specifying its unique resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_key" "test_oracle_db_azure_key" {
	#Required
	oracle_db_azure_key_id = oci_dbmulticloud_oracle_db_azure_key.test_oracle_db_azure_key.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_azure_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault Key resource.


## Attributes Reference

The following attributes are exported:

* `azure_key_id` - The ID of the Azure Key resource.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Vault Key resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB Azure Vault Key resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault Key resource.
* `key_properties` - Key properties
* `last_modification` - Description of the latest modification of the Oracle DB Azure Vault Key resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `oracle_db_azure_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault resource.
* `resource_type` - Key Resource type.
* `state` - The current lifecycle state of the Oracle DB Azure Vault Key resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Vault Key resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Vault Key resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

