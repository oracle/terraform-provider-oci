---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_identity_connector"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_aws_identity_connector"
description: |-
  Provides the Oracle Db Aws Identity Connector resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_aws_identity_connector
This resource provides the Oracle Db Aws Identity Connector resource in Oracle Cloud Infrastructure Dbmulticloud service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/latest/OracleDbAwsIdentityConnector

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dbmulticloud

Creates Oracle DB AWS Identity Connector resource.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_aws_identity_connector" "test_oracle_db_aws_identity_connector" {
	#Required
	aws_location = var.oracle_db_aws_identity_connector_aws_location
	compartment_id = var.compartment_id
	display_name = var.oracle_db_aws_identity_connector_display_name
	issuer_url = var.oracle_db_aws_identity_connector_issuer_url
	oidc_scope = var.oracle_db_aws_identity_connector_oidc_scope
	resource_id = oci_cloud_guard_resource.test_resource.id
	service_role_details {
		#Required
		role_arn = var.oracle_db_aws_identity_connector_service_role_details_role_arn
		service_private_endpoint = var.oracle_db_aws_identity_connector_service_role_details_service_private_endpoint
		service_type = var.oracle_db_aws_identity_connector_service_role_details_service_type
	}

	#Optional
	aws_account_id = oci_dbmulticloud_aws_account.test_aws_account.id
	aws_sts_private_endpoint = var.oracle_db_aws_identity_connector_aws_sts_private_endpoint
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `aws_account_id` - (Optional) (Updatable) AWS Account ID.
* `aws_location` - (Required) (Updatable) AWS resource location.
* `aws_sts_private_endpoint` - (Optional) (Updatable) Private endpoint of AWS Security Token Service.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB AWS Identity Connector resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Oracle DB Google AWS Identity Connector resource name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `issuer_url` - (Required) (Updatable) OIDC token issuer Url.
* `oidc_scope` - (Required) (Updatable) Oracle Cloud Infrastructure IAM Domain scope for issuer URL.
* `resource_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the AWS VM Cluster resource.
* `service_role_details` - (Required) (Updatable) Service role details and respective Amazon resource nam of Role. 
	* `role_arn` - (Required) (Updatable) Amazon resource name AWSof the IAM role.
	* `service_private_endpoint` - (Required) (Updatable) Private endpoint of the AWS service.
	* `service_type` - (Required) (Updatable) Type of service.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `aws_account_id` - AWS Account ID.
* `aws_location` - AWS resource location.
* `aws_sts_private_endpoint` - Private endpoint of AWS Security Token Service.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB AWS Identity Connector resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB AWS Identity Connector resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Identity Connector resource.
* `issuer_url` - OIDC token issuer Url.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `oidc_scope` - Oracle Cloud Infrastructure IAM Domain scope for issuer URL.
* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the AWS VM Cluster resource.
* `service_role_details` - Service role details and respective Amazon resource nam of Role. 
	* `assume_role_status` - Assume role  status.
	* `aws_nodes` - List of all VMs where AWS Identity Connector is configured for Oracle DB Cloud VM Cluster.
		* `host_id` - AWS host ID.
		* `host_name` - AWS Host name or Identity Connector name.
		* `status` - The current status of the AWS Identity Connector resource.
		* `time_last_checked` - Time when the AWS Identity Connector's status was checked [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'. 
	* `role_arn` - Amazon resource name AWSof the IAM role.
	* `service_private_endpoint` - Private endpoint of the AWS service.
	* `service_type` - Type of service.
* `state` - The current lifecycle state of the AWS Identity Connector resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB AWS Identity Connector resource was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'. 
* `time_updated` - Time when the Oracle DB AWS Identity Connector resource was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Aws Identity Connector
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Aws Identity Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Aws Identity Connector


## Import

OracleDbAwsIdentityConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector "id"
```

