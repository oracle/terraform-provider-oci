variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "subnet_id" {}

provider "oci" {
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region = "us-ashburn-1"
}

resource "oci_ai_vision_project" "test_project" {
  compartment_id = var.compartment_ocid
}

resource "oci_ai_vision_model" "test_model" {
  compartment_id = var.compartment_ocid
  is_quick_mode = "false"
  max_training_duration_in_hours = 0.6
  model_type = "IMAGE_CLASSIFICATION"
  project_id = oci_ai_vision_project.test_project.id
  training_dataset {
    bucket = "Test"
    dataset_type = "OBJECT_STORAGE"
    namespace_name = "axfelw9p2fyr"
    object = "0_ic_multi-2labels_combined.jsonl"
  }
}

resource "oci_ai_vision_vision_private_endpoint" "test_vision_private_endpoint" {
  compartment_id = var.compartment_ocid
  subnet_id      = var.subnet_id
}

resource "oci_ai_vision_stream_source" "test_stream_source" {
  compartment_id = var.compartment_ocid
  stream_source_details {
    camera_url = "rtsp://64.2.1.212"
    source_type = "RTSP"
    stream_network_access_details {
      stream_access_type = "PRIVATE"
      private_endpoint_id = oci_ai_vision_vision_private_endpoint.test_vision_private_endpoint.id
    }
  }
}

resource "oci_ai_vision_stream_job" "test_stream_job" {
  compartment_id = var.compartment_ocid
  features {
    feature_type = "FACE_DETECTION"
  }
  stream_output_location {
    bucket = "Test"
    namespace = "axfelw9p2fyr"
    output_location_type = "OBJECT_STORAGE"
    prefix = "prefix"
  }
  stream_source_id = oci_ai_vision_stream_source.test_stream_source.id
}

resource "oci_ai_vision_stream_group" "test_stream_group" {
  compartment_id = var.compartment_ocid
  stream_source_ids = [oci_ai_vision_stream_source.test_stream_source.id]
}

data "oci_ai_vision_vision_private_endpoints" "test_vision_private_endpoints" {
  compartment_id = var.compartment_ocid
}

data "oci_ai_vision_vision_private_endpoint" "test_vision_private_endpoints" {
  vision_private_endpoint_id = oci_ai_vision_vision_private_endpoint.test_vision_private_endpoint.id
}

data "oci_ai_vision_stream_source" "test_stream_source" {
  stream_source_id = oci_ai_vision_stream_source.test_stream_source.id
}

data "oci_ai_vision_stream_sources" "test_stream_sources" {
  compartment_id = var.compartment_ocid
}

data "oci_ai_vision_stream_job" "test_stream_job" {
  stream_job_id = oci_ai_vision_stream_job.test_stream_job.id
}

data "oci_ai_vision_stream_jobs" "test_stream_jobs" {
  compartment_id = var.compartment_ocid
}

data "oci_ai_vision_stream_group" "test_stream_groups" {
  stream_group_id = oci_ai_vision_stream_group.test_stream_group.id
}

data "oci_ai_vision_stream_groups" "test_stream_groups" {
  compartment_id = var.compartment_ocid
}
