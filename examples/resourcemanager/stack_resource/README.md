# Resource Manager Stack Resource Example

This example shows how to create and read back an OCI Resource Manager stack with
`oci_resourcemanager_stack`.

## Notes

- The current resource implementation supports `ZIP_UPLOAD` config sources only.
- `stack_zip_path` must point to a ZIP archive that already exists on disk.
- The stack and singular data source both persist the downloaded ZIP content in Terraform state so refresh and import can preserve stack configuration. For large archives, that can noticeably increase state size.
