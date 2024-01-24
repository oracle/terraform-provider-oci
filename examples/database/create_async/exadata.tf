// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "test_vcn" {
  compartment_id = var.compartment_ocid
  cidr_block     = "10.1.0.0/16"
  display_name   = "TestVcn"
  dns_label      = "examplevcn"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.compartment_ocid
  ad_number      = 1
}

resource "oci_core_security_list" "exadata_shapes_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "ExadataSecurityList"

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "6"
  }

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "1"
  }

  egress_security_rules {
    destination = "10.1.22.0/24"
    protocol    = "6"
  }

  egress_security_rules {
    destination = "10.1.22.0/24"
    protocol    = "1"
  }
}

resource "oci_core_subnet" "exadata_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.22.0/24"
  display_name        = "TestExadataSubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id, oci_core_security_list.exadata_shapes_security_list.id]
  dns_label           = "subnetexadata"
}

resource "oci_database_autonomous_exadata_infrastructure" "test_autonomous_exadata_infrastructure" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestExadata11"
  domain              = var.autonomous_exadata_infrastructure_domain
  freeform_tags       = var.autonomous_database_freeform_tags
  license_model       = "LICENSE_INCLUDED"

  create_async = true
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

  nsg_ids   = [oci_core_network_security_group.test_network_security_group.id]
  shape     = "Exadata.Quarter2.92"
  subnet_id = oci_core_subnet.exadata_subnet.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

data "oci_database_autonomous_exadata_infrastructures" "test_autonomous_exadata_infrastructures" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestExadata"
}

data "oci_database_autonomous_exadata_infrastructure" "test_autonomous_exadata_infrastructure" {
  autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
}

output "test_autonomous_exadata_infrastructures" {
  value = [data.oci_database_autonomous_exadata_infrastructures.test_autonomous_exadata_infrastructures.autonomous_exadata_infrastructures]
}

