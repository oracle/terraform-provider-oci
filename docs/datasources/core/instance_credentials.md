# oci\_core\_instance_credentials

[InstanceCredentials Reference][5d7b7cd3]

  [5d7b7cd3]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InstanceCredentials/ "InstanceCredentialsReference"

Gets the initial user name and password for a Windows instance.

## Example Usage

```
data "oci_core_instance_credentials" "s" {
    instance_id = "instanceId"
}
```

## Argument Reference

The following argument is supported:

* `instance_id` - (Required) The OCID of the instance.


## Attributes Reference

The following attributes are exported:

* `username` - The administrator username for the Windows instance.
* `password` - The initial password for the Windows instance.

## Instance Credential Reference
* `username` - The username.
* `password` - The password for the username.
