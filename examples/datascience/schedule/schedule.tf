// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_id" {}
variable "data_science_project_id" {}
variable "data_science_job_id" {}
variable "data_science_ml_app_instance_view_id" {}
variable "log_group_id" {}
variable "log_id" {}

variable "data_science_job_run_display_name" {
  default = "Job Run"
}

variable "schedule_description" {
  default = "description"
}

variable "schedule_display_name" {
  default = "Schedule"
}

variable "schedule_state" {
  default = "ACTIVE"
}

variable "schedule_id" {
  default = "id"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


//CRON Schedule for creating Job Run
resource "oci_datascience_schedule" "cron_schedule_example" {

  compartment_id = var.compartment_id
  display_name   = var.schedule_display_name
  project_id     = var.data_science_project_id
  description    = var.schedule_description

  #Required
  action {
    action_type = "HTTP"
    action_details {
      http_action_type = "CREATE_JOB_RUN"
      create_job_run_details {
        project_id = var.data_science_project_id
        compartment_id = var.compartment_id
        job_id = var.data_science_job_id
        display_name = var.data_science_job_run_display_name
      }
    }
  }

  #Required
  trigger {
    trigger_type = "CRON"
    cron_expression = "11 11 * * *"
    time_start = "2055-01-01T05:00:00.000Z"
    time_end = "2065-01-01T05:00:00.000Z"
  }

  #Optional
  log_details {
    log_group_id = var.log_group_id
    log_id = var.log_id
  }
}

//Interval Schedule for involing Ml Application Instance View Trigger
resource "oci_datascience_schedule" "interval_schedule_example" {

  compartment_id = var.compartment_id
  display_name   = var.schedule_display_name
  project_id     = var.data_science_project_id
  description    = var.schedule_description

  #Required
  action {
    action_type = "HTTP"
    action_details {
      http_action_type = "INVOKE_ML_APPLICATION_PROVIDER_TRIGGER"
      ml_application_instance_view_id = var.data_science_ml_app_instance_view_id
      trigger_ml_application_instance_view_flow_details {
        trigger_name = "TrainingTrigger"
      }
    }
  }

  #Required
  trigger {
    trigger_type = "INTERVAL"
    frequency = "HOURLY"
    interval = "12"
    is_random_start_time = "true"
  }
}


//iCal Schedule for creating Job Runs
resource "oci_datascience_schedule" "ical_schedule_example" {

  compartment_id = var.compartment_id
  display_name   = var.schedule_display_name
  project_id     = var.data_science_project_id
  description    = var.schedule_description

  #Required
  action {
    action_type = "HTTP"
    action_details {
      http_action_type = "CREATE_JOB_RUN"
      create_job_run_details {
        project_id = var.data_science_project_id
        compartment_id = var.compartment_id
        job_id = var.data_science_job_id
        display_name = var.data_science_job_run_display_name
      }
    }
  }

  #Required
  trigger {
    trigger_type = "ICAL"
    recurrence = "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;INTERVAL=1"
  }
}


# Plural data source
data "oci_datascience_schedules" "test_schedules" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.schedule_display_name
  state        = var.schedule_state
}

# Singular data source
data "oci_datascience_schedule" "test_schedule" {
  #Required
  schedule_id  = var.schedule_id
}