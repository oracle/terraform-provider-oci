# Batch Service Full Workflow Example

This example demonstrates how to create and configure Batch service resources including:
- Batch Context
- Batch Job Pool
- Batch Task Environment
- Batch Task Profile

## Prerequisites

- OCI CLI configured with appropriate credentials
- A container image URL for the Batch Task Environment

## Quick Start

This example creates a new Batch Context along with other Batch resources:

```bash
export TF_VAR_image_url=""
export TF_VAR_compartment_ocid=""

terraform init
terraform plan
terraform apply
```

**Note:** Batch Context creation takes 30-45 minutes. The example includes appropriate timeouts (120 minutes) to handle this.

## Variables

### Required Variables

- `compartment_ocid`: OCID of the compartment where resources will be created
- `image_url`: Container image URL for the Batch Task Environment

### Optional Variables

- `display_name_suffix`: Suffix for resource display names (default: "demo")
- `task_env_mount_target_fqdn`: NFS mount target FQDN (leave blank to skip NFS volumes)
- `task_env_mount_target_export_path`: NFS export path (leave blank to skip NFS volumes)

## Notes

- Batch Context creation takes 30-45 minutes. The example includes appropriate timeouts (120 minutes) to handle this.
- The example creates networking resources (VCN, Subnet, NSG) which are quick to create and destroy.
- All resources include appropriate timeouts for long-running operations.

