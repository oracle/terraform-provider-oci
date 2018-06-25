# oci_load_balancer_certificate

## Certificate Resource

### Certificate Reference

The following attributes are exported:

* `ca_certificate` - The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.  Example:      -----BEGIN CERTIFICATE-----     MIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD..AkGA1UEBhMCR0Ix     EzARBgNVBAgTClNvbWUtU3RhdGUxFDASBgNVBAoTC0..0EgTHRkMTcwNQYD     VQQLEy5DbGFzcyAxIFB1YmxpYyBQcmltYXJ5IENlcn..XRpb24gQXV0aG9y     aXR5MRQwEgYDVQQDEwtCZXN0IENBIEx0ZDAeFw0wMD..TUwMTZaFw0wMTAy     ...     -----END CERTIFICATE----- 
* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
* `load_balancer_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate.
* `passphrase` - A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.  Example: `Mysecretunlockingcode42!1!`
* `private_key` - The SSL private key for your certificate, in PEM format.  Example:      -----BEGIN RSA PRIVATE KEY-----     jO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca..JZNUgYYwNS0dP2UK     tmyN+XqVcAKw4HqVmChXy5b5msu8eIq3uc2NqNVtR..2ksSLukP8pxXcHyb     +sEwvM4uf8qbnHAqwnOnP9+KV9vds6BaH1eRA4CHz..n+NVZlzBsTxTlS16     /Umr7wJzVrMqK5sDiSu4WuaaBdqMGfL5hLsTjcBFD..Da2iyQmSKuVD4lIZ     ...     -----END RSA PRIVATE KEY-----
* `public_certificate` - The public certificate, in PEM format, that you received from your SSL certificate provider.  Example:      -----BEGIN CERTIFICATE-----     MIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbMQswCQYDVQQGEwJKUDEOMAwG     A1UECBMFVG9reW8xEDAOBgNVBAcTB0NodW8ta3UxETAPBgNVBAoTCEZyYW5rNERE     MRgwFgYDVQQLEw9XZWJDZXJ0IFN1cHBvcnQxGDAWBgNVBAMTD0ZyYW5rNEREIFdl     YiBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBmcmFuazRkZC5jb20wHhcNMTIw     ...     -----END CERTIFICATE----- 



### Create Operation
Creates an asynchronous request to add an SSL certificate bundle.

Set the terraform flag `lifecycle { create_before_destroy = true }` in your certificate to facilitate rotating certificates. 
A certificate cannot be deleted if it is attached to another resource (a listener or a backend set for example).
Because certificate_name in the listener is an updatable parameter, terraform will attempt to recreate the certificate first and then update the listener but the certificate cannot be deleted while it is attached to a listener so it will fail.
Setting the flag makes it so that when a certificate is recreated, the new certificate will be created first before the old one gets deleted.
Whenever you change any values on a certificate that causes it to be recreated the certificate_name MUST also change. Otherwise you will get an error saying that a certificate with that name already exists.

The following arguments are supported:

* `ca_certificate` - (Optional) The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.  Example:      -----BEGIN CERTIFICATE-----     MIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD..AkGA1UEBhMCR0Ix     EzARBgNVBAgTClNvbWUtU3RhdGUxFDASBgNVBAoTC0..0EgTHRkMTcwNQYD     VQQLEy5DbGFzcyAxIFB1YmxpYyBQcmltYXJ5IENlcn..XRpb24gQXV0aG9y     aXR5MRQwEgYDVQQDEwtCZXN0IENBIEx0ZDAeFw0wMD..TUwMTZaFw0wMTAy     ...     -----END CERTIFICATE----- 
* `certificate_name` - (Required) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate bundle.
* `passphrase` - (Optional) A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase. 
* `private_key` - (Optional) The SSL private key for your certificate, in PEM format.  Example:      -----BEGIN RSA PRIVATE KEY-----     jO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca..JZNUgYYwNS0dP2UK     tmyN+XqVcAKw4HqVmChXy5b5msu8eIq3uc2NqNVtR..2ksSLukP8pxXcHyb     +sEwvM4uf8qbnHAqwnOnP9+KV9vds6BaH1eRA4CHz..n+NVZlzBsTxTlS16     /Umr7wJzVrMqK5sDiSu4WuaaBdqMGfL5hLsTjcBFD..Da2iyQmSKuVD4lIZ     ...     -----END RSA PRIVATE KEY----- 
* `public_certificate` - (Optional) The public certificate, in PEM format, that you received from your SSL certificate provider.  Example:      -----BEGIN CERTIFICATE-----     MIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM..QswCQYDVQQGEwJKU     A1UECBMFVG9reW8xEDAOBgNVBAcTB0NodW8ta3UxE..TAPBgNVBAoTCEZyY     MRgwFgYDVQQLEw9XZWJDZXJ0IFN1cHBvcnQxGDAWB..gNVBAMTD0ZyYW5rN     YiBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBmc..mFuazRkZC5jb20wH     ...     -----END CERTIFICATE----- 


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_load_balancer_certificate" "test_certificate" {
	#Required
	certificate_name = "${var.certificate_certificate_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

	#Optional
	ca_certificate = "${var.certificate_ca_certificate}"
	passphrase = "${var.certificate_passphrase}"
	private_key = "${var.certificate_private_key}"
	public_certificate = "${var.certificate_public_certificate}"

	lifecycle {
	    create_before_destroy = true
	}
}
```

# oci_load_balancer_certificates

## Certificate DataSource

Gets a list of certificates.

### List Operation
Lists all SSL certificates bundles associated with a given load balancer.
The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the certificate bundles to be listed. 


The following attributes are exported:

* `certificates` - The list of certificates.

### Example Usage

```hcl
data "oci_load_balancer_certificates" "test_certificates" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```