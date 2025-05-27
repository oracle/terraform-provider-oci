---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_named_credentials"
sidebar_current: "docs-oci-datasource-management_agent-named_credentials"
description: |-
  Provides the list of Named Credentials in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_named_credentials
This data source provides the list of Named Credentials in Oracle Cloud Infrastructure Management Agent service.

A list of Management Agent Data Sources for the given Management Agent Id.


## Example Usage

```hcl
data "oci_management_agent_named_credentials" "test_named_credentials" {
	#Required
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id

	#Optional
	id = [var.named_credential_id]
	name = [var.named_credential_name]
	state = [var.named_credential_state]
	type = [var.named_credential_type]
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) Filter list for these Named credentials identifiers (ocid) values.
* `management_agent_id` - (Required) The ManagementAgentID of the agent from which the named credentials are associated.
* `name` - (Optional) Filter list for these name items.
* `state` - (Optional) Filter list to return only Management Agents in the particular lifecycle state.
* `type` - (Optional) Filter list for these type values.


## Attributes Reference

The following attributes are exported:

* `named_credential_collection` - The list of named_credential_collection.

### NamedCredential Reference

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

