output "bootVolumeFromInstance" {
  value = ["${oci_core_instance.TFInstance.boot_volume_id}"]
}

output "bootVolumeFromSourceBootVolume" {
  value = ["${oci_core_boot_volume.TFBootVolumeFromSourceBootVolume.id}"]
}

output "bootVolumeBackupFromSourceBootVolume" {
  value = ["${oci_core_boot_volume_backup.TFBootVolumeBackupFromSourceBootVolume.id}"]
}

output "bootVolumeFromSourceBootVolumeBackup" {
  value = ["${oci_core_boot_volume.TFBootVolumeFromSourceBootVolumeBackup.id}"]
}

output "bootVolumeFromSourceBootVolumeDatasource" {
  value = ["${data.oci_core_boot_volumes.TFBootVolumeFromSourceBootVolumeDatasource.boot_volumes}"]
}

output "bootVolumeBackupFromSourceBootVolumeDatasource" {
  value = ["${data.oci_core_boot_volume_backups.TFBootVolumeBackupFromSourceBootVolumeDatasource.boot_volume_backups}"]
}

output "bootVolumeFromSourceBootVolumeBackupDatasource" {
  value = ["${data.oci_core_boot_volumes.TFBootVolumeFromSourceBootVolumeBackupDatasource.boot_volumes}"]
}
