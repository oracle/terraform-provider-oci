---
subcategory: "Apiaccesscontrol"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apiaccesscontrol_api_metadatas"
sidebar_current: "docs-oci-datasource-apiaccesscontrol-api_metadatas"
description: |-
  Provides the list of Api Metadatas in Oracle Cloud Infrastructure Apiaccesscontrol service
---

# Data Source: oci_apiaccesscontrol_api_metadatas
This data source provides the list of Api Metadatas in Oracle Cloud Infrastructure Apiaccesscontrol service.

Gets a list of ApiMetadata.


## Example Usage

```hcl
data "oci_apiaccesscontrol_api_metadatas" "test_api_metadatas" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.api_metadata_display_name
	resource_type = var.api_metadata_resource_type
	state = var.api_metadata_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `resource_type` - (Optional) A filter to return only lists of resources that match the entire given service type.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `api_metadata_collection` - The list of api_metadata_collection.

### ApiMetadata Reference

The following attributes are exported:

* `api_name` - The name of the api to execute the api request.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The operation Name of the api. The name must be unique.
* `entity_type` - ResourceType to which the apiMetadata belongs to.
* `fields` - List of the fields that is use while calling post or put for the data.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ApiDetail.
* `lifecycle_details` - A message that describes the current state of the ApiMetadata in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `path` - rest path of the api.
* `service_name` - The service Name to which the api belongs to.
* `state` - The current state of the ApiMetadata.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the PrivilegedApiControl was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_deleted` - The date and time the PrivilegedApiControl was marked for delete, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the PrivilegedApiControl was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

