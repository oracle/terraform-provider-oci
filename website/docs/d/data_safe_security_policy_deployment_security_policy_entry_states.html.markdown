---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_deployment_security_policy_entry_states"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_deployment_security_policy_entry_states"
description: |-
  Provides the list of Security Policy Deployment Security Policy Entry States in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_deployment_security_policy_entry_states
This data source provides the list of Security Policy Deployment Security Policy Entry States in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all security policy entry states in Data Safe.

The ListSecurityPolicyEntryStates operation returns only the security policy entry states for the specified security policy entry.


## Example Usage

```hcl
data "oci_data_safe_security_policy_deployment_security_policy_entry_states" "test_security_policy_deployment_security_policy_entry_states" {
	#Required
	security_policy_deployment_id = oci_data_safe_security_policy_deployment.test_security_policy_deployment.id

	#Optional
	deployment_status = var.security_policy_deployment_security_policy_entry_state_deployment_status
	security_policy_entry_id = oci_data_safe_security_policy_entry.test_security_policy_entry.id
}
```

## Argument Reference

The following arguments are supported:

* `deployment_status` - (Optional) The current state of the security policy deployment.
* `security_policy_deployment_id` - (Required) The OCID of the security policy deployment resource.
* `security_policy_entry_id` - (Optional) An optional filter to return only resources that match the specified security policy entry OCID.


## Attributes Reference

The following attributes are exported:

* `security_policy_entry_state_collection` - The list of security_policy_entry_state_collection.

### SecurityPolicyDeploymentSecurityPolicyEntryState Reference

The following attributes are exported:

* `deployment_status` - The current deployment status of the security policy deployment and the security policy entry associated.
* `entry_details` - Details specific to the security policy entry.
	* `entry_type` - The security policy entry type. Allowed values:
		* FIREWALL_POLICY - The SQL Firewall policy entry type. 
	* `time_generated` - The time the the SQL Firewall policy was generated on the target database, in the format defined by RFC3339.
	* `time_status_updated` - The last date and time the status of the SQL Firewall policy was updated on the target database, in the format defined by RFC3339.
* `id` - Unique id of the security policy entry state.
* `security_policy_deployment_id` - The OCID of the security policy deployment associated.
* `security_policy_entry_id` - The OCID of the security policy entry type associated.

