// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_file_system" "file_storage_file_system_rd" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"

  #Optional
  display_name = "fileStorageFileSystemRD"
  defined_tags = "${map("example-tag-namespace-all.example-tag", "value")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_file_storage_mount_target" "file_storage_mount_target_rd" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  subnet_id           = "${oci_core_subnet.subnet_rd.id}"

  #Optional
  display_name = "fileStorageMountTargetRD"
  defined_tags = "${map("example-tag-namespace-all.example-tag", "value")}"

  freeform_tags = {
    "Department" = "Finance"
  }

  nsg_ids = ["${oci_core_network_security_group.test_network_security_group_file_storage_rd.id}"]
}

resource "oci_core_network_security_group" "test_network_security_group_file_storage_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"
}

resource "oci_core_security_list" "my_security_list_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "mySecurityListRD"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"

  // Allow all outbound requests
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
  }

  # See https://docs.us-phoenix-1.oraclecloud.com/Content/File/Tasks/creatingfilesystems.htm.
  # Specific security list rules are required to allow mount targets to work properly.
  ingress_security_rules {
    protocol = "6"
    source   = "10.1.0.0/16"

    tcp_options {
      min = 2048
      max = 2050
    }
  }

  ingress_security_rules {
    protocol = "6"
    source   = "10.1.0.0/16"

    tcp_options {
      source_port_range {
        min = 2048
        max = 2050
      }
    }
  }

  ingress_security_rules {
    protocol = "6"
    source   = "10.1.0.0/16"

    tcp_options {
      min = 111
      max = 111
    }
  }

  ingress_security_rules {
    // Allowing inbound SSH traffic to instances in the subnet from any source
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      min = 22
      max = 22
    }
  }

  ingress_security_rules {
    // Allowing inbound ICMP traffic of a specific type and code from any source
    protocol = 1
    source   = "0.0.0.0/0"

    icmp_options {
      type = 3
      code = 4
    }
  }

  ingress_security_rules {
    // Allowing inbound ICMP traffic of a specific type from within our VCN
    protocol = 1
    source   = "10.1.0.0/16"

    icmp_options {
      type = 3
    }
  }
}

resource "oci_file_storage_export" "file_storage_export_rd" {
  #Required
  export_set_id  = "${oci_file_storage_export_set.file_storage_export_set_rd.id}"
  file_system_id = "${oci_file_storage_file_system.file_storage_file_system_rd.id}"
  path           = "${var.export_path_fs1_mt1}"

  export_options {
    source                         = "${var.export_read_write_access_source}"
    access                         = "READ_WRITE"
    identity_squash                = "NONE"
    require_privileged_source_port = true
  }

  export_options {
    source                         = "${var.export_read_only_access_source}"
    access                         = "READ_ONLY"
    identity_squash                = "ALL"
    require_privileged_source_port = true
  }
}

resource "oci_file_storage_export_set" "file_storage_export_set_rd" {
  # Required
  mount_target_id = "${oci_file_storage_mount_target.file_storage_mount_target_rd.id}"

  # Optional
  display_name      = "fileStorageExportSetRD"
  max_fs_stat_bytes = "${var.max_byte}"
  max_fs_stat_files = "${var.max_files}"
}

variable "export_read_write_access_source" {
  default = "10.0.0.0/8"
}

variable "export_read_only_access_source" {
  default = "0.0.0.0/0"
}

variable "export_path_fs1_mt1" {
  default = "/myfsspaths/fs1/path1"
}

variable "max_byte" {
  default = 23843202333
}

variable "max_files" {
  default = 223442
}

resource "oci_file_storage_snapshot" "file_storage_snapshot_rd" {
  #Required
  file_system_id = "${oci_file_storage_file_system.file_storage_file_system_rd.id}"
  name           = "fileStorageSnapshotNameRD}"
  defined_tags   = "${map("example-tag-namespace-all.example-tag", "value")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}
