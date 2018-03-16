output "Bastion Hosts Public IP" {
  value = ["${module.dmz_hosts.public_ip}"]
}

output "Bastion Hosts Private IP" {
  value = ["${module.dmz_hosts.private_ip}"]
}

output "Active Directory Servers IP" {
  value = ["${module.admin_hosts.private_ip}"]
}

output "SQL Servers IP" {
  value = ["${module.sql_hosts.private_ip}"]
}

output "SQL Servers Additional IP 1" {
  value = ["${module.secondaryIPs.ip2}"]
}

output "SQL Servers Additional IP 2" {
  value = ["${module.secondaryIPs.ip3}"]
}

output "Witness Servers IP" {
  value = ["${module.witness_hosts.private_ip}"]
}

output "DB ISCSI attachment IP" {
  value = ["${module.sql_hosts.iscsi_attachment_db}"]
}

output "LOG ISCSI attachment IP" {
  value = ["${module.sql_hosts.iscsi_attachment_log}"]
}

output "BACKUP ISCSI attachment IP" {
  value = ["${module.sql_hosts.iscsi_attachment_backup}"]
}

output "Witness ISCSI attachment IP" {
  value = ["${module.witness_hosts.iscsi_attachment}"]
}
