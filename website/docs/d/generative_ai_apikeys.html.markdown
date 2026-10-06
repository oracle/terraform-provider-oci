---
subcategory: "Generative Ai"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_apikeys"
sidebar_current: "docs-oci-datasource-generative_ai-apikeys"
description: |-
  Provides the list of Apikeys in Oracle Cloud Infrastructure Generative Ai service
---

# Data Source: oci_generative_ai_apikeys
This data source provides the list of Apikeys in Oracle Cloud Infrastructure Generative Ai service.

Lists the ApiKeys of a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_apikeys" "test_apikeys" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.apikey_display_name
	id = var.apikey_id
	state = var.apikey_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the APIKey.
* `state` - (Optional) A filter to return only resources that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `api_key_collection` - The list of api_key_collection.

### Apikey Reference

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

