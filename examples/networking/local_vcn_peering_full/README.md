# Local VCN Peering Example

This example demonstrates how to do a VCN local peering connection using Local Peering Gateway (LPG) resources when you have different administrators of the VCNs in the connection.

** IMPORTANT **
You would not want to use this example the way it is written as it uses multiple users. This example is there to demonstrate the workflow of establishing a local peering connection when the 2 VCNs are administered by different users.

This example creates policies so it should be run in the home region.

One of the users will have the `requestor` LPG that will request a local peering connection to the `acceptor` LPG that is managed by a different user. See [Local VCN Peering](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/localVCNpeering.htm) for more details.

* policies.tf show the policies that are needed for each of the users.
* requestor.tf shows what the requestor config would look like, including the LPG, the Route Table and the Security List.
* acceptor.tf shows what the acceptor config would look like, including the LPG, the Route Table and the Security List.

An instance is created on the requestor side and the acceptor side so that you can test the connection. 
You can SSH to one of the instances using its public IP and try to PING from there the other instance using its Private IP.
