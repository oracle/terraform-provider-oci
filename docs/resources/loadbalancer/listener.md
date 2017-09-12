# oci\_load\_balancer\_backendset

Provide a load balancer listener resource.

## Example Usage

```
resource "oci_load_balancer_listener" "t" {
  load_balancer_id         = "stub_load_balancer_id"
  name                     = "stub_name"
  default_backend_set_name = "stub_backend_set_name"
  port                     = 1234
  protocol                 = "stub_protocol"

  ssl_configuration {
      certificate_name        = "stub_certificate_name"
      verify_depth            = 6
      verify_peer_certificate = false
  }
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed.
* `default_backend_set_name` - (Required) The name of the associated backend set.
* `port` - (Required) The communication port for the listener.
* `protocol` - (Required) The protocol on which the listener accepts connection requests. 
* `ssl_configuration` - (Optional) An SSL Configuration


## Attributes Reference
None