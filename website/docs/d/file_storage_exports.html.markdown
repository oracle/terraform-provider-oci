---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_exports"
sidebar_current: "docs-oci-datasource-file_storage-exports"
description: |-
  Provides the list of Exports in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_exports
This data source provides the list of Exports in Oracle Cloud Infrastructure File Storage service.

Lists export resources by compartment, file system, or export
set. You must specify an export set ID, a file system ID, and
/ or a compartment ID.


## Example Usage

```hcl
data "oci_file_storage_exports" "test_exports" {

	#Optional
	compartment_id = var.compartment_id
	export_set_id = oci_file_storage_export_set.test_export_set.id
	file_system_id = oci_file_storage_file_system.test_file_system.id
	id = var.export_id
	state = var.export_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `export_set_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export set.
* `file_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `exports` - The list of exports.

### Export Reference

The following attributes are exported:

* `export_options` - Policies that apply to NFS requests made through this export. `exportOptions` contains a sequential list of `ClientOptions`. Each `ClientOptions` item defines the export options that are applied to a specified set of clients.

	For each NFS request, the first `ClientOptions` option in the list whose `source` attribute matches the source IP address of the request is applied.

	If a client source IP address does not match the `source` property of any `ClientOptions` in the list, then the export will be invisible to that client. This export will not be returned by `MOUNTPROC_EXPORT` calls made by the client and any attempt to mount or access the file system through this export will result in an error.

	**Exports without defined `ClientOptions` are invisible to all clients.**

	If one export is invisible to a particular client, associated file systems may still be accessible through other exports on the same or different mount targets. To completely deny client access to a file system, be sure that the client source IP address is not included in any export for any mount target associated with the file system. 
	* `access` - Type of access to grant clients using the file system through this export. If unspecified defaults to `READ_WRITE`. 
	* `allowed_auth` - Array of allowed NFS authentication types.
	* `anonymous_gid` - GID value to remap to when squashing a client GID (see identitySquash for more details.) If unspecified defaults to `65534`. 
	* `anonymous_uid` - UID value to remap to when squashing a client UID (see identitySquash for more details.) If unspecified, defaults to `65534`. 
	* `identity_squash` - Used when clients accessing the file system through this export have their UID and GID remapped to 'anonymousUid' and 'anonymousGid'. If `ALL`, all users and groups are remapped; if `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `ROOT`. 
	* `is_anonymous_access_allowed` - Whether or not to enable anonymous access to the file system through this export in cases where a user isn't found in the LDAP server used for ID mapping. If true, and the user is not found in the LDAP directory, the operation uses the Squash UID and Squash GID. 
	* `require_privileged_source_port` - If `true`, clients accessing the file system through this export must connect from a privileged source port. If unspecified, defaults to `true`. 
	* `source` - Clients these options should apply to. Must be a either single IPv4 address or single IPv4 CIDR block.

		**Note:** Access will also be limited by any applicable VCN security rules and the ability to route IP packets to the mount target. Mount targets do not have Internet-routable IP addresses. 
* `export_set_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this export's export set.
* `file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this export's file system.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this export.
* `is_idmap_groups_for_sys_auth` - Whether or not the export should use ID mapping for Unix groups rather than the group list provided within an NFS request's RPC header. When this flag is true the Unix UID from the RPC header is used to retrieve the list of secondary groups from a the ID mapping subsystem. The primary GID is always taken from the RPC header. If ID mapping is not configured, incorrectly configured, unavailable, or cannot be used to determine a list of secondary groups then an empty secondary group list is used for authorization. If the number of groups exceeds the limit of 256 groups, the list retrieved from LDAP is truncated to the first 256 groups read.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `path` - Path used to access the associated file system.

	Avoid entering confidential information.

	Example: `/accounting` 
* `state` - The current state of this export.
* `time_created` - The date and time the export was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

