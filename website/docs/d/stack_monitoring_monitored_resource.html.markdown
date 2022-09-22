---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitored_resource"
description: |-
  Provides details about a specific Monitored Resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitored_resource
This data source provides details about a specific Monitored Resource resource in Oracle Cloud Infrastructure Stack Monitoring service.

Gets a monitored resource by identifier

## Example Usage

```hcl
data "oci_stack_monitoring_monitored_resource" "test_monitored_resource" {
	#Required
	monitored_resource_id = oci_stack_monitoring_monitored_resource.test_monitored_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `monitored_resource_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource.


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

