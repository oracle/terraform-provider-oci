---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_registration_policy"
sidebar_current: "docs-oci-datasource-data_safe-registration_policy"
description: |-
  Provides details about a specific Registration Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_registration_policy
This data source provides details about a specific Registration Policy resource in Oracle Cloud Infrastructure Data Safe service.

Returns the details of the specified Registration Policy.


## Example Usage

```hcl
data "oci_data_safe_registration_policy" "test_registration_policy" {
	#Required
	registration_policy_id = oci_data_safe_registration_policy.test_registration_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `registration_policy_id` - (Required) The OCID of the registration policy to be used for identification


## Attributes Reference

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
