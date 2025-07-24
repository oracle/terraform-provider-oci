---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_named_credential"
sidebar_current: "docs-oci-datasource-management_agent-named_credential"
description: |-
  Provides details about a specific Named Credential in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_named_credential
This data source provides details about a specific Named Credential resource in Oracle Cloud Infrastructure Management Agent service.

Get Named credential details for given Id and given Management Agent.


## Example Usage

```hcl
data "oci_management_agent_named_credential" "test_named_credential" {
	#Required
	named_credential_id = oci_management_agent_named_credential.test_named_credential.id
}
```

## Argument Reference

The following arguments are supported:

* `named_credential_id` - (Required) Named credential ID


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the Named Credential.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Identifier for Named Credential.
* `management_agent_id` - The Management Agent parent resource to associated with this named credential. This is the ManagementAgent resource OCID.
* `name` - Name for Named Credential. This is unique for the Management Agent.
* `properties` - Properties for the named credential
	* `name` - Name of the property
	* `value` - Value of the property
	* `value_category` - The category of the Named credential property value. CLEAR_TEXT indicates the value field contains a clear text value. SECRET_IDENTIFIER indicates the value field contains a vault secret ocid identifier. ADB_IDENTIFIER indicates the value field contains an Autonomous database ocid identifier. ALLOWED_VALUE indicates the value should be selected from the options in the allowedValues field. 
* `state` - The current state of the named credential
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Named Credential was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Named Credential data was last updated. An RFC3339 formatted datetime string
* `type` - The type of the Named Credential.

