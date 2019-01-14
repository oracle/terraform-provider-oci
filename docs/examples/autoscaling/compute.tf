resource "oci_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstanceForInstancePool"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.ExampleSubnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleinstance${count.index}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_instance_configuration" "TFInstanceConfiguration" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    launch_details {
      compartment_id = "${var.compartment_ocid}"
      ipxe_script    = "ipxeScript"
      shape          = "${var.instance_shape}"
      display_name   = "TFExampleInstanceConfigurationLaunchDetails"

      create_vnic_details {
        assign_public_ip       = true
        display_name           = "TFExampleInstanceConfigurationVNIC"
        skip_source_dest_check = false
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details = {
        source_type = "image"
        image_id    = "${var.instance_image_ocid[var.region]}"
      }
    }
  }
}

resource "oci_core_instance_pool" "TFInstancePool" {
  compartment_id            = "${var.compartment_ocid}"
  instance_configuration_id = "${oci_core_instance_configuration.TFInstanceConfiguration.id}"
  size                      = 2
  state                     = "RUNNING"
  display_name              = "TFInstancePool"

  placement_configurations {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
    primary_subnet_id   = "${oci_core_subnet.ExampleSubnet.id}"
  }
}

resource "oci_auto_scaling_auto_scaling_configuration" "TFAutoScalingConfiguration" {
  compartment_id       = "${var.compartment_ocid}"
  cool_down_in_seconds = "300"
  display_name         = "TFAutoScalingConfiguration"
  is_enabled           = "true"

  policies {
    capacity {
      initial = "4"
      max     = "4"
      min     = "2"
    }

    display_name = "TFPolicy"
    policy_type  = "threshold"

    rules {
      action {
        type  = "CHANGE_COUNT_BY"
        value = "1"
      }

      display_name = "TFScaleOutRule"

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

      display_name = "TFScaleInRule"

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
    id   = "${oci_core_instance_pool.TFInstancePool.id}"
    type = "instancePool"
  }
}
