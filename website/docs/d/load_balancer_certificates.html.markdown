---
layout: "oci"
page_title: "OCI: oci_load_balancer_certificates"
sidebar_current: "docs-oci-datasource-load_balancer-certificates"
description: |-
  Provides a list of Certificates
---

# Data Source: oci_load_balancer_certificates
The Certificates data source allows access to the list of OCI certificates

Lists all SSL certificates bundles associated with a given load balancer.

## Example Usage

```hcl
data "oci_load_balancer_certificates" "test_certificates" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the certificate bundles to be listed. 


## Attributes Reference

The following attributes are exported:

* `certificates` - The list of certificates.

### Certificate Reference

The following attributes are exported:

* `ca_certificate` - The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.  Example:      -----BEGIN CERTIFICATE-----     MIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD..AkGA1UEBhMCR0Ix     EzARBgNVBAgTClNvbWUtU3RhdGUxFDASBgNVBAoTC0..0EgTHRkMTcwNQYD     VQQLEy5DbGFzcyAxIFB1YmxpYyBQcmltYXJ5IENlcn..XRpb24gQXV0aG9y     aXR5MRQwEgYDVQQDEwtCZXN0IENBIEx0ZDAeFw0wMD..TUwMTZaFw0wMTAy     ...     -----END CERTIFICATE----- 
* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
* `load_balancer_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate.
* `passphrase` - A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.  Example: `Mysecretunlockingcode42!1!`
* `private_key` - The SSL private key for your certificate, in PEM format.  Example:      -----BEGIN RSA PRIVATE KEY-----     jO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca..JZNUgYYwNS0dP2UK     tmyN+XqVcAKw4HqVmChXy5b5msu8eIq3uc2NqNVtR..2ksSLukP8pxXcHyb     +sEwvM4uf8qbnHAqwnOnP9+KV9vds6BaH1eRA4CHz..n+NVZlzBsTxTlS16     /Umr7wJzVrMqK5sDiSu4WuaaBdqMGfL5hLsTjcBFD..Da2iyQmSKuVD4lIZ     ...     -----END RSA PRIVATE KEY-----
* `public_certificate` - The public certificate, in PEM format, that you received from your SSL certificate provider.  Example:      -----BEGIN CERTIFICATE-----     MIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbMQswCQYDVQQGEwJKUDEOMAwG     A1UECBMFVG9reW8xEDAOBgNVBAcTB0NodW8ta3UxETAPBgNVBAoTCEZyYW5rNERE     MRgwFgYDVQQLEw9XZWJDZXJ0IFN1cHBvcnQxGDAWBgNVBAMTD0ZyYW5rNEREIFdl     YiBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBmcmFuazRkZC5jb20wHhcNMTIw     ...     -----END CERTIFICATE----- 

