# oci\_identity\_api\_keys

[ApiKey Reference][5c500506]

  [5c500506]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/ApiKey/ "ApiKeyReference"

Lists API keys.

## Example Usage

```
data "oci_identity_api_keys" "t" {
  user_id = "user_id"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.

## Attribute Reference
* `api_keys` - A list of API keys.

## ApiKey Reference
* `key_id` - An Oracle-assigned identifier for the key, in this format: TENANCY_OCID/USER_OCID/KEY_FINGERPRINT.
* `key_value` - The key's value.
* `fingerprint` - The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
* `user_id` - The OCID of the user the key belongs to.
* `time_created` - Date and time the ApiKey was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `state` - The compartment's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
