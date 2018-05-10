resource "null_resource" "remote-exec1" {
    depends_on = ["null_resource.mongokey","oci_core_instance.MongoP","null_resource.remote-exec2","null_resource.remote-exec3"]
    provisioner "file" {
      source = "mongo.sh"
      destination = "/home/opc/mongo.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongokey"
      destination = "/home/opc/mongokey"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    } 
    }
    provisioner "file" {
      source = "mongo.exec"
      destination = "/home/opc/mongo.exec"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongod.conf"
      destination = "/home/opc/mongod.conf"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
	"chown opc:opc /home/opc/mongo.sh",
	"chmod +x /home/opc/mongo.sh",
      	"/home/opc/mongo.sh"
	]
    }
}

resource "null_resource" "mongo-master" {
    depends_on = ["null_resource.remote-exec1"]
    provisioner "file" {
      source = "mongo_clusteradmin.exec"
      destination = "/home/opc/mongo_clusteradmin.exec"
    connection {
      agent = false
      timeout = "30m"
      host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
      user = "opc" 
      private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongo_admin.exec"
      destination = "/home/opc/mongo_admin.exec"
    connection {
      agent = false
      timeout = "30m"
      host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
      user = "opc"
      private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongo_admin.sh"
      destination = "/home/opc/mongo_admin.sh"
    connection {
      agent = false
      timeout = "30m"
      host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
      user = "opc"
      private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongo_pritunl.exec"
      destination = "/home/opc/mongo_pritunl.exec"
    connection {
      agent = false
      timeout = "30m"
      host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
      user = "opc"
      private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MPVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
	"chmod +x /home/opc/mongo_admin.sh",
        "/home/opc/mongo_admin.sh"
        ]
    }
}


resource "null_resource" "remote-exec2" {
    depends_on = ["null_resource.mongokey","oci_core_instance.MongoR1"]
    provisioner "file" {
      source = "mongo.sh"
      destination = "/home/opc/mongo.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongokey"
      destination = "/home/opc/mongokey"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongod.conf"
      destination = "/home/opc/mongod.conf"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chown opc:opc /home/opc/mongo.sh",
        "chmod +x /home/opc/mongo.sh",
        "/home/opc/mongo.sh",
	"/usr/bin/mongo < /home/opc/mongo.exec"
        ]
    }
}

resource "null_resource" "remote-exec3" {
    depends_on = ["null_resource.mongokey","oci_core_instance.MongoR2"]
    provisioner "file" {
      source = "mongo.sh"
      destination = "/home/opc/mongo.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongokey"
      destination = "/home/opc/mongokey"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "mongod.conf"
      destination = "/home/opc/mongod.conf"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.MR2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chown opc:opc /home/opc/mongo.sh",
        "chmod +x /home/opc/mongo.sh",
        "/home/opc/mongo.sh"
        ]
    }
}

resource "null_resource" "remote-exec4" {
    depends_on = ["oci_core_instance.Pritunl1"]
    provisioner "file" {
      source = "pritunl.sh"
      destination = "/home/opc/pritunl.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chown opc:opc /home/opc/pritunl.sh",
        "chmod +x /home/opc/pritunl.sh",
        "/home/opc/pritunl.sh"
        ]
    }
}

resource "null_resource" "remote-exec5" {
    depends_on = ["oci_core_instance.Pritunl2"]
    provisioner "file" {
      source = "pritunl.sh"
      destination = "/home/opc/pritunl.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chown opc:opc /home/opc/pritunl.sh",
        "chmod +x /home/opc/pritunl.sh",
        "/home/opc/pritunl.sh"
        ]
    }
}

resource "null_resource" "pritunl1-key" {
    depends_on = ["null_resource.mongo-master"]
    provisioner "file" {
      source = "pritunl.conf"
      destination = "/home/opc/pritunl.conf"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "pritunl2.sh"
      destination = "/home/opc/pritunl2.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P1VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chmod +x /home/opc/pritunl2.sh",
	"/home/opc/pritunl2.sh"
        ]
    }
}

resource "null_resource" "pritunl2-key" {
    depends_on = ["null_resource.mongo-master"]
    provisioner "file" {
      source = "pritunl.conf"
      destination = "/home/opc/pritunl.conf"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "file" {
      source = "pritunl2.sh"
      destination = "/home/opc/pritunl2.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.P2VNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chmod +x /home/opc/pritunl2.sh",
	"/home/opc/pritunl2.sh"
        ]
    }
}

resource "null_resource" "pritunl-link" {
depends_on = ["oci_core_instance.Pritunllink"]
    provisioner "file" {
      source = "pritunl-link.sh"
      destination = "/home/opc/pritunl-link.sh"
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.PTLVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.PTLVNIC.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
        "chmod +x /home/opc/pritunl-link.sh",
        "/home/opc/pritunl-link.sh"
        ]
    }
}


