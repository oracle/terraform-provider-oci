resource "oci_core_boot_volume" "TFBootVolumeFromSourceBootVolume" {
  availability_domain = "${oci_core_instance.TFInstance.availability_domain}"
  compartment_id      = "${oci_core_instance.TFInstance.compartment_id}"

  source_details {
    #Required
    id   = "${oci_core_instance.TFInstance.boot_volume_id}"
    type = "bootVolume"
  }
}

resource "oci_core_boot_volume_backup" "TFBootVolumeBackupFromSourceBootVolume" {
  #Required
  boot_volume_id = "${oci_core_boot_volume.TFBootVolumeFromSourceBootVolume.id}"
}

resource "oci_core_boot_volume" "TFBootVolumeFromSourceBootVolumeBackup" {
  availability_domain = "${oci_core_instance.TFInstance.availability_domain}"
  compartment_id      = "${oci_core_instance.TFInstance.compartment_id}"

  source_details {
    #Required
    id   = "${oci_core_boot_volume_backup.TFBootVolumeBackupFromSourceBootVolume.id}"
    type = "bootVolumeBackup"
  }
}

data "oci_core_boot_volume_backups" "TFBootVolumeBackupFromSourceBootVolumeDatasource" {
  compartment_id = "${oci_core_instance.TFInstance.compartment_id}"

  filter {
    name   = "id"
    values = ["${oci_core_boot_volume_backup.TFBootVolumeBackupFromSourceBootVolume.id}"]
  }
}

data "oci_core_boot_volumes" "TFBootVolumeFromSourceBootVolumeDatasource" {
  #Required
  availability_domain = "${oci_core_boot_volume.TFBootVolumeFromSourceBootVolume.availability_domain}"
  compartment_id      = "${oci_core_boot_volume.TFBootVolumeFromSourceBootVolume.compartment_id}"

  filter {
    name   = "id"
    values = ["${oci_core_boot_volume.TFBootVolumeFromSourceBootVolume.id}"]
  }
}

data "oci_core_boot_volumes" "TFBootVolumeFromSourceBootVolumeBackupDatasource" {
  #Required
  availability_domain = "${oci_core_boot_volume.TFBootVolumeFromSourceBootVolumeBackup.availability_domain}"
  compartment_id      = "${oci_core_boot_volume.TFBootVolumeFromSourceBootVolumeBackup.compartment_id}"

  filter {
    name   = "id"
    values = ["${oci_core_boot_volume.TFBootVolumeFromSourceBootVolumeBackup.id}"]
  }
}
