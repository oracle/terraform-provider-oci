resource "oci_resourcemanager_stack" "example" {
  compartment_id = var.compartment_ocid
  display_name   = var.stack_display_name
  description    = var.stack_description

  config_source {
    config_source_type     = "ZIP_UPLOAD"
    zip_file_base64encoded = filebase64(var.stack_zip_path)
    working_directory      = var.stack_working_directory
  }

  freeform_tags = {
    Example = "terraform-provider-oci"
  }

  variables = {
    compartment_ocid = var.compartment_ocid
  }
}

data "oci_resourcemanager_stack" "example" {
  stack_id = oci_resourcemanager_stack.example.id
}

output "stack_id" {
  value = oci_resourcemanager_stack.example.id
}

output "downloaded_config_source_type" {
  value = data.oci_resourcemanager_stack.example.config_source[0].config_source_type
}

output "downloaded_working_directory" {
  value = data.oci_resourcemanager_stack.example.config_source[0].working_directory
}
