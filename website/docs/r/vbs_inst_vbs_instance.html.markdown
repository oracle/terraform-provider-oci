---
subcategory: "Vbs Inst"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vbs_inst_vbs_instance"
sidebar_current: "docs-oci-resource-vbs_inst-vbs_instance"
description: |-
  Provides the Visual Builder Studio Instance resource in Oracle Cloud Infrastructure Vbs Inst service
---

# oci_vbs_inst_vbs_instance
This resource provides the Vbs Instance resource in Oracle Cloud Infrastructure Vbs Inst service.

Creates a new VbsInstance.


## Example Usage

```hcl
resource "oci_vbs_inst_vbs_instance" "test_vbs_instance" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.vbs_instance_display_name
	name = var.vbs_instance_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	idcs_access_token = var.vbs_instance_idcs_access_token
	is_resource_usage_agreement_granted = var.vbs_instance_is_resource_usage_agreement_granted
	resource_compartment_id = var.resource_compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier. It can only be the root compartment
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Display Name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_access_token` - (Optional) IDCS personal acceess token identifying IDCS user and stripe for the VBS service
* `is_resource_usage_agreement_granted` - (Optional) (Updatable) Whether VBS is authorized to create and use resources in the customer tenancy
* `name` - (Required) Service Instance Name
* `resource_compartment_id` - (Optional) (Updatable) Compartment where VBS may create additional resources for the service instance


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vbs Instance
	* `update` - (Defaults to 20 minutes), when updating the Vbs Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Vbs Instance


## Import

VbsInstances can be imported using the `id`, e.g.

```
$ terraform import oci_vbs_inst_vbs_instance.test_vbs_instance "id"
```

