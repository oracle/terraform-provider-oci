resource "oci_lustre_file_storage_lustre_file_system" "test_lustre_file_system_import_job" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  capacity_in_gbs     = var.lustre_file_system_capacity_in_gbs
  compartment_id      = var.compartment_ocid
  file_system_name    = var.lustre_file_system_name
  performance_tier    = var.lustre_file_system_performance_tier
  root_squash_configuration {

    #Optional
    client_exceptions = var.lustre_file_system_root_squash_configuration_client_exceptions
    identity_squash   = var.lustre_file_system_root_squash_configuration_identity_squash
    squash_gid        = var.lustre_file_system_root_squash_configuration_squash_gid
    squash_uid        = var.lustre_file_system_root_squash_configuration_squash_uid
  }
  subnet_id = oci_core_subnet.my_subnet.id

  #Optional
  cluster_placement_group_id = oci_cluster_placement_groups_cluster_placement_group.test_cpg.id
#   defined_tags               = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.lustre_file_system_defined_tags_value)
  display_name               = var.lustre_file_system_display_name_import_job
  file_system_description    = var.lustre_file_system_import_job_file_system_description

  freeform_tags = {
    "Department" = "Finance"
  }
}