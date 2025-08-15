---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_gcp_key_ring"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_gcp_key_ring"
description: |-
  Provides details about a specific Oracle Db Gcp Key Ring in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_gcp_key_ring
This data source provides details about a specific Oracle Db Gcp Key Ring resource in Oracle Cloud Infrastructure Dbmulticloud service.

Retrieves the Oracle GCP Key Ring details using a specific Container resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_gcp_key_ring" "test_oracle_db_gcp_key_ring" {
	#Required
	oracle_db_gcp_key_ring_id = oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_gcp_key_ring_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB GCP Key-Ring resource.


## Attributes Reference

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

