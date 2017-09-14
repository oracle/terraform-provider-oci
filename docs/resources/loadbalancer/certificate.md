# oci\_load\_balancer\_certificate

Provide a load balancer certificate set resource.

## Example Usage

```
resource "oci_load_balancer_certificate" "t" {
  load_balancer_id   = "ocid1.loadbalancer.stub_id"
  ca_certificate     = "stub_ca_certificate"
  certificate_name   = "stub_certificate_name"
  passphrase         = "stub_passphrase"
  private_key        = "stub_private_key"
  public_certificate = "stub_public_certificate"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `ca_certificate` - (Required) The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
* `private_key` - (Required) The SSL private key for your certificate, in PEM format.
* `public_certificate` - (Required) The public certificate, in PEM format, that you received from your SSL certificate provider.
* `passphrase` - (Optional) A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.
* `certificate_name` - (Optional) A friendly name for the certificate bundle. It must be unique and it cannot be changed.


## Attributes Reference
None