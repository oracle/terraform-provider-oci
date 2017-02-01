# baremetal\_core\_ipsec\_connection

Provide an ipsec connection resource.

## Example Usage

```
resource "baremetal_core_ipsec" "t" {
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
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.


## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `cpe_id` - The OCID of the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `drg_id` - The OCID of the DRG.
* `id` - The IPSec connection's Oracle ID (OCID).
* `state` - The IPSec connection's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `static_routes` - Static routes to the CPE. At least one route must be included.
* `time_created` - The date and time the IPSec connection was created.
