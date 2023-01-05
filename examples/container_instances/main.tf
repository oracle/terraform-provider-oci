provider "oci" {
}

variable "tenancy_ocid" {
  default = ""
}

variable "ssh_public_key" {
  default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
}

variable "region" {
  default = ""
}

variable "compartment_ocid" {
  default = ""
}

variable "defined_tag_namespace_name" {
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
  cidr_block        = "10.0.0.0/24"
  compartment_id    = var.compartment_ocid
  dns_label         = "testsubnet"
  route_table_id    = oci_core_route_table.test_route_table.id
  security_list_ids = [
    "${oci_core_security_list.test_sec_list.id}"]
  vcn_id            = oci_core_vcn.test_vcn.id
  lifecycle {
    ignore_changes = [
      "defined_tags"]
  }
}

resource "oci_core_security_list" "test_sec_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  egress_security_rules {
    destination  = "0.0.0.0/0"
    protocol     = "all"
    stateless    = "false"
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

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired     = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired       = false
}

resource "oci_container_instances_container_instance" "test_container_instance" {
  #Required
  availability_domain      = data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name
  compartment_id           = var.compartment_ocid
  containers {
    #Required
    image_url                      = "busybox"

    #Optional
    additional_capabilities = [
      "CAP_NET_ADMIN"]
    arguments = [
      "-c",
      "sleep 24h"]
    command = [
      "/bin/sh"]
    display_name = "displayName"
    environment_variables = {
      "environment" = "variable"
    }
    health_checks {
      #Required
      health_check_type        = "HTTP"

      #Optional
      failure_action           = "KILL"
      failure_threshold        = "10"
      headers {

        #Optional
        name = "name"
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
      mount_path   = "/mnt"
      volume_name  = "volumeName"

      #Optional
      is_read_only = "false"
      partition    = "10"
      sub_path     = "/subPath"
    }
    working_directory = "/mnt"
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
    defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")
    display_name = "displayName"
    freeform_tags = {
      "freeformTag" = "freeformTags"
    }
    hostname_label = "hostnamelabel"
    is_public_ip_assigned = "true"
    nsg_ids = []
    private_ip = "10.0.0.7"
    skip_source_dest_check = "false"
  }

  #Optional
  container_restart_policy = "ALWAYS"
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")
  display_name = "displayName"
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
    ignore_changes = [
      "defined_tags"]
  }
  state           = "ACTIVE"
  volumes {
    #Required
    name = "volumeName"
    volume_type = "EMPTYDIR"

    #Optional
    backing_store = "EPHEMERAL_STORAGE"
  }
}