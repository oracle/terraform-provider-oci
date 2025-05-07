---
subcategory: "Api Platform"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_api_platform_api_platform_instances"
sidebar_current: "docs-oci-datasource-api_platform-api_platform_instances"
description: |-
  Provides the list of Api Platform Instances in Oracle Cloud Infrastructure Api Platform service
---

# Data Source: oci_api_platform_api_platform_instances
This data source provides the list of Api Platform Instances in Oracle Cloud Infrastructure Api Platform service.

Gets a list of API Platform Instances


## Example Usage

```hcl
data "oci_api_platform_api_platform_instances" "test_api_platform_instances" {

	#Optional
	compartment_id = var.compartment_id
	id = var.api_platform_instance_id
	name = var.api_platform_instance_name
	state = var.api_platform_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance
* `name` - (Optional) A filter to return only resources that match the given name exactly
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `api_platform_instance_collection` - The list of api_platform_instance_collection.

### ApiPlatformInstance Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User-provided changeable and non-unique description of the instance
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance
* `idcs_app` - IDCS app associated with the instance, that can be used to manage the roles of the users
	* `url` - IDCS URL of the app
* `lifecycle_details` - A message that describes the current state of the instance in more detail. For example, can be used to provide actionable information for a resource in the Failed state 
* `name` - A regionally unique, non-changeable instance name provided by the user during instance creation
* `state` - The current state of the instance
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the instance was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the instance was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `uris` - Service URIs pertaining to the instance
	* `developers_portal_uri` - Developer's Portal URI of the instance (/developers)
	* `management_portal_uri` - Management Portal URI of the instance (/apiplatform)

