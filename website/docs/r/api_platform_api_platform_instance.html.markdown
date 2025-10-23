---
subcategory: "Api Platform"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_api_platform_api_platform_instance"
sidebar_current: "docs-oci-resource-api_platform-api_platform_instance"
description: |-
  Provides the Api Platform Instance resource in Oracle Cloud Infrastructure Api Platform service
---

# oci_api_platform_api_platform_instance
This resource provides the Api Platform Instance resource in Oracle Cloud Infrastructure Api Platform service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/api_platform

Creates an API Platform Instance


## Example Usage

```hcl
resource "oci_api_platform_api_platform_instance" "test_api_platform_instance" {
	#Required
	compartment_id = var.compartment_id
	name = var.api_platform_instance_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.api_platform_instance_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the instance in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) User-provided changeable and non-unique description of the instance
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) A regionally unique, non-changeable instance name provided by the user during creation


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Api Platform Instance
	* `update` - (Defaults to 20 minutes), when updating the Api Platform Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Api Platform Instance


## Import

ApiPlatformInstances can be imported using the `id`, e.g.

```
$ terraform import oci_api_platform_api_platform_instance.test_api_platform_instance "id"
```

