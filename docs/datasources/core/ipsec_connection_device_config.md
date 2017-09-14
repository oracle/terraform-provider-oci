# oci\_core\_ipsec\_config

Gets a list of internet gateways.

## Example Usage

```
data "oci_core_ipsec_config" "s" {
  ipsec_id = "ipsecid"
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The OCID of the IPSec connection.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `id` - The IPSec connection's Oracle ID (OCID).
* `time_created` - The date and time the IPSec connection was created.
* `tunnels` - A list of tunnel objects
