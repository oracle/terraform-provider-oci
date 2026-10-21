# OpenSearch remote-reindex trust with an OCI CA bundle

This example creates an OCI CA bundle and an OpenSearch cluster configured to trust one remote OpenSearch or Elasticsearch HTTPS endpoint. It creates a VCN and subnet for the cluster.

The example configures remote-reindex trust only. It does not create the remote source, issue an OpenSearch `_reindex` request, or demonstrate master-user authentication through OCI Vault.

## Migrating from `certificate_id`

`remote_sources.certificate_id` is a deprecated alias of `ca_bundle_id`. Each source must configure exactly one of these names with a non-empty OCI Certificates service **CA bundle OCID**; certificate OCIDs are not supported by this alias. Rename the argument in the source block to `ca_bundle_id`, retaining the same value. This migration updates the existing cluster without replacing it. Resource refresh preserves the configured name, while data sources expose both computed aliases during the deprecation period. The old name remains supported until its removal in a major provider release.

## Prerequisites

- An OCI provider build or release supporting `remote_reindex_trust_config` and `ca_bundle_id`.
- An authenticated OCI CLI security-token profile for the target tenancy, and permission to create networking, CA bundle, and OpenSearch resources. The OpenSearch service must be permitted to use the CA bundle.
- A target compartment and region supporting the selected OpenSearch software version and node configuration. The coordinator shape is `VM.Standard.E3.Flex`; adjust it if your environment requires a different shape.
- An existing remote HTTPS endpoint and a PEM file containing its trusted CA certificates, not private keys.
- A master username and an OpenSearch-compatible password hash for the new cluster. Do not supply a plaintext password as the hash.

No tenancy, compartment, private endpoint override, or developer-specific filesystem path is hardcoded.

## Inputs

Set these variables in the terminal that launches Terraform or the workflow. Replace the placeholders with your values:

```sh
export TF_VAR_region="<oci-region>"
export TF_VAR_config_file_profile="<authenticated-oci-profile>"
export TF_VAR_compartment_id="<compartment-ocid>"
export TF_VAR_security_master_user_name="<master-username>"
export TF_VAR_security_master_user_password_hash='<password-hash>'
export TF_VAR_remote_reindex_ca_bundle_pem_path="/absolute/path/to/ca-bundle.pem"
export TF_VAR_remote_reindex_endpoint="https://source.example.com:9200"
```

Optional inputs:

- `TF_VAR_remote_reindex_endpoint_resolution_ip`: an IP used to resolve the remote endpoint hostname from the cluster. Leave it unset to omit the override. It does not disable hostname verification or establish network connectivity.
- `TF_VAR_display_name`: defaults to `opensearch-remote-reindex-example`.
- `TF_VAR_software_version`: defaults to `2.15.0`.

The CA certificates must validate the remote server certificate, including its chain; the certificate must match the endpoint hostname. Actual reindexing also requires a reachable source and any source credentials required by that server. This example does not provision all networking needed for that connectivity.

Keep credentials and Terraform state out of version control. Marking the password hash sensitive hides it in normal CLI output but does not encrypt Terraform state.

## Run with tfworkflow

Build the feature provider and configure the workflow's standard prerequisites, including `provider_bin_path`. The workflow refreshes the `terraform-federation-test` profile; use that authenticated profile for `TF_VAR_config_file_profile` when relying on workflow-managed refresh.

Select **7 — Example Test**, answer `y`, and enter:

```text
opensearch/remote_reindex
```

The workflow supplies `required_providers_block.tf`, initializes with the local provider, applies the example, and destroys the resources. Do not add another `required_providers` block to `main.tf`.

For the existing basic backward-compatibility example, select `opensearch/opensearch_cluster` instead. The baseline released provider must support all fields in whichever example is selected.

## Run directly with Terraform

From this directory, create a local `required_providers_block.tf` containing the following declaration. It is supplied automatically when using the workflow instead:

```hcl
terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}
```

With a released provider that supports the feature:

```sh
terraform init
terraform fmt -check
terraform validate
terraform plan
terraform apply
terraform plan -detailed-exitcode
terraform destroy
```

To use a local feature provider instead, replace `terraform init` with `terraform init -plugin-dir="$provider_bin_path"`, using the provider installation prepared by the workflow. A post-apply plan with exit code `0` confirms there are no changes.

## Cleanup

`terraform destroy` removes the cluster, CA bundle, subnet, and VCN. The CA bundle reference orders cluster deletion before bundle deletion. The local PEM file and existing remote source are not deleted.

If apply or cleanup fails, retain the state and correct the error before retrying `terraform destroy` with valid credentials. Do not discard state or rerun the workflow over resources left behind by a failed run.
