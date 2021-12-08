---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet"
sidebar_current: "docs-oci-resource-jms-fleet"
description: |-
  Provides the Fleet resource in Oracle Cloud Infrastructure Jms service
---

# oci_jms_fleet
This resource provides the Fleet resource in Oracle Cloud Infrastructure Jms service.

Create a new Fleet using the information provided.


## Example Usage

```hcl
resource "oci_jms_fleet" "test_fleet" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.fleet_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.fleet_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `description` - (Optional) (Updatable) The Fleet's description. If nothing is provided, the Fleet description will be null.
* `display_name` - (Required) (Updatable) The name of the Fleet. The displayName must be unique for Fleets in the same compartment.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `approximate_application_count` - The approximate count of all unique applications in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag. 
* `approximate_installation_count` - The approximate count of all unique Java installations in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag. 
* `approximate_jre_count` - The approximate count of all unique Java Runtimes in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag. 
* `approximate_managed_instance_count` - The approximate count of all unique managed instances in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `description` - The Fleet's description.
* `display_name` - The name of the Fleet.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `state` - The lifecycle state of the Fleet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The creation date and time of the Fleet (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fleet
	* `update` - (Defaults to 20 minutes), when updating the Fleet
	* `delete` - (Defaults to 20 minutes), when destroying the Fleet


## Import

Fleets can be imported using the `id`, e.g.

```
$ terraform import oci_jms_fleet.test_fleet "id"
```

