# oci\_core\_internet\_gateways

[InternetGateway Reference][0162d0a8]

  [0162d0a8]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/ "InternetGatewayReference"

Gets a list of Internet Gateways.

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
* `enabled` - (Optional) Whether the gateway is enabled upon creation. Default is `true`.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the internet gateway.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `state` - The route table's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `enabled` - Whether the gateway is enabled. When the gateway is disabled, traffic is not routed to/from the Internet, regardless of route rules. Example: `true`
* `time_created` - The date and time the Internet Gateway was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `vcn_id` - The OCID of the VCN the Internet Gateway belongs to.
