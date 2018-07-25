---
layout: "oci"
page_title: "OCI: oci_identity_policies"
sidebar_current: "docs-oci-datasource-identity-policies"
description: |-
  Provides a list of Policies
---

# Data Source: oci_identity_policies
The `oci_identity_policies` data source allows access to the list of OCI policies

Lists the policies in the specified compartment (either the tenancy or another of your compartments).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).

To determine which policies apply to a particular group or compartment, you must view the individual
statements inside all your policies. There isn't a way to automatically obtain that information via the API.


## Example Usage

```hcl
data "oci_identity_policies" "test_policies" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `policies` - The list of policies.

### Policy Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the policy (either the tenancy or another compartment). 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the policy. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the policy.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the policy during creation. The name must be unique across all policies in the tenancy and cannot be changed. 
* `state` - The policy's current state. After creating a policy, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `statements` - An array of one or more policy statements written in the policy language.
* `time_created` - Date and time the policy was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `version_date` - The version of the policy. If null or set to an empty string, when a request comes in for authorization, the policy will be evaluated according to the current behavior of the services at that moment. If set to a particular date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date. 

