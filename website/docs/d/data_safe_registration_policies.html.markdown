---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_registration_policies"
sidebar_current: "docs-oci-datasource-data_safe-registration_policies"
description: |-
  Provides the list of Registration Policies in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_registration_policies
This data source provides the list of Registration Policies in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of registration policies according to the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_registration_policies" "test_registration_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.registration_policy_access_level
	compartment_id_in_subtree = var.registration_policy_compartment_id_in_subtree
	connection_id = oci_database_migration_connection.test_connection.id
	connection_type = var.registration_policy_connection_type
	display_name = var.registration_policy_display_name
	enablement_level = var.registration_policy_enablement_level
	registration_policy_id = oci_data_safe_registration_policy.test_registration_policy.id
	resource_id = oci_cloud_guard_resource.test_resource.id
	state = var.registration_policy_state
	time_created_greater_than_or_equal_to = var.registration_policy_time_created_greater_than_or_equal_to
	time_created_less_than = var.registration_policy_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
* `connection_id` - (Optional) Filter to return the registration policies matching the specified connection ID.
* `connection_type` - (Optional) Filter to return the registration policies matching the specified connectionType i.e ONPREM_CONNECTOR or PRIVATE_ENDPOINT.
* `display_name` - (Optional) A filter to return only resources that match the specified display name.
* `enablement_level` - (Optional) Filter registration policies by resource type.
* `registration_policy_id` - (Optional) Filter to return the registration policy matching the specified OCID.
* `resource_id` - (Optional) Filter to return the registration policy matching the specified resource OCID.
* `state` - (Optional) Filter registration policies by their lifecycle state.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z


## Attributes Reference

The following attributes are exported:

* `registration_policy_collection` - The list of registration_policy_collection.

### RegistrationPolicy Reference

The following attributes are exported:

* `can_override_features` - Indicates whether features will be overridden.
* `compartment_id` - The OCID for the compartment containing the registration policy.
* `connection_option` - Types of connection supported by Data Safe.
	* `connection_type` - The connection type used to connect to the database. Allowed values:
		* PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
		* ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe.
	* `identifiers` - List of OCIDs required to establish the connection.
		* For `PRIVATE_ENDPOINT`, provide the OCID(s) of Data Safe private endpoint(s).
		* For `ONPREM_CONNECTOR`, provide the OCID(s) of on-premises connector(s).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}`
* `description` - A description of the registration policy.
* `display_name` - The display name of the registration policy.
* `enablement_level` - The resource type which has been opted in for the registration policy.
	* DATABASE - The registration policy will be opted in at the Container Database level.
* `features` - The Data Safe features granted to the databases registering under the registration policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the registration policy.
* `lifecycle_state_details` - Details about the lifecycle state of the registration policy
* `resource_id` - The OCID of the resource used in the registration policy.
* `state` - The lifecycle state of the registration policy.
	* CREATING - The registration policy is getting created.
	* ACTIVE  - The registration policy has been successfully created.
	* UPDATING - The registration policy is getting updated.
	* NEEDS_ATTENTION - The registration policy needs attention.
	* FAILED - The registration policy creation or update has failed.
	* DELETING - The registration policy is getting deleted.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time when the registration policy was created.
* `time_updated` - The date and time when the registration policy was last updated.
