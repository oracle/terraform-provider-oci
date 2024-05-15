---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_events"
sidebar_current: "docs-oci-datasource-os_management_hub-events"
description: |-
  Provides the list of Events in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_events
This data source provides the list of Events in Oracle Cloud Infrastructure Os Management Hub service.

Lists events that match the specified criteria, such as compartment, state, and event type.


## Example Usage

```hcl
data "oci_os_management_hub_events" "test_events" {

	#Optional
	compartment_id = var.compartment_id
	event_fingerprint = var.event_event_fingerprint
	event_summary = var.event_event_summary
	event_summary_contains = var.event_event_summary_contains
	id = var.event_id
	is_managed_by_autonomous_linux = var.event_is_managed_by_autonomous_linux
	resource_id = oci_usage_proxy_resource.test_resource.id
	state = var.event_state
	time_created_greater_than_or_equal_to = var.event_time_created_greater_than_or_equal_to
	time_created_less_than = var.event_time_created_less_than
	type = var.event_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `event_fingerprint` - (Optional) The eventFingerprint of the KernelEventData.
* `event_summary` - (Optional) A filter to return only events whose summary matches the given value.
* `event_summary_contains` - (Optional) A filter to return only events with a summary that contains the value provided.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the event.
* `is_managed_by_autonomous_linux` - (Optional) Indicates whether to list only resources managed by the Autonomous Linux service. 
* `resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. This filter returns resources associated with the specified resource.
* `state` - (Optional) A filter to return only events that match the state provided. The state value is case-insensitive. 
* `time_created_greater_than_or_equal_to` - (Optional) A filter that returns events that occurred on or after the date provided.       Example: `2016-08-25T21:10:29.600Z` 
* `time_created_less_than` - (Optional) A filter that returns events that occurred on or before the date provided.       Example: `2016-08-25T21:10:29.600Z` 
* `type` - (Optional) A filter to return only resources whose type matches the given value.


## Attributes Reference

The following attributes are exported:

* `event_collection` - The list of event_collection.

### Event Reference

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

