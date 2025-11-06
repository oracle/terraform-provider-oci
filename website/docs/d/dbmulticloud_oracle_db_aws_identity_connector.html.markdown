---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_identity_connector"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_aws_identity_connector"
description: |-
  Provides details about a specific Oracle Db Aws Identity Connector in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_aws_identity_connector
This data source provides details about a specific Oracle Db Aws Identity Connector resource in Oracle Cloud Infrastructure Dbmulticloud service.

Retrieves detailed information about a Oracle DB AWS Identity Connector resource by specifying its unique resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_aws_identity_connector" "test_oracle_db_aws_identity_connector" {
	#Required
	oracle_db_aws_identity_connector_id = oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_aws_identity_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)  of the Oracle DB AWS Identity Connector resource.


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

