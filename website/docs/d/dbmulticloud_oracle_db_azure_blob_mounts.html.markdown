---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_blob_mounts"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_blob_mounts"
description: |-
  Provides the list of Oracle Db Azure Blob Mounts in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_blob_mounts
This data source provides the list of Oracle Db Azure Blob Mounts in Oracle Cloud Infrastructure Dbmulticloud service.

Lists all Oracle DB Azure Blob Mount resources based on the specified filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_blob_mounts" "test_oracle_db_azure_blob_mounts" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_azure_blob_mount_display_name
	oracle_db_azure_blob_container_id = oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id
	oracle_db_azure_blob_mount_id = oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount.id
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
	state = var.oracle_db_azure_blob_mount_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB Azure Blob Mount resources that match the specified display name.
* `oracle_db_azure_blob_container_id` - (Optional) A filter to return Oracle DB Azure Blob Container resource.
* `oracle_db_azure_blob_mount_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Blob Mount resource.
* `oracle_db_azure_connector_id` - (Optional) A filter to return Oracle DB Azure Azure Identity Connector resources.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_azure_blob_mount_summary_collection` - The list of oracle_db_azure_blob_mount_summary_collection.

### OracleDbAzureBlobMount Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Blob Mount resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB Azure Blob Mount resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Blob Mount resource.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Blob Mount resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `mount_path` - Oracle DB Azure Blob Mount path.
* `oracle_db_azure_blob_container_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Blob Container resource that contains Oracle DB Azure Blob Mount resource.
* `oracle_db_azure_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector resource that contains Oracle DB Azure Blob Mount resource.
* `state` - The current lifecycle state of the Oracle DB Azure Blob Mount resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Blob Mount was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Blob Mount was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

