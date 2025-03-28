// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_exadata_infrastructure" "test_exadata_infrastructure" {
  #Required
  admin_network_cidr = "192.168.0.0/16"
  cloud_control_plane_server1 = "10.32.88.1"
  cloud_control_plane_server2 = "10.32.88.3"
  compartment_id = var.compartment_id
  display_name = "tstExaInfra"
  dns_server = [
    "10.231.225.65"]
  gateway = "10.32.88.5"
  infini_band_network_cidr = "10.31.8.0/21"
  netmask = "255.255.255.0"
  ntp_server = [
    "10.231.225.76"]
  shape = "ExadataCC.X7"
  time_zone = "US/Pacific"
  activation_file = "activation.zip"
  storage_count = 3
  compute_count = 2
  is_cps_offline_report_enabled = false
  is_multi_rack_deployment = false

  #Optional
  network_bonding_mode_details {
    backup_network_bonding_mode = "ACTIVE_BACKUP"
    client_network_bonding_mode = "ACTIVE_BACKUP"
  }

  #Optional
  contacts {
    email = "testuser2@testdomain.com"
    is_primary = "true"
    name = "name"
    phone_number = "1234567891"
  }

  corporate_proxy = "http://192.168.19.1:80"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  maintenance_window {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "TUESDAY"
    }

    hours_of_day = [
      "4"]
    lead_time_in_weeks = 2

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

    weeks_of_month = [
      "2"]

    patching_mode = "ROLLING"

    is_custom_action_timeout_enabled = true

    custom_action_timeout_in_mins = 30
  }
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "tagNamespace1"
  name           = "testexamples-tag-namespace1"
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}