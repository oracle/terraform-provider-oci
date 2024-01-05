// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_instance_configuration" "instance_autoscaling_configuration_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "instanceAutoscalingConfigurationRD"

  instance_details {
    instance_type = "compute"

    launch_details {
      compartment_id = "${var.compartment_ocid}"
      ipxe_script    = "ipxeScript"
      shape          = "${var.instance_shape}"
      display_name   = "instanceAutoscalingConfigurationLaunchDetailsRD"

      create_vnic_details {
        assign_public_ip       = true
        display_name           = "instanceAutoscalingConfigurationVNICRD"
        skip_source_dest_check = false
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = "${var.instance_image_ocid[var.region]}"
      }
    }
  }
}

resource "oci_core_instance_pool" "oci_core_instance_pool_rd" {
  compartment_id            = "${var.compartment_ocid}"
  instance_configuration_id = "${oci_core_instance_configuration.instance_autoscaling_configuration_rd.id}"
  size                      = 2
  state                     = "RUNNING"
  display_name              = "instancePoolRD"

  placement_configurations {
    availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
    primary_subnet_id   = "${oci_core_subnet.subnet_rd.id}"
  }
}

resource "oci_autoscaling_auto_scaling_configuration" "autoscaling_auto_scaling_configuration_rd" {
  compartment_id       = "${var.compartment_ocid}"
  cool_down_in_seconds = "300"
  display_name         = "autoScalingConfigurationRD"
  is_enabled           = "true"

  policies {
    capacity {
      initial = "4"
      max     = "4"
      min     = "2"
    }

    display_name = "autoScalingPolicyRD"
    policy_type  = "threshold"

    execution_schedule {
      expression = "0 15 10 ? * *"
      timezone   = "UTC"
      type       = "cron"
    }

    rules {
      action {
        type  = "CHANGE_COUNT_BY"
        value = "1"
      }

      display_name = "autoScaleOutRuleRD"

      metric {
        metric_type = "CPU_UTILIZATION"

        threshold {
          operator = "GT"
          value    = "1"
        }
      }
    }

    rules {
      action {
        type  = "CHANGE_COUNT_BY"
        value = "-1"
      }

      display_name = "autoscalingScaleInRuleED"

      metric {
        metric_type = "CPU_UTILIZATION"

        threshold {
          operator = "LT"
          value    = "1"
        }
      }
    }
  }

  auto_scaling_resources {
    id   = "${oci_core_instance_pool.oci_core_instance_pool_rd.id}"
    type = "instancePool"
  }
}
