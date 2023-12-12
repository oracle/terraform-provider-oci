---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_deployment_security_policy_entry_state"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_deployment_security_policy_entry_state"
description: |-
  Provides details about a specific Security Policy Deployment Security Policy Entry State in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_deployment_security_policy_entry_state
This data source provides details about a specific Security Policy Deployment Security Policy Entry State resource in Oracle Cloud Infrastructure Data Safe service.

Gets a security policy entity states by identifier.

## Example Usage

```hcl
data "oci_data_safe_security_policy_deployment_security_policy_entry_state" "test_security_policy_deployment_security_policy_entry_state" {
	#Required
	security_policy_deployment_id = oci_data_safe_security_policy_deployment.test_security_policy_deployment.id
	security_policy_entry_state_id = oci_data_safe_security_policy_entry_state.test_security_policy_entry_state.id
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_deployment_id` - (Required) The OCID of the security policy deployment resource.
* `security_policy_entry_state_id` - (Required) Unique security policy entry state identifier. The unique id for a given security policy entry state can be obtained  from the list api by passing the OCID of the corresponding  security policy deployment resource as the query parameter. 


## Attributes Reference

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

