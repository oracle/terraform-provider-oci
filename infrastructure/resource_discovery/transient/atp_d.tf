// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_security_list" "exadata_shapes_security_list_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"
  display_name   = "ExadataSecurityListRD"

  ingress_security_rules {
    source   = "10.1.23.0/24"
    protocol = "6"
  }

  ingress_security_rules {
    source   = "10.1.23.0/24"
    protocol = "1"
  }

  egress_security_rules {
    destination = "10.1.23.0/24"
    protocol    = "6"
  }

  egress_security_rules {
    destination = "10.1.23.0/24"
    protocol    = "1"
  }
}

resource "oci_core_subnet" "exadata_subnet_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.1.23.0/24"
  display_name        = "TestExadataSubnetRD"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn_rd.id}"
  route_table_id      = "${oci_core_vcn.vcn_rd.default_route_table_id}"
  dhcp_options_id     = "${oci_core_vcn.vcn_rd.default_dhcp_options_id}"
  security_list_ids   = ["${oci_core_vcn.vcn_rd.default_security_list_id}", "${oci_core_security_list.exadata_shapes_security_list_rd.id}"]
  dns_label           = "subnetexadatard"
}

resource "oci_database_autonomous_exadata_infrastructure" "test_autonomous_exadata_infrastructure_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Exadata11RD"
  domain              = "${var.autonomous_exadata_infrastructure_domain}"
  freeform_tags       = "${var.autonomous_database_freeform_tags}"
  license_model       = "LICENSE_INCLUDED"

  maintenance_window_details {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "MONDAY"
    }

    hours_of_day = ["4"]

    months {
      name = "JANUARY"
    }

    months {
      name = "APRIL"
    }

    months {
      name = "JULY"
    }

    months {
      name = "OCTOBER"
    }

    weeks_of_month = ["2"]
  }

  nsg_ids   = ["${oci_core_network_security_group.network_security_group_rd.id}"]
  shape     = "Exadata.Quarter2.92"
  subnet_id = "${oci_core_subnet.exadata_subnet_rd.id}"
}

resource "oci_database_autonomous_container_database" "test_autonomous_container_database_rd" {
  #Required
  autonomous_exadata_infrastructure_id = "${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure_rd.id}"
  display_name                         = "container-database-rd"
  patch_model                          = "RELEASE_UPDATES"

  #Optional
  backup_config {
    #Optional
    recovery_window_in_days = "${var.autonomous_container_database_backup_config_recovery_window_in_days}"
  }

  compartment_id               = "${var.compartment_ocid}"
  freeform_tags                = "${var.autonomous_database_freeform_tags}"
  service_level_agreement_type = "STANDARD"

  maintenance_window_details {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "MONDAY"
    }

    hours_of_day = ["4"]

    months {
      name = "JANUARY"
    }

    months {
      name = "APRIL"
    }

    months {
      name = "JULY"
    }

    months {
      name = "OCTOBER"
    }

    weeks_of_month = ["2"]
  }
}
