provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

variable "tenancy_ocid" {
  default = ""
}



variable "compartment_ocid" {
  default = ""
}


resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  lifecycle {
    ignore_changes = [
    "defined_tags"]
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  dns_label      = "testvcn"
  lifecycle {
    ignore_changes = [
    "defined_tags"]
  }
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  dns_label      = "testsubnet"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [
  "${oci_core_security_list.test_sec_list.id}"]
  vcn_id = oci_core_vcn.test_vcn.id
  lifecycle {
    ignore_changes = [
    "defined_tags"]
  }

  provisioner "local-exec" {
    command = "sleep 180"
  }
}

resource "oci_core_security_list" "test_sec_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
    stateless   = "false"
  }

  ingress_security_rules {
    protocol  = "6"
    source    = "0.0.0.0/0"
    stateless = "false"
    tcp_options {
      max = "22"
      min = "22"
    }
  }

  ingress_security_rules {
    protocol  = "6"
    source    = "0.0.0.0/0"
    stateless = "false"
    tcp_options {
      max = "80"
      min = "80"
    }
  }

  ingress_security_rules {
    protocol  = "1"
    source    = "0.0.0.0/0"
    stateless = "true"
    icmp_options {
      code = "4"
      type = "3"
    }
  }

  ingress_security_rules {
    protocol  = "1"
    source    = "10.0.0.0/16"
    stateless = "false"
  }

  ingress_security_rules {
    protocol  = "6"
    source    = "10.0.0.0/16"
    stateless = "false"
    tcp_options {
      max = "111"
      min = "111"
    }
  }

  ingress_security_rules {
    protocol  = "6"
    source    = "10.0.0.0/16"
    stateless = "false"
    tcp_options {
      max = "2051"
      min = "2048"
    }
  }

  ingress_security_rules {
    protocol  = "17"
    source    = "10.0.0.0/16"
    stateless = "false"
    udp_options {
      max = "111"
      min = "111"
    }
  }

  ingress_security_rules {
    protocol  = "17"
    source    = "10.0.0.0/16"
    stateless = "false"
    udp_options {
      max = "2051"
      min = "2048"
    }
  }
}

resource "oci_core_internet_gateway" "test_ig" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  route_rules {
    destination       = "0.0.0.0/0"
    network_entity_id = oci_core_internet_gateway.test_ig.id
  }
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

resource "oci_file_storage_file_system" "test_file_system" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name
  compartment_id      = var.compartment_ocid
  display_name        = "ci-fss-file-system"
}

resource "oci_file_storage_mount_target" "test_mount_target" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name
  compartment_id      = var.compartment_ocid
  subnet_id           = oci_core_subnet.test_subnet.id
  display_name        = "ci-fss-mount-target"
}

resource "oci_file_storage_export" "test_export" {
  export_set_id  = oci_file_storage_mount_target.test_mount_target.export_set_id
  file_system_id = oci_file_storage_file_system.test_file_system.id
  path           = "/ci-fss"

  export_options {
    source                         = "10.0.0.0/16"
    access                         = "READ_WRITE"
    allowed_auth                   = ["SYS"]
    identity_squash                = "NONE"
    require_privileged_source_port = true
  }
}

resource "oci_container_instances_container_instance" "test_container_instance" {
  depends_on = [oci_core_subnet.test_subnet, oci_file_storage_export.test_export]

  #Required
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name
  compartment_id      = var.compartment_ocid
  containers {
    #Required
    image_url = "busybox"

    #Optional
    arguments = [
      "-c",
    "cat /mnt/my_file"]
    command = [
    "/bin/sh"]
    display_name = "displayName"
    environment_variables = {
      "environment" = "variable"
    }
    health_checks {
      #Required
      health_check_type = "HTTP"

      #Optional
      failure_action    = "KILL"
      failure_threshold = "10"
      headers {

        #Optional
        name  = "name"
        value = "value"
      }
      initial_delay_in_seconds = "10"
      interval_in_seconds      = "10"
      name                     = "name"
      path                     = "path"
      port                     = "10"
      success_threshold        = "1"
      timeout_in_seconds       = "10"
    }
    is_resource_principal_disabled = "false"
    resource_config {

      #Optional
      memory_limit_in_gbs = "1.0"
      vcpus_limit         = "1.0"
    }
    volume_mounts {
      #Required
      mount_path  = "/mnt"
      volume_name = "volumeName1"

      #Optional
      is_read_only = "false"
      partition    = "10"
      sub_path     = "subPath"
    }
    volume_mounts {
      #Required
      mount_path  = "/mnt"
      volume_name = "volumeName2"
    }
    volume_mounts {
      #Required
      mount_path  = "/fss"
      volume_name = "volumeName3"

      #Optional
      is_read_only = "false"
    }
    working_directory = "/mnt"
    security_context {
      security_context_type          = "LINUX"
      is_non_root_user_check_enabled = false
      is_root_file_system_readonly   = true
      run_as_group                   = 10
      run_as_user                    = 10
      capabilities {
        add_capabilities  = ["CAP_CHOWN", "CAP_KILL"]
        drop_capabilities = ["ALL"]
      }
    }
  }
  shape = "CI.Standard.E4.Flex"
  shape_config {
    memory_in_gbs = "8"
    ocpus         = "4"
  }
  vnics {
    #Required
    subnet_id = oci_core_subnet.test_subnet.id

    #Optional
    defined_tags = tomap({ "tf_test_namespace.test_tag" = "value" })
    display_name = "displayName"
    freeform_tags = {
      "freeformTag" = "freeformTags"
    }
    hostname_label         = "hostnamelabel"
    is_public_ip_assigned  = "true"
    nsg_ids                = []
    private_ip             = "10.0.0.7"
    skip_source_dest_check = "false"
  }

  #Optional
  container_restart_policy = "ALWAYS"
  defined_tags             = tomap({ "tf_test_namespace.test_tag" = "value" })
  display_name             = "displayName"
  dns_config {

    #Optional
    nameservers = [
    "8.8.8.8"]
    options = [
    "options"]
    searches = [
    "search domain"]
  }
  freeform_tags = {
    "bar-key" = "foo-value"
  }
  graceful_shutdown_timeout_in_seconds = "10"
  lifecycle {
    ignore_changes = ["defined_tags", "vnics[0].defined_tags"]
  }
  state = "ACTIVE"
  volumes {
    #Required
    name        = "volumeName1"
    volume_type = "EMPTYDIR"

    #Optional
    backing_store = "EPHEMERAL_STORAGE"
  }
  volumes {
    #Required
    name        = "volumeName2"
    volume_type = "CONFIGFILE"

    configs {
      data      = "SGFyc2hpdA=="
      file_name = "my_file"
    }
  }
  volumes {
    #Required
    name        = "volumeName3"
    volume_type = "OCI_FSS_FILE_SYSTEM"

    export {
      #Required
      id = oci_file_storage_export.test_export.id

      #Optional
      oci_fss_export_type = "OCID"
    }
    mount_command {
      mount_options {
        option = "retrans"
        value  = "3"
      }
    }
    mount_target {
      #Required
      id = oci_file_storage_mount_target.test_mount_target.id

      #Optional
      oci_fss_mount_target_type = "OCID"
    }
    security {
      auth                    = "SYS"
      is_encrypted_in_transit = "false"
    }
    subnet_id = oci_core_subnet.test_subnet.id
  }
}

data "oci_container_instances_container_instance" "test_container_instance" {
  container_instance_id = oci_container_instances_container_instance.test_container_instance.id
}

