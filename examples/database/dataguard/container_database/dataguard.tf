// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "autonomous_database_defined_tags_value" {
  default = "value"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_exadata_infrastructure_domain" {
  default = "subnetexadata.examplevcn.oraclevcn.com"
}

variable "autonomous_container_database_backup_config_recovery_window_in_days" {
  default = 10
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


# You can only apply one oci_database_autonomous_container_database_dataguard_association_operation once at a time. After you successfully done a switchover or failover, the role of
# primary and standby will switch, so you will need to change the database_id and dataguard_association_id accordingly (autonomous_container_database_id <-> peer_autonomous_container_database_id)
# (id <-> peer_autonomous_container_database_dataguard_association_id)
resource "oci_database_autonomous_container_database_dataguard_association_operation" "switchover" {
  operation = "switchover" # "failover" or "reinstate"
  autonomous_container_database_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["autonomous_container_database_id"]
  autonomous_container_database_dataguard_association_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]
}
/*
resource "oci_database_autonomous_container_database_dataguard_association_operation" "failover" {
  operation = "failover" # "failover" or "reinstate"
  autonomous_container_database_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_id"]
  autonomous_container_database_dataguard_association_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_dataguard_association_id"]
}

resource "oci_database_autonomous_container_database_dataguard_association_operation" "reninstate" {
  operation = "reinstate" # "failover" or "reinstate"
  autonomous_container_database_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_id"]
  autonomous_container_database_dataguard_association_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_dataguard_association_id"]
}
*/
resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
  #Required
  autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
  display_name                         = "example-container-database"
  patch_model                          = "RELEASE_UPDATES"
  timeouts {
    create = "60h"
    delete = "60h"
  }
  #Optional
  backup_config {
    #Optional
    recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
  }
  compartment_id               = var.compartment_ocid

  peer_autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.peer_autonomous_exadata_infrastructure.id
  protection_mode = "MAXIMUM_AVAILABILITY"
  freeform_tags                = var.autonomous_database_freeform_tags
  service_level_agreement_type = "AUTONOMOUS_DATAGUARD"

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

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "atpdb1"
  timeouts {
    create = "60h"
    delete = "60h"
  }
  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  db_workload                      = "OLTP"
  display_name                     = "example_autonomous_database-009"
  freeform_tags                    = var.autonomous_database_freeform_tags
  is_dedicated                     = "true"
}

data "oci_database_autonomous_container_databases" "test_autonomous_container_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
  availability_domain                  = data.oci_identity_availability_domain.ad.name
  display_name                         = "example-container-database"
  state                                = "AVAILABLE"
}

data "oci_database_autonomous_container_databases" "peer_autonomous_container_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.peer_autonomous_exadata_infrastructure.id
}

data "oci_database_autonomous_container_database_dataguard_associations" "dataguard_associations" {
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  display_name                     = oci_database_autonomous_database.test_autonomous_database.display_name
  db_workload                      = "OLTP"
}

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

  timeouts {
    create = "60h"
    delete = "60h"
  }

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

resource "oci_database_autonomous_exadata_infrastructure" "peer_autonomous_exadata_infrastructure" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestExadata12"
  domain              = var.autonomous_exadata_infrastructure_domain
  freeform_tags       = var.autonomous_database_freeform_tags
  license_model       = "LICENSE_INCLUDED"
  timeouts {
    create = "60h"
    delete = "60h"
  }

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
  state               = "AVAILABLE"
}

data "oci_database_autonomous_exadata_infrastructure" "test_autonomous_exadata_infrastructure" {
  autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
}