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
This is the most powerful pre-configured option, providing high density and performance for EDH on OCI. This environment provides high availability, and is an appropriate entry point for scaling up a Production Big Data practice. For larger scale deployments, see Custom (N-Node).

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
