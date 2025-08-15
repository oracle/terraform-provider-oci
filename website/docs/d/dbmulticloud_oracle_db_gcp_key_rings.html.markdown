---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_gcp_key_rings"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_gcp_key_rings"
description: |-
  Provides the list of Oracle Db Gcp Key Rings in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_gcp_key_rings
This data source provides the list of Oracle Db Gcp Key Rings in Oracle Cloud Infrastructure Dbmulticloud service.

Lists the all DB GCP Key Rings based on filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_gcp_key_rings" "test_oracle_db_gcp_key_rings" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_gcp_key_ring_display_name
	oracle_db_gcp_connector_id = oci_dbmulticloud_oracle_db_gcp_connector.test_oracle_db_gcp_connector.id
	oracle_db_gcp_key_ring_id = oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id
	state = var.oracle_db_gcp_key_ring_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB GCP Key Ring resources that match the specified display name.
* `oracle_db_gcp_connector_id` - (Optional) A filter to return Oracle DB GCP Identity Connector resources that match the specified resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `oracle_db_gcp_key_ring_id` - (Optional) A filter to return Oracle DB GCP Key Rings.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_gcp_key_ring_summary_collection` - The list of oracle_db_gcp_key_ring_summary_collection.

### OracleDbGcpKeyRing Reference

The following attributes are exported:

* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the Oracle DB GCP Key Ring resource resides.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Display name of DB GCP Key Ring resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gcp_key_ring_id` - GCP Key Ring ID.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB GCP Key Ring resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `location` - Location of the GCP Key Ring resource.
* `oracle_db_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Oracle DB GCP Identity Connector resource resides.
* `properties` - Oracle DB GCP Key Ring resource's properties.
* `state` - The lifecycle state of the Oracle DB GCP Key Ring resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the DB GCP Key Ring resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z' 
* `time_updated` - Time when the DB GCP Key Ring resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z' 
* `type` - Oracle DB GCP Key Ring resource Type.

