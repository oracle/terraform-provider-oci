# Cloudera Docker VM Version
Note that this template specifically uses wget to fetch a specific version of the Cloudera Docker VM - this is because the version currently availble in public Docker registry is an older version.   The wget command should be updated to fetch the latest Cloudera Docker VM when new versions are released.

# Usage Guide
  
Note that this installation is a stand-alone instance running the Cloudera VM Docker image.  This is a self contained environment.   Access to the Sandbox is done via post deployment URLs.   SSH access is also possible, but because this is running inside docker, shell commands to the container require attaching to the Docker container first:

        ssh -i ~/.ssh/id_rsa opc@<sandbox_public_ip>
        sudo docker ps

Output will show a CONTAINER ID - use that in the following command

        sudo docker exec -it <container_id> bash

## PREREQUISITES

Installation has a dependency on Terraform being installed and configured for the user tenancy.   As such an "env-vars" file is included with this package which contains all the necessary environment variables.  This file should be updated with the appropriate values prior to installation.  To source this file prior to installation, either reference it in your .rc file for your shell or run the following:

        source env-vars

## Deployment

Deploy using standard terraform commands

        terraform init && terraform plan && terraform apply

## Post Deployment

All post deployment for the Sandbox instance is done in remote-exec as part of the Terraform apply process.  You will see output on the screen as part of this process.  Once complete, URLs for access to the Sandbox will be displayed.
