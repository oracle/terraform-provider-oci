# Example: Autonomous Distributed Database (ADB_D) + Private Endpoint

This example provisions:

1. `oci_distributed_database_distributed_database_private_endpoint`
2. `oci_distributed_database_distributed_autonomous_database` (deployment type `ADB_D`)

It uses the **variables.tf + terraform.tfvars** workflow.

## Prerequisites

- OCI Terraform provider configured (auth, tenancy, region) via `provider.tf` in your root, or environment variables.
- An existing **Cloud Autonomous VM Cluster** OCID (for `cloud_autonomous_vm_cluster_id`).
- Networking: an existing **subnet OCID** where the Distributed DB private endpoint will be created.

## Usage

```bash
cd examples/adb_d
cp terraform.tfvars.example terraform.tfvars
# edit terraform.tfvars with your OCIDs and settings
terraform init
terraform plan
terraform apply
```

## Notes

- `admin_password` is **input-only** for some API shapes. Keep it in `terraform.tfvars` and treat it as sensitive.
- The resource schema in this provider generation requires `distributed_autonomous_database_id` as an input. Provide it via tfvars.
