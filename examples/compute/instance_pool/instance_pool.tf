// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

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

variable "instance_shape" {
  default = "VM.Standard.E3.Flex"
}

variable "instance_ocpus" {
  default = 1
}

variable "instance_shape_config_memory_in_gbs" {
  default = 8
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "flex_instance_image_ocid" {
  type = map(string)
  default = {
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaa6hooptnlbfwr5lwemqjbu3uqidntrlhnt45yihfj222zahe7p3wq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa6tp7lhyrcokdtf7vrbmxyp2pctgg4uxvt4jz4vc47qoc2ec4anha"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaadvi77prh3vjijhwe5xbd6kjg3n5ndxjcpod6om6qaiqeu3csof7a"
    uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaw5gvriwzjhzt2tnylrfnpanz5ndztyrv3zpwhlzxdbkqsjfkwxaq"
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_route_table.test_route_table.id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "TestRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_load_balancer" "test_load_balancer" {
  shape          = "100Mbps"
  compartment_id = var.compartment_ocid

  subnet_ids = [
    oci_core_subnet.test_subnet.id,
  ]

  display_name = "TestLoadBalancer"
  is_private   = true
}

resource "oci_load_balancer_backend_set" "test_backend_set" {
  name             = "lb-bes1"
  load_balancer_id = oci_load_balancer.test_load_balancer.id
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "80"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstanceForInstancePool"
  shape               = var.instance_shape
  shape_config {
    ocpus = var.instance_ocpus
    memory_in_gbs = var.instance_shape_config_memory_in_gbs
  }

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "PrimaryVnic"
    assign_public_ip = true
    hostname_label   = "testinstanceforinstancepool"
  }

  source_details {
    source_type = "image"
    source_id   = var.flex_instance_image_ocid[var.region]
  }

  timeouts {
    create = "60m"
  }

  fault_domain = "FAULT-DOMAIN-1"
}

resource "oci_core_instance_configuration" "test_instance_configuration" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    /*
      Attach multiple block volumes
    */
    block_volumes {
      create_details {
        compartment_id       = var.compartment_ocid
        display_name         = "TestCreateVolumeDetails-1"
        availability_domain  = data.oci_identity_availability_domain.ad.name
        size_in_gbs          = 50
        vpus_per_gb          = 20 // min vpus
        is_auto_tune_enabled = false
        block_volume_replicas {
          display_name        = "TestCreateVolumeDetails-1"
          availability_domain = data.oci_identity_availability_domain.ad.name
        }
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails-1"
        is_read_only                        = true
        is_shareable                        = true
      }
    }

    block_volumes {
      create_details {
        compartment_id      = var.compartment_ocid
        display_name        = "TestCreateVolumeDetails-2"
        availability_domain = data.oci_identity_availability_domain.ad.name
        size_in_gbs         = 50
        vpus_per_gb         = 20 // min vpus
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails-2"
        is_read_only                        = true
        is_shareable                        = true
      }
    }

    launch_details {
      compartment_id                      = var.compartment_ocid
      ipxe_script                         = "ipxeScript"
      shape                               = var.instance_shape
      display_name                        = "TestInstanceConfigurationLaunchDetails"
      is_pv_encryption_in_transit_enabled = false
      preferred_maintenance_action        = "LIVE_MIGRATE"
      launch_mode                         = "NATIVE"

      agent_config {
        is_management_disabled = false
        is_monitoring_disabled = false
      }

      availability_config {
        recovery_action             = "RESTORE_INSTANCE"
        is_live_migration_preferred = false
      }

      launch_options {
        network_type = "PARAVIRTUALIZED"
      }

      instance_options {
        are_legacy_imds_endpoints_disabled = false
      }

      shape_config {
        ocpus = var.instance_ocpus
        memory_in_gbs = var.instance_shape_config_memory_in_gbs
      }

      create_vnic_details {
        assign_private_dns_record = true
        assign_public_ip       = true
        display_name           = "TestInstanceConfigurationVNIC"
        skip_source_dest_check = false
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = var.flex_instance_image_ocid[var.region]
      }
    }
  }
}

resource "oci_core_instance_pool" "test_instance_pool" {
  compartment_id = var.compartment_ocid
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
  size = 2
  state = "RUNNING"
  display_name = "TestInstancePool"
  instance_display_name_formatter = "host-$${launchCount}"
  instance_hostname_formatter = "host-$${launchCount}"

  placement_configurations {
    availability_domain = data.oci_identity_availability_domain.ad.name
    fault_domains = [
      "FAULT-DOMAIN-1"]
    primary_subnet_id = oci_core_subnet.test_subnet.id
  }

  load_balancers {
    backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
    load_balancer_id = oci_load_balancer.test_load_balancer.id
    port = 80
    vnic_selection = "primaryvnic"
  }

  lifecycle {
    ignore_changes = [size]
  }
}

resource "oci_core_instance_pool_instance" "test_instance_pool_instance" {
  instance_pool_id = oci_core_instance_pool.test_instance_pool.id
  instance_id = oci_core_instance.test_instance.id
  decrement_size_on_delete = true
  auto_terminate_instance_on_delete = false
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_core_instance_configuration" "test_instance_configuration_datasource" {
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
}

data "oci_core_instance_configurations" "test_instance_configurations_datasource" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_instance_configuration.test_instance_configuration.id]
  }
}

data "oci_core_instance_pool" "test_instance_pool_datasource" {
  instance_pool_id = oci_core_instance_pool.test_instance_pool.id
}

data "oci_core_instance_pools" "test_instance_pools_datasource" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstancePool"
  state          = "RUNNING"

  filter {
    name   = "id"
    values = [oci_core_instance_pool.test_instance_pool.id]
  }
}

data "oci_core_instance_pool_instances" "test_instance_pool_instances_datasource" {
  compartment_id   = var.compartment_ocid
  instance_pool_id = oci_core_instance_pool.test_instance_pool.id
  depends_on = [oci_core_instance_pool_instance.test_instance_pool_instance]
}

# Usage of singular instance datasources to show the public_ips, private_ips, and hostname_labels for the instances in the pool
data "oci_core_instance" "test_instance_pool_instance_singular_datasource" {
  count       = 3
  instance_id = data.oci_core_instance_pool_instances.test_instance_pool_instances_datasource.instances[count.index]["id"]
}

output "pooled_instances_private_ips" {
  value = [data.oci_core_instance.test_instance_pool_instance_singular_datasource.*.private_ip]
}

output "pooled_instances_public_ips" {
  value = [data.oci_core_instance.test_instance_pool_instance_singular_datasource.*.public_ip]
}

output "pooled_instances_hostname_labels" {
  value = [data.oci_core_instance.test_instance_pool_instance_singular_datasource.*.hostname_label]
}

