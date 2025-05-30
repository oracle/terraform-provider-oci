---
subcategory: "Apiaccesscontrol"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apiaccesscontrol_api_metadata"
sidebar_current: "docs-oci-datasource-apiaccesscontrol-api_metadata"
description: |-
  Provides details about a specific Api Metadata in Oracle Cloud Infrastructure Apiaccesscontrol service
---

# Data Source: oci_apiaccesscontrol_api_metadata
This data source provides details about a specific Api Metadata resource in Oracle Cloud Infrastructure Apiaccesscontrol service.

Gets information about a ApiMetadata.

## Example Usage

```hcl
data "oci_apiaccesscontrol_api_metadata" "test_api_metadata" {
	#Required
	api_metadata_id = oci_apiaccesscontrol_api_metadata.test_api_metadata.id
}
```

## Argument Reference

The following arguments are supported:

* `api_metadata_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiControl.


## Attributes Reference

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

