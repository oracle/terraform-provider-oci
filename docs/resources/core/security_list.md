# oci\_core\_security\_lists

[SecurityList Reference][b38fec4c]

  [b38fec4c]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/ "SecurityListReference"

Provides a security list resource.
See the [Security Lists](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm)
overview for more information.

For more information on configuring a VCN's default security list, 
see [Managing Default VCN Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Managing%20Default%20Resources.md)

## Example Usage

Protocols are specified as protocol numbers. For information about protocol numbers, see
http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml

```
resource "oci_core_security_list" "t" {
    compartment_id = "compartment_id"
    vcn_id = "vcn_id"
    display_name = "display_name"

    egress_security_rules {
        protocol = "1"
        destination = "0.0.0.0/0"

        icmp_options {
            "type" = 3
            "code" = 4
        }
    }

    ingress_security_rules {
        protocol = "6"
        source = "0.0.0.0/0"
        stateless = true

        tcp_options {
            "min" = 80
            "max" = 82
        }
    }

    ingress_security_rules {
        protocol = "17"
        source = "0.0.0.0/0"
        stateless = true

        udp_options {
            "min" = 319
            "max" = 320
        }
    }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the security list.
* `vcn_id` - (Required) The OCID of the VCN the security list belongs to.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `egress_security_rules` - (Required) Rules for allowing egress IP packets. [EgressSecurityRule API Docs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/EgressSecurityRule/)
* `ingress_security_rules` - (Required) Rules for allowing ingress IP packets. [IngressSecurityRule API Docs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IngressSecurityRule/)

## Attributes Reference

* `compartment_id` - The OCID of the compartment containing the security list.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `egress_security_rules` - Rules for allowing egress IP packets.
* `id` - The security list's Oracle Cloud ID (OCID).
* `ingress_security_rules` - Rules for allowing ingress IP packets.
* `state` - The security list's current state. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the security list was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `vcn_id` - The OCID of the VCN the security list belongs to.
