---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_named_credentials_metadata"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_named_credentials_metadata"
description: |-
  Provides details about a specific Management Agent Named Credentials Metadata in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_named_credentials_metadata
This data source provides details about a specific Management Agent Named Credentials Metadata resource in Oracle Cloud Infrastructure Management Agent service.

Return the Metadata definition for Named Credentials supported by Management Agent.


## Example Usage

```hcl
data "oci_management_agent_management_agent_named_credentials_metadata" "test_management_agent_named_credentials_metadata" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `management_agent_id` - (Optional) Filter the named credential metadata which is compatible with the given Management Agent identifier.


## Attributes Reference

The following attributes are exported:

* `metadata` - List of supported metadata definitions.
	* `display_name` - Display name for this type of Named Credential
	* `minimum_agent_version` - This Named Credential type is supported on management agents at this version or above.
	* `properties` - The property definitions for this named credential metadata
		* `allowed_values` - List of values which can be applied to the value when valueCategory is ALLOWED_VALUES
		* `default_value` - The default value which will be used if no value is set.  If defaultValue is empty, then no default will be set.
		* `display_name` - The field display name
		* `is_required` - Set to true if the field must be defined
		* `name` - The field name
		* `regex` - Optional regular expression definition which will be applied to the value when valueCategory is CLEAR_TEXT
		* `value_category` - List of value categories of field allowed for this property
	* `type` - The type of the Named Credential.

