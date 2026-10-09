---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_artifact_signature"
sidebar_current: "docs-oci-resource-datascience-model_artifact_signature"
description: |-
  Provides the Model Artifact Signature resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model_artifact_signature
This resource provides the Model Artifact Signature resource in Oracle Cloud Infrastructure Data Science service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-science/latest/ArtifactSignature

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datascience

Create model artifact signature for the specified model

## Example Usage

```hcl
resource "oci_datascience_model_artifact_signature" "test_model_artifact_signature" {
	#Required
	compartment_id = var.compartment_id
	kms_key_id = oci_kms_key.test_key.id
	kms_key_version_id = oci_kms_key_version.test_key_version.id
	model_id = oci_datascience_model.test_model.id
	signing_algorithm = var.model_artifact_signature_signing_algorithm

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.model_artifact_signature_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My Model Artifact Signature`
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `kms_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key from Oracle Cloud Infrastructure KMS.
* `kms_key_version_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key version from Oracle Cloud Infrastructure KMS.
* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `signing_algorithm` - (Required) The algorithm used for signing.
* `verify_trigger` - (Optional) (Updatable) An optional property when incremented triggers Verify. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model Artifact Signature
	* `update` - (Defaults to 20 minutes), when updating the Model Artifact Signature
	* `delete` - (Defaults to 20 minutes), when destroying the Model Artifact Signature


## Import

ModelArtifactSignatures can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_artifact_signature.test_model_artifact_signature "models/{modelId}/artifactSignatures/{artifactSignatureId}"
```
