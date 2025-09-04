---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_gcp_key_ring"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_gcp_key_ring"
description: |-
  Provides the Oracle Db Gcp Key Ring resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_gcp_key_ring
This resource provides the Oracle Db Gcp Key Ring resource in Oracle Cloud Infrastructure Dbmulticloud service.

Creates DB GCP Key Rings based on the provided information and retrieves the associated keys.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_gcp_key_ring" "test_oracle_db_gcp_key_ring" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.oracle_db_gcp_key_ring_display_name
	oracle_db_connector_id = oci_dbmulticloud_oracle_db_connector.test_oracle_db_connector.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	gcp_key_ring_id = oci_dbmulticloud_gcp_key_ring.test_gcp_key_ring.id
	location = var.oracle_db_gcp_key_ring_location
	properties = var.oracle_db_gcp_key_ring_properties
	type = var.oracle_db_gcp_key_ring_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the Oracle DB GCP Key Ring resource resides.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Display name of DB GCP Key Ring resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gcp_key_ring_id` - (Optional) GCP Key Ring ID.
* `location` - (Optional) Location of the GCP Key Ring resource.
* `oracle_db_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Oracle DB GCP Identity Connector resource resides.
* `properties` - (Optional) Oracle DB GCP Key Ring resource's properties.
* `type` - (Optional) Oracle DB GCP Key Ring resource Type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Gcp Key Ring
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Gcp Key Ring
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Gcp Key Ring


## Import

OracleDbGcpKeyRings can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring "id"
```

