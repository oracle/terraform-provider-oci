---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_image_signature"
sidebar_current: "docs-oci-datasource-artifacts-container_image_signature"
description: |-
  Provides details about a specific Container Image Signature in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_image_signature
This data source provides details about a specific Container Image Signature resource in Oracle Cloud Infrastructure Artifacts service.

Get container image signature metadata.

## Example Usage

```hcl
data "oci_artifacts_container_image_signature" "test_container_image_signature" {
	#Required
	image_signature_id = oci_artifacts_image_signature.test_image_signature.id
}
```

## Argument Reference

The following arguments are supported:

* `image_signature_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image signature.  Example: `ocid1.containersignature.oc1..exampleuniqueID` 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the container repository exists.
* `created_by` - The id of the user or principal that created the resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The last 10 characters of the kmsKeyId, the last 10 characters of the kmsKeyVersionId, the signingAlgorithm, and the last 10 characters of the signatureId.  Example: `wrmz22sixa::qdwyc2ptun::SHA_256_RSA_PKCS_PSS::2vwmobasva` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image signature.  Example: `ocid1.containerimagesignature.oc1..exampleuniqueID` 
* `image_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image.  Example: `ocid1.containerimage.oc1..exampleuniqueID` 
* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyId used to sign the container image.  Example: `ocid1.key.oc1..exampleuniqueID` 
* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.  Example: `ocid1.keyversion.oc1..exampleuniqueID` 
* `message` - The base64 encoded signature payload that was signed.
* `signature` - The signature of the message field using the kmsKeyId, the kmsKeyVersionId, and the signingAlgorithm.
* `signing_algorithm` - The algorithm to be used for signing. These are the only supported signing algorithms for container images.
* `state` - The current state of the container image signature.
* `system_tags` - The system tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - An RFC 3339 timestamp indicating when the image was created.

