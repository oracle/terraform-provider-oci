# Public Safety and Security

The [instance example](instance/main.tf) creates a PSS instance and reads it through the singular and plural data sources. It uses an existing identity domain and an administrator user belonging to that domain.

Use a service-owned test compartment with PSS enabled and a current OCI session in the `terraform-federation-test` profile. The caller needs PSS permissions and the identity-domain permissions required for PSS provisioning.

Set `TF_VAR_region`, `TF_VAR_compartment_ocid`, `TF_VAR_pss_identity_domain_url`, and `TF_VAR_pss_admin_identity_domain_user_id`. Override `TF_VAR_config_file_profile` when using a different session profile.

Alternatively, copy `instance/terraform.tfvars.example` to `instance/terraform.tfvars` and fill in its placeholders before the later live run. Keep real environment values in your local copy.

For preview testing, build the provider from the TERSI feature branch and configure Terraform to use that local binary through the workflow tool or a development override. An installed public provider may not contain this service yet.

For the workflow's **Example Test** step (option 7 in the current tool), enter `pss/instance` when asked for the directory relative to `examples/`. The tool supplies `required_providers_block.tf`, runs init/apply/destroy, and uploads successful example evidence to the TERSI. Do not add a second `required_providers` block to `main.tf`. The tool authenticates at the start but does not continuously refresh the session during the example; keep the same profile refreshed from a separate terminal throughout apply and destroy.

For standalone use, first copy `required_providers_block.tf.example` to `required_providers_block.tf` inside `instance/`. Then run `terraform init`, `terraform plan`, and `terraform apply` from that directory, using the locally built provider for preview testing. Change `display_name` or tags and apply again to exercise updates. Run `terraform destroy` when finished; the existing identity domain and administrator are inputs and are not managed by this example. The workflow removes `required_providers_block.tf` after its run, so recreate it from the template if switching back to standalone commands afterward.

The acceptance test `TestPssInstanceResource_basic` uses the same two PSS environment variables in addition to the standard provider acceptance-test configuration. It covers create, update, compartment move, data-source reads, import, and deletion. Set distinct service-owned `TF_VAR_compartment_ocid` and `TF_VAR_compartment_id_for_update` values to exercise an actual compartment move.

Both the acceptance configurations and the standalone example use `var.compartment_ocid` for the create/list compartment, matching `TF_VAR_compartment_ocid` and the test assertions. An inherited `TF_VAR_compartment_id` does not select the PSS test compartment. The move step continues to use `TF_VAR_compartment_id_for_update`.

For acceptance-test defined tags, supply `TF_VAR_home_region_defined_tag_namespace_name` and `TF_VAR_home_region_defined_tag_name` with existing names provisioned in the tenancy home region, following the DDFS test setup. The test applies and updates tag values on the PSS instance; it does not create or delete tag namespaces or tag definitions. The caller needs permission to use the supplied tags. The standalone example uses freeform tags and does not require these two additional inputs.

The test and example provider configurations ignore only the automatically added `Oracle-Tags.CreatedBy` and `Oracle-Tags.CreatedOn` keys. Custom defined-tag changes remain managed; the acceptance test checks both the initial and updated custom-tag values.

Following DDFS, the acceptance test creates one instance with required fields, applies optional fields to that same instance, moves it, updates it and moves it back, checks both data sources, imports it, and finally destroys it. There are seven test steps plus final cleanup. Each update checks that the instance ID is unchanged. The list data source has its own synthetic ID; the returned collection item's ID is checked against the instance.
