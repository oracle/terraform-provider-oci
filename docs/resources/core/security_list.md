# baremetal\_core\_security\_lists

Provides a security list resource.
See the [Security Lists](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm)
overview for more information

## Example Usage

Protocols are specified as protocol numbers. For protocol numbers see
http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml

```
resource "baremetal_core_security_list" "t" {
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
        stateful = true

        tcp_options {
            "min" = 80
            "max" = 82
        }
    }

    ingress_security_rules {
        protocol = "17"
        source = "0.0.0.0/0"
        stateful = true

        upd_options {
            "min" = 319
            "max" = 320
        }
    }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the security list.
* `display_name` - (Required) The OCID of the VCN.
* `egress_security_rules` - (Required) Rules for allowing egress IP packets. [EgressSecurityRule API Docs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/EgressSecurityRule/)
* `ingress_security_rules` - (Required) Rules for allowing ingress IP packets. [IngressSecurityRule API Docs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IngressSecurityRule/)
* `vcn_id` - (Optional) The OCID of the VCN the security list belongs to.

## Attributes Reference

* `compartment_id` - The OCID of the compartment containing the security list.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `egress_security_rules` - Rules for allowing egress IP packets.
* `id` - The security list's Oracle Cloud ID (OCID).
* `ingress_security_rules` - Rules for allowing ingress IP packets.
* `state` - The security list's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the security list was created.
* `vcn_id` - The OCID of the VCN the security list belongs to.
