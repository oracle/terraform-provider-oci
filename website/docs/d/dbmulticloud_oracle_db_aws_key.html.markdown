---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_aws_key"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_aws_key"
description: |-
  Provides details about a specific Oracle Db Aws Key in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_aws_key
This data source provides details about a specific Oracle Db Aws Key resource in Oracle Cloud Infrastructure Dbmulticloud service.

Retrieves detailed information about a Oracle AWS Key resource by specifying its unique resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_aws_key" "test_oracle_db_aws_key" {
	#Required
	oracle_db_aws_key_id = oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_aws_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Key resource.


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

