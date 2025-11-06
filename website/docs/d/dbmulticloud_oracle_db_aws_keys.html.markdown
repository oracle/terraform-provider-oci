---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_keys"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_aws_keys"
description: |-
  Provides the list of Oracle Db Aws Keys in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_aws_keys
This data source provides the list of Oracle Db Aws Keys in Oracle Cloud Infrastructure Dbmulticloud service.

Lists all DB AWS Keys based on the specified filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_aws_keys" "test_oracle_db_aws_keys" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_aws_key_display_name
	oracle_db_aws_connector_id = oci_dbmulticloud_oracle_db_aws_connector.test_oracle_db_aws_connector.id
	oracle_db_aws_key_id = oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key.id
	state = var.oracle_db_aws_key_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB AWS Key Resource that match the given display name.
* `oracle_db_aws_connector_id` - (Optional) A filter to return Oracle DB AWS Identity Connector resources that match the specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Identity Connector resource.
* `oracle_db_aws_key_id` - (Optional) A filter to return Oracle DB AWS Identity Connector Resource that match the given OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Key resource.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_aws_key_summary_collection` - The list of oracle_db_aws_key_summary_collection.

### OracleDbAwsKey Reference

The following attributes are exported:

* `aws_account_id` - AWS Account ID.
* `aws_key_arn` - Amazon resource name of AWS Key.
* `compartment_id` - The Compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that contains this Oracle DB AWS Key resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB AWS Key Ring resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB AWS Key Ring resource.
* `is_aws_key_enabled` - The Oracle AWS Key resource is enabled or disabled at AWS.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `location` - AWS Key resource location.
* `oracle_db_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Connector resource.
* `properties` - AWS Key resource's properties.
* `state` - The lifecycle state of the Oracle DB AWS Key resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the DB AWS Key resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'. 
* `time_updated` - Time when the DB AWS Key resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'. 
* `type` - Key resource type.

