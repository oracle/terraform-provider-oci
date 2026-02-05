---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_readiness_check"
sidebar_current: "docs-oci-datasource-fleet_software_update-fsu_readiness_check"
description: |-
  Provides details about a specific Fsu Readiness Check in Oracle Cloud Infrastructure Fleet Software Update service
---

# Data Source: oci_fleet_software_update_fsu_readiness_check
This data source provides details about a specific Fsu Readiness Check resource in Oracle Cloud Infrastructure Fleet Software Update service.

Gets a Exadata Fleet Update Readiness Check by identifier.


## Example Usage

```hcl
data "oci_fleet_software_update_fsu_readiness_check" "test_fsu_readiness_check" {
	#Required
	fsu_readiness_check_id = oci_fleet_software_update_fsu_readiness_check.test_fsu_readiness_check.id
}
```

## Argument Reference

The following arguments are supported:

* `fsu_readiness_check_id` - (Required) Unique Exadata Fleet Update Readiness Check identifier. 


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

