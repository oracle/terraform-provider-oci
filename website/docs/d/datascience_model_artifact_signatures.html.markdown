---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_artifact_signatures"
sidebar_current: "docs-oci-datasource-datascience-model_artifact_signatures"
description: |-
  Provides the list of Model Artifact Signatures in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_artifact_signatures
This data source provides the list of Model Artifact Signatures in Oracle Cloud Infrastructure Data Science service.

List model artifact signatures of a model in the specified compartment

## Example Usage

```hcl
data "oci_datascience_model_artifact_signatures" "test_model_artifact_signatures" {
	#Required
	compartment_id = var.compartment_id
	model_id = oci_datascience_model.test_model.id

	#Optional
	digest = var.model_artifact_signature_digest
	display_name = var.model_artifact_signature_display_name
	state = var.model_artifact_signature_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `digest` - (Optional) The digest of the model artifact.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type.


## Attributes Reference

The following attributes are exported:

* `model_artifact_signatures` - The list of model_artifact_signatures.

### ModelArtifactSignature Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resource.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `digest` - Base64 encoded string representing the model artifact digest.
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model artifact signature.
* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key from Oracle Cloud Infrastructure KMS.
* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key version from Oracle Cloud Infrastructure KMS.
* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model whose artifact has to be signed.
* `signature` - The signature of the digest generated using kmsKeyId, kmsKeyVersion Id and signingAlgorithm.
* `signing_algorithm` - The algorithm used for signing.
* `state` - The lifecycle state of the model artifact signature
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z
