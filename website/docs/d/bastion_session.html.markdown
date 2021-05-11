---
subcategory: "Bastion"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bastion_session"
sidebar_current: "docs-oci-datasource-bastion-session"
description: |-
  Provides details about a specific Session in Oracle Cloud Infrastructure Bastion service
---

# Data Source: oci_bastion_session
This data source provides details about a specific Session resource in Oracle Cloud Infrastructure Bastion service.

Retrieves a session identified by the session ID. A bastion session lets authorized users connect to a target resource for a predetermined amount of time.

## Example Usage

```hcl
data "oci_bastion_session" "test_session" {
	#Required
	session_id = oci_bastion_session.test_session.id
}
```

## Argument Reference

The following arguments are supported:

* `session_id` - (Required) The unique identifier (OCID) of the session.


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

