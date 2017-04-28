# baremetal\_core\_instance_credentials

Gets the initial user name and password for a Windows instance.

## Example Usage

```
data "baremetal_core_instance_credentials" "s" {
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
