---
subcategory: "Generative Ai"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_apikey"
sidebar_current: "docs-oci-datasource-generative_ai-apikey"
description: |-
  Provides details about a specific Apikey in Oracle Cloud Infrastructure Generative Ai service
---

# Data Source: oci_generative_ai_apikey
This data source provides details about a specific Apikey resource in Oracle Cloud Infrastructure Generative Ai service.

Gets information about an API key.

## Example Usage

```hcl
data "oci_generative_ai_apikey" "test_apikey" {
	#Required
	api_key_id = oci_identity_api_key.test_api_key.id
}
```

## Argument Reference

The following arguments are supported:

* `api_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the APIKey.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID to create the apiKey in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the Api key.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - the ApiKey id.
* `keys` - The list of keys.
	* `key` - The key.
	* `key_mask` - The masked key.
	* `key_name` - The key name.
	* `state` - The current state of the API key item.
	* `time_activated` - The date and time that the key is activated in the format of an RFC3339 datetime string.
	* `time_created` - The date and time that the key was created in the format of an RFC3339 datetime string.
	* `time_deactivated` - The date and time that the key is deactivated in the format of an RFC3339 datetime string.
	* `time_expiry` - The date and time when the key would be expired, if not provided it would be 90 days, in the format defined by RFC 3339.
	* `time_last_used` - The date and time that the key is last used in the format of an RFC3339 datetime string.
	* `time_revoked` - The date and time that the key is revoked in the format of an RFC3339 datetime string.
* `lifecycle_details` - A message describing the current state with detail that can provide actionable information.
* `state` - The current state of the API key.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the ApiKey was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time the ApiKey was updated, in the format defined by RFC 3339.

