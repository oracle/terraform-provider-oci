---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_dynamic_set"
sidebar_current: "docs-oci-datasource-os_management_hub-dynamic_set"
description: |-
  Provides details about a specific Dynamic Set in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_dynamic_set
This data source provides details about a specific Dynamic Set resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified dynamic set.

## Example Usage

```hcl
data "oci_os_management_hub_dynamic_set" "test_dynamic_set" {
	#Required
	dynamic_set_id = oci_os_management_hub_dynamic_set.test_dynamic_set.id
}
```

## Argument Reference

The following arguments are supported:

* `dynamic_set_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. This filter returns resources associated with this dynamic set.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the dynamic set. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified description for the dynamic set.
* `display_name` - User-friendly name for the dynamic set.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. 
* `match_type` - Include either any or all attributes.
* `matching_rule` - An object that defines the set of rules that identifies the target instances in a dynamic set.
	* `architectures` - The list of managed instance architectures.
	* `display_names` - The list of managed instance display names.
	* `is_reboot_required` - Indicates if the managed instance needs to be rebooted.
	* `locations` - The list of managed instance locations.
	* `managed_instance_group_ids` - The list of managed instance group IDs.
	* `managed_instance_ids` - The list of managed instance ids.
	* `os_families` - The list of managed instance OS families.
	* `os_names` - The list of managed instance OS names.
	* `statuses` - The list of managed instance statuses.
	* `tags` - The list of the managed instance tags.
		* `key` - The key of the tag. 
		* `namespace` - The namespace of the tag. 
		* `type` - The type of the tag. Common values include `defined` or `freeform`. 
		* `value` - The value associated with the tag key. 
* `scheduled_job_count` - Number of scheduled jobs currently targeting this dynamic set.
* `state` - The current state of the event.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_compartments` - The list of compartment details.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `does_include_children` - Indicates if the child compartments are included in the matching rule. 
* `time_created` - The date and time the dynamic set was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `time_updated` - The date and time the dynamic set was last updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 

