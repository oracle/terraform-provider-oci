# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml

resource "oci_core_security_list" "my_security_list" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "my_security_list"
  vcn_id         = "${oci_core_virtual_network.my_vcn.id}"

  // Allow all outbound requests
  egress_security_rules = [
    {
      destination = "0.0.0.0/0"
      protocol    = "all"
    },
  ]

  // See https://docs.us-phoenix-1.oraclecloud.com/Content/File/Tasks/creatingfilesystems.htm.
  // Specific security list rules are required to allow mount targets to work properly.
  ingress_security_rules = [
    {
      protocol = "6"
      source   = "${var.my_vcn-cidr}"

      tcp_options {
        "min" = 2048
        "max" = 2050
      }
    },
    {
      protocol = "6"
      source   = "${var.my_vcn-cidr}"

      tcp_options {
        source_port_range {
          "min" = 2048
          "max" = 2050
        }
      }
    },
    {
      protocol = "6"
      source   = "${var.my_vcn-cidr}"

      tcp_options {
        "min" = 111
        "max" = 111
      }
    },
    {
      // Allowing inbound SSH traffic to instances in the subnet from any source
      protocol = "6"
      source   = "0.0.0.0/0"

      tcp_options {
        "min" = 22
        "max" = 22
      }
    },
    {
      // Allowing inbound ICMP traffic of a specific type and code from any source
      protocol = 1
      source   = "0.0.0.0/0"

      icmp_options {
        "type" = 3
        "code" = 4
      }
    },
    {
      // Allowing inbound ICMP traffic of a specific type from within our VCN
      protocol = 1
      source   = "${var.my_vcn-cidr}"

      icmp_options {
        "type" = 3
      }
    },
  ]
}
