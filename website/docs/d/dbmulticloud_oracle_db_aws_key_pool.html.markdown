---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_key_pool"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_aws_key_pool"
description: |-
  Provides details about a specific Oracle Db Aws Key Pool in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_aws_key_pool
This data source provides details about a specific Oracle Db Aws Key Pool resource in Oracle Cloud Infrastructure Dbmulticloud service.

Retrieves detailed information about a Oracle AWS Key Pool resource by specifying its unique resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_aws_key_pool" "test_oracle_db_aws_key_pool" {
	#Required
	oracle_db_aws_key_pool_id = oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_aws_key_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Key Pool resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The Compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that contains this Oracle DB AWS Pool resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the Oracle DB AWS key pool resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB AWS Key Pool resource.
* `is_multi_region_key_pool` - Indicates whether this key pool accepts only multi-region AWS KMS keys. When false, the key pool accepts only single-region AWS KMS keys.
* `is_reuse_allowed` - Indicates whether AWS KMS keys in this key pool can be reused.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `replication_metadata` - Replication metadata, it has information about replication and target region. 
	* `replication_details` - List of all regions and their respective resource ID.
		* `replication_state` - The lifecycle state of the Oracle DB Azure Vault resource.
		* `target_region` - The target region, where resource is replicated.
		* `time_created` - Time when the Replication was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z' 
		* `time_updated` - Time when the Replication was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z' 
* `state` - The lifecycle state of the Oracle DB AWS Key Pool resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the DB AWS Key Pool resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'. 
* `time_updated` - Time when the DB AWS Key Pool resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'. 

