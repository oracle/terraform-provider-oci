---
subcategory: "Ai Data Platform"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_data_platform_ai_data_platform"
sidebar_current: "docs-oci-datasource-ai_data_platform-ai_data_platform"
description: |-
  Provides details about a specific Ai Data Platform in Oracle Cloud Infrastructure Ai Data Platform service
---

# Data Source: oci_ai_data_platform_ai_data_platform
This data source provides details about a specific Ai Data Platform resource in Oracle Cloud Infrastructure Ai Data Platform service.

Gets information about a AiDataPlatform.

## Example Usage

```hcl
data "oci_ai_data_platform_ai_data_platform" "test_ai_data_platform" {
	#Required
	ai_data_platform_id = oci_ai_data_platform_ai_data_platform.test_ai_data_platform.id
}
```

## Argument Reference

The following arguments are supported:

* `ai_data_platform_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the AiDataPlatform.


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

