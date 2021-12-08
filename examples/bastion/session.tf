// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "session_display_name" {
  default = "bastionSessionExample"
}

variable "session_key_details_public_key_content" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC9HLTG/GtjHeFkZ9TCY6we0rK68oqAd2WkUj181N8vpt5ejYG7Oz6NpR4H8NSI5Vpra9DpGZMJRcMD56hj8vl3fNDcZp9j9D4ZmIuNBJ6rqOTOspnLoknT/c+5HILMzSgvFVFBvVQ4ftM030ldMsaGuCIyo8O6b3eB5SxSvoXZBvupaKhhrfQGilrNuVqtt7GJNB6aVXPfL/DOhX0DXOxwq/FVRMH1ovIXgoeg0rzzX5q3KYGVq2ZNWFYpVzFihf7ecJyvQ1SYOF6okrN+bKq3iQ+oKcveAfvpjma+LEI+TrZEJ7Xd0A/b1gh4HJsSa+pKCP+MYqjkFkEE3zqoDWYB bastions"
}

variable "session_key_type" {
  default = "PUB"
}

variable "session_session_lifecycle_state" {
  default = "ACTIVE"
}

variable "session_session_ttl_in_seconds" {
  default = 1800
}

variable "session_target_resource_details_session_type_managed_ssh" {
  default = "MANAGED_SSH"
}

variable "session_target_resource_details_session_type_port_forwarding" {
  default = "PORT_FORWARDING"
}

variable "session_target_resource_details_target_resource_port" {
  default = 22
}

resource "time_sleep" "wait_3_minutes_for_bastion_plugin" {
  depends_on = [oci_core_instance.test_instance]

  create_duration = "3m"
}

resource "oci_bastion_session" "test_session_managed_ssh" {
  #Required
  bastion_id = oci_bastion_bastion.test_bastion.id
  key_details {
    #Required
    public_key_content = var.session_key_details_public_key_content
  }
  target_resource_details {
    #Required
    session_type       = var.session_target_resource_details_session_type_managed_ssh
    // target_resource_id is required for managed ssh session
    target_resource_id = oci_core_instance.test_instance.id

    #Optional
    target_resource_operating_system_user_name = "opc"
    target_resource_port                       = var.session_target_resource_details_target_resource_port
    target_resource_private_ip_address         = oci_core_instance.test_instance.private_ip
  }

  display_name           = var.session_display_name
  key_type               = var.session_key_type
  session_ttl_in_seconds = var.session_session_ttl_in_seconds

  depends_on = [time_sleep.wait_3_minutes_for_bastion_plugin]
}

resource "oci_bastion_session" "test_session_port_forwarding" {
  #Required
  bastion_id = oci_bastion_bastion.test_bastion.id
  key_details {
    #Required
    public_key_content = var.session_key_details_public_key_content
  }
  target_resource_details {
    #Required
    session_type       = var.session_target_resource_details_session_type_port_forwarding

    #Optional
    // You should either use target_resource_id or target_resource_private_ip_address in port forwarding session
    target_resource_id                         = oci_core_instance.test_instance.id
    target_resource_private_ip_address         = oci_core_instance.test_instance.private_ip
    target_resource_port                       = var.session_target_resource_details_target_resource_port
  }

  display_name           = var.session_display_name
  key_type               = var.session_key_type
  session_ttl_in_seconds = var.session_session_ttl_in_seconds
}

data "oci_bastion_sessions" "test_sessions_managed_ssh" {
  #Required
  bastion_id = oci_bastion_bastion.test_bastion.id

  #Optional
  display_name            = var.session_display_name
  session_id              = oci_bastion_session.test_session_managed_ssh.id
  session_lifecycle_state = var.session_session_lifecycle_state
}

data "oci_bastion_sessions" "test_sessions_port_forwarding" {
  #Required
  bastion_id = oci_bastion_bastion.test_bastion.id

  #Optional
  display_name            = var.session_display_name
  session_id              = oci_bastion_session.test_session_port_forwarding.id
  session_lifecycle_state = var.session_session_lifecycle_state
}