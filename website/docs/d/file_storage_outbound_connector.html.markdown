---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_outbound_connector"
sidebar_current: "docs-oci-datasource-file_storage-outbound_connector"
description: |-
  Provides details about a specific Outbound Connector in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_outbound_connector
This data source provides details about a specific Outbound Connector resource in Oracle Cloud Infrastructure File Storage service.

Gets the specified outbound connector's information.

## Example Usage

```hcl
data "oci_file_storage_outbound_connector" "test_outbound_connector" {
	#Required
	outbound_connector_id = oci_file_storage_outbound_connector.test_outbound_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `outbound_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the outbound connector.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the outbound connector is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `bind_distinguished_name` - The LDAP Distinguished Name of the account.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
* `connector_type` - The account type of this outbound connector.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My outbound connector` 
* `endpoints` - Array of server endpoints to use when connecting with the LDAP bind account. 
	* `hostname` - Name of the DNS server.
	* `port` - Port of the DNS server.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the outbound connector.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the password for the LDAP bind account in the Vault.
* `password_secret_version` - Version of the password secret in the Vault to use.
* `state` - The current state of this outbound connector.
* `time_created` - The date and time the outbound connector was created in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

