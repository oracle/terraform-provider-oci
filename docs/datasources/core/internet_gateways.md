# oci\_core\_internet\_gateways

[InternetGateway Reference][873db635]

  [873db635]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/ "InternetGatewayReference"

Gets a list of Internet Gateways. An Internet Gateway represents a router that connects the edge of a VCN with the Internet.

## Example Usage

```
    data "oci_core_internet_gateways" "s" {
      compartment_id = "compartmentid"
      vcn_id = "vcnid"
    }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `gateways` - The list of internet gateways.

## Internet Gateway reference
* `compartment_id` - The OCID of the compartment containing the internet gateway.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `state` - The route table's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `is_enabled` - Whether the gateway is enabled. When the gateway is disabled, traffic is not routed to/from the Internet, regardless of route rules. Example: `true`
* `time_created` - The date and time the Internet Gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `vcn_id` - The OCID of the VCN the security list belongs to.
