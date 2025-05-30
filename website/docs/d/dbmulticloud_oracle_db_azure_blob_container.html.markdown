---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_blob_container"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_blob_container"
description: |-
  Provides details about a specific Oracle Db Azure Blob Container in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_blob_container
This data source provides details about a specific Oracle Db Azure Blob Container resource in Oracle Cloud Infrastructure Dbmulticloud service.

Get Oracle DB Azure Blob Container Details form a particular Container Resource ID.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_blob_container" "test_oracle_db_azure_blob_container" {
	#Required
	oracle_db_azure_blob_container_id = oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_azure_blob_container_id` - (Required) The ID of the Oracle DB Azure Blob Container Resource.


## Attributes Reference

The following attributes are exported:

* `azure_storage_account_name` - Azure Storage Account Name.
* `azure_storage_container_name` - Azure Storage Container Name.
* `compartment_id` - The ID of the compartment that contains Oracle DB Azure Blob Container Resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Display name of Oracle DB Azure Blob Container.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The ID of the compartment that contains Oracle DB Azure Blob Container Resource.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Blob Container Resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `private_endpoint_dns_alias` - Private endpoint DNS Alias.
* `private_endpoint_ip_address` - Private endpoint IP.
* `state` - The current lifecycle state of the Oracle DB Azure Blob Container Resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Blob Container was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Blob Container was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

