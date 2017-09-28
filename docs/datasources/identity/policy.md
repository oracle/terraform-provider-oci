# oci\_identity\_policies

[Policy Reference][1418777c]

  [1418777c]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/ "PolicyReference"

Lists policies. A policy is a document that specifies the type of access a group has to the resources in a compartment.

## Example Usage

```
  data "oci_identity_policies" "p" {
    compartment_id = "compartment ocid"
  }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `policies` - A list of policies.

## Policy Reference
* `id` - The OCID of the policy.
* `compartment_id` - The OCID of the compartment containing the policy (either the tenancy or another compartment).
* `name` - The name you assign to the policy during creation. The name must be unique across all policies in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - The description you assign to the policy. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `time_created` - Date and time the policy was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `state` - The group's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
* `version_date` - The version of the policy. If null or set to an empty string, when a request comes in for authorization, the policy will be evaluated according to the current behavior of the services at that moment. If set to a particular date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date.
