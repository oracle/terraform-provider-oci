---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_outbound_connector"
sidebar_current: "docs-oci-resource-file_storage-outbound_connector"
description: |-
  Provides the Outbound Connector resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_outbound_connector
This resource provides the Outbound Connector resource in Oracle Cloud Infrastructure File Storage service.

Creates a new outbound connector in the specified compartment.
You can associate an outbound connector with a mount target only when
they exist in the same availability domain.

For information about access control and compartments, see
[Overview of the IAM
Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

For information about availability domains, see [Regions and
Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the
`ListAvailabilityDomains` operation in the Identity and Access
Management Service API.

All Oracle Cloud Infrastructure Services resources, including
outbound connectors, get an Oracle-assigned, unique ID called an
Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource
type, or by viewing the resource in the Console.


## Example Usage

```hcl
resource "oci_file_storage_outbound_connector" "test_outbound_connector" {
	#Required
	availability_domain = var.outbound_connector_availability_domain
	bind_distinguished_name = var.outbound_connector_bind_distinguished_name
	compartment_id = var.compartment_id
	connector_type = var.outbound_connector_connector_type
	endpoints {
		#Required
		hostname = var.outbound_connector_endpoints_hostname
		port = var.outbound_connector_endpoints_port
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.outbound_connector_display_name
	freeform_tags = {"Department"= "Finance"}
	locks {
		#Required
		type = var.outbound_connector_locks_type

		#Optional
		message = var.outbound_connector_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.outbound_connector_locks_time_created
	}
	password_secret_id = oci_vault_secret.test_secret.id
	password_secret_version = var.outbound_connector_password_secret_version
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain the outbound connector is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `bind_distinguished_name` - (Required) The LDAP Distinguished Name of the bind account. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
* `connector_type` - (Required) The account type of this outbound connector.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My outbound connector` 
* `endpoints` - (Required) Array of server endpoints to use when connecting with the LDAP bind account. 
	* `hostname` - (Required) Name of the DNS server.
	* `port` - (Required) Port of the DNS server.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the password for the LDAP bind account in the Vault.
* `password_secret_version` - (Optional) Version of the password secret in the Vault to use.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Outbound Connector
	* `update` - (Defaults to 20 minutes), when updating the Outbound Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Outbound Connector


## Import

OutboundConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_outbound_connector.test_outbound_connector "id"
```

