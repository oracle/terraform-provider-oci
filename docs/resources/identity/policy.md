# oci\_identity\_policy

Provides a policy resource.

## Example Usage

```
  resource "oci_identity_policy" "p" {
    compartment_id = <Compartment or Tenancy OCID>
    name = "pol"
    description = "desc"
    statements = ["statementX","statementY"]
  }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the policy (either the tenancy or another compartment).
* `name` - (Required) The name you assign to the policy during creation. The name must be unique across all policies in the tenancy and cannot be changed.
* `statements` - (Required) An array of policy statements written in the policy language.
* `descriptions` - (Required) The description you assign to the policy during creation. Does not have to be unique, and it's changeable.


## Attributes Reference
* `id` - The OCID of the policy.
* `compartment_id` - The OCID of the compartment containing the policy (either the tenancy or another compartment).
* `name` - The name you assign to the policy during creation. The name must be unique across all policies in the tenancy and cannot be changed.
* `statements` - An array of one or more policy statements written in the policy language.
* `descriptions` - The description you assign to the policy during creation. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the policy was created.
* `state` - The policy's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
* `version_date` - The version of the policy. If null or set to an empty string, when a request comes in for authorization, the policy will be evaluated according to the current behavior of the services at that moment. If set to a particular date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date.
