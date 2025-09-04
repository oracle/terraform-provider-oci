---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_blob_mount"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_azure_blob_mount"
description: |-
  Provides the Oracle Db Azure Blob Mount resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_azure_blob_mount
This resource provides the Oracle Db Azure Blob Mount resource in Oracle Cloud Infrastructure Dbmulticloud service.

Creates Oracle DB Azure Blob Mount resource.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_azure_blob_mount" "test_oracle_db_azure_blob_mount" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.oracle_db_azure_blob_mount_display_name
	oracle_db_azure_blob_container_id = oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) TThe [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Blob Mount resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Oracle DB Azure Blob Mount resource name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `oracle_db_azure_blob_container_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Blob Container resource that contains Oracle DB Azure Blob Mount resource.
* `oracle_db_azure_connector_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector resource that contains Oracle DB Azure Blob Mount resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Azure Blob Mount
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Azure Blob Mount
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Azure Blob Mount


## Import

OracleDbAzureBlobMounts can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount "id"
```

