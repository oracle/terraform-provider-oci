---
subcategory: "Ai Data Platform"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_data_platform_ai_data_platform"
sidebar_current: "docs-oci-resource-ai_data_platform-ai_data_platform"
description: |-
  Provides the Ai Data Platform resource in Oracle Cloud Infrastructure Ai Data Platform service
---

# oci_ai_data_platform_ai_data_platform
This resource provides the Ai Data Platform resource in Oracle Cloud Infrastructure Ai Data Platform service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/ai-data-platform/latest/AiDataPlatform

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a AiDataPlatform.


## Example Usage

```hcl
resource "oci_ai_data_platform_ai_data_platform" "test_ai_data_platform" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	ai_data_platform_type = var.ai_data_platform_ai_data_platform_type
	default_workspace_name = oci_dataintegration_workspace.test_workspace.name
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.ai_data_platform_display_name
	freeform_tags = {"Department"= "Finance"}
	system_tags = var.ai_data_platform_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `ai_data_platform_type` - (Optional) (Updatable) The AiDataPlatform type.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the AiDataPlatform in. 
* `default_workspace_name` - (Optional) The name for the default workspace for the AiDataPlatform
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `system_tags` - (Optional) (Updatable) System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ai_data_platform_type` - The AiDataPlatform type.
* `alias_key` - The alias Id of the AiDataPlatform which is the short form of OCID.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IAM user.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the AiDataPlatform.
* `lifecycle_details` - A message that describes the current state of the AiDataPlatform in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the AiDataPlatform.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the AiDataPlatform was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the AiDataPlatform was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `web_socket_endpoint` - The WebSocket URL of the AiDataPlatform.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ai Data Platform
	* `update` - (Defaults to 20 minutes), when updating the Ai Data Platform
	* `delete` - (Defaults to 20 minutes), when destroying the Ai Data Platform


## Import

AiDataPlatforms can be imported using the `id`, e.g.

```
$ terraform import oci_ai_data_platform_ai_data_platform.test_ai_data_platform "id"
```

