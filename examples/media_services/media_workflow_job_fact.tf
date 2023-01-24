// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_workflow_job_fact_key" {
  default = "key"
}

variable "media_workflow_job_fact_type" {
  default = "runnableJob"
}

data "oci_media_services_media_workflow_job_facts" "test_media_workflow_job_facts" {
  #Required
  media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id

  #Optional
  key  = var.media_workflow_job_fact_key
  type = var.media_workflow_job_fact_type
}

