---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_mount_targets"
sidebar_current: "docs-oci-datasource-file_storage-mount_targets"
description: |-
  Provides the list of Mount Targets in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_mount_targets
This data source provides the list of Mount Targets in Oracle Cloud Infrastructure File Storage service.

Lists the mount target resources in the specified compartment.


## Example Usage

```hcl
data "oci_file_storage_mount_targets" "test_mount_targets" {
	#Required
	availability_domain = var.mount_target_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.mount_target_display_name
	export_set_id = oci_file_storage_export_set.test_export_set.id
	id = var.mount_target_id
	state = var.mount_target_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `export_set_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export set.
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `mount_targets` - The list of mount_targets.

### MountTarget Reference

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

