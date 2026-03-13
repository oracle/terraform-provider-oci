# Example: Distributed Database Private Endpoint

This example provisions:

1. `oci_distributed_database_distributed_database_private_endpoint`

It also demonstrates:

1. `data.oci_distributed_database_distributed_database_private_endpoint`
2. `data.oci_distributed_database_distributed_database_private_endpoints`

It uses the **variables.tf + terraform.tfvars** workflow.

## Usage

1. Copy the sample vars file:

```bash
cp terraform.tfvars-template terraform.tfvars
```

2. Edit `terraform.tfvars` and fill in OCIDs and values.

3. Run:

```bash
terraform init
terraform plan
terraform apply
```

## Notes

- `subnet_id` must be in the same region/realm where you run this example.
- `reinstate_proxy_instance_trigger` is an action trigger. Increase it only when you need to run reinstate.
