# oci\_load\_balancer\_backend

[Backend Reference][45922089]

  [45922089]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/ "BackendReference"

Provide a load balancer backend server resource.

## Example Usage

```
resource "oci_load_balancer_backend" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
  ip_address       = "1.2.3.4"
  port             = 1234
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `backendset_name` - (Required) The name of the backend set to add the backend server to. Must be unique and is not changeable. Shows the IP address and port. Example: `10.10.10.4:8080`
* `ip_address` - (Required) The IP address of the backend server.
* `port` - (Required) The communication port for the backend server.
* `backup` - (Optional) Whether the load balancer should treat this server as a backup unit. Example: `true`
* `drain` - (Optional) Whether the load balancer should drain this server. Example: `true`
* `offline` - (Optional) Whether the load balancer should treat this server as offline. Example: `true`
* `weight` - (Optional) The load balancing policy weight assigned to the server.


## Attributes Reference
None
