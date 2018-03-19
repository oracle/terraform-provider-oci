# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml

resource "oci_core_security_list" "my_security_list" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "my_security_list"
  vcn_id = "${oci_core_virtual_network.my_vcn.id}"

  // Allow all outbound requests
  egress_security_rules = [
    {
      destination = "0.0.0.0/0"
      protocol = "6"
    }]

  // See https://docs.us-phoenix-1.oraclecloud.com/Content/File/Tasks/creatingfilesystems.htm.
  // Specific security list rules are required to allow mount targets to work properly.
  ingress_security_rules = [
    {
      protocol = "6"
      source = "${var.my_vcn-cidr}"

      tcp_options {
        "min" = 2048
        "max" = 2050
      }
    },
    {
      protocol = "6"
      source = "${var.my_vcn-cidr}"

      tcp_options {
        source_port_range {
          "min" = 2048
          "max" = 2050
        }
      }
    },
    {
      protocol = "6"
      source = "${var.my_vcn-cidr}"

      tcp_options {
        "max" = 111
        "min" = 111
      }
    }]
}