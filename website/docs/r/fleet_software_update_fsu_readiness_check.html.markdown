---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_readiness_check"
sidebar_current: "docs-oci-resource-fleet_software_update-fsu_readiness_check"
description: |-
  Provides the Fsu Readiness Check resource in Oracle Cloud Infrastructure Fleet Software Update service
---

# oci_fleet_software_update_fsu_readiness_check
This resource provides the Fsu Readiness Check resource in Oracle Cloud Infrastructure Fleet Software Update service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/edsfu/latest/FsuReadinessCheck

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleetsoftwareupdate

Creates a new Exadata Fleet Update Readiness Check.


## Example Usage

```hcl
resource "oci_fleet_software_update_fsu_readiness_check" "test_fsu_readiness_check" {
	#Required
	compartment_id = var.compartment_id
	type = var.fsu_readiness_check_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.fsu_readiness_check_display_name
	freeform_tags = {"bar-key"= "value"}
	targets {
		#Required
		entity_id = oci_fleet_software_update_entity.test_entity.id
		entity_type = var.fsu_readiness_check_targets_entity_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Exadata Fleet Update Readiness Check resource. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `targets` - (Optional) List of targets that will run the Exadata Fleet Update Readiness Check. The targets have to be of the same entity type. 
	* `entity_id` - (Required) Resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) 
	* `entity_type` - (Required) Resource entity type 
* `type` - (Required) Type of Exadata Fleet Update Readiness Check. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The user-friendly name for the Exadata Fleet Update Readiness Check resource. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Readiness Check. 
* `issue_count` - Number of issues found during the Exadata Fleet Update Readiness Check run. 
* `issues` - Issues found during the Exadata Fleet Update Readiness Check run. 
	* `description` - Description of the patching issue. 
	* `impacted_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource impacted by the patching issue. 
	* `name` - Name of the patching issue. 
	* `recommended_action` - Recommended action to perform to address or further triage the patching issue. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `state` - Possible lifecycle states for the Exadata Fleet Update Readiness Check resource. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `targets` - List of targets that will run the Exadata Fleet Update Readiness Check. The targets have to be of the same entity type. 
	* `entity_id` - Resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) 
	* `entity_type` - Resource entity type 
* `time_created` - The date and time the Exadata Fleet Update Readiness Check was created, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `time_finished` - The date and time the Exadata Fleet Update Readiness Check was finished, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The date and time the Exadata Fleet Update Readiness Check was updated, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `type` - Possible Exadata Fleet Update Readiness Check types. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fsu Readiness Check
	* `update` - (Defaults to 20 minutes), when updating the Fsu Readiness Check
	* `delete` - (Defaults to 20 minutes), when destroying the Fsu Readiness Check


## Import

FsuReadinessChecks can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_software_update_fsu_readiness_check.test_fsu_readiness_check "id"
```

