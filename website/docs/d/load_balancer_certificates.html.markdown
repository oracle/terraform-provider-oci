---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_certificates"
sidebar_current: "docs-oci-datasource-load_balancer-certificates"
description: |-
  Provides the list of Certificates in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_certificates
This data source provides the list of Certificates in Oracle Cloud Infrastructure Load Balancer service.

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

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the certificate bundles to be listed. 


## Attributes Reference

The following attributes are exported:

* `certificates` - The list of certificates.

### Certificate Reference

The following attributes are exported:

* `ca_certificate` - The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.

	Example:

	    -----BEGIN CERTIFICATE-----
	    MIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD..AkGA1UEBhMCR0Ix
	    EzARBgNVBAgTClNvbWUtU3RhdGUxFDASBgNVBAoTC0..0EgTHRkMTcwNQYD
	    VQQLEy5DbGFzcyAxIFB1YmxpYyBQcmltYXJ5IENlcn..XRpb24gQXV0aG9y
	    aXR5MRQwEgYDVQQDEwtCZXN0IENBIEx0ZDAeFw0wMD..TUwMTZaFw0wMTAy
	    ...
	    -----END CERTIFICATE-----
	
* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
* `public_certificate` - The public certificate, in PEM format, that you received from your SSL certificate provider.

	Example:

	    -----BEGIN CERTIFICATE-----
	    MIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbMQswCQYDVQQGEwJKUDEOMAwG
	    A1UECBMFVG9reW8xEDAOBgNVBAcTB0NodW8ta3UxETAPBgNVBAoTCEZyYW5rNERE
	    MRgwFgYDVQQLEw9XZWJDZXJ0IFN1cHBvcnQxGDAWBgNVBAMTD0ZyYW5rNEREIFdl
	    YiBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBmcmFuazRkZC5jb20wHhcNMTIw
	    ...
	    -----END CERTIFICATE-----
	

