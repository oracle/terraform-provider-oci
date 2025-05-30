---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_multi_cloud_resource_discoveries"
sidebar_current: "docs-oci-datasource-dbmulticloud-multi_cloud_resource_discoveries"
description: |-
  Provides the list of Multi Cloud Resource Discoveries in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_multi_cloud_resource_discoveries
This data source provides the list of Multi Cloud Resource Discoveries in Oracle Cloud Infrastructure Dbmulticloud service.

Lists the all Multi Cloud Resource Discovery based on filters.


## Example Usage

```hcl
data "oci_dbmulticloud_multi_cloud_resource_discoveries" "test_multi_cloud_resource_discoveries" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.multi_cloud_resource_discovery_display_name
	multi_cloud_resource_discovery_id = oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery.id
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
	resource_type = var.multi_cloud_resource_discovery_resource_type
	state = var.multi_cloud_resource_discovery_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) Display Name of the Multi Cloud Discovery Resource.
* `multi_cloud_resource_discovery_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multi Cloud Discovery Resource.
* `oracle_db_azure_connector_id` - (Optional) A filter to return Oracle DB Azure Blob Mount Resources.
* `resource_type` - (Optional) The type of Multi Cloud Resource.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `multi_cloud_resource_discovery_summary_collection` - The list of multi_cloud_resource_discovery_summary_collection.

### MultiCloudResourceDiscovery Reference

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

