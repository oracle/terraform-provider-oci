---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_multi_cloud_resource_discovery"
sidebar_current: "docs-oci-resource-dbmulticloud-multi_cloud_resource_discovery"
description: |-
  Provides the Multi Cloud Resource Discovery resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_multi_cloud_resource_discovery
This resource provides the Multi Cloud Resource Discovery resource in Oracle Cloud Infrastructure Dbmulticloud service.

Discover Azure Vaults and Keys based on the provided information.


## Example Usage

```hcl
resource "oci_dbmulticloud_multi_cloud_resource_discovery" "test_multi_cloud_resource_discovery" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.multi_cloud_resource_discovery_display_name
	oracle_db_connector_id = oci_dbmulticloud_oracle_db_connector.test_oracle_db_connector.id
	resource_type = var.multi_cloud_resource_discovery_resource_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Discovered Resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Display name of Discovered Resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `oracle_db_connector_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Oracle DB Connector.
* `resource_type` - (Required) (Updatable) Resource Type to discover.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Multi Cloud Discovery Resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Display name of Multi Cloud Discovery Resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multi Cloud Discovery Resource.
* `last_modification` - Description of the latest modification of the Multi Cloud Discovery Resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `oracle_db_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Connector Resource.
* `resource_type` - Resource Type to discover.
* `resources` - List of All Discovered resources.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Discovered Resource.
	* `location` - Discovered Resource Location.
	* `name` - Discovered Resource Name.
	* `properties` - Discovered Resource's properties.
	* `resource_group` - Discovered Resource Group Name.
	* `type` - Discovered Resource Type.
* `state` - The current lifecycle state of the discovered resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Multi Cloud Discovery Resource was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Multi Cloud Discovery Resource was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Multi Cloud Resource Discovery
	* `update` - (Defaults to 20 minutes), when updating the Multi Cloud Resource Discovery
	* `delete` - (Defaults to 20 minutes), when destroying the Multi Cloud Resource Discovery


## Import

MultiCloudResourceDiscoveries can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery "id"
```

