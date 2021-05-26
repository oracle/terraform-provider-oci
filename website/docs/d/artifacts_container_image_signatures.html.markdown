---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_image_signatures"
sidebar_current: "docs-oci-datasource-artifacts-container_image_signatures"
description: |-
  Provides the list of Container Image Signatures in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_image_signatures
This data source provides the list of Container Image Signatures in Oracle Cloud Infrastructure Artifacts service.

List container image signatures in an image.

## Example Usage

```hcl
data "oci_artifacts_container_image_signatures" "test_container_image_signatures" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.container_image_signature_compartment_id_in_subtree
	display_name = var.container_image_signature_display_name
	image_digest = var.container_image_signature_image_digest
	image_id = oci_core_image.test_image.id
	kms_key_id = oci_kms_key.test_key.id
	kms_key_version_id = oci_kms_key_version.test_key_version.id
	repository_id = oci_artifacts_repository.test_repository.id
	repository_name = oci_artifacts_repository.test_repository.name
	signing_algorithm = var.container_image_signature_signing_algorithm
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are inspected depending on the the setting of `accessLevel`. Default is false. Can only be set to true when calling the API on the tenancy (root compartment). 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `image_digest` - (Optional) The digest of the container image.  Example: `sha256:e7d38b3517548a1c71e41bffe9c8ae6d6d29546ce46bf62159837aad072c90aa` 
* `image_id` - (Optional) A filter to return a container image summary only for the specified container image OCID. 
* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.  Example: `ocid1.keyversion.oc1..exampleuniqueID` 
* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.  Example: `ocid1.keyversion.oc1..exampleuniqueID` 
* `repository_id` - (Optional) A filter to return container images only for the specified container repository OCID. 
* `repository_name` - (Optional) A filter to return container images or container image signatures that match the repository name.  Example: `foo` or `foo*` 
* `signing_algorithm` - (Optional) The algorithm to be used for signing. These are the only supported signing algorithms for container images.


## Attributes Reference

The following attributes are exported:

* `container_image_signature_collection` - The list of container_image_signature_collection.

### ContainerImageSignature Reference

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

