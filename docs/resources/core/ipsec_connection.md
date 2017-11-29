# oci\_core\_ipsec\_connection

[IPSecConnection Reference][90077a20]

  [90077a20]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/ "IPSecConnectionReference"

Provide an IPSec connection resource.

## Example Usage

```
resource "oci_core_ipsec" "t" {
    compartment_id = "compartmentid"
    cpe_id = "cpeid"
    drg_id = "drgid"
    display_name = "display_name"
    static_routes = ["route1","route2"]
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `drg_id` - (Required) The OCID of the DRG.
* `cpe_id` - (Required) The OCID of the CPE.
* `static_routes` - (Required) Static routes to the CPE. At least one route must be included. The CIDR must not be a multicast address or class E address.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.


## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `cpe_id` - The OCID of the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `drg_id` - The OCID of the DRG.
* `id` - The IPSec connection's Oracle ID (OCID).
* `state` - The IPSec connection's current state. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `static_routes` - Static routes to the CPE. At least one route must be included.
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
