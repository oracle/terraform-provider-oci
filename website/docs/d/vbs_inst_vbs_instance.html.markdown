---
subcategory: "Vbs Inst"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vbs_inst_vbs_instance"
sidebar_current: "docs-oci-datasource-vbs_inst-vbs_instance"
description: |-
  Provides details about a specific Visual Builder Studio Instance in Oracle Cloud Infrastructure Vbs Inst service
---

# Data Source: oci_vbs_inst_vbs_instance
This data source provides details about a specific Vbs Instance resource in Oracle Cloud Infrastructure Vbs Inst service.

Gets a VbsInstance by identifier

## Example Usage

```hcl
data "oci_vbs_inst_vbs_instance" "test_vbs_instance" {
	#Required
	vbs_instance_id = oci_vbs_inst_vbs_instance.test_vbs_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `vbs_instance_id` - (Required) unique VbsInstance identifier


## Attributes Reference

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

