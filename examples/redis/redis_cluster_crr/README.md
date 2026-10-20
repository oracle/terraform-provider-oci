# Redis cross-region replication secondary cluster

This example creates a Redis cluster as a cross-region replication (CRR) secondary. It creates the secondary-region VCN, security list, and subnet, then creates the secondary Redis cluster.

It deliberately does **not** create, modify, or delete the primary cluster. The primary is an external prerequisite because it belongs in the paired primary region.

## Prerequisites

Before applying this configuration:

1. Create an active primary Redis cluster in the paired primary region.
2. Create or identify an OCI Cache config set in the secondary region.
3. Verify the primary and secondary config sets meet CRR compatibility requirements:
   - `reserved-memory-percentage` has the same value on both config sets and is at least `35`.
   - `maxmemory-policy` has the same value on both config sets.
   - `databases` has the same value on both config sets when it differs from the default.
4. Copy `terraform.tfvars.example` to `terraform.tfvars` and provide the compartment OCID, primary cluster OCID, and secondary-region config-set OCID.

## Run

```sh
terraform init
terraform plan
terraform apply
```

The default values create a non-sharded `REDIS_7_0` secondary with one 8 GB node. Override these values in `terraform.tfvars` when they must match the primary cluster's topology or software version.

To remove the resources created by this example, run `terraform destroy`. This does not affect the external primary cluster or its config set.

## Lifecycle behavior

- Adding `primary_cluster_id` to an existing standalone cluster converts it to a secondary cluster.
- Removing `primary_cluster_id` from a secondary cluster converts it to a standalone cluster.
- Changing directly from one populated `primary_cluster_id` to another is not supported. Remove it and apply first, then set the new primary and apply again.
- Switchover is intentionally not supported through this Terraform resource.
