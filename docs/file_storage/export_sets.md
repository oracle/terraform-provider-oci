# oci_file_storage_export_set

## ExportSet Resource

### ExportSet Reference

The following attributes are exported:

* `availability_domain` - The availability domain the export set is in. May be unset as a blank or NULL value.  Example: `kIdk:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the export set.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My export set` 
* `id` - The OCID of the export set.
* `max_fs_stat_bytes` - Controls the maximum `tbytes`, `fbytes`, and `abytes`, values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tbytes` value reported by `FSSTAT` will be `maxFsStatBytes`. The value of `fbytes` and `abytes` will be `maxFsStatBytes` minus the metered size of the file system. If the metered size is larger than `maxFsStatBytes`, then `fbytes` and `abytes` will both be '0'. 
* `max_fs_stat_files` - Controls the maximum `tfiles`, `ffiles`, and `afiles` values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tfiles` value reported by `FSSTAT` will be `maxFsStatFiles`. The value of `ffiles` and `afiles` will be `maxFsStatFiles` minus the metered size of the file system. If the metered size is larger than `maxFsStatFiles`, then `ffiles` and `afiles` will both be '0'. 
* `state` - The current state of the export set.
* `time_created` - The date and time the export set was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the virtual cloud network (VCN) the export set is in.



### Create Operation

The export set resource can neither be directly created, nor destroyed.

An export set is created by the service automatically when a mount target is created.
When a mount target is deleted, the export set associated with it is also deleted automatically.

However, export sets expose a few attributes that can be updated.

Hence we provide this resource for managing the already created export set from within Terraform.

The following arguments are supported:

* `mount_target_id` - (Required) The OCID of the mount target that the export set is associated with
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My export set` 
* `max_fs_stat_bytes` - Controls the maximum `tbytes`, `fbytes`, and `abytes`, values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tbytes` value reported by `FSSTAT` will be `maxFsStatBytes`. The value of `fbytes` and `abytes` will be `maxFsStatBytes` minus the metered size of the file system. If the metered size is larger than `maxFsStatBytes`, then `fbytes` and `abytes` will both be '0'. 
* `max_fs_stat_files` - Controls the maximum `tfiles`, `ffiles`, and `afiles` values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tfiles` value reported by `FSSTAT` will be `maxFsStatFiles`. The value of `ffiles` and `afiles` will be `maxFsStatFiles` minus the metered size of the file system. If the metered size is larger than `maxFsStatFiles`, then `ffiles` and `afiles` will both be '0'. 


### Update Operation
Updates the specified export set's information.

The following arguments support updates:
* `mount_target_id` - (Required) The OCID of the mount target that the export set is associated with
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My export set` 
* `max_fs_stat_bytes` - Controls the maximum `tbytes`, `fbytes`, and `abytes`, values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tbytes` value reported by `FSSTAT` will be `maxFsStatBytes`. The value of `fbytes` and `abytes` will be `maxFsStatBytes` minus the metered size of the file system. If the metered size is larger than `maxFsStatBytes`, then `fbytes` and `abytes` will both be '0'. 
* `max_fs_stat_files` - Controls the maximum `tfiles`, `ffiles`, and `afiles` values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tfiles` value reported by `FSSTAT` will be `maxFsStatFiles`. The value of `ffiles` and `afiles` will be `maxFsStatFiles` minus the metered size of the file system. If the metered size is larger than `maxFsStatFiles`, then `ffiles` and `afiles` will both be '0'. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_file_storage_export_set" "test_export_set" {
    #Required
    mount_target_id = "${oci_file_storage_mount_target.test_mount_target.id}"
  
    #Optional
    display_name = "${var.export_set_name}"
    max_fs_stat_bytes = 23843202333
    max_fs_stat_files = 223442
}
```

# oci_file_storage_export_sets

## ExportSet DataSource

Gets a list of export_sets.

### List Operation
Lists the export set resources in the specified compartment.

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `kIdk:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


The following attributes are exported:

* `export_sets` - The list of export_sets.

### Example Usage

```hcl
data "oci_file_storage_export_sets" "test_export_sets" {
	#Required
	availability_domain = "${var.export_set_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.export_set_display_name}"
	id = "${var.export_set_id}"
	state = "${var.export_set_state}"
}
```