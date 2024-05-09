
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "osmh_managed_instance_windows_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

#################################################################
## Prepare Scheduled Job Targets
#################################################################
# Reference OL8 software source
data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
  arch_type = ["X86_64"]
  availability = ["SELECTED"]
  compartment_id = var.compartment_id
  display_name = "ol8_baseos_latest-x86_64"
  os_family = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state = ["ACTIVE"]
  vendor_name = "ORACLE"
}

# Managed instance - Windows 2022
resource "oci_os_management_hub_managed_instance" "test_managed_instance_windows" {
  managed_instance_id = var.osmh_managed_instance_windows_ocid
}

# Managed instance group - OL8
resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
  #Required
  arch_type = "X86_64"
  compartment_id = var.compartment_id
  display_name = "displayNameExample"
  os_family = "ORACLE_LINUX_8"
  software_source_ids = [
    data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
  ]
  vendor_name = "ORACLE"
}

# Lifecycle environment - OL8
resource "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
  #Required
  arch_type = "X86_64"
  compartment_id = var.compartment_id
  display_name = "displayName"
  os_family = "ORACLE_LINUX_8"
  stages {
    #Required
    display_name = "test"
    rank = "1"

    #Optional
    defined_tags = {"Operations.CostCenter"= "42"}
    freeform_tags = {"Department"= "Finance"}
  }
  stages {
    #Required
    display_name = "prod"
    rank = "2"

    #Optional
    defined_tags = {"Operations.CostCenter"= "42"}
    freeform_tags = {"Department"= "Finance"}
  }
  vendor_name = "ORACLE"
}

#################################################################
## Create Scheduled Jobs
#################################################################
# Scheduled job on a windows instance - install windows updates
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_windows_update" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "INSTALL_OTHER_WINDOWS_UPDATES"
  }
  schedule_type = "ONETIME"
  time_next_execution = "2024-03-27T23:00:49.382Z"

  # Optional
  managed_instance_ids = [
    oci_os_management_hub_managed_instance.test_managed_instance_windows.id
  ]
}

# Scheduled job on a windows instance - recurring install windows updates
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_windows_update_recurring" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "INSTALL_OTHER_WINDOWS_UPDATES"
  }
  schedule_type = "RECURRING"
  recurring_rule = "FREQ=DAILY;INTERVAL=1"
  time_next_execution = "2024-03-27T23:00:49.382Z"

  # Optional
  managed_instance_ids = [
    oci_os_management_hub_managed_instance.test_managed_instance_windows.id
  ]
}

# Scheduled job on a compartment - install other updates
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_compartment_other_update" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "UPDATE_OTHER"
  }
  schedule_type = "ONETIME"
  time_next_execution = "2030-03-27T23:00:49.382Z"

  # Optional
  managed_compartment_ids = [
    var.compartment_id
  ]
}

# Scheduled job on managed instance group - install packages
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_install_packages_on_group" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "INSTALL_PACKAGES"
    package_names = ["InvalidPackage"]
  }
  schedule_type = "ONETIME"
  time_next_execution = "2024-03-27T23:00:49.382Z"

  # Optional
  managed_instance_group_ids = [
    oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  ]
}

# Scheduled job on managed instance group - update ksplice userspace
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_update_ksplice_userspace_on_group" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "UPDATE_KSPLICE_USERSPACE"
  }
  schedule_type = "ONETIME"
  time_next_execution = "2024-03-27T23:00:49.382Z"

  # Optional
  managed_instance_group_ids = [
    oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  ]
}

# Scheduled job on managed instance group - switch module streams
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_switch_module_stream_on_group" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "MANAGE_MODULE_STREAMS"
    manage_module_streams_details {
      enable {
        module_name = "fakeModule"
        stream_name = "fakeStream"
      }
    }
  }
  schedule_type = "ONETIME"
  time_next_execution = "2024-03-27T23:00:49.382Z"

  # Optional
  managed_instance_group_ids = [
    oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  ]
}

# Scheduled job on lifecycle stage - promote
resource "oci_os_management_hub_scheduled_job" "test_scheduled_job_promote_lifecycle_stage" {
  # Required
  compartment_id = var.compartment_id
  operations {
    operation_type = "PROMOTE_LIFECYCLE"
  }
  schedule_type = "ONETIME"
  time_next_execution = "2024-03-27T23:00:49.382Z"

  # Optional
  lifecycle_stage_ids = [
    oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id
  ]
}
#################################################################
## List Scheduled Jobs
#################################################################
data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_compartment" {
  # Required
  compartment_id = var.compartment_id
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_job_id" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  id = oci_os_management_hub_scheduled_job.test_scheduled_job_install_packages_on_group.id
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_operation_type" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  operation_type = "INSTALL_PACKAGES"
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_schedule_type" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  schedule_type = "RECURRING"
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_location" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  location = ["OCI_COMPUTE"]
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_mig_id" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_lcs_id" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  lifecycle_stage_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id
}

data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs_mi_id" {
  # Required
  compartment_id = var.compartment_id
  # Optional
  managed_instance_id = var.osmh_managed_instance_windows_ocid
}

#################################################################
## Get Scheduled Job
#################################################################
data "oci_os_management_hub_scheduled_job" "test_scheduled_job" {
  # Required
  scheduled_job_id = oci_os_management_hub_scheduled_job.test_scheduled_job_install_packages_on_group.id
}