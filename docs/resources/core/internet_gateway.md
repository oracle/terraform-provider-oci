# oci\_core\_internet\_gateways

Gets a list of internet gateways.

## Example Usage

```
resource "oci_core_internet_gateway" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    vcn_id = "vcnid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `is_enabled` - (Optional) Whether the gateway is enabled upon creation. Default false.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the internet gateway.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `state` - The route table's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `is_enabled` - Whether the gateway is enabled. When the gateway is disabled, traffic is not routed to/from the Internet, regardless of route rules.
* `time_created` - The date and time the security list was created.
* `vcn_id` - The OCID of the VCN the security list belongs to.
