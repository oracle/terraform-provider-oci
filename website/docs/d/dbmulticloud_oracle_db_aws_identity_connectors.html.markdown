---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_identity_connectors"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_aws_identity_connectors"
description: |-
  Provides the list of Oracle Db Aws Identity Connectors in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_aws_identity_connectors
This data source provides the list of Oracle Db Aws Identity Connectors in Oracle Cloud Infrastructure Dbmulticloud service.

Lists all Oracle DB AWS Identity Connectors based on the specified filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_aws_identity_connectors" "test_oracle_db_aws_identity_connectors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_aws_identity_connector_display_name
	resource_id = oci_cloud_guard_resource.test_resource.id
	state = var.oracle_db_aws_identity_connector_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB AWS Identity Connector Resource that match the given display name.
* `resource_id` - (Optional) A filter to return Oracle DB Identity Connector resource that match the given resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_aws_identity_connector_summary_collection` - The list of oracle_db_aws_identity_connector_summary_collection.

### OracleDbAwsIdentityConnector Reference

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

