
# Create LB backend VMs and add them to LB backendset

resource "oci_core_instance" "instance1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "instance1"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    subnet_id = "${oci_core_subnet.BESubnet1.id}"
    hostname_label = "instance1"
    metadata {
      ssh_authorized_keys = "${var.ssh_public_key}"
      user_data = "${base64encode(var.user-data)}"
    }
}

resource "oci_core_instance" "instance2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "instance2"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    subnet_id = "${oci_core_subnet.BESubnet2.id}"
    hostname_label = "instance2"
    metadata {
      ssh_authorized_keys = "${var.ssh_public_key}"
      user_data = "${base64encode(var.user-data)}"
    }
}

# Create a default html file on the backend VM
variable "user-data" {
  default = <<EOF
#!/bin/bash -x
echo '################### webserver userdata begins #####################'
touch ~opc/userdata.`date +%s`.start

# echo '########## yum update all ###############'
# yum update -y

echo '########## basic webserver ##############'
yum install -y httpd
systemctl enable  httpd.service
systemctl start  httpd.service


echo '<html><head></head><body><pre><code>' > /var/www/html/index.html
hostname >> /var/www/html/index.html
echo '' >> /var/www/html/index.html
cat /etc/os-release >> /var/www/html/index.html
echo '</code></pre></body></html>' >> /var/www/html/index.html

systemctl enable  firewalld

firewall-offline-cmd --add-service=http
systemctl restart  firewalld

touch ~opc/userdata.`date +%s`.finish
echo '################### webserver userdata ends #######################'
EOF
}


# Add VMs to lb1 backendset

resource "oci_load_balancer_backend" "lb1-be1" {
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  backendset_name  = "${oci_load_balancer_backendset.lb1-bes1.id}"
  ip_address       = "${oci_core_instance.instance1.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "oci_load_balancer_backend" "lb1-be2" {
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  backendset_name  = "${oci_load_balancer_backendset.lb1-bes1.id}"
  ip_address       = "${oci_core_instance.instance2.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

# Add VMs to lb2 backendset

resource "oci_load_balancer_backend" "lb2-be1" {
  load_balancer_id = "${oci_load_balancer.lb2.id}"
  backendset_name  = "${oci_load_balancer_backendset.lb2-bes1.id}"
  ip_address       = "${oci_core_instance.instance1.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "oci_load_balancer_backend" "lb2-be2" {
  load_balancer_id = "${oci_load_balancer.lb2.id}"
  backendset_name  = "${oci_load_balancer_backendset.lb2-bes1.id}"
  ip_address       = "${oci_core_instance.instance2.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

