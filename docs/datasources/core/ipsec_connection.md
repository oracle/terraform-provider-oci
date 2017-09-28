# oci\_core\_ipsec\_connection

[IPSecConnection Reference][70301f58]

  [70301f58]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/ "IPSecConnectionReference"

Gets a list of IPSec connections.

## Example Usage

```
data "oci_core_ipsec_connections" "s" {
  compartment_id = "compartmentid"
  cpe_id = "cpeid"
  drg_id = "drgid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `drg_id` - (Required) The OCID of the DRG.
* `cpe_id` - (Required) The OCID of the CPE.
* `limit` - (Required) The maximum number of items to return in a paginated "List" call.
* `page` - (Required) The page number to fetch.


## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `cpe_id` - The OCID of the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `drg_id` - The OCID of the DRG.
* `id` - The IPSec connection's Oracle ID (OCID).
* `state` - The IPSec connection's current state. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `static_routes` - Static routes to the CPE. At least one route must be included.
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
