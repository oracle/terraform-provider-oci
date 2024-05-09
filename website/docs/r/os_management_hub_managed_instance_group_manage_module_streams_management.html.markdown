---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_manage_module_streams_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group_manage_module_streams_management"
description: |-
  Provides the Managed Instance Group Manage Module Streams Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group_manage_module_streams_management
This resource provides the Managed Instance Group Manage Module Streams Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Enables or disables module streams and installs or removes module stream profiles. Once complete, the state of the modules, streams, and profiles will match the state indicated in the operation. See [ManageModuleStreamsOnManagedInstanceGroupDetails](https://docs.cloud.oracle.com/iaas/api/#/en/osmh/latest/datatypes/ManageModuleStreamsOnManagedInstanceGroupDetails) for more information.
You can preform this operation as a dry run. For a dry run, the service evaluates the operation against the current module, stream, and profile state on the managed instance, but does not commit the changes. Instead, the service returns work request log or error entries indicating the impact of the operation.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group_manage_module_streams_management" "test_managed_instance_group_manage_module_streams_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	disable {
		#Required
		module_name = var.managed_instance_group_manage_module_streams_management_disable_module_name
		stream_name = oci_streaming_stream.test_stream.name

		#Optional
		software_source_id = oci_os_management_hub_software_source.test_software_source.id
	}
	enable {
		#Required
		module_name = var.managed_instance_group_manage_module_streams_management_enable_module_name
		stream_name = oci_streaming_stream.test_stream.name

		#Optional
		software_source_id = oci_os_management_hub_software_source.test_software_source.id
	}
	install {
		#Required
		module_name = var.managed_instance_group_manage_module_streams_management_install_module_name
		profile_name = oci_os_management_hub_profile.test_profile.name
		stream_name = oci_streaming_stream.test_stream.name

		#Optional
		software_source_id = oci_os_management_hub_software_source.test_software_source.id
	}
	is_dry_run = var.managed_instance_group_manage_module_streams_management_is_dry_run
	remove {
		#Required
		module_name = var.managed_instance_group_manage_module_streams_management_remove_module_name
		profile_name = oci_os_management_hub_profile.test_profile.name
		stream_name = oci_streaming_stream.test_stream.name

		#Optional
		software_source_id = oci_os_management_hub_software_source.test_software_source.id
	}
	work_request_details {

		#Optional
		description = var.managed_instance_group_manage_module_streams_management_work_request_details_description
		display_name = var.managed_instance_group_manage_module_streams_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `disable` - (Optional) The set of module streams to disable.
	* `module_name` - (Required) The name of a module.
	* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
	* `stream_name` - (Required) The name of a stream of the specified module.
* `enable` - (Optional) The set of module streams to enable.
	* `module_name` - (Required) The name of a module.
	* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
	* `stream_name` - (Required) The name of a stream of the specified module.
* `install` - (Optional) The set of module stream profiles to install.
	* `module_name` - (Required) The name of a module.
	* `profile_name` - (Required) The name of a profile of the specified module stream.
	* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
	* `stream_name` - (Required) The name of a stream of the specified module.
* `is_dry_run` - (Optional) Indicates if this operation is a dry run or if the operation should be committed.  If set to true, the result of the operation will be evaluated but not committed.  If set to false, the operation is committed to the managed instance(s).  The default is false. 
* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `remove` - (Optional) The set of module stream profiles to remove.
	* `module_name` - (Required) The name of a module.
	* `profile_name` - (Required) The name of a profile of the specified module stream.
	* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
	* `stream_name` - (Required) The name of a stream of the specified module.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group Manage Module Streams Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group Manage Module Streams Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group Manage Module Streams Management


## Import

ManagedInstanceGroupManageModuleStreamsManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group_manage_module_streams_management.test_managed_instance_group_manage_module_streams_management "id"
```

