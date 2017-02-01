# baremetal\_core\_security\_lists

Provides a security list resource.

## Example Usage

```
resource "baremetal_core_security_list" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    egress_security_rules {
        destination = "destination"
        icmp_options {
            "code" = 1
            "type" = 2
        }
        protocol = "protocol"
    }
    ingress_security_rules {
        tcp_options {
            "max" = 2
            "min" = 1
        }
        protocol = "protocol"
        source = "source"
    }
    vcn_id = "vcn_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the security list.
* `display_name` - (Required) The OCID of the VCN.
* `egress_security_rules` - (Required) Rules for allowing egress IP packets.
* `ingress_security_rules` - (Required) Rules for allowing ingress IP packets.
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
