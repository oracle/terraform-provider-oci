variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "SubnetOCID" {}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

# DBSystem specific 
variable "DBNodeShape" {
    default = "VM.Standard1.2"
}

variable "CPUCoreCount" {
    default = "2"
}

variable "DBEdition" {
    default = "ENTERPRISE_EDITION"
}

variable "DBAdminPassword" {
    default = "BEstrO0ng_#11"
}

variable "DBName" {
    default = "aTFdb"
}

variable "DBVersion" {
    default = "12.1.0.2"
}

variable "DBDisplayName" {
    default = "MyTFDB"
}

variable "DBDiskRedundancy" {
    default = "HIGH"
}

variable "DBNodeDisplayName" {
    default = "MyTFDatabaseNode0"
}

variable "DBNodeDomainName" {
    default = "mycompany.com"
}

variable "DBNodeHostName" {
    default = "myOracleDB"
}

# Define existing bastion host
variable "BastionHost" {
    default = "129.146.26.52"
}

variable "HostUserName" {
    default = "opc"
}

variable "NCharacterSet" {
	default = "AL16UTF16"
}

variable "CharacterSet" {
	default = "AL32UTF8"
}

variable "DBWorkload" {
	default = "OLTP"
}

variable "PDBName" {
	default = "pdbName"
}

variable "InitialDataStorageSizeInGB" {
	default = "256"
}

variable "LicenseModel" {
	default = "BRING_YOUR_OWN_LICENSE"
}

variable "NodeCount" {
	default = "1"
}
