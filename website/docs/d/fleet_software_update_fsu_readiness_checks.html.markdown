---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_readiness_checks"
sidebar_current: "docs-oci-datasource-fleet_software_update-fsu_readiness_checks"
description: |-
  Provides the list of Fsu Readiness Checks in Oracle Cloud Infrastructure Fleet Software Update service
---

# Data Source: oci_fleet_software_update_fsu_readiness_checks
This data source provides the list of Fsu Readiness Checks in Oracle Cloud Infrastructure Fleet Software Update service.

Returns a list of Exadata Fleet Update Readiness Checks resources in the specified compartment.


## Example Usage

```hcl
data "oci_fleet_software_update_fsu_readiness_checks" "test_fsu_readiness_checks" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.fsu_readiness_check_display_name
	resource_id = oci_cloud_guard_resource.test_resource.id
	state = var.fsu_readiness_check_state
	type = var.fsu_readiness_check_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource related to the Exadata Fleet Update Readiness Check. 
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the specified lifecycleState. 
* `type` - (Optional) A filter to return only resources whose type matches the specified type. 


## Attributes Reference

The following attributes are exported:

* `fsu_readiness_check_collection` - The list of fsu_readiness_check_collection.

### FsuReadinessCheck Reference

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

