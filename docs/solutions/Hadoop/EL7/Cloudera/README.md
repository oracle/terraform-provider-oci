# Cloudera on OCI automation with Terraform
Included here are Terraform templates for deploying a fully configured Cloudera EDH instance or cluster on OCI.  Descriptions follow:

## Sandbox
This is a great starting point for customers wanting to explore the power and functionality of EDH on OCI.  This deployment consists of a single instance running the Cloudera Docker Container.  This is a good fit for individuals who want to explore Cloudera on OCI, while maintaining a cost effective bottom line.  This is not a good fit for multiple users, development efforts, or large datasets.

Minimum Instance : VM.Standard1.8

Suggested Instance: VM.Standard2.8

## Development
A small implementation, this is the next step up for running EDH on OCI.  This deployment consists of 5 instances - a Bastion host, Utility Host, and 3 Workers.  This environment provides a much higher HDFS storage capacity, along with a good size pool of compute and memory resources for use with a variety of Big Data workloads.   This environment is not a good fit for customers who want highly available services, as the reduced infrastructure footprint does not support high availability.

| Minimum Worker Instance | Minimum Bastion Instance | Minimum Utility Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.Standard1.16 with (3) 700GB Block Storage devices per worker | VM.Standard1.4 | VM.Standard1.8  |                   


| Suggested Worker Instance | Suggested Bastion Instance | Suggested Utility Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.Standard2.24 with (3) 1TB Block Storage devices per worker | VM.Standard2.4 | VM.Standard2.8 |

## Production Starter
The largest pre-set configuration for EDH on OCI, this deployment contains 10 instances - a Bastion host, Utility Host, 2 Master hosts, and 6 workers.  This is the best pre-configured option for the most density and performance for EDH on OCI.  This environment provides high availability, and is an appropriate entry point for scaling up a Production Big Data practice.

| Minimum Worker Instance | Minimum Bastion Instance | Minimum Utility & Master Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO1.36 | VM.Standard1.4 | VM.Standard1.8 |                                

| Suggested Worker Instance | Suggested Bastion Instance | Suggested Utility & Master Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO2.52 | VM.Standard2.4 | VM.Standard2.8 |                                   

## Custom (N-Node)
OCI also supports N-Node EDH implementations for customers whose needs may exceed the performance or capacity limitations of the largest pre-set cluster configuration.   Please contact OCI for more information, we are happy to work with you to provide guidance on the optimal cluster deployment for your needs, and have an automated solution to support dynamic cluster sizes scaling into the thousands of nodes.

| Minimum Worker Instance | Minimum Bastion Instance | Minimum Utility & Master Instance |
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO1.36 | VM.Standard1.4 | VM.Standard1.8 |    

| Suggested Worker Instance | Suggested Bastion Instance | Suggested Utility & Master Instance |
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO2.52 | VM.Standard2.4 | VM.Standard2.16 |         

## Security Auditing

Note that Terraform Automation templates currently disable the local firewall as part of the installation process.   It is recommended to run a security audit as part of the post-deployment process when choosing to use these automation templates.   This should include re-enabling host firewalls as needed, and adjusting Security List rules to be more restrictive to access requirements - this should include restricting access to publicly available services on Bastion & Utility hosts to allowed IP sources/segments only.  Oracle will be working to provide automation with local firewalls enabled in the future as part of updates to these Terraform templates.

Also note that hosts on the Private Subnet currently require internet access as part of the deployment process to fetch installation packages.  This can also be disabled as part of the post-deployment process, but is currently required for setup.   Oracle is working with Cloudera to enhance this deployment process to provide additional security as part of the automation with Terraform.
