# oci\_load\_balancer\_certificates

Provide a list of load balancer certificates.

## Example Usage

```
data "oci_load_balancer_certificates" "t" {
  load_balancer_id   = "ocid1.loadbalancer.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.


## Attributes Reference
* `certificates` - The list of certificates.

## Certificate Reference
* `ca_certificate` - The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
* `public_certificate` - The public certificate, in PEM format, that you received from your SSL certificate provider.
* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed.
