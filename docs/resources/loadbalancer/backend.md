# baremetal\_load\_balancer\_backend

Provide a load balancer backend resource.

## Example Usage

```
resource "baremetal_load_balancer_backend" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
  name             = "stub_backend_name"
  ip_address       = "1.2.3.4"
  port             = 1234
  backup           = true
  drain            = true
  offline          = true
  weight           = 1
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `backendset_name` - (Required) The public IP address of the on-premise router.
* `name` - (Required) A name to uniquely identify this backend server in the backend set.
* `ip_address` - (Required) The IP address of the backend server.
* `port` - (Required) The communication port for the backend server.
* `backup` - (Optional) Whether the load balancer should treat this server as a backup unit.
* `drain` - (Optional) Whether the load balancer should drain this server.
* `offline` - (Optional) Whether the load balancer should treat this server as offline. 
* `weight` - (Optional) The load balancing policy weight assigned to the server.


## Attributes Reference
None
