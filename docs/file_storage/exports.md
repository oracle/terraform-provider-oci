# oci_file_storage_export

## Export Resource

### Export Reference

The following attributes are exported:

* `export_options` - Policies that apply to NFS requests made through this export. `exportOptions` contains a sequential list of `ClientOptions`. Each `ClientOptions` item defines the export options that are applied to a specified set of clients.

	For each NFS request, the first `ClientOptions` option in the list whose `source` attribute matches the source IP address of the request is applied.

	If a client source IP address does not match the `source` property of any `ClientOptions` in the list, then the export will be invisible to that client. This export will not be returned by `MOUNTPROC_EXPORT` calls made by the client and any attempt to mount or access the file system through this export will result in an error.

	**Exports without defined `ClientOptions` are invisible to all clients.**

	If one export is invisible to a particular client, associated file systems may still be accessible through other exports on the same or different mount targets. To completely deny client access to a file system, be sure that the client source IP address is not included in any export for any mount target associated with the file system. 
	* `access` - Type of access to grant clients using the file system through this export. If unspecified defaults to `READ_ONLY`. 
	* `anonymous_gid` - GID value to remap to when squashing a client GID (see identitySquash for more details.) If unspecified defaults to `65534`. 
	* `anonymous_uid` - UID value to remap to when squashing a client UID (see identitySquash for more details.) If unspecified, defaults to `65534`. 
	* `identity_squash` - Used when clients accessing the file system through this export have their UID and GID remapped to 'anonymousUid' and 'anonymousGid'. If `ALL`, all users and groups are remapped; if `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `ROOT`. 
	* `require_privileged_source_port` - If `true`, clients accessing the file system through this export must connect from a privileged source port. If unspecified, defaults to `true`. 
	* `source` - Clients these options should apply to. Must be a either single IPv4 address or single IPv4 CIDR block.

		**Note:** Access will also be limited by any applicable VCN security rules and the ability to route IP packets to the mount target. Mount targets do not have Internet-routable IP addresses. 
* `export_set_id` - The OCID of this export's export set.
* `file_system_id` - The OCID of this export's file system.
* `id` - The OCID of this export.
* `path` - Path used to access the associated file system.

	Avoid entering confidential information.  Example: `/accounting` 
* `state` - The current state of this export.
* `time_created` - The date and time the export was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new export in the specified export set, path, and
file system.


The following arguments are supported:

* `export_options` - (Optional) Export options for the new export. If left unspecified, defaults to:

	[          {             "source" : "0.0.0.0/0",             "requirePrivilegedSourcePort" : false,             "access" : "READ_WRITE",             "identitySquash" : "NONE"           }        ]

	**Note:** Mount targets do not have Internet-routable IP   addresses.  Therefore they will not be reachable from the   Internet, even if an associated `ClientOptions` item has   a source of `0.0.0.0/0`.

	**If set to the empty array then the export will not be   visible to any clients.**

	The export's `exportOptions` can be changed after creation   using the `UpdateExport` operation. 
	* `access` - (Optional) Type of access to grant clients using the file system through this export. If unspecified defaults to `READ_ONLY`. 
	* `anonymous_gid` - (Optional) GID value to remap to when squashing a client GID (see identitySquash for more details.) If unspecified defaults to `65534`. 
	* `anonymous_uid` - (Optional) UID value to remap to when squashing a client UID (see identitySquash for more details.) If unspecified, defaults to `65534`. 
	* `identity_squash` - (Optional) Used when clients accessing the file system through this export have their UID and GID remapped to 'anonymousUid' and 'anonymousGid'. If `ALL`, all users and groups are remapped; if `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `ROOT`. 
	* `require_privileged_source_port` - (Optional) If `true`, clients accessing the file system through this export must connect from a privileged source port. If unspecified, defaults to `true`. 
	* `source` - (Required) Clients these options should apply to. Must be a either single IPv4 address or single IPv4 CIDR block.

		**Note:** Access will also be limited by any applicable VCN security rules and the ability to route IP packets to the mount target. Mount targets do not have Internet-routable IP addresses. 
* `export_set_id` - (Required) The OCID of this export's export set.
* `file_system_id` - (Required) The OCID of this export's file system.
* `path` - (Required) Path used to access the associated file system.

	Avoid entering confidential information.  Example: `/mediafiles` 


### Update Operation
Updates the specified export's information.

The following arguments support updates:
* `export_options` - Export options for the new export. If left unspecified, defaults to:

	[          {             "source" : "0.0.0.0/0",             "requirePrivilegedSourcePort" : false,             "access" : "READ_WRITE",             "identitySquash" : "NONE"           }        ]

	**Note:** Mount targets do not have Internet-routable IP   addresses.  Therefore they will not be reachable from the   Internet, even if an associated `ClientOptions` item has   a source of `0.0.0.0/0`.

	**If set to the empty array then the export will not be   visible to any clients.**

	The export's `exportOptions` can be changed after creation   using the `UpdateExport` operation. 
	* `access` - Type of access to grant clients using the file system through this export. If unspecified defaults to `READ_ONLY`. 
	* `anonymous_gid` - GID value to remap to when squashing a client GID (see identitySquash for more details.) If unspecified defaults to `65534`. 
	* `anonymous_uid` - UID value to remap to when squashing a client UID (see identitySquash for more details.) If unspecified, defaults to `65534`. 
	* `identity_squash` - Used when clients accessing the file system through this export have their UID and GID remapped to 'anonymousUid' and 'anonymousGid'. If `ALL`, all users and groups are remapped; if `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `ROOT`. 
	* `require_privileged_source_port` - If `true`, clients accessing the file system through this export must connect from a privileged source port. If unspecified, defaults to `true`. 
	* `source` - Clients these options should apply to. Must be a either single IPv4 address or single IPv4 CIDR block.

		**Note:** Access will also be limited by any applicable VCN security rules and the ability to route IP packets to the mount target. Mount targets do not have Internet-routable IP addresses. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_file_storage_export" "test_export" {
	#Required
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	path = "${var.export_path}"

	#Optional
	export_options {
		#Required
		source = "${var.export_export_options_source}"

		#Optional
		access = "${var.export_export_options_access}"
		anonymous_gid = "${var.export_export_options_anonymous_gid}"
		anonymous_uid = "${var.export_export_options_anonymous_uid}"
		identity_squash = "${var.export_export_options_identity_squash}"
		require_privileged_source_port = "${var.export_export_options_require_privileged_source_port}"
	}
}
```

# oci_file_storage_exports

## Export DataSource

Gets a list of exports.

### List Operation
Lists export resources by compartment, file system, or export
set. You must specify an export set ID, a file system ID, and
/ or a compartment ID.

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment.
* `export_set_id` - (Optional) The OCID of the export set.
* `file_system_id` - (Optional) The OCID of the file system.
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


The following attributes are exported:

* `exports` - The list of exports.

### Example Usage

```hcl
data "oci_file_storage_exports" "test_exports" {

	#Optional
	compartment_id = "${var.compartment_id}"
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	id = "${var.export_id}"
	state = "${var.export_state}"
}
```
