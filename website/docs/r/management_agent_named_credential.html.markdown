---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_named_credential"
sidebar_current: "docs-oci-resource-management_agent-named_credential"
description: |-
  Provides the Named Credential resource in Oracle Cloud Infrastructure Management Agent service
---

# oci_management_agent_named_credential
This resource provides the Named Credential resource in Oracle Cloud Infrastructure Management Agent service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/management-agent/latest/NamedCredential

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/management_agent

Named credential creation request to given Management Agent.


## Example Usage

```hcl
resource "oci_management_agent_named_credential" "test_named_credential" {
	#Required
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	name = var.named_credential_name
	properties {
		#Required
		name = var.named_credential_properties_name
		value = var.named_credential_properties_value
		value_category = var.named_credential_properties_value_category
	}
	properties {
		#Optional
		name = var.named_credential_properties_name2
		value = var.named_credential_properties_value2
		value_category = var.named_credential_properties_value_category2
	}
	type = var.named_credential_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.named_credential_description
	freeform_tags = {"bar-key"= "value"}
}

resource "oci_management_agent_named_credential" "example_named_credential" {
	#Required
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	name = "Example1"
	properties {
		name = "DBUserName"
		value = var.vault_secret_id
		value_category = "SECRET_IDENTIFIER"
	}
	properties {
		name = "DBPassword"
		value = var.vault_secret2_id
		value_category = "SECRET_IDENTIFIER"
	}
	type = "DBCREDS"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = "Example DBCREDS named credential for management agent"
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the Named Credential.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `management_agent_id` - (Required) The Management Agent parent resource to associate this named credential with.  This is the ManagementAgent resource OCID.
* `name` - (Required) Identifier for Named Credential. This is unique for the Management Agent.
* `properties` - (Required) (Updatable) Properties for the named credential
	* `name` - (Required) (Updatable) Name of the property
	* `value` - (Required) (Updatable) Value of the property
	* `value_category` - (Required) (Updatable) The category of the Named credential property value. CLEAR_TEXT indicates the value field contains a clear text value. SECRET_IDENTIFIER indicates the value field contains a vault secret ocid identifier. ADB_IDENTIFIER indicates the value field contains an Autonomous database ocid identifier. ALLOWED_VALUE indicates the value should be selected from the options in the allowedValues field. 
* `type` - (Required) The type of the Named Credential.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Named Credential
	* `update` - (Defaults to 20 minutes), when updating the Named Credential
	* `delete` - (Defaults to 20 minutes), when destroying the Named Credential


## Import

NamedCredentials can be imported using the `id`, e.g.

```
$ terraform import oci_management_agent_named_credential.test_named_credential "id"
```

