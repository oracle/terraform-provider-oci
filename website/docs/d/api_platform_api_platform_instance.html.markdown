---
subcategory: "Api Platform"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_api_platform_api_platform_instance"
sidebar_current: "docs-oci-datasource-api_platform-api_platform_instance"
description: |-
  Provides details about a specific Api Platform Instance in Oracle Cloud Infrastructure Api Platform service
---

# Data Source: oci_api_platform_api_platform_instance
This data source provides details about a specific Api Platform Instance resource in Oracle Cloud Infrastructure Api Platform service.

Gets information about an API Platform Instance

## Example Usage

```hcl
data "oci_api_platform_api_platform_instance" "test_api_platform_instance" {
	#Required
	api_platform_instance_id = oci_api_platform_api_platform_instance.test_api_platform_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `api_platform_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance


## Attributes Reference

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

