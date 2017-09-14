# oci\_load\_balancer\_backendset

Provide a load balancer backend set resource.

## Example Usage

```
resource "oci_load_balancer_backendset" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  name             = "stub_backendset_name"
  policy           = "stub_policy"

  health_checker {
    interval_ms         = 30001
    port                = 1234
    protocol            = "stub_protocol"
    response_body_regex = "stub_regex"
  }

  ssl_configuration {
    certificate_name        = "stub_certificate_name"
    verify_depth            = 6
    verify_peer_certificate = false
  }
  
  session_persistence_configuration {
    cookie_name      = "cookiename"
    disable_fallback = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `name` - (Required) A friendly name for the backend set. It must be unique and it cannot be changed.
* `policy` - (Optional) The load balancer policy for the backend set. The default load balancing policy is 'ROUND_ROBIN'.
* `health_checker` - (Optional) Health Checker Settings
* `ssl_configuration` - (Optional) SSL Configuration Settings
* `session_persistence_configuration` - (Optional) Session persistence enables the Load Balancing Service to direct any number of requests that originate from a single logical client to a single backend web server.


## Attributes Reference
* `backend` - The list of backends


