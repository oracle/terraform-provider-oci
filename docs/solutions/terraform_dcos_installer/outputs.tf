output "master_public_ips" {
  value = "${compact(concat(module.dcos_master_ad1.public_ips,module.dcos_master_ad2.public_ips,module.dcos_master_ad3.public_ips))}"
}

output "master_private_ips" {
  value = "${concat(module.dcos_master_ad1.private_ips,module.dcos_master_ad2.private_ips,module.dcos_master_ad3.private_ips )}"
}

