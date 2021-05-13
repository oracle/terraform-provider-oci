// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

variable "external_container_database_display_name" {
    default = "myTestExternalContainerDatabase"
}

variable "external_non_container_database_display_name" {
    default = "myTestExternalNonContainerDatabase"
}

variable "external_pluggable_database_display_name" {
    default = "myTestExternalPluggableDatabase"
}

variable "credential_type" {
    default = "NAME_REFERENCE"
}

variable "credential_name" {
    default = "credential.name"
}

variable "username" {
    default = "username"
}

variable "password" {
    default = "password"
}

variable "role" {
    default = "SYSDBA"
}

variable "hostname" {
    default = "host.name"
}

variable "port" {
    default = 10
}

variable "protocol" {
    default = "TCP"
}

variable "service" {
    default = "myService"
}

variable "connector_agent_id" {
    default = "ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy"
}

variable "external_database_connector_display_name" {
    default = "myTestConnector"
}

variable "connector_type" {
    default = "MACS"
}

variable "license_model" {
    default = "BRING_YOUR_OWN_LICENSE"
}

variable "external_database_state" {
    default = "AVAILABLE"
}