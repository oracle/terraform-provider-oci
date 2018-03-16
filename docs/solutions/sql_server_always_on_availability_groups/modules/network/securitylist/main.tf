resource "oci_core_security_list" "dmz" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "DMZ.sl"
  vcn_id         = "${var.vcn_id}"

  egress_security_rules = [
    {
      protocol    = "${var.proto_number_All}"
      destination = "0.0.0.0/0"
    },
  ]

  ingress_security_rules = [
    {
      tcp_options {
        "max" = 3389
        "min" = 3389
      }

      protocol = "${var.proto_number_tcp}"
      source   = "0.0.0.0/0"
    },
  ]
}

resource "oci_core_security_list" "admin" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "ADMIN.sl"
  vcn_id         = "${var.vcn_id}"

  egress_security_rules = [
    {
      protocol    = "${var.proto_number_All}"
      destination = "${var.DMZ_prefix}"
    },
    {
      protocol    = "${var.proto_number_All}"
      destination = "${var.ADMIN_prefix}"
    },
    {
      protocol    = "${var.proto_number_All}"
      destination = "${var.SQL_prefix}"
    },
  ]

  ingress_security_rules = [
    {
      tcp_options {
        "max" = 53
        "min" = 53
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 53
        "min" = 53
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 389
        "min" = 389
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 389
        "min" = 389
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 636
        "min" = 636
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 3269
        "min" = 3268
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 88
        "min" = 88
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 88
        "min" = 88
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 135
        "min" = 135
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 135
        "min" = 135
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 137
        "min" = 137
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 137
        "min" = 137
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 139
        "min" = 139
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 138
        "min" = 138
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 445
        "min" = 445
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 445
        "min" = 445
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 464
        "min" = 464
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 464
        "min" = 464
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 65535
        "min" = 49152
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.SQL_prefix}"
    },
    {
      udp_options {
        "max" = 65535
        "min" = 49152
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.SQL_prefix}"
    },
    {
      protocol = "${var.proto_number_All}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 1433
        "min" = 1433
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.DMZ_prefix}"
    },
    {
      tcp_options {
        "max" = 3389
        "min" = 3389
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.DMZ_prefix}"
    },
  ]
}

resource "oci_core_security_list" "sql" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "SQL.sl"
  vcn_id         = "${var.vcn_id}"

  egress_security_rules = [
    {
      protocol    = "${var.proto_number_All}"
      destination = "${var.DMZ_prefix}"
    },
    {
      protocol    = "${var.proto_number_All}"
      destination = "${var.ADMIN_prefix}"
    },
    {
      protocol    = "${var.proto_number_All}"
      destination = "${var.SQL_prefix}"
    },
  ]

  ingress_security_rules = [
    {
      icmp_options {
        "type" = 1
      }

      protocol = "${var.proto_number_icmp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 135
        "min" = 135
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 137
        "min" = 137
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      udp_options {
        "max" = 137
        "min" = 137
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 445
        "min" = 445
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 1434
        "min" = 1433
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      udp_options {
        "max" = 1434
        "min" = 1434
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 3343
        "min" = 3343
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      udp_options {
        "max" = 3343
        "min" = 3343
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 5022
        "min" = 5022
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 5985
        "min" = 5985
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 1433
        "min" = 1433
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 464
        "min" = 464
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      udp_options {
        "max" = 464
        "min" = 464
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      tcp_options {
        "max" = 65535
        "min" = 49152
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      udp_options {
        "max" = 65535
        "min" = 49152
      }

      protocol = "${var.proto_number_udp}"
      source   = "${var.ADMIN_prefix}"
    },
    {
      protocol = "${var.proto_number_All}"
      source   = "${var.SQL_prefix}"
    },
    {
      tcp_options {
        "max" = 3389
        "min" = 3389
      }

      protocol = "${var.proto_number_tcp}"
      source   = "${var.DMZ_prefix}"
    },
  ]
}
