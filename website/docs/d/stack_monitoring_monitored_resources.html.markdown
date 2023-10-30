---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resources"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitored_resources"
description: |-
  Provides the list of Monitored Resources in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitored_resources
This data source provides the list of Monitored Resources in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of monitored resources.

## Example Usage

```hcl
data "oci_stack_monitoring_monitored_resources" "test_monitored_resources" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.monitored_resource_name
	work_request_id = oci_containerengine_work_request.test_work_request.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which data is listed.
* `name` - (Optional) A filter to return resources that match exact resource name.
* `work_request_id` - (Optional) A filter to return resources which were impacted as part of this work request identifier.


## Attributes Reference

The following attributes are exported:

* `monitored_resource_collection` - The list of monitored_resource_collection.

### MonitoredResource Reference

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

