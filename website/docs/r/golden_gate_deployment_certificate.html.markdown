---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_certificate"
sidebar_current: "docs-oci-resource-golden_gate-deployment_certificate"
description: |-
  Provides the Deployment Certificate resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_deployment_certificate
This resource provides the Deployment Certificate resource in Oracle Cloud Infrastructure Golden Gate service.

Creates a new certificate to truststore.


## Example Usage

```hcl
resource "oci_golden_gate_deployment_certificate" "test_deployment_certificate" {
	#Required
	certificate_content = var.deployment_certificate_certificate_content
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	key = var.deployment_certificate_key
}
```

## Argument Reference

The following arguments are supported:

* `certificate_content` - (Required) A PEM-encoded SSL certificate. 
* `deployment_id` - (Required) A unique Deployment identifier. 
* `key` - (Required) The identifier key (unique name in the scope of the deployment) of the certificate being referenced.  It must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `authority_key_id` - The Certificate authority key id. 
* `certificate_content` - A PEM-encoded SSL certificate. 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `is_ca` - Indicates if the certificate is ca. 
* `is_self_signed` - Indicates if the certificate is self signed. 
* `issuer` - The Certificate issuer. 
* `key` - The identifier key (unique name in the scope of the deployment) of the certificate being referenced.  It must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter. 
* `md5hash` - The Certificate md5Hash. 
* `public_key` - The Certificate public key. 
* `public_key_algorithm` - The Certificate public key algorithm. 
* `public_key_size` - The Certificate public key size. 
* `serial` - The Certificate serial. 
* `sha1hash` - The Certificate sha1 hash. 
* `state` - Possible certificate lifecycle states. 
* `subject` - The Certificate subject. 
* `subject_key_id` - The Certificate subject key id. 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_valid_from` - The time the certificate is valid from. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_valid_to` - The time the certificate is valid to. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `version` - The Certificate version. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Deployment Certificate
	* `update` - (Defaults to 20 minutes), when updating the Deployment Certificate
	* `delete` - (Defaults to 20 minutes), when destroying the Deployment Certificate


## Import

DeploymentCertificates can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_deployment_certificate.test_deployment_certificate "deployments/{deploymentId}/certificates/{certificateKey}" 
```

