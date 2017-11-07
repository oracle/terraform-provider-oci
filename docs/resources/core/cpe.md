# oci\_core\_cpe

[Cpe Reference][ede04078]

  [ede04078]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/ "CpeReference"

Provide a CPE resource. The `cpe` is a virtual representation of your Customer-Premises Equipment, which is the actual router on-premises at your site at your end of the IPSec VPN connection.

## Example Usage

```
resource "oci_core_cpe" "t" {
    compartment_id = "compartmentid"
    display_name = "displayname"
    ip_address = "123.123.123.123"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `ip_address` - (Required) The public IP address of the on-premises router.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premises router.
* `time_created` - The date and time the CPE was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
