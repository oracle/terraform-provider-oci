# sql.tf - Settings specific to the configuration of SQL Always On deployment
# Set all variables here to meet your needs within your particular environment
# Check this file and configuration.tf!


# SQL Always On Witness Server configuration
# Use this section to adjust the configuration of the witness server

# witness_deployment - AD in which to deploy the Witness servers
# Provide a single number only
variable "witness_deployment" {
  default = "2"
}

# Size of volume used for witness server (in GB)
variable "witness_volume_size" {
  default = "256"
}

# witness_shape - Shape to use for witness server
variable "witness_shape" {
  default = "VM.Standard1.4"
}

# SQL Server configuration section
# These resourcese are specified PER SERVER! 

# Server specification
# sql_shape - Shape to use for EACH SQL Server deployed
variable "sql_shape" {
  default = "VM.Standard1.8"
}

# Volume sizes
# This section adjusts the sizes of the various volumes that are 
# created as part of this process.  Adjust as needed.

# sql_db_size - Sets size of the volume used for the db partitions.
variable "sql_db_size" {
  default = "2048"
}

# sql_backup_size - Sets size of volume used for backups
variable "sql_backup_size" {
  default = "256"
}

# sql_log_size - Sets size of log volume used for redo logs
variable "sql_log_size" {
  default = "512"
}




