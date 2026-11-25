# Resolver endpoint and rule dependencies

This example creates two VCNs, each with a subnet, a listening endpoint and a
forwarding endpoint. Root-module resolver rules demonstrate forwarding in both
directions using distinct example suffixes.

Use OCI provider v8.16.0 or later. That release added the VCN resolver association
availability wait, allowing endpoints to consume the association's resolver ID
after asynchronous resolver provisioning completes.

`network.tf` calls the two network modules, `modules/network/main.tf` creates each
network and its endpoints, and `resolver_rules.tf` connects them with rules.
`provider.tf` configures the provider and authentication.

The dependency order is:

1. VCN and subnet, then VCN resolver association.
2. Listening and forwarding endpoints, using the association's resolver ID.
3. Resolver configuration and rules, using endpoint attributes.

An endpoint must not depend on an `oci_dns_resolver` configuration whose rules
reference that endpoint: those references would form a cycle. Read the resolver
ID directly from the VCN association to avoid that dependency.

This example composes rules in the root module using endpoint outputs from both
network modules. Endpoint references establish the dependencies needed for safe
creation and teardown.

Terraform reverses these dependencies on destroy: resolver rules are cleared
before endpoints are deleted, then subnets and VCNs are removed. Destroying an
`oci_dns_resolver` configuration clears its managed settings; the underlying
resolver belongs to the VCN and is removed with the VCN.

Set `TF_VAR_compartment_ocid` and `TF_VAR_region` for a test compartment.

Authentication defaults to `APIKey` with the `DEFAULT` OCI configuration
profile. Set `TF_VAR_auth` and `TF_VAR_config_file_profile` to use a different
authentication method or profile. See the [provider authentication guide](https://registry.terraform.io/providers/oracle/oci/latest/docs#authentication)
for supported options.

Then run:

```sh
terraform init
terraform apply -var='name_prefix=resolver-demo'
terraform plan -var='name_prefix=resolver-demo'
terraform destroy -var='name_prefix=resolver-demo'
```

This example demonstrates configuration and lifecycle ordering. It does not
cover connectivity between the VCNs.

See the [resolver resource](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/dns_resolver),
[endpoint resource](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/dns_resolver_endpoint),
and [VCN association data source](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_vcn_dns_resolver_association).
