---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_channel"
sidebar_current: "docs-oci-datasource-mysql-channel"
description: |-
  Provides details about a specific Channel in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_channel
This data source provides details about a specific Channel resource in Oracle Cloud Infrastructure MySQL Database service.

Gets the full details of the specified Channel, including the user-specified
configuration parameters (passwords are omitted), as well as information about
the state of the Channel, its sources and targets.


## Example Usage

```hcl
data "oci_mysql_channel" "test_channel" {
	#Required
	channel_id = oci_mysql_channel.test_channel.id
}
```

## Argument Reference

The following arguments are supported:

* `channel_id` - (Required) The Channel [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User provided description of the Channel.
* `display_name` - The user-friendly name for the Channel. It does not have to be unique.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_enabled` - Whether the Channel has been enabled by the user.
* `lifecycle_details` - A message describing the state of the Channel.
* `source` - Parameters detailing how to provision the source for the given Channel.
	* `hostname` - The network address of the MySQL instance.
	* `port` - The port the source MySQL instance listens on.
	* `source_type` - The specific source identifier.
	* `ssl_ca_certificate` - The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
		* `certificate_type` - The type of CA certificate.
		* `contents` - The string containing the CA certificate in PEM format.
	* `ssl_mode` - The SSL mode of the Channel.
	* `username` - The name of the replication user on the source MySQL instance. The username has a maximum length of 96 characters. For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html) 
* `state` - The state of the Channel.
* `target` - Details about the Channel target.
	* `applier_username` - The username for the replication applier of the target MySQL DB System.
	* `channel_name` - The case-insensitive name that identifies the replication channel. Channel names must follow the rules defined for [MySQL identifiers](https://dev.mysql.com/doc/refman/8.0/en/identifiers.html). The names of non-Deleted Channels must be unique for each DB System. 
	* `db_system_id` - The OCID of the source DB System.
	* `target_type` - The specific target identifier.
* `time_created` - The date and time the Channel was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The time the Channel was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

