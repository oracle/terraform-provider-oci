# oci\_core\_ipsec\_status

**API:** [IPSecConnectionDeviceStatus Reference][6c32ac75]

  [6c32ac75]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceStatus/ "IPSecConnectionDeviceStatusReference"

~~Gets a list of internet gateways.~~ Gets information about the status of the IPSec connection.

## Example Usage

```
data "oci_core_ipsec_status" "s" {
  ipsec_id = "ipsecid"
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The OCID of the IPSec connection.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `id` - The IPSec connection's Oracle ID (OCID).
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `tunnels` - A list of tunnel statuses.
