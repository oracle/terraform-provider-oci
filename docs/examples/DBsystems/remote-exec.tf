resource "null_resource" "remote-exec" {
    depends_on = ["baremetal_database_db_system.TFDBNode"]
    provisioner "file" {
      connection {
        agent = false
        timeout = "10m"
        host = "${data.baremetal_core_vnic.DBNodeVnic.private_ip_address}"
        user = "${var.HostUserName}"
        private_key = "${var.ssh_private_key}"
        # Bastion details
        bastion_host = "${var.BastionHost}"
        bastion_user = "${var.HostUserName}"
        bastion_key = "${var.ssh_private_key}"        
    }
      source = "./scripts/bootstrap"
      destination = "~/bootstrap.sh"
 
    }    
    
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "10m"
        host = "${data.baremetal_core_vnic.DBNodeVnic.private_ip_address}"
        user = "${var.HostUserName}"
        private_key = "${var.ssh_private_key}"
        # Bastion details
        bastion_host = "${var.BastionHost}"       
        bastion_user = "${var.HostUserName}"
        bastion_key = "${var.ssh_private_key}"
    }
      inline = [
        "chmod +x ~/bootstrap.sh",
	"~/bootstrap.sh",
      ]
    }
}
