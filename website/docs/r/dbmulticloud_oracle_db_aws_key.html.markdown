---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_key"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_aws_key"
description: |-
  Provides the Oracle Db Aws Key resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_aws_key
This resource provides the Oracle Db Aws Key resource in Oracle Cloud Infrastructure Dbmulticloud service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/latest/OracleDbAwsKey

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dbmulticloud

Create DB AWS Key resource.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_aws_key" "test_oracle_db_aws_key" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.oracle_db_aws_key_display_name
	oracle_db_connector_id = oci_dbmulticloud_oracle_db_connector.test_oracle_db_connector.id

	#Optional
	aws_account_id = oci_dbmulticloud_aws_account.test_aws_account.id
	aws_key_arn = var.oracle_db_aws_key_aws_key_arn
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_aws_key_enabled = var.oracle_db_aws_key_is_aws_key_enabled
	location = var.oracle_db_aws_key_location
	properties = var.oracle_db_aws_key_properties
	type = var.oracle_db_aws_key_type
}
```

## Argument Reference

The following arguments are supported:

* `aws_account_id` - (Optional) AWS Account ID.
* `aws_key_arn` - (Optional) Amazon resource name of AWS Key.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB AWS Key resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Oracle DB AWS Key resource name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_aws_key_enabled` - (Optional) The Oracle AWS Key resource is enabled or disabled at AWS.
* `location` - (Optional) AWS Key resource Location.
* `oracle_db_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Connector resource.
* `properties` - (Optional) AWS Key resource's properties.
* `type` - (Optional) AWS Key resource type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Aws Key
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Aws Key
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Aws Key


## Import

OracleDbAwsKeys can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key "id"
```

