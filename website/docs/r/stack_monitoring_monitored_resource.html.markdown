---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resource"
description: |-
  Provides the Monitored Resource resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resource
This resource provides the Monitored Resource resource in Oracle Cloud Infrastructure Stack Monitoring service.

Creates a new monitored resource for the given resource type with the details and submits 
a work request for promoting the resource to agent. Once the resource is successfully 
added to agent, resource state will be marked active.


## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resource" "test_monitored_resource" {
	#Required
	compartment_id = var.compartment_id
	name = var.monitored_resource_name
	type = var.monitored_resource_type

	#Optional
	additional_aliases {
		#Required
		credential {
			#Required
			name = var.monitored_resource_additional_aliases_credential_name
			service = var.monitored_resource_additional_aliases_credential_service
			source = var.monitored_resource_additional_aliases_credential_source
		}
		name = var.monitored_resource_additional_aliases_name
		source = var.monitored_resource_additional_aliases_source
	}
	additional_credentials {

		#Optional
		credential_type = var.monitored_resource_additional_credentials_credential_type
		description = var.monitored_resource_additional_credentials_description
		key_id = oci_kms_key.test_key.id
		name = var.monitored_resource_additional_credentials_name
		properties {

			#Optional
			name = var.monitored_resource_additional_credentials_properties_name
			value = var.monitored_resource_additional_credentials_properties_value
		}
		source = var.monitored_resource_additional_credentials_source
		type = var.monitored_resource_additional_credentials_type
	}
	aliases {
		#Required
		credential {
			#Required
			name = var.monitored_resource_aliases_credential_name
			service = var.monitored_resource_aliases_credential_service
			source = var.monitored_resource_aliases_credential_source
		}
		name = var.monitored_resource_aliases_name
		source = var.monitored_resource_aliases_source
	}
	credentials {

		#Optional
		credential_type = var.monitored_resource_credentials_credential_type
		description = var.monitored_resource_credentials_description
		key_id = var.monitored_resource_credentials_key_id
		name = var.monitored_resource_credentials_name
		properties {

			#Optional
			name = var.monitored_resource_credentials_properties_name
			value = var.monitored_resource_credentials_properties_value
		}
		source = var.monitored_resource_credentials_source
		type = var.monitored_resource_credentials_type
	}
	database_connection_details {
		#Required
		port = var.monitored_resource_database_connection_details_port
		protocol = var.monitored_resource_database_connection_details_protocol
		service_name = var.monitored_resource_database_service_name

		#Optional
		connector_id = var.monitored_resource_database_connector_id
		db_id = var.monitored_resource_database_id
		db_unique_name = var.monitored_resource_database_connection_details_db_unique_name
		ssl_secret_id = oci_vault_secret.test_secret.id
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.monitored_resource_display_name
	external_resource_id = var.monitored_resource_external_resource_id
	external_id = oci_stack_monitoring_external.test_external.id
	freeform_tags = {"bar-key"= "value"}
	host_name = var.monitored_resource_host_name
	license = var.monitored_resource_license
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	properties {

		#Optional
		name = var.monitored_resource_properties_name
		value = var.monitored_resource_properties_value
	}
	resource_time_zone = var.monitored_resource_resource_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `additional_aliases` - (Optional) (Updatable) List of MonitoredResourceAliasCredentials. This property complements the existing  "aliases" property by allowing user to specify more than one credential alias.  If both "aliases" and "additionalAliases" are specified, union of the  values is used as list of aliases applicable for this resource. If any duplicate found in the combined list of "alias" and "additionalAliases",  an error will be thrown. 
	* `credential` - (Required) (Updatable) Monitored Resource Alias Reference Source Credential. 
		* `name` - (Required) (Updatable) The name of the pre-existing source credential which alias cred should point to. This should refer to the pre-existing source attribute which is bound to credential name. 
		* `service` - (Required) (Updatable) The name of the service owning the credential.  Example: stack-monitoring or dbmgmt 
		* `source` - (Required) (Updatable) The source type and source name combination,delimited with (.) separator. This refers to the pre-existing source which alias cred should point to. Ex. {source type}.{source name} and source type max char limit is 63. 
	* `name` - (Required) (Updatable) The name of the alias, within the context of the source.
	* `source` - (Required) (Updatable) The source type and source name combination,delimited with (.) separator. Example: {source type}.{source name} and source type max char limit is 63. 
* `additional_credentials` - (Optional) (Updatable) List of MonitoredResourceCredentials. This property complements the existing  "credentials" property by allowing user to specify more than one credential.  If both "credential" and "additionalCredentials" are specified, union of the  values is used as list of credentials applicable for this resource. If any duplicate found in the combined list of "credentials" and "additionalCredentials",  an error will be thrown. 
	* `credential_type` - (Optional) (Updatable) Type of credentials specified in the credentials element. Three possible values - EXISTING, PLAINTEXT and ENCRYPTED.
		* EXISTING  - Credential is already stored in agent and only credential name need to be passed for existing credential.
		* PLAINTEXT - The credential properties will have credentials in plain text format.
		* ENCRYPTED - The credential properties will have credentials stored in vault in encrypted format using KMS client which uses master key for encryption. The same master key will be used to decrypt the credentials before passing on to the management agent. 
	* `description` - (Optional) (Updatable) The user-specified textual description of the credential.
	* `key_id` - (Required when credential_type=ENCRYPTED) (Updatable) The master key should be created in Oracle Cloud Infrastructure Vault owned by the client of this API.  The user should have permission to access the vault key. 
	* `name` - (Optional) (Updatable) The name of the credential, within the context of the source.
	* `properties` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The credential properties list. Credential property values will be either  in plain text format or encrypted for encrypted credentials. 
		* `name` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The name of the credential property, should confirm with names of properties of this credential's type.  Example: For JMXCreds type, credential property name for weblogic user is 'Username'. 
		* `value` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The value of the credential property name. Example: For JMXCreds type, credential property value for 'Username' property is 'weblogic'. 
	* `source` - (Optional) (Updatable) The source type and source name combination, delimited with (.) separator. {source type}.{source name} and source type max char limit is 63. 
	* `type` - (Optional) (Updatable) The type of the credential ( ex. JMXCreds,DBCreds).
* `aliases` - (Optional) (Updatable) Monitored Resource Alias Credential Details
	* `credential` - (Required) (Updatable) Monitored Resource Alias Reference Source Credential. 
		* `name` - (Required) (Updatable) The name of the pre-existing source credential which alias cred should point to. This should refer to the pre-existing source attribute which is bound to credential name. 
		* `service` - (Required) (Updatable) The name of the service owning the credential.  Example: stack-monitoring or dbmgmt 
		* `source` - (Required) (Updatable) The source type and source name combination,delimited with (.) separator. This refers to the pre-existing source which alias cred should point to. Ex. {source type}.{source name} and source type max char limit is 63. 
	* `name` - (Required) (Updatable) The name of the alias, within the context of the source.
	* `source` - (Required) (Updatable) The source type and source name combination,delimited with (.) separator. Example: {source type}.{source name} and source type max char limit is 63. 
* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `credentials` - (Optional) (Updatable) Monitored Resource Credential Details. 
	* `credential_type` - (Optional) (Updatable) Type of credentials specified in the credentials element. Three possible values - EXISTING, PLAINTEXT and ENCRYPTED.
		* EXISTING  - Credential is already stored in agent and only credential name need to be passed for existing credential.
		* PLAINTEXT - The credential properties will have credentials in plain text format.
		* ENCRYPTED - The credential properties will have credentials stored in vault in encrypted format using KMS client which uses master key for encryption. The same master key will be used to decrypt the credentials before passing on to the management agent. 
	* `description` - (Optional) (Updatable) The user-specified textual description of the credential.
	* `key_id` - (Required when credential_type=ENCRYPTED) (Updatable) The master key should be created in Oracle Cloud Infrastructure Vault owned by the client of this API.  The user should have permission to access the vault key. 
	* `name` - (Optional) (Updatable) The name of the credential, within the context of the source.
	* `properties` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The credential properties list. Credential property values will be either  in plain text format or encrypted for encrypted credentials. 
		* `name` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The name of the credential property, should confirm with names of properties of this credential's type.  Example: For JMXCreds type, credential property name for weblogic user is 'Username'. 
		* `value` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The value of the credential property name. Example: For JMXCreds type, credential property value for 'Username' property is 'weblogic'. 
	* `source` - (Optional) (Updatable) The source type and source name combination, delimited with (.) separator. {source type}.{source name} and source type max char limit is 63. 
	* `type` - (Optional) (Updatable) The type of the credential ( ex. JMXCreds,DBCreds).
* `database_connection_details` - (Optional) (Updatable) Connection details for the database. 
	* `connector_id` - (Optional) (Updatable) Database connector Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `db_id` - (Optional) (Updatable) dbId of the database. 
	* `db_unique_name` - (Optional) (Updatable) UniqueName used for database connection requests.
	* `port` - (Required) (Updatable) Listener Port number used for connection requests.
	* `protocol` - (Required) (Updatable) Protocol used in DB connection string when connecting to external database service.
	* `service_name` - (Required) (Updatable) Service name used for connection requests.
	* `ssl_secret_id` - (Optional) (Updatable) SSL Secret Identifier for TCPS connector in Oracle Cloud Infrastructure Vault[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Monitored resource display name.
* `external_resource_id` - (Optional) Generally used by DBaaS to send the Database OCID stored on the DBaaS. The same will be passed to resource service to enable Stack Monitoring Service on DBM. This will be stored in Stack Monitoring Resource Service data store as identifier for monitored resource. If this header is not set as part of the request, then an id will be generated and stored for the resource. 
* `external_id` - (Optional) External resource is any Oracle Cloud Infrastructure resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) which is not a Stack Monitoring service resource. Currently supports only Oracle Cloud Infrastructure compute instance. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_name` - (Optional) (Updatable) Host name of the monitored resource. 
* `license` - (Optional) (Updatable) License edition of the monitored resource. If not provided  the default license type for the compartment will be used. 
* `management_agent_id` - (Optional) Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - (Required) Monitored Resource Name. 
* `properties` - (Optional) (Updatable) List of monitored resource properties. 
	* `name` - (Optional) (Updatable) Property Name. 
	* `value` - (Optional) (Updatable) Property Value. 
* `resource_time_zone` - (Optional) (Updatable) Time zone in the form of tz database canonical zone ID. Specifies the preference with a value that uses the IANA Time Zone Database format (x-obmcs-time-zone). For example - America/Los_Angeles 
* `type` - (Required) Monitored Resource Type. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `aliases` - Monitored Resource Alias Credential Details
	* `credential` - Monitored Resource Alias Reference Source Credential. 
		* `name` - The name of the pre-existing source credential which alias cred should point to. This should refer to the pre-existing source attribute which is bound to credential name. 
		* `service` - The name of the service owning the credential.  Example: stack-monitoring or dbmgmt 
		* `source` - The source type and source name combination,delimited with (.) separator. This refers to the pre-existing source which alias cred should point to. Ex. {source type}.{source name} and source type max char limit is 63. 
	* `name` - The name of the alias, within the context of the source.
	* `source` - The source type and source name combination,delimited with (.) separator. Example: {source type}.{source name} and source type max char limit is 63. 
* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `credentials` - Monitored Resource Credential Details. 
	* `credential_type` - Type of credentials specified in the credentials element. Three possible values - EXISTING, PLAINTEXT and ENCRYPTED.
		* EXISTING  - Credential is already stored in agent and only credential name need to be passed for existing credential.
		* PLAINTEXT - The credential properties will have credentials in plain text format.
		* ENCRYPTED - The credential properties will have credentials stored in vault in encrypted format using KMS client which uses master key for encryption. The same master key will be used to decrypt the credentials before passing on to the management agent. 
	* `description` - The user-specified textual description of the credential.
	* `key_id` - The master key should be created in Oracle Cloud Infrastructure Vault owned by the client of this API.  The user should have permission to access the vault key. 
	* `name` - The name of the credential, within the context of the source.
	* `properties` - The credential properties list. Credential property values will be either  in plain text format or encrypted for encrypted credentials. 
		* `name` - The name of the credential property, should confirm with names of properties of this credential's type.  Example: For JMXCreds type, credential property name for weblogic user is 'Username'. 
		* `value` - The value of the credential property name. Example: For JMXCreds type, credential property value for 'Username' property is 'weblogic'. 
	* `source` - The source type and source name combination, delimited with (.) separator. {source type}.{source name} and source type max char limit is 63. 
	* `type` - The type of the credential ( ex. JMXCreds,DBCreds).
* `database_connection_details` - Connection details for the database. 
	* `connector_id` - Database connector Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `db_id` - dbId of the database. 
	* `db_unique_name` - UniqueName used for database connection requests.
	* `port` - Listener Port number used for connection requests.
	* `protocol` - Protocol used in DB connection string when connecting to external database service.
	* `service_name` - Service name used for connection requests.
	* `ssl_secret_id` - SSL Secret Identifier for TCPS connector in Oracle Cloud Infrastructure Vault[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Monitored resource display name.
* `external_id` - The external resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). External resource is any Oracle Cloud Infrastructure resource which is not a Stack Monitoring service resource. Currently supports only following resource types - Container database, non-container database,  pluggable database and Oracle Cloud Infrastructure compute instance. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_name` - Monitored resource host name.
* `id` - Monitored resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `license` - License edition of the monitored resource.
* `management_agent_id` - Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - Monitored resource name.
* `properties` - List of monitored resource properties. 
	* `name` - Property Name. 
	* `value` - Property Value. 
* `resource_time_zone` - Time zone in the form of tz database canonical zone ID.
* `state` - Lifecycle state of the monitored resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenant_id` - Tenancy Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `time_created` - The date and time when the monitored resource was created, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time when the monitored resource was last updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `type` - Monitored Resource Type. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resource
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resource
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resource


## Import

MonitoredResources can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resource.test_monitored_resource "id"
```

