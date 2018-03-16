resource "oci_core_volume" "db_block" {
  count               = "${var.ad_count}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${count.index}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "SQL${"${count.index}" + 1}-DB.blk"
  size_in_gbs         = "${var.sql_db_size}"
}

resource "oci_core_volume" "db_log" {
  count               = "${var.ad_count}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${count.index}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "SQL${"${count.index}" + 1}-LOG.blk"
  size_in_gbs         = "${var.sql_log_size}"
}

resource "oci_core_volume" "db_backup" {
  count               = "${var.ad_count}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${count.index}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "SQL${"${count.index}" + 1}-BACKUP.blk"
  size_in_gbs         = "${var.sql_backup_size}"
}

resource "oci_core_volume" "witness_volume" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${var.ad_deployment}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "SQL${"${var.ad_deployment}" + 1}-WITNESS.blk"
  size_in_gbs         = "${var.witness_volume_size}"
}
