variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "cluster_endpoint" {
  description = "Spark master URL used for job submission."
  type        = string
  default     = "spark://spark-master.example.com:7077"
}

variable "objectstorage_namespace" {
  description = "Object Storage namespace containing the artifact bucket."
  type        = string
}

variable "objectstorage_bucket_name" {
  description = "Object Storage bucket used for Spark application artifacts."
  type        = string
}

variable "artifact_path" {
  description = "Path within the artifact bucket."
  type        = string
  default     = "applications"
}

variable "display_name" {
  type    = string
  default = "Spark_TerraformExample"
}

variable "freeform_tags" {
  type    = map(string)
  default = { "bar-key" = "value" }
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_connection" "spark" {
  compartment_id  = var.compartment_id
  connection_type = "SPARK"
  technology_type = "APACHE_SPARK"
  display_name    = var.display_name
  description     = "Spark connection created by Terraform"
  freeform_tags   = var.freeform_tags
  routing_method  = "SHARED_SERVICE_ENDPOINT"

  submit_mode         = "SUBMIT"
  cluster_endpoint    = var.cluster_endpoint
  authentication_type = "NONE"
  security_protocol   = "PLAINTEXT"

  artifact_location {
    location_type             = "OCI_OBJECT_STORAGE"
    should_overwrite_existing = true

    configuration {
      authentication_type = "RESOURCE_PRINCIPAL"
      region              = var.region
      namespace           = var.objectstorage_namespace
      bucket              = var.objectstorage_bucket_name
      path                = var.artifact_path
    }
  }
}

data "oci_golden_gate_connection" "spark" {
  connection_id = oci_golden_gate_connection.spark.id
}

data "oci_golden_gate_connections" "spark" {
  compartment_id  = var.compartment_id
  connection_type = ["SPARK"]
  technology_type = ["APACHE_SPARK"]

  filter {
    name   = "id"
    values = [oci_golden_gate_connection.spark.id]
  }
}

output "connection_id" {
  value = oci_golden_gate_connection.spark.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.spark.display_name
}

output "connection_count" {
  value = length(data.oci_golden_gate_connections.spark.connection_collection[0].items)
}
