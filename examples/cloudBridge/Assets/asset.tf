// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

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

variable "asset_defined_tags_value" {
  default = "value"
}

variable "asset_display_name" {
  default = "displayName"
}

variable "asset_external_asset_key" {
  default = "externalAssetKey"
}

variable "asset_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "asset_source_key" {
  default = "sourceKey"
}

variable "asset_state" {
  default = "AVAILABLE"
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



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_asset" "test_asset" {
  #Required
  asset_type         = var.asset_asset_type
  compartment_id     = var.compartment_id
  external_asset_key = var.asset_external_asset_key
  inventory_id       = oci_cloud_bridge_inventory.test_inventory.id
  source_key         = var.asset_source_key

  #Optional
  asset_source_ids = var.asset_asset_source_ids
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
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.asset_defined_tags_value)
  display_name  = var.asset_display_name
  freeform_tags = var.asset_freeform_tags
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
  #Required
  compartment_id = var.compartment_id

  #Optional
  asset_id           = oci_cloud_bridge_asset.test_asset.id
  asset_type         = var.asset_asset_type
  display_name       = var.asset_display_name
  external_asset_key = var.asset_external_asset_key
  inventory_id       = oci_cloud_bridge_inventory.test_inventory.id
  source_key         = var.asset_source_key
  state              = var.asset_state
}

