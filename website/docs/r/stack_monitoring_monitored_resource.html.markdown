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

Creates a new monitored resource for the given resource type


## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resource" "test_monitored_resource" {
	#Required
	compartment_id = var.compartment_id
	name = var.monitored_resource_name
	type = var.monitored_resource_type

	#Optional
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
	display_name = var.monitored_resource_display_name
	external_resource_id = var.monitored_resource_external_resource_id
	external_id = oci_stack_monitoring_external.test_external.id
	host_name = var.monitored_resource_host_name
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

* `aliases` - (Optional) (Updatable) Monitored Resource Alias Credential Details
	* `credential` - (Required) (Updatable) Monitored Resource Alias Reference Source Credential
		* `name` - (Required) (Updatable) The name of the pre-existing source credential which alias cred should point to. This should refer to the pre-existing source attribute binded credential name.
		* `service` - (Required) (Updatable) The name of the service owning the credential. Ex stack-monitoring or dbmgmt
		* `source` - (Required) (Updatable) The source type and source name combination,delimited with (.) separator. This refers to the pre-existing source which alias cred should point to. Ex. {source type}.{source name} and source type max char limit is 63.
	* `name` - (Required) (Updatable) The name of the alias, within the context of the source.
	* `source` - (Required) (Updatable) The source type and source name combination,delimited with (.) separator. Ex. {source type}.{source name} and source type max char limit is 63.
* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `credentials` - (Optional) (Updatable) Monitored Resource Credential Details
	* `credential_type` - (Optional) (Updatable) Type of credentials specified in the credentials element. Three possible values - EXISTING, PLAINTEXT and ENCRYPTED. * EXISTING  - Credential is already stored in agent and only credential name need to be passed for existing credential. * PLAINTEXT - The credential properties will have credentials in plain text format. * ENCRYPTED - The credential properties will have credentials stored in vault in encrypted format using KMS client which uses master key for encryption. The same master key will be used to decrypt the credentials before passing on to the management agent.
	* `description` - (Optional) (Updatable) The user-specified textual description of the credential.
	* `key_id` - (Required when credential_type=ENCRYPTED) (Updatable) The master key OCID and applicable only for property value type ENCRYPTION. Key OCID is passed as input to Key management service decrypt API to retrieve the encrypted property value text.
	* `name` - (Optional) (Updatable) The name of the credential, within the context of the source.
	* `properties` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The credential properties list. Credential property values will be either in plain text format.
		* `name` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The name of the credential property, should confirm with names of properties of this credential's type. Ex. For JMXCreds type , credential property name for weblogic user is 'Username'.
		* `value` - (Required when credential_type=ENCRYPTED | PLAINTEXT) (Updatable) The value of the credential property name. Ex. For JMXCreds type, credential property value for 'Username' property is 'weblogic'.
	* `source` - (Optional) (Updatable) The source type and source name combination,delimited with (.) separator. {source type}.{source name} and source type max char limit is 63.
	* `type` - (Optional) (Updatable) The type of the credential ( ex. JMXCreds,DBCreds).
* `database_connection_details` - (Optional) (Updatable) Connection details to connect to the database. HostName, protocol, and port should be specified.
	* `connector_id` - (Optional) (Updatable) Database connector Identifier
	* `db_id` - (Optional) (Updatable) dbId of the database
	* `db_unique_name` - (Optional) (Updatable) UniqueName used for database connection requests.
	* `port` - (Required) (Updatable) Listener Port number used for connection requests.
	* `protocol` - (Required) (Updatable) Protocol used in DB connection string when connecting to external database service.
	* `service_name` - (Required) (Updatable) Service name used for connection requests.
	* `ssl_secret_id` - (Optional) (Updatable) SSL Secret Identifier for TCPS connector in Oracle Cloud Infrastructure Vault[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `display_name` - (Optional) (Updatable) Monitored resource display name.
* `external_resource_id` - (Optional) Generally used by DBaaS to send the Database OCID stored on the DBaaS. The same will be passed to resource service to enable Stack Monitoring Service on DBM. This will be stored in Stack Monitoring Resource Service data store as identifier for monitored resource. If this header is not set as part of the request, then an id will be generated and stored for the resource. 
* `external_id` - (Optional) External resource is any Oracle Cloud Infrastructure resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) which is not a Stack Monitoring service resource. Currently supports only Oracle Cloud Infrastructure compute instance. 
* `host_name` - (Optional) (Updatable) Host name of the monitored resource
* `management_agent_id` - (Optional) Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - (Required) Monitored resource name
* `properties` - (Optional) (Updatable) List of monitored resource properties
	* `name` - (Optional) (Updatable) property name
	* `value` - (Optional) (Updatable) property value
* `resource_time_zone` - (Optional) (Updatable) Time zone in the form of tz database canonical zone ID.
* `type` - (Required) Monitored resource type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `aliases` - Monitored Resource Alias Credential Details
	* `credential` - Monitored Resource Alias Reference Source Credential
		* `name` - The name of the pre-existing source credential which alias cred should point to. This should refer to the pre-existing source attribute binded credential name.
		* `service` - The name of the service owning the credential. Ex stack-monitoring or dbmgmt
		* `source` - The source type and source name combination,delimited with (.) separator. This refers to the pre-existing source which alias cred should point to. Ex. {source type}.{source name} and source type max char limit is 63.
	* `name` - The name of the alias, within the context of the source.
	* `source` - The source type and source name combination,delimited with (.) separator. Ex. {source type}.{source name} and source type max char limit is 63.
* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `credentials` - Monitored Resource Credential Details
	* `credential_type` - Type of credentials specified in the credentials element. Three possible values - EXISTING, PLAINTEXT and ENCRYPTED. * EXISTING  - Credential is already stored in agent and only credential name need to be passed for existing credential. * PLAINTEXT - The credential properties will have credentials in plain text format. * ENCRYPTED - The credential properties will have credentials stored in vault in encrypted format using KMS client which uses master key for encryption. The same master key will be used to decrypt the credentials before passing on to the management agent.
	* `description` - The user-specified textual description of the credential.
	* `key_id` - The master key OCID and applicable only for property value type ENCRYPTION. Key OCID is passed as input to Key management service decrypt API to retrieve the encrypted property value text.
	* `name` - The name of the credential, within the context of the source.
	* `properties` - The credential properties list. Credential property values will be either in plain text format.
		* `name` - The name of the credential property, should confirm with names of properties of this credential's type. Ex. For JMXCreds type , credential property name for weblogic user is 'Username'.
		* `value` - The value of the credential property name. Ex. For JMXCreds type, credential property value for 'Username' property is 'weblogic'.
	* `source` - The source type and source name combination,delimited with (.) separator. {source type}.{source name} and source type max char limit is 63.
	* `type` - The type of the credential ( ex. JMXCreds,DBCreds).
* `database_connection_details` - Connection details to connect to the database. HostName, protocol, and port should be specified.
	* `connector_id` - Database connector Identifier
	* `db_id` - dbId of the database
	* `db_unique_name` - UniqueName used for database connection requests.
	* `port` - Listener Port number used for connection requests.
	* `protocol` - Protocol used in DB connection string when connecting to external database service.
	* `service_name` - Service name used for connection requests.
	* `ssl_secret_id` - SSL Secret Identifier for TCPS connector in Oracle Cloud Infrastructure Vault[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Monitored resource display name.
* `external_id` - External resource is any Oracle Cloud Infrastructure resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) which is not a Stack Monitoring service resource. Currently supports only following resource type identifiers - externalcontainerdatabase, externalnoncontainerdatabase, externalpluggabledatabase and Oracle Cloud Infrastructure compute instance. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_name` - Monitored resource host name.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource.
* `management_agent_id` - Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - Monitored resource name.
* `properties` - List of monitored resource properties
	* `name` - property name
	* `value` - property value
* `resource_time_zone` - Time zone in the form of tz database canonical zone ID.
* `state` - Lifecycle state of the monitored resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenant_id` - Tenancy Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `time_created` - The time the the resource was created. An RFC3339 formatted datetime string
* `time_updated` - The time the the resource was updated. An RFC3339 formatted datetime string
* `type` - Monitored resource type

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resource
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resource
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resource


## Import

MonitoredResources can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resource.test_monitored_resource "id"
```

