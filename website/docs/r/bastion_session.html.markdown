---
subcategory: "Bastion"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bastion_session"
sidebar_current: "docs-oci-resource-bastion-session"
description: |-
  Provides the Session resource in Oracle Cloud Infrastructure Bastion service
---

# oci_bastion_session
This resource provides the Session resource in Oracle Cloud Infrastructure Bastion service.

Creates a new session in a bastion. A bastion session lets authorized users connect to a target resource for a predetermined amount of time. The Bastion service recognizes two types of sessions, managed SSH sessions and SSH port forwarding sessions. Managed SSH sessions require that the target resource has an OpenSSH server and the Oracle Cloud Agent both running.


## Example Usage

```hcl
resource "oci_bastion_session" "test_session" {
	#Required
	bastion_id = oci_bastion_bastion.test_bastion.id
	key_details {
		#Required
		public_key_content = var.session_key_details_public_key_content
	}
	target_resource_details {
		#Required
		session_type = var.session_target_resource_details_session_type

		#Optional
        target_resource_id = oci_bastion_target_resource.test_target_resource.id
		target_resource_operating_system_user_name = oci_identity_user.test_user.name
		target_resource_port = var.session_target_resource_details_target_resource_port
		target_resource_private_ip_address = var.session_target_resource_details_target_resource_private_ip_address
	}

	#Optional
	display_name = var.session_display_name
	key_type = var.session_key_type
	session_ttl_in_seconds = var.session_session_ttl_in_seconds
}
```

## Argument Reference

The following arguments are supported:

* `bastion_id` - (Required) The unique identifier (OCID) of the bastion on which to create this session.
* `display_name` - (Optional) (Updatable) The name of the session.
* `key_details` - (Required) Public key details for a bastion session.
	* `public_key_content` - (Required) The public key in OpenSSH format of the SSH key pair for the session. When you connect to the session, you must provide the private key of the same SSH key pair.
* `key_type` - (Optional) The type of the key used to connect to the session. PUB is a standard public key in OpenSSH format.
* `session_ttl_in_seconds` - (Optional) The amount of time the session can remain active.
* `target_resource_details` - (Required) Details about a bastion session's target resource.
	* `session_type` - (Required) The session type.
	* `target_resource_id` - (Optional) The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to. It's optional depends on the type of session you want to create.
		* (Required) For MANAGED_SSH session type, we can only use target_resource_id to create session.
		* (Optional) For PORT_FORWARDING session type, you must either use target_resource_id or target_resource_private_ip_address
	* `target_resource_operating_system_user_name` - (Required when session_type=MANAGED_SSH) The name of the user on the target resource operating system that the session uses for the connection.
	* `target_resource_port` - (Optional) The port number to connect to on the target resource.
	* `target_resource_private_ip_address` - (Optional) The private IP address of the target resource that the session connects to. For PORT_FORWARDING session type, you must either use target_resource_id or target_resource_private_ip_address


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bastion_id` - The unique identifier (OCID) of the bastion that is hosting this session.
* `bastion_name` - The name of the bastion that is hosting this session.
* `bastion_public_host_key_info` - The public key of the bastion host. You can use this to verify that you're connecting to the correct bastion.
* `bastion_user_name` - The username that the session uses to connect to the target resource.
* `display_name` - The name of the session.
* `id` - The unique identifier (OCID) of the session, which can't be changed after creation.
* `key_details` - Public key details for a bastion session.
	* `public_key_content` - The public key in OpenSSH format of the SSH key pair for the session. When you connect to the session, you must provide the private key of the same SSH key pair.
* `key_type` - The type of the key used to connect to the session. PUB is a standard public key in OpenSSH format.
* `lifecycle_details` - A message describing the current session state in more detail.
* `session_ttl_in_seconds` - The amount of time the session can remain active.
* `ssh_metadata` - The connection message for the session.
* `state` - The current state of the session.
* `target_resource_details` - Details about a bastion session's target resource.
	* `session_type` - The Bastion service recognizes two types of sessions, managed SSH sessions and SSH port forwarding sessions. Managed SSH sessions require that the target resource has an OpenSSH server and the Oracle Cloud Agent both running.
	* `target_resource_display_name` - The display name of the target Compute instance that the session connects to.
	* `target_resource_id` - The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to.
	* `target_resource_operating_system_user_name` - The name of the user on the target resource operating system that the session uses for the connection.
	* `target_resource_port` - The port number to connect to on the target resource.
	* `target_resource_private_ip_address` - The private IP address of the target resource that the session connects to.
* `time_created` - The time the session was created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 
* `time_updated` - The time the session was updated. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Session
	* `update` - (Defaults to 20 minutes), when updating the Session
	* `delete` - (Defaults to 20 minutes), when destroying the Session


## Import

Sessions can be imported using the `id`, e.g.

```
$ terraform import oci_bastion_session.test_session "id"
```

