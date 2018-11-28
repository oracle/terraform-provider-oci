### Use Route Table Attachment to avoid cyclic dependency between Subnet and Route Table

This example uses the `oci_core_route_table_attachment` resource to resolve this dependency cycle problem:
 * oci_core_vnic_attachment.ExampleVnicAttachment
 * oci_core_private_ip.TFPrivateIP
 * oci_core_route_table.ExampleRouteTable 
 * oci_core_subnet.ExampleSubnet
 * oci_core_instance.ExampleInstance