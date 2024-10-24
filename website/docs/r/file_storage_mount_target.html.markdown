---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_mount_target"
sidebar_current: "docs-oci-resource-file_storage-mount_target"
description: |-
  Provides the Mount Target resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_mount_target
This resource provides the Mount Target resource in Oracle Cloud Infrastructure File Storage service.

Creates a new mount target in the specified compartment and
subnet. You can associate a file system with a mount
target only when they exist in the same availability domain. Instances
can connect to mount targets in another availablity domain, but
you might see higher latency than with instances in the same
availability domain as the mount target.

Mount targets have one or more private IP addresses that you can
provide as the host portion of remote target parameters in
client mount commands. These private IP addresses are listed
in the privateIpIds property of the mount target and are highly available. Mount
targets also consume additional IP addresses in their subnet.
Do not use /30 or smaller subnets for mount target creation because they
do not have sufficient available IP addresses.
Allow at least three IP addresses for each mount target.

For information about access control and compartments, see
[Overview of the IAM
Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

For information about availability domains, see [Regions and
Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the
`ListAvailabilityDomains` operation in the Identity and Access
Management Service API.

All Oracle Cloud Infrastructure Services resources, including
mount targets, get an Oracle-assigned, unique ID called an
Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource
type, or by viewing the resource in the Console.


## Example Usage

```hcl
resource "oci_file_storage_mount_target" "test_mount_target" {
	#Required
	availability_domain = var.mount_target_availability_domain
	compartment_id = var.compartment_id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.mount_target_display_name
	freeform_tags = {"Department"= "Finance"}
	hostname_label = var.mount_target_hostname_label
	idmap_type = var.mount_target_idmap_type
	ip_address = var.mount_target_ip_address
	is_lock_override = var.mount_target_is_lock_override
	kerberos {
		#Required
		kerberos_realm = var.mount_target_kerberos_kerberos_realm

		#Optional
		backup_key_tab_secret_version = var.mount_target_kerberos_backup_key_tab_secret_version
		current_key_tab_secret_version = var.mount_target_kerberos_current_key_tab_secret_version
		is_kerberos_enabled = var.mount_target_kerberos_is_kerberos_enabled
		key_tab_secret_id = oci_vault_secret.test_secret.id
	}
	ldap_idmap {

		#Optional
		cache_lifetime_seconds = var.mount_target_ldap_idmap_cache_lifetime_seconds
		cache_refresh_interval_seconds = var.mount_target_ldap_idmap_cache_refresh_interval_seconds
		group_search_base = var.mount_target_ldap_idmap_group_search_base
		negative_cache_lifetime_seconds = var.mount_target_ldap_idmap_negative_cache_lifetime_seconds
		outbound_connector1id = oci_file_storage_outbound_connector1.test_outbound_connector1.id
		outbound_connector2id = oci_file_storage_outbound_connector2.test_outbound_connector2.id
		schema_type = var.mount_target_ldap_idmap_schema_type
		user_search_base = var.mount_target_ldap_idmap_user_search_base
	}
	locks {
		#Required
		type = var.mount_target_locks_type

		#Optional
		message = var.mount_target_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.mount_target_locks_time_created
	}
	nsg_ids = var.mount_target_nsg_ids
	requested_throughput = var.mount_target_requested_throughput
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain in which to create the mount target.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the mount target.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My mount target` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `hostname_label` - (Optional) The hostname for the mount target's IP address, used for DNS resolution. The value is the hostname portion of the private IP address's fully qualified domain name (FQDN). For example, `files-1` in the FQDN `files-1.subnet123.vcn1.oraclevcn.com`. Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).

	Note: This attribute value is stored in the [PrivateIp](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/PrivateIp/) resource, not in the `mountTarget` resource. To update the `hostnameLabel`, use `GetMountTarget` to obtain the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the mount target's private IPs (`privateIpIds`). Then, you can use [UpdatePrivateIp](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/PrivateIp/UpdatePrivateIp) to update the `hostnameLabel` value.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `files-1` 
* `idmap_type` - (Optional) (Updatable) The method used to map a Unix UID to secondary groups, if any.
* `ip_address` - (Optional) A private IP address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.  Example: `10.0.3.3` 
* `is_lock_override` - (Optional) (Updatable) Whether to override locks (if any exist).
* `kerberos` - (Optional) (Updatable) Kerberos details needed to create configuration. 
	* `backup_key_tab_secret_version` - (Optional) (Updatable) Version of the keytab Secret in the Vault to use as a backup.
	* `current_key_tab_secret_version` - (Optional) (Updatable) Version of the keytab Secret in the Vault to use.
	* `is_kerberos_enabled` - (Optional) (Updatable) Specifies whether to enable or disable Kerberos.
	* `kerberos_realm` - (Required) (Updatable) The Kerberos realm that the mount target will join.
	* `key_tab_secret_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the keytab Secret in the Vault.
* `ldap_idmap` - (Optional) (Updatable) Mount target details about the LDAP ID mapping configuration. 
	* `cache_lifetime_seconds` - (Optional) (Updatable) The maximum amount of time the mount target is allowed to use a cached entry.
	* `cache_refresh_interval_seconds` - (Optional) (Updatable) The amount of time that the mount target should allow an entry to persist in its cache before attempting to refresh the entry.
	* `group_search_base` - (Optional) (Updatable) All LDAP searches are recursive starting at this group.  Example: `CN=Group,DC=domain,DC=com` 
	* `negative_cache_lifetime_seconds` - (Optional) (Updatable) The amount of time that a mount target will maintain information that a user is not found in the ID mapping configuration.
	* `outbound_connector1id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the first connector to use to communicate with the LDAP server.
	* `outbound_connector2id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second connector to use to communicate with the LDAP server.
	* `schema_type` - (Optional) (Updatable) Schema type of the LDAP account.
	* `user_search_base` - (Optional) (Updatable) All LDAP searches are recursive starting at this user.  Example: `CN=User,DC=domain,DC=com` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `nsg_ids` - (Optional) (Updatable) A list of Network Security Group [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this mount target. A maximum of 5 is allowed. Setting this to an empty array after the list is created removes the mount target from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
* `requested_throughput` - (Optional) (Updatable) Throughput for mount target in Gbps. Currently only 1 Gbps of requestedThroughput is supported during create MountTarget. Available shapes and corresponding throughput are listed at [Mount Target Performance](https://docs.oracle.com/iaas/Content/File/Tasks/managingmounttargets.htm#performance). 
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which to create the mount target. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the mount target is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the mount target.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My mount target` 
* `export_set_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated export set. Controls what file systems will be exported through Network File System (NFS) protocol on this mount target. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the mount target.
* `idmap_type` - The method used to map a Unix UID to secondary groups. If NONE, the mount target will not use the Unix UID for ID mapping.
* `kerberos` - Allows administrator to configure a mount target to interact with the administrator's Kerberos infrastructure. 
	* `backup_key_tab_secret_version` - Version of the keytab secert in the Vault to use as a backup.
	* `current_key_tab_secret_version` - Version of the keytab secret in the Vault to use.
	* `is_kerberos_enabled` - Specifies whether to enable or disable Kerberos.
	* `kerberos_realm` - The Kerberos realm that the mount target will join.
	* `key_tab_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the keytab secret in the Vault.
* `ldap_idmap` - Mount target details about the LDAP ID mapping configuration. 
	* `cache_lifetime_seconds` - The maximum amount of time the mount target is allowed to use a cached entry.
	* `cache_refresh_interval_seconds` - The amount of time that the mount target should allow an entry to persist in its cache before attempting to refresh the entry.
	* `group_search_base` - All LDAP searches are recursive starting at this group.  Example: `CN=Group,DC=domain,DC=com` 
	* `negative_cache_lifetime_seconds` - The amount of time that a mount target will maintain information that a user is not found in the ID mapping configuration.
	* `outbound_connector1id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the first connector to use to communicate with the LDAP server.
	* `outbound_connector2id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second connector to use to communicate with the LDAP server.
	* `schema_type` - Schema type of the LDAP account.
	* `user_search_base` - All LDAP searches are recursive starting at this user.  Example: `CN=User,DC=domain,DC=com` 
* `lifecycle_details` - Additional information about the current 'lifecycleState'.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `nsg_ids` - A list of Network Security Group [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this mount target. A maximum of 5 is allowed. Setting this to an empty array after the list is created removes the mount target from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
* `observed_throughput` - Current billed throughput for mount target in Gbps. This corresponds to shape of mount target. Available shapes and corresponding throughput are listed at [Mount Target Performance](https://docs.oracle.com/iaas/Content/File/Tasks/managingmounttargets.htm#performance). 
* `private_ip_ids` - The OCIDs of the private IP addresses associated with this mount target.
* `requested_throughput` - 
	* New throughput for mount target at the end of billing cycle in Gbps. 
* `reserved_storage_capacity` - 
	* Reserved capacity (GB) associated with this mount target. Reserved capacity depends on observedThroughput value of mount target. Value is listed at [Mount Target Performance](https://docs.oracle.com/iaas/Content/File/Tasks/managingmounttargets.htm#performance). 
* `state` - The current state of the mount target.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the mount target is in.
* `time_billing_cycle_end` - The date and time the mount target current billing cycle will end, expressed in  [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. Once a cycle ends, it is updated  automatically to next timestamp which is after 30 days.  Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the mount target was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Mount Target
	* `update` - (Defaults to 20 minutes), when updating the Mount Target
	* `delete` - (Defaults to 20 minutes), when destroying the Mount Target


## Import

MountTargets can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_mount_target.test_mount_target "id"
```

