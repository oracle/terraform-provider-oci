// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {
  default = "ocidv1:tenancy:oc1:phx:1458753575596:aaaaaaaavary4yqe4ljpv5wzp74eflkwpu"
}

variable "user_ocid" {
  default = "ocid1.user.oc1..aaaaaaaaegg7wrwmpy2hbq7lcj2qjtnnfirbnrvcerzdzmlqoxxnyzy5okqq"
}

variable "fingerprint" {
  default = "29:20:66:1a:b7:59:6b:e2:bf:bb:65:94:04:75:2e:31"
}

variable "private_key_path" {
  default = "/Users/rajeevthakur/.ssh/private.pem"
}

variable "region" {
  default = "ap-hyderabad-1"
}

variable "compartment_ocid" {
  default = "ocidv1:tenancy:oc1:phx:1458753575596:aaaaaaaavary4yqe4ljpv5wzp74eflkwpu"
}

variable "ssh_public_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDwudUD3jENVXRrqUzTlqdOTlxl29b9wyw+gTxfZQt9ze3kIynu0cjLnkHS9/NPk1fGaoodu67aJ0TPb4ZazWCE/ib9ppHyn2yGFr+8bjmJdyIvhEdPZbvOHp5b9qlxQ6jX7KnlGcWD49uIaGypG9RYUVNg+OYhmClOn2aJvys9uo/TPkIBAQiXCjwZz94hlKvecxbRuYNhDT6KLj52ILjM019s3YngixLFeDFNvW7pSMC6cRKqJJzv4K+l+xfrB4bfZHqgVWw4Oh6XBIORkDgaf7sx4r82Q54YuNTEhTvYt/LXPSy4ywkigsQeqbV8PHDCCRGL7FYP92lzHnnpxYKz4bR57QUJJ9+Vdj36em+SZi4xiRwUY5LcIxCw8S7tTPVHlBUA9bvEeoJgzSxWIdi5zpwIGQkoE0lr0MVHca56ipRYOly2sxmTwNTNu6nk2JAA1lk7SvO0R10g/c+Y36kxY52oqg5ZOh9utanLwiVOhDPPHcgS6QgI+XAkAu0w+SU1qHcDGHzvVTXgwN0K9qgxEn81y4P8WXEysnc9u5BGtOSBSD6d/TRtxNgwuCmlYHKrxulDQk0tGOprCT24BkUlhmMZBfJh5vHEnI7t21/Y7tr8GCd8kO0e+3nxMxzHx1+dNT1mHQ/aEyfNH29iN1pOZE/u8wQWSllHT50FsE6ZEw== rajeev.thakur@oracle.com"
}

variable "ssh_private_key" {
  default = "-----BEGIN OPENSSH PRIVATE KEY-----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 -----END OPENSSH PRIVATE KEY-----"
}

# DBSystem specific
variable "db_system_shape" {
  default = "VM.Standard2.1"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION"
}

variable "db_admin_password" {
  default = "FIpassword12##"
}

variable "db_version" {
  default = "19.0.0.0"
}

variable "db_disk_redundancy" {
  default = "NORMAL"
}

variable "sparse_diskgroup" {
  default = true
}

variable "hostname" {
  default = "myoracledb"
}

variable "host_user_name" {
  default = "opc"
}

variable "n_character_set" {
  default = "AL16UTF16"
}

variable "character_set" {
  default = "AL32UTF8"
}

variable "db_workload" {
  default = "OLTP"
}

variable "pdb_name" {
  default = "pdbName"
}

variable "data_storage_size_in_gb" {
  default = "256"
}

variable "license_model" {
  default = "LICENSE_INCLUDED"
}

variable "node_count" {
  default = "1"
}