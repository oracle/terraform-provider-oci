---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_registration_policy_target_databases"
sidebar_current: "docs-oci-datasource-data_safe-registration_policy_target_databases"
description: |-
  Provides the list of Registration Policy Target Databases in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_registration_policy_target_databases
This data source provides the list of Registration Policy Target Databases in Oracle Cloud Infrastructure Data Safe service.

Retrieves the OCIDs of target databases registered via the specified registration policy. Supports optional filtering by registration status (OPTIN/OPTOUT) and by target database OCID.


## Example Usage

```hcl
data "oci_data_safe_registration_policy_target_databases" "test_registration_policy_target_databases" {
	#Required
	compartment_id = var.compartment_id
	registration_policy_id = oci_data_safe_registration_policy.test_registration_policy.id

	#Optional
	membership_status = var.registration_policy_target_database_membership_status
	target_database_id = oci_data_safe_target_database.test_target_database.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `membership_status` - (Optional) Filters registered targets associated with this registration policy by membership status.
	* OPTIN: returns targets that are included (opted in) by the policy.
	* OPTOUT: returns targets that are explicitly excluded (opted out) by the policy. 
* `registration_policy_id` - (Required) The OCID of the registration policy to be used for identification
* `target_database_id` - (Optional) A filter to return the target database only if it is registered via the registration policy.


## Attributes Reference

The following attributes are exported:

* `registration_policy_target_database_summary_collection` - The list of registration_policy_target_database_summary_collection.

### RegistrationPolicyTargetDatabase Reference

The following attributes are exported:

* `items` - Array of RegistrationPolicyTargetDatabaseSummary items.
	* `discovered_resource_id` - The ID of the discovered database resource (for example, a Database or Pluggable Database) that is part of discovery.
	* `discovered_resource_type` - The type of the discovered database resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `target_database_id` - The ID of the Data Safe Target Database associated with the discovered resource.

