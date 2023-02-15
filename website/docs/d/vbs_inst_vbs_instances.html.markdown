---
subcategory: "Vbs Inst"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vbs_inst_vbs_instances"
sidebar_current: "docs-oci-datasource-vbs_inst-vbs_instances"
description: |-
  Provides the list of Visual Builder Studio Instances in Oracle Cloud Infrastructure Vbs Inst service
---

# Data Source: oci_vbs_inst_vbs_instances
This data source provides the list of Vbs Instances in Oracle Cloud Infrastructure Vbs Inst service.

Returns a list of VbsInstances.


## Example Usage

```hcl
data "oci_vbs_inst_vbs_instances" "test_vbs_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.vbs_instance_id
	name = var.vbs_instance_name
	state = var.vbs_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `id` - (Optional) unique VbsInstance identifier
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `vbs_instance_summary_collection` - The list of vbs_instance_summary_collection.

### VbsInstance Reference

The following attributes are exported:

* `compartment_id` - Compartment of the service instance
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Service instance display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `is_resource_usage_agreement_granted` - Whether the VBS service instance owner explicitly approved VBS to create and use resources in the customer tenancy
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - Service instance name (unique identifier)
* `resource_compartment_id` - Compartment where VBS may create additional resources for the service instance
* `state` - The current state of the VbsInstance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the VbsInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the VbsInstance was updated. An RFC3339 formatted datetime string
* `vbs_access_url` - Public web URL for accessing the VBS service instance

