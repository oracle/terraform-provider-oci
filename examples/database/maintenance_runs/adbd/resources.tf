
resource "oci_core_vcn" "test_vcn" {
  compartment_id = var.compartment_ocid
  cidr_block     = "10.1.0.0/16"
  display_name   = "TestVcn"
  dns_label      = "examplevcn"
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

resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TFATPD"
  shape               = var.cloud_exadata_infrastructure_shape

  #Optional
  compute_count = var.cloud_exadata_infrastructure_compute_count
  storage_count = var.cloud_exadata_infrastructure_storage_count
}


resource "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster" {
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
  compartment_id                  = var.compartment_ocid
  display_name                    = "TestCloudAutonomousVmCluster"
  license_model                   = "LICENSE_INCLUDED"
  subnet_id                       = oci_core_subnet.exadata_subnet.id
  compute_model                   = "ECPU"

  //To ignore changes to autonomous_data_storage_size_in_tbs and db_servers
  lifecycle {
    ignore_changes = [
      autonomous_data_storage_size_in_tbs,
      db_servers,
    ]
  }
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
  #Required
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  display_name                         = "example-container-database"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = "19.26.0.1.0"
  db_name                              = "ACDNAME"

  #Optional
  db_split_threshold           = 12
  vm_failover_reservation      = 25
  distribution_affinity        = "MINIMUM_DISTRIBUTION"
  net_services_architecture    = "DEDICATED"

  compartment_id               = var.compartment_ocid
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
  rotate_key_trigger = "true"
  version_preference = "LATEST_RELEASE_UPDATE"
  is_dst_file_update_enabled = false
}

resource "oci_database_maintenance_run" "test_maintenance_run" {
  target_resource_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  patch_type = "QUARTERLY"
  time_scheduled = var.time_for_schedule_mr
}