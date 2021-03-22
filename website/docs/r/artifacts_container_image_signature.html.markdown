---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_image_signature"
sidebar_current: "docs-oci-resource-artifacts-container_image_signature"
description: |-
  Provides the Container Image Signature resource in Oracle Cloud Infrastructure Artifacts service
---

# oci_artifacts_container_image_signature
This resource provides the Container Image Signature resource in Oracle Cloud Infrastructure Artifacts service.

Upload a signature to an image.

## Example Usage

```hcl
resource "oci_artifacts_container_image_signature" "test_container_image_signature" {
	#Required
	compartment_id = var.compartment_id
	image_id = oci_core_image.test_image.id
	kms_key_id = oci_kms_key.test_key.id
	kms_key_version_id = oci_kms_key_version.test_key_version.id
	message = var.container_image_signature_message
	signature = var.container_image_signature_signature
	signing_algorithm = var.container_image_signature_signing_algorithm
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the container repository exists.
* `image_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image.  Example: `ocid1.containerimage.oc1..exampleuniqueID` 
* `kms_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyId used to sign the container image.  Example: `ocid1.key.oc1..exampleuniqueID` 
* `kms_key_version_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.  Example: `ocid1.keyversion.oc1..exampleuniqueID` 
* `message` - (Required) The base64 encoded signature payload that was signed.
* `signature` - (Required) The signature of the message field using the kmsKeyId, the kmsKeyVersionId, and the signingAlgorithm.
* `signing_algorithm` - (Required) The algorithm to be used for signing. These are the only supported signing algorithms for container images.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the container repository exists.
* `created_by` - The id of the user or principal that created the resource.
* `display_name` - The last 10 characters of the kmsKeyId, the last 10 characters of the kmsKeyVersionId, the signingAlgorithm, and the last 10 characters of the signatureId.  Example: `wrmz22sixa::qdwyc2ptun::SHA_256_RSA_PKCS_PSS::2vwmobasva` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image signature.  Example: `ocid1.containerimagesignature.oc1..exampleuniqueID` 
* `image_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image.  Example: `ocid1.containerimage.oc1..exampleuniqueID` 
* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyId used to sign the container image.  Example: `ocid1.key.oc1..exampleuniqueID` 
* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.  Example: `ocid1.keyversion.oc1..exampleuniqueID` 
* `message` - The base64 encoded signature payload that was signed.
* `signature` - The signature of the message field using the kmsKeyId, the kmsKeyVersionId, and the signingAlgorithm.
* `signing_algorithm` - The algorithm to be used for signing. These are the only supported signing algorithms for container images.
* `time_created` - An RFC 3339 timestamp indicating when the image was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Container Image Signature
	* `update` - (Defaults to 20 minutes), when updating the Container Image Signature
	* `delete` - (Defaults to 20 minutes), when destroying the Container Image Signature


## Import

ContainerImageSignatures can be imported using the `id`, e.g.

```
$ terraform import oci_artifacts_container_image_signature.test_container_image_signature "container/imageSignatures/{imageSignatureId}" 
```

