---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_gcp_keys"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_gcp_keys"
description: |-
  Provides the list of Oracle Db Gcp Keys in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_gcp_keys
This data source provides the list of Oracle Db Gcp Keys in Oracle Cloud Infrastructure Dbmulticloud service.

Lists all Oracle DB Google Cloud Keys based on the specified filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_gcp_keys" "test_oracle_db_gcp_keys" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_gcp_key_display_name
	oracle_db_gcp_key_id = oci_dbmulticloud_oracle_db_gcp_key.test_oracle_db_gcp_key.id
	oracle_db_gcp_key_ring_id = oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id
	state = var.oracle_db_gcp_key_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB Google Cloud Key resources that match the specified display name.
* `oracle_db_gcp_key_id` - (Optional) A filter to return Oracle DB Google Cloud Key resources.
* `oracle_db_gcp_key_ring_id` - (Optional) A filter to return Oracle DB GCP Key Rings.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_gcp_key_summary_collection` - The list of oracle_db_gcp_key_summary_collection.

### OracleDbGcpKey Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Google Key resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Display name of Oracle DB Google Key resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gcp_key_id` - TThe Google Cloud Key ID and Key URL associated with the Google Key under the specified Key Ring resource.
* `gcp_key_properties` - Gcp Key properties
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Google Key resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `oracle_db_gcp_key_ring_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Google Cloud Key-Ring resource.
* `resource_type` - Key resource type.
* `state` - The current lifecycle state of the Oracle DB Google Key resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Google Key resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Google Key resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

