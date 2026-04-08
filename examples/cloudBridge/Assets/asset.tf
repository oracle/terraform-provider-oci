// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}
variable "compartment_id" {
  default = "OCID"
}

variable "inventory_id" {
  default = "OCID"
}

variable "asset_asset_source_ids" {
  default = []
}

variable "asset_asset_type" {
  default = "VMWARE_VM"
}

variable "asset_compute_connected_networks" {
  default = 10
}

variable "asset_compute_cores_count" {
  default = 10
}

variable "asset_compute_cpu_model" {
  default = "cpuModel"
}

variable "asset_compute_description" {
  default = "description"
}

variable "asset_compute_disks_boot_order" {
  default = 10
}

variable "asset_compute_disks_location" {
  default = "location"
}

variable "asset_compute_disks_name" {
  default = "name"
}

variable "asset_compute_disks_persistent_mode" {
  default = "persistentMode"
}

variable "asset_compute_disks_size_in_mbs" {
  default = 10
}

variable "asset_compute_disks_uuid" {
  default = "uuid"
}

variable "asset_compute_disks_uuid_lun" {
  default = "uuidLun"
}

variable "asset_compute_disks_count" {
  default = 10
}

variable "asset_compute_dns_name" {
  default = "dnsName"
}

variable "asset_compute_firmware" {
  default = "firmware"
}

variable "asset_compute_gpu_devices_cores_count" {
  default = 10
}

variable "asset_compute_gpu_devices_description" {
  default = "description"
}

variable "asset_compute_gpu_devices_manufacturer" {
  default = "manufacturer"
}

variable "asset_compute_gpu_devices_memory_in_mbs" {
  default = 10
}

variable "asset_compute_gpu_devices_name" {
  default = "name"
}

variable "asset_compute_gpu_devices_count" {
  default = 10
}

variable "asset_compute_guest_state" {
  default = "guestState"
}

variable "asset_compute_hardware_version" {
  default = "hardwareVersion"
}

variable "asset_compute_host_name" {
  default = "hostName"
}

variable "asset_compute_is_pmem_enabled" {
  default = false
}

variable "asset_compute_is_tpm_enabled" {
  default = false
}

variable "asset_compute_latency_sensitivity" {
  default = "latencySensitivity"
}

variable "asset_compute_memory_in_mbs" {
  default = 10
}

variable "asset_compute_nics_ip_addresses" {
  default = []
}

variable "asset_compute_nics_label" {
  default = "label"
}

variable "asset_compute_nics_mac_address" {
  default = "macAddress"
}

variable "asset_compute_nics_mac_address_type" {
  default = "macAddressType"
}

variable "asset_compute_nics_network_name" {
  default = "networkName"
}

variable "asset_compute_nics_switch_name" {
  default = "switchName"
}

variable "asset_compute_nics_count" {
  default = 10
}

variable "asset_compute_nvdimm_controller_bus_number" {
  default = 10
}

variable "asset_compute_nvdimm_controller_label" {
  default = "label"
}

variable "asset_compute_nvdimms_controller_key" {
  default = 10
}

variable "asset_compute_nvdimms_label" {
  default = "label"
}

variable "asset_compute_nvdimms_unit_number" {
  default = 10
}

variable "asset_compute_operating_system" {
  default = "operatingSystem"
}

variable "asset_compute_operating_system_version" {
  default = "operatingSystemVersion"
}

variable "asset_compute_pmem_in_mbs" {
  default = 10
}

variable "asset_compute_power_state" {
  default = "powerState"
}

variable "asset_compute_primary_ip" {
  default = "primaryIp"
}

variable "asset_compute_scsi_controller_label" {
  default = "label"
}

variable "asset_compute_scsi_controller_shared_bus" {
  default = "sharedBus"
}

variable "asset_compute_scsi_controller_unit_number" {
  default = 10
}

variable "asset_compute_storage_provisioned_in_mbs" {
  default = 10
}

variable "asset_compute_threads_per_core_count" {
  default = 10
}

variable "asset_display_name" {
  default = "displayName"
}

variable "asset_inventory_asset_type" {
  default = "INVENTORY_ASSET"
}

variable "asset_inventory_asset_class_name" {
  default = "com.oracle.pic.ocb.discovery.model.OlvmStorageDomainAssetDetails"
}

variable "asset_inventory_asset_class_version" {
  default = "0"
}

variable "asset_inventory_display_name" {
  default = "kvm-storagedomain1"
}

variable "asset_inventory_external_asset_key" {
  default = "1a10b288-f688-47d2-b38d-22dede44ba8a"
}

variable "asset_inventory_source_key" {
  default = "https://11.0.11.131:443/ovirt-engine/api"
}

variable "asset_external_asset_key" {
  default = "externalAssetKey"
}

variable "asset_source_key" {
  default = "sourceKey"
}

variable "asset_state" {
  default = "ACTIVE"
}

variable "asset_vm_hypervisor_host" {
  default = "hypervisorHost"
}

variable "asset_vm_hypervisor_vendor" {
  default = "hypervisorVendor"
}

variable "asset_vm_hypervisor_version" {
  default = "hypervisorVersion"
}

variable "asset_vmware_vcenter_data_center" {
  default = "dataCenter"
}

variable "asset_vmware_vcenter_vcenter_key" {
  default = "vcenterKey"
}

variable "asset_vmware_vcenter_vcenter_version" {
  default = "vcenterVersion"
}

variable "asset_vmware_vm_cluster" {
  default = "cluster"
}

variable "asset_vmware_vm_customer_fields" {
  default = []
}

variable "asset_vmware_vm_customer_tags_description" {
  default = "description"
}

variable "asset_vmware_vm_customer_tags_name" {
  default = "name"
}

variable "asset_vmware_vm_fault_tolerance_bandwidth" {
  default = 10
}

variable "asset_vmware_vm_fault_tolerance_secondary_latency" {
  default = 10
}

variable "asset_vmware_vm_fault_tolerance_state" {
  default = "faultToleranceState"
}

variable "asset_vmware_vm_instance_uuid" {
  default = "instanceUuid"
}

variable "asset_vmware_vm_is_disks_cbt_enabled" {
  default = false
}

variable "asset_vmware_vm_is_disks_uuid_enabled" {
  default = false
}

variable "asset_vmware_vm_path" {
  default = "path"
}

variable "asset_vmware_vm_vmware_tools_status" {
  default = "vmwareToolsStatus"
}

variable "asset_aws_ec2_asset_type" {
  default = "AWS_EC2"
}

variable "asset_aws_ebs_asset_type" {
  default = "AWS_EBS"
}

variable "asset_aws_ec2_source_key" {
  default = "ec2SourceKey"
}

variable "asset_aws_ec2_external_asset_key" {
  default = "ec2ExternalAssetKey"
}

variable "asset_aws_ebs_source_key" {
  default = "ebs_source_key"
}

variable "asset_aws_ebs_external_asset_key" {
  default = "ebsExternalAssetKey"
}

variable "asset_aws_architecture" { default = "architecture" }

variable "asset_aws_are_elastic_inference_accelerators_present" { default = false }

variable "asset_aws_boot_mode" { default = "bootMode" }

variable "asset_aws_capacity_reservation_key" { default = "capacityReservationKey" }

variable "asset_aws_image_key" { default = "imageKey" }

variable "asset_aws_instance_key" { default = "instanceKey" }

variable "asset_aws_instance_lifecycle" { default = "instanceLifecycle" }

variable "asset_aws_instance_type" { default = "instanceType" }

variable "asset_aws_ip_address" { default = "ipAddress" }

variable "asset_aws_ipv6address" { default = "ipv6Address" }

variable "asset_aws_is_enclave_options" { default = false }

variable "asset_aws_is_hibernation_options" { default = false }

variable "asset_aws_is_source_dest_check" { default = false }

variable "asset_aws_is_spot_instance" { default = false }

variable "asset_aws_kernel_key" { default = false }

variable "asset_aws_licenses" { default = ["licenses"] }

variable "asset_aws_maintenance_options" { default = "maintenanceOptions" }

variable "asset_aws_monitoring" { default = "monitoring" }

variable "asset_aws_private_dns_name" { default = "privateDnsName" }

variable "asset_aws_private_ip_address" { default = "privateIpAddress" }

variable "asset_aws_root_device_name" { default = "rootDeviceName" }

variable "asset_aws_root_device_type" { default = "rootDeviceType" }

variable "asset_aws_sriov_net_support" { default = "sriovNetSupport" }

variable "asset_aws_subnet_key" { default = "subnetKey" }

variable "asset_aws_tpm_support" { default = "tpmSupport" }

variable "asset_aws_virtualization_type" { default = "virtualizationType" }

variable "asset_aws_vpc_key" { default = "vpcKey" }

variable "asset_aws_description" { default = "description" }

variable "asset_aws_interface_type" { default = "interfaceType" }

variable "asset_aws_ipv4prefixes" { default = ["ipv4Prefixes"] }

variable "asset_aws_ipv6addresses" { default = ["ipv6Addresses"] }

variable "asset_aws_ipv6prefixes" { default = ["ipv6Prefixes"] }

variable "asset_aws_mac_address" { default = "macAddress" }

variable "asset_aws_network_interface_key" { default = "networkInterfaceKey" }

variable "asset_aws_owner_key" { default = "ownerKey" }

variable "asset_aws_status" { default = "status" }

variable "asset_aws_carrier_ip" { default = "carrierIp" }

variable "asset_aws_customer_owned_ip" { default = "customerOwnedIp" }

variable "asset_aws_ip_owner_key" { default = "ipOwnerKey" }

variable "asset_aws_public_dns_name" { default = "publicDnsName" }

variable "asset_aws_public_ip" { default = "publicIp" }

variable "asset_aws_attachment_key" { default = "attachmentKey" }

variable "asset_aws_device_index" { default = 10 }

variable "asset_aws_is_delete_on_termination" { default = false }

variable "asset_aws_network_card_index" { default = 10 }

variable "asset_aws_is_primary" { default = false }

variable "asset_aws_group_key" { default = "groupKey" }

variable "asset_aws_group_name" { default = "groupName" }

variable "asset_aws_affinity" { default = "affinity" }

variable "asset_aws_availability_zone" { default = "availabilityZone" }

variable "asset_aws_host_key" { default = "hostKey" }

variable "asset_aws_host_resource_group_arn" { default = "hostResourceGroupArn" }

variable "asset_aws_partition_number" { default = 10 }

variable "asset_aws_spread_domain" { default = "spreadDomain" }

variable "asset_aws_tenancy" { default = "tenancy" }

variable "asset_aws_code" { default = 10 }

variable "asset_aws_name" { default = "name" }

variable "asset_aws_key" { default = "key" }

variable "asset_aws_value" { default = "value" }

variable "asset_aws_ec2cost_amount" { default = 10 }

variable "asset_aws_ec2cost_currency_code" { default = "USD" }

variable "asset_aws_iops" { default = 10 }

variable "asset_aws_is_encrypted" { default = false }

variable "asset_aws_is_multi_attach_enabled" { default = false }

variable "asset_aws_size_in_gi_bs" { default = 10 }

variable "asset_aws_throughput" { default = 10 }

variable "asset_aws_volume_key" { default = "volumeKey" }

variable "asset_aws_volume_type" { default = "volumeType" }

variable "asset_aws_device" { default = "device" }


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  # version             = "8.3.0"
}

resource "oci_cloud_bridge_asset" "test_asset" {
  asset_type         = var.asset_asset_type
  compartment_id     = var.compartment_id
  external_asset_key = var.asset_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_source_key
  asset_source_ids   = var.asset_asset_source_ids
  display_name       = var.asset_display_name
  compute {

    #Optional
    connected_networks = var.asset_compute_connected_networks
    cores_count        = var.asset_compute_cores_count
    cpu_model          = var.asset_compute_cpu_model
    description        = var.asset_compute_description
    disks {

      #Optional
      boot_order      = var.asset_compute_disks_boot_order
      location        = var.asset_compute_disks_location
      name            = var.asset_compute_disks_name
      persistent_mode = var.asset_compute_disks_persistent_mode
      size_in_mbs     = var.asset_compute_disks_size_in_mbs
      uuid            = var.asset_compute_disks_uuid
      uuid_lun        = var.asset_compute_disks_uuid_lun
    }
    disks_count = var.asset_compute_disks_count
    dns_name    = var.asset_compute_dns_name
    firmware    = var.asset_compute_firmware
    gpu_devices {

      #Optional
      cores_count   = var.asset_compute_gpu_devices_cores_count
      description   = var.asset_compute_gpu_devices_description
      manufacturer  = var.asset_compute_gpu_devices_manufacturer
      memory_in_mbs = var.asset_compute_gpu_devices_memory_in_mbs
      name          = var.asset_compute_gpu_devices_name
    }
    gpu_devices_count   = var.asset_compute_gpu_devices_count
    guest_state         = var.asset_compute_guest_state
    hardware_version    = var.asset_compute_hardware_version
    host_name           = var.asset_compute_host_name
    is_pmem_enabled     = var.asset_compute_is_pmem_enabled
    is_tpm_enabled      = var.asset_compute_is_tpm_enabled
    latency_sensitivity = var.asset_compute_latency_sensitivity
    memory_in_mbs       = var.asset_compute_memory_in_mbs
    nics {

      #Optional
      ip_addresses     = var.asset_compute_nics_ip_addresses
      label            = var.asset_compute_nics_label
      mac_address      = var.asset_compute_nics_mac_address
      mac_address_type = var.asset_compute_nics_mac_address_type
      network_name     = var.asset_compute_nics_network_name
      switch_name      = var.asset_compute_nics_switch_name
    }
    nics_count = var.asset_compute_nics_count
    nvdimm_controller {

      #Optional
      bus_number = var.asset_compute_nvdimm_controller_bus_number
      label      = var.asset_compute_nvdimm_controller_label
    }
    nvdimms {

      #Optional
      controller_key = var.asset_compute_nvdimms_controller_key
      label          = var.asset_compute_nvdimms_label
      unit_number    = var.asset_compute_nvdimms_unit_number
    }
    operating_system         = var.asset_compute_operating_system
    operating_system_version = var.asset_compute_operating_system_version
    pmem_in_mbs              = var.asset_compute_pmem_in_mbs
    power_state              = var.asset_compute_power_state
    primary_ip               = var.asset_compute_primary_ip
    scsi_controller {

      #Optional
      label       = var.asset_compute_scsi_controller_label
      shared_bus  = var.asset_compute_scsi_controller_shared_bus
      unit_number = var.asset_compute_scsi_controller_unit_number
    }
    storage_provisioned_in_mbs = var.asset_compute_storage_provisioned_in_mbs
    threads_per_core_count     = var.asset_compute_threads_per_core_count
  }
  vm {

    #Optional
    hypervisor_host    = var.asset_vm_hypervisor_host
    hypervisor_vendor  = var.asset_vm_hypervisor_vendor
    hypervisor_version = var.asset_vm_hypervisor_version
  }
  vmware_vcenter {

    #Optional
    data_center     = var.asset_vmware_vcenter_data_center
    vcenter_key     = var.asset_vmware_vcenter_vcenter_key
    vcenter_version = var.asset_vmware_vcenter_vcenter_version
  }
  vmware_vm {

    #Optional
    cluster         = var.asset_vmware_vm_cluster
    customer_fields = var.asset_vmware_vm_customer_fields
    customer_tags {

      #Optional
      description = var.asset_vmware_vm_customer_tags_description
      name        = var.asset_vmware_vm_customer_tags_name
    }
    fault_tolerance_bandwidth         = var.asset_vmware_vm_fault_tolerance_bandwidth
    fault_tolerance_secondary_latency = var.asset_vmware_vm_fault_tolerance_secondary_latency
    fault_tolerance_state             = var.asset_vmware_vm_fault_tolerance_state
    instance_uuid                     = var.asset_vmware_vm_instance_uuid
    is_disks_cbt_enabled              = var.asset_vmware_vm_is_disks_cbt_enabled
    is_disks_uuid_enabled             = var.asset_vmware_vm_is_disks_uuid_enabled
    path                              = var.asset_vmware_vm_path
    vmware_tools_status               = var.asset_vmware_vm_vmware_tools_status
  }
}

data "oci_cloud_bridge_assets" "test_assets" {
  compartment_id     = var.compartment_id
  asset_id           = oci_cloud_bridge_asset.test_asset.id
  asset_type         = var.asset_asset_type
  display_name       = var.asset_display_name
  external_asset_key = var.asset_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_source_key
  state              = var.asset_state
}

resource "oci_cloud_bridge_asset" "test_aws_ec2_asset" {
  asset_type         = var.asset_aws_ec2_asset_type
  compartment_id     = var.compartment_id
  external_asset_key = var.asset_aws_ec2_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_aws_ec2_source_key
  asset_source_ids   = var.asset_asset_source_ids
  display_name       = var.asset_display_name
  compute {

    #Optional
    connected_networks = var.asset_compute_connected_networks
    cores_count        = var.asset_compute_cores_count
    cpu_model          = var.asset_compute_cpu_model
    description        = var.asset_compute_description
    disks {

      #Optional
      boot_order      = var.asset_compute_disks_boot_order
      location        = var.asset_compute_disks_location
      name            = var.asset_compute_disks_name
      persistent_mode = var.asset_compute_disks_persistent_mode
      size_in_mbs     = var.asset_compute_disks_size_in_mbs
      uuid            = var.asset_compute_disks_uuid
      uuid_lun        = var.asset_compute_disks_uuid_lun
    }
    disks_count = var.asset_compute_disks_count
    dns_name    = var.asset_compute_dns_name
    firmware    = var.asset_compute_firmware
    gpu_devices {

      #Optional
      cores_count   = var.asset_compute_gpu_devices_cores_count
      description   = var.asset_compute_gpu_devices_description
      manufacturer  = var.asset_compute_gpu_devices_manufacturer
      memory_in_mbs = var.asset_compute_gpu_devices_memory_in_mbs
      name          = var.asset_compute_gpu_devices_name
    }
    gpu_devices_count   = var.asset_compute_gpu_devices_count
    guest_state         = var.asset_compute_guest_state
    hardware_version    = var.asset_compute_hardware_version
    host_name           = var.asset_compute_host_name
    is_pmem_enabled     = var.asset_compute_is_pmem_enabled
    is_tpm_enabled      = var.asset_compute_is_tpm_enabled
    latency_sensitivity = var.asset_compute_latency_sensitivity
    memory_in_mbs       = var.asset_compute_memory_in_mbs
    nics {

      #Optional
      ip_addresses     = var.asset_compute_nics_ip_addresses
      label            = var.asset_compute_nics_label
      mac_address      = var.asset_compute_nics_mac_address
      mac_address_type = var.asset_compute_nics_mac_address_type
      network_name     = var.asset_compute_nics_network_name
      switch_name      = var.asset_compute_nics_switch_name
    }
    nics_count = var.asset_compute_nics_count
    nvdimm_controller {

      #Optional
      bus_number = var.asset_compute_nvdimm_controller_bus_number
      label      = var.asset_compute_nvdimm_controller_label
    }
    nvdimms {

      #Optional
      controller_key = var.asset_compute_nvdimms_controller_key
      label          = var.asset_compute_nvdimms_label
      unit_number    = var.asset_compute_nvdimms_unit_number
    }
    operating_system         = var.asset_compute_operating_system
    operating_system_version = var.asset_compute_operating_system_version
    pmem_in_mbs              = var.asset_compute_pmem_in_mbs
    power_state              = var.asset_compute_power_state
    primary_ip               = var.asset_compute_primary_ip
    scsi_controller {

      #Optional
      label       = var.asset_compute_scsi_controller_label
      shared_bus  = var.asset_compute_scsi_controller_shared_bus
      unit_number = var.asset_compute_scsi_controller_unit_number
    }
    storage_provisioned_in_mbs = var.asset_compute_storage_provisioned_in_mbs
    threads_per_core_count     = var.asset_compute_threads_per_core_count
  }
  vm {

    #Optional
    hypervisor_host    = var.asset_vm_hypervisor_host
    hypervisor_vendor  = var.asset_vm_hypervisor_vendor
    hypervisor_version = var.asset_vm_hypervisor_version
  }
  aws_ec2 {
    architecture                               = var.asset_aws_architecture
    are_elastic_inference_accelerators_present = var.asset_aws_are_elastic_inference_accelerators_present
    boot_mode                                  = var.asset_aws_boot_mode
    capacity_reservation_key                   = var.asset_aws_capacity_reservation_key
    image_key                                  = var.asset_aws_image_key
    instance_key                               = var.asset_aws_instance_key
    instance_lifecycle                         = var.asset_aws_instance_lifecycle
    instance_type                              = var.asset_aws_instance_type
    ip_address                                 = var.asset_aws_ip_address
    ipv6address                                = var.asset_aws_ipv6address
    is_enclave_options                         = var.asset_aws_is_enclave_options
    is_hibernation_options                     = var.asset_aws_is_hibernation_options
    is_source_dest_check                       = var.asset_aws_is_source_dest_check
    is_spot_instance                           = var.asset_aws_is_spot_instance
    kernel_key                                 = var.asset_aws_kernel_key
    licenses                                   = var.asset_aws_licenses
    maintenance_options                        = var.asset_aws_maintenance_options
    monitoring                                 = var.asset_aws_monitoring
    network_interfaces {
      association {
        carrier_ip        = var.asset_aws_carrier_ip
        customer_owned_ip = var.asset_aws_customer_owned_ip
        ip_owner_key      = var.asset_aws_ip_owner_key
        public_dns_name   = var.asset_aws_public_dns_name
        public_ip         = var.asset_aws_public_ip
      }
      attachment {
        attachment_key           = var.asset_aws_attachment_key
        device_index             = var.asset_aws_device_index
        is_delete_on_termination = var.asset_aws_is_delete_on_termination
        network_card_index       = var.asset_aws_network_card_index
        status                   = var.asset_aws_status
      }
      description           = var.asset_aws_description
      interface_type        = var.asset_aws_interface_type
      ipv4prefixes          = var.asset_aws_ipv4prefixes
      ipv6addresses         = var.asset_aws_ipv6addresses
      ipv6prefixes          = var.asset_aws_ipv6prefixes
      is_source_dest_check  = var.asset_aws_is_source_dest_check
      mac_address           = var.asset_aws_mac_address
      network_interface_key = var.asset_aws_network_interface_key
      owner_key             = var.asset_aws_owner_key
      private_ip_addresses {
        association {
          carrier_ip        = var.asset_aws_carrier_ip
          customer_owned_ip = var.asset_aws_customer_owned_ip
          ip_owner_key      = var.asset_aws_ip_owner_key
          public_dns_name   = var.asset_aws_public_dns_name
          public_ip         = var.asset_aws_public_ip
        }
        is_primary         = var.asset_aws_is_primary
        private_dns_name   = var.asset_aws_private_dns_name
        private_ip_address = var.asset_aws_private_ip_address
      }
      security_groups {
        group_key  = var.asset_aws_group_key
        group_name = var.asset_aws_group_name
      }
      status     = var.asset_aws_status
      subnet_key = var.asset_aws_subnet_key
    }
    placement {
      affinity                = var.asset_aws_affinity
      availability_zone       = var.asset_aws_availability_zone
      group_name              = var.asset_aws_group_name
      host_key                = var.asset_aws_host_key
      host_resource_group_arn = var.asset_aws_host_resource_group_arn
      partition_number        = var.asset_aws_partition_number
      spread_domain           = var.asset_aws_spread_domain
      tenancy                 = var.asset_aws_tenancy
    }
    private_dns_name   = var.asset_aws_private_dns_name
    private_ip_address = var.asset_aws_private_ip_address
    root_device_name   = var.asset_aws_root_device_name
    root_device_type   = var.asset_aws_root_device_type
    security_groups {
      group_key  = var.asset_aws_group_key
      group_name = var.asset_aws_group_name
    }
    sriov_net_support = var.asset_aws_sriov_net_support
    state {
      code = var.asset_aws_code
      name = var.asset_aws_name
    }
    subnet_key = var.asset_aws_subnet_key
    tags {
      key   = var.asset_aws_key
      value = var.asset_aws_value
    }
    tpm_support         = var.asset_aws_tpm_support
    virtualization_type = var.asset_aws_virtualization_type
    vpc_key             = var.asset_aws_vpc_key
  }
  aws_ec2cost {
    amount        = var.asset_aws_ec2cost_amount
    currency_code = var.asset_aws_ec2cost_currency_code
  }
  attached_ebs_volumes_cost {
    amount        = var.asset_aws_ec2cost_amount
    currency_code = var.asset_aws_ec2cost_currency_code
  }
}

data "oci_cloud_bridge_assets" "test_aws_ec2_assets" {
  compartment_id     = var.compartment_id
  asset_id           = oci_cloud_bridge_asset.test_aws_ec2_asset.id
  asset_type         = var.asset_aws_ec2_asset_type
  display_name       = var.asset_display_name
  external_asset_key = var.asset_aws_ec2_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_aws_ec2_source_key
  state              = var.asset_state
}

resource "oci_cloud_bridge_asset" "test_aws_ebs_asset" {
  asset_type         = var.asset_aws_ebs_asset_type
  compartment_id     = var.compartment_id
  external_asset_key = var.asset_aws_ebs_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_aws_ebs_source_key
  asset_source_ids   = var.asset_asset_source_ids
  display_name       = var.asset_display_name
  aws_ebs {
    attachments {
      device                   = var.asset_aws_device
      instance_key             = var.asset_aws_instance_key
      is_delete_on_termination = var.asset_aws_is_delete_on_termination
      status                   = var.asset_aws_status
      volume_key               = var.asset_aws_volume_key
    }
    availability_zone       = var.asset_aws_availability_zone
    iops                    = var.asset_aws_iops
    is_encrypted            = var.asset_aws_is_encrypted
    is_multi_attach_enabled = var.asset_aws_is_multi_attach_enabled
    size_in_gi_bs           = var.asset_aws_size_in_gi_bs
    status                  = var.asset_aws_status
    tags {
      key   = var.asset_aws_key
      value = var.asset_aws_value
    }
    throughput  = var.asset_aws_throughput
    volume_key  = var.asset_aws_volume_key
    volume_type = var.asset_aws_volume_type
  }
}

data "oci_cloud_bridge_assets" "test_aws_ebs_assets" {
  compartment_id     = var.compartment_id
  asset_id           = oci_cloud_bridge_asset.test_aws_ebs_asset.id
  asset_type         = var.asset_aws_ebs_asset_type
  display_name       = var.asset_display_name
  external_asset_key = var.asset_aws_ebs_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_aws_ebs_source_key
  state              = var.asset_state
}

resource "oci_cloud_bridge_asset" "test_inventory_asset" {
  asset_type          = var.asset_inventory_asset_type
  asset_class_name    = var.asset_inventory_asset_class_name
  asset_class_version = var.asset_inventory_asset_class_version
  asset_details = jsonencode({
    olvmStorageDomain = {
      availableSpaceInBytes = 643171352576
      storageDomainName     = "kvm-storagedomain1"
    }
  })
  compartment_id     = var.compartment_id
  external_asset_key = var.asset_inventory_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_inventory_source_key
  asset_source_ids   = var.asset_asset_source_ids
  display_name       = var.asset_inventory_display_name
}

data "oci_cloud_bridge_assets" "test_inventory_assets" {
  compartment_id     = var.compartment_id
  asset_id           = oci_cloud_bridge_asset.test_inventory_asset.id
  asset_type         = var.asset_inventory_asset_type
  display_name       = var.asset_inventory_display_name
  external_asset_key = var.asset_inventory_external_asset_key
  inventory_id       = var.inventory_id
  source_key         = var.asset_inventory_source_key
  state              = var.asset_state
}
