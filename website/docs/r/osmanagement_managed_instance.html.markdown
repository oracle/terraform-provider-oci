---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance"
sidebar_current: "docs-oci-resource-osmanagement-managed_instance"
description: |-
  Provides the Managed Instance resource in Oracle Cloud Infrastructure OS Management service
---

# oci_osmanagement_managed_instance
This resource provides the Managed Instance resource in Oracle Cloud Infrastructure OS Management service.

Updates a specific Managed Instance.


## Example Usage

```hcl
resource "oci_osmanagement_managed_instance" "test_managed_instance" {
	#Required
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id

	#Optional
	is_data_collection_authorized = var.managed_instance_is_data_collection_authorized
	notification_topic_id = oci_ons_notification_topic.test_notification_topic.id
}
```

## Argument Reference

The following arguments are supported:

* `is_data_collection_authorized` - (Optional) (Updatable) True if user allow data collection for this instance
* `managed_instance_id` - (Required) OCID for the managed instance
* `notification_topic_id` - (Optional) (Updatable) OCID of the ONS topic used to send notification to users


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous` - if present, indicates the Managed Instance is an autonomous instance. Holds all the Autonomous specific information
	* `is_auto_update_enabled` - True if daily updates are enabled
* `bug_updates_available` - Number of bug fix type updates available to be installed
* `child_software_sources` - list of child Software Sources attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `compartment_id` - OCID for the Compartment
* `description` - Information specified by the user about the managed instance
* `display_name` - Managed Instance identifier
* `enhancement_updates_available` - Number of enhancement type updates available to be installed
* `id` - OCID for the managed instance
* `is_data_collection_authorized` - True if user allow data collection for this instance
* `is_reboot_required` - Indicates whether a reboot is required to complete installation of updates.
* `ksplice_effective_kernel_version` - The ksplice effective kernel version
* `last_boot` - Time at which the instance last booted
* `last_checkin` - Time at which the instance last checked in
* `managed_instance_groups` - The ids of the managed instance groups of which this instance is a member. 
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `notification_topic_id` - OCID of the ONS topic used to send notification to users
* `os_family` - The Operating System type of the managed instance.
* `os_kernel_version` - Operating System Kernel Version
* `os_name` - Operating System Name
* `os_version` - Operating System Version
* `other_updates_available` - Number of non-classified updates available to be installed
* `parent_software_source` - the parent (base) Software Source attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `scheduled_job_count` - Number of scheduled jobs associated with this instance
* `security_updates_available` - Number of security type updates available to be installed
* `status` - status of the managed instance.
* `updates_available` - Number of updates available to be installed
* `work_request_count` - Number of work requests associated with this instance

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance


## Import

ManagedInstances can be imported using the `id`, e.g.

```
$ terraform import oci_osmanagement_managed_instance.test_managed_instance "id"
```

