---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_event"
sidebar_current: "docs-oci-resource-os_management_hub-event"
description: |-
  Provides the Event resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_event
This resource provides the Event resource in Oracle Cloud Infrastructure Os Management Hub service.

Updates the tags for an event.

## Example Usage

```hcl
resource "oci_os_management_hub_event" "test_event" {
	#Required
	event_id = oci_os_management_hub_event.test_event.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `event_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the event.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `data` - Provides additional information for a management station event.
	* `additional_details` - Provides additional information for the work request associated with an event.
		* `exploit_cves` - List of CVEs in the exploit.
		* `initiator_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that triggered the event, such as scheduled job id.
		* `vmcore` - Kernel event vmcore details
			* `backtrace` - Kernel vmcore backtrace.
			* `component` - Kernel vmcore component.
		* `work_request_ids` - List of all work request [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the event.
	* `content` - Provides information collected for the exploit attempt event.
		* `content_availability` - Crash content availability status:
			* 'NOT_AVAILABLE' indicates the content is not available on the instance nor in the service
			* 'AVAILABLE_ON_INSTANCE' indicates the content is only available on the instance.
			* 'AVAILABLE_ON_SERVICE' indicates the content is only available on the service.
			* 'AVAILABLE_ON_INSTANCE_AND_SERVICE' indicates the content is available both on the instance and the service
			* 'AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS' indicates the content is available on the instance and its upload to the service is in progress. 
		* `content_location` - Location of the Kernel event content.
		* `exploit_detection_log_content` - The content of the exploit detection log.
		* `exploit_object_store_location` - The location of the exploit detection log within object storage.
		* `size` - Size of the event content.
		* `type` - Event type:
			* `KERNEL` - Used to identify a kernel oops/crash content
			* `EXPLOIT_ATTEMPT` - Used to identify a known exploit detection content 
	* `event_fingerprint` - Fingerprint of the event.
	* `event_count` - Number of times the event has occurred.
	* `operation_type` - Type of management station operation.
	* `reason` - Reason for the event.
	* `status` - Status of the management station operation.
	* `time_first_occurred` - The date and time that the event first occurred.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `event_details` - Details of an event.
* `event_summary` - Summary of the event.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the event.
* `is_managed_by_autonomous_linux` - Indicates whether the event occurred on a resource that is managed by the Autonomous Linux service.
* `lifecycle_details` - Describes the current state of the event in more detail. For example, the  message can provide actionable information for a resource in the 'FAILED' state. 
* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance or resource where the event occurred.
* `state` - The current state of the event.
* `system_details` - Provides information about the system architecture and operating system.
	* `architecture` - Architecture type.
	* `ksplice_effective_kernel_version` - Version of the Ksplice effective kernel.
	* `os_family` - Operating system type.
	* `os_kernel_release` - Release of the kernel.
	* `os_kernel_version` - Version of the kernel.
	* `os_name` - Name of the operating system.
	* `os_system_version` - Version of the operating system.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Event was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_occurred` - The date and time that the event occurred.
* `time_updated` - The date and time that the event was updated (in [RFC 3339](https://tools.ietf.org/html/rfc3339) format). Example: `2016-08-25T21:10:29.600Z` 
* `type` - Event type:
	* `KERNEL_OOPS` - Used to identify a kernel panic condition event
	* `KERNEL_CRASH` - Used to identify an internal fatal kernel error that cannot be safely recovered from
	* `EXPLOIT_ATTEMPT` - Used to identify a known exploit detection as identified by Ksplice
	* `SOFTWARE_UPDATE` - Software updates - Packages
	* `KSPLICE_UPDATE` - Ksplice updates
	* `SOFTWARE_SOURCE` - Software source
	* `AGENT` - Agent
	* `MANAGEMENT_STATION` - Management Station 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Event
	* `update` - (Defaults to 20 minutes), when updating the Event
	* `delete` - (Defaults to 20 minutes), when destroying the Event


## Import

Events can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_event.test_event "id"
```

