---
subcategory: "Bastion"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bastion_bastion"
sidebar_current: "docs-oci-datasource-bastion-bastion"
description: |-
  Provides details about a specific Bastion in Oracle Cloud Infrastructure Bastion service
---

# Data Source: oci_bastion_bastion
This data source provides details about a specific Bastion resource in Oracle Cloud Infrastructure Bastion service.

Retrieves a bastion identified by the bastion ID. A bastion provides secured, public access to target resources in the cloud that you cannot otherwise reach from the internet.

## Example Usage

```hcl
data "oci_bastion_bastion" "test_bastion" {
	#Required
	bastion_id = oci_bastion_bastion.test_bastion.id
}
```

## Argument Reference

The following arguments are supported:

* `bastion_id` - (Required) The unique identifier (OCID) of the bastion.


## Attributes Reference

The following attributes are exported:

* `bastion_type` - The type of bastion.
* `client_cidr_block_allow_list` - A list of address ranges in CIDR notation that you want to allow to connect to sessions hosted by this bastion.
* `compartment_id` - The unique identifier (OCID) of the compartment where the bastion is located.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The unique identifier (OCID) of the bastion, which can't be changed after creation.
* `lifecycle_details` - A message describing the current state in more detail.
* `max_session_ttl_in_seconds` - The maximum amount of time that any session on the bastion can remain active.
* `max_sessions_allowed` - The maximum number of active sessions allowed on the bastion.
* `name` - The name of the bastion, which can't be changed after creation.
* `phone_book_entry` - The phonebook entry of the customer's team, which can't be changed after creation. Not applicable to `standard` bastions. 
* `private_endpoint_ip_address` - The private IP address of the created private endpoint.
* `state` - The current state of the bastion.
* `static_jump_host_ip_addresses` - A list of IP addresses of the hosts that the bastion has access to. Not applicable to `standard` bastions. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_subnet_id` - The unique identifier (OCID) of the subnet that the bastion connects to.
* `target_vcn_id` - The unique identifier (OCID) of the virtual cloud network (VCN) that the bastion connects to.
* `time_created` - The time the bastion was created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 
* `time_updated` - The time the bastion was updated. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 

