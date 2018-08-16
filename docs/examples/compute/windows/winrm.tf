# Use the setup.ps1 as a template and pass the block volume ipv4 and iqn for ISCSI
data "template_file" "setup_ps1" {
  vars {
    volume_ipv4 = "${oci_core_volume_attachment.TFVolumeAttachment.ipv4}"
    volume_iqn  = "${oci_core_volume_attachment.TFVolumeAttachment.iqn}"
  }

  template = "${file("${var.userdata}/${var.setup_ps1}")}"
}

####################################
# Execute Windows Commands Remotely
####################################
/*
  This resource depends on instance creation and volume attachment. After both of them are ready, will perform scripts execution on remote instances.
  This block will help upload the scripts to remote instances and execute them.

  Troubleshooting tips:

  To check if winrm is listening:
  curl --header "Content-Type: application/soap+xml;charset=UTF-8" --header "WSMANIDENTIFY: unauthenticated" --insecure https://<ip-address>:5986/wsman --data '&lt;s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsmid="http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"&gt;&lt;s:Header/&gt;&lt;s:Body&gt;&lt;wsmid:Identify/&gt;&lt;/s:Body&gt;&lt;/s:Envelope&gt;'

  To check winrm with authentication:
  curl --header "Content-Type: application/soap+xml;charset=UTF-8" --insecure https://<ip-address>:5986/wsman --basic  -u opc:password --data '&lt;s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsmid="http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"&gt;&lt;s:Header/&gt;&lt;s:Body&gt;&lt;wsmid:Identify/&gt;&lt;/s:Body&gt;&lt;/s:Envelope&gt;'

  Using winrm-cli from https://github.com/masterzen/winrm-cli that uses same underlying library that Terraform uses: https://github.com/masterzen/winrm
  ./winrm -hostname <ip-address> -username "opc" -password "password" -https -insecure "ipconfig /all"

  The winrm Go library uses BasicAuth, please refer to the cloudinit.yml to configure winrm with that option otherwise you may get a http response error

  Because runcmd and script are not yet supported via #cloud-config in the above images, one can use #ps1_sysnative to change password and configure winrm with a https listener

  When using your own images:
  You may alternatively prepare your images by configuring winrm using the same commands as listed in cloudinit.yml

  winrm quickconfig
  Enable-PSRemoting

  winrm set winrm/config/client/auth '@{Basic="true"}'
  winrm set winrm/config/service/auth '@{Basic="true"}'
  winrm set winrm/config/service '@{AllowUnencrypted="true"}'
  winrm set winrm/config/winrs '@{MaxMemoryPerShellMB="300"}'
  winrm set winrm/config '@{MaxTimeoutms="1800000"}'

  netsh advfirewall firewall add rule name="WinRM HTTP" protocol=TCP dir=in profile=any localport=5985 remoteip=any localip=any action=allow
  netsh advfirewall firewall add rule name="WinRM HTTPS" protocol=TCP dir=in profile=any localport=5986 remoteip=any localip=any action=allow

  net stop winrm
  sc.exe config winrm start=auto
  net start winrm
*/

# If your image does not have winrm enabled, please follow steps above to configure it
# The remote-exec-windows resource will only execute if this is set to true
variable "IsWinRMConfiguredForImage" {
  default = "true"
}

# Use the https 5986 port for winrm by default
# If that fails with a http response error: 401 - invalid content type, the SSL may not be configured correctly
# In that case you can switch to http 5985 by setting this to false, and configuring winrm to AllowUnencrypted traffic
variable "IsWinRMConfiguredForSSL" {
  default = "true"
}

resource "null_resource" "wait_for_cloudinit" {
  depends_on = ["oci_core_instance.TFInstance", "oci_core_volume_attachment.TFVolumeAttachment"]

  count = "${var.IsWinRMConfiguredForImage == "true" ? 1 : 0}"

  # WinRM configuration through cloudinit may not have completed and password may not have changed yet, so sleep before doing remote-exec
  provisioner "local-exec" {
    command = "sleep 60"
  }
}

resource "null_resource" "remote-exec-windows" {
  # Although we wait on the wait_for_cloudinit resource, the configuration may not be complete, if this step fails please retry
  depends_on = ["oci_core_instance.TFInstance", "oci_core_volume_attachment.TFVolumeAttachment", "null_resource.wait_for_cloudinit"]

  # WinRM connections via Terraform are going to fail if it is not configured correctly as mentioned in comment section above
  count = "${var.IsWinRMConfiguredForImage == "true" ? 1 : 0}"

  provisioner "file" {
    connection = {
      type     = "winrm"
      agent    = false
      timeout  = "1m"
      host     = "${oci_core_instance.TFInstance.public_ip}"
      user     = "${data.oci_core_instance_credentials.InstanceCredentials.username}"
      password = "${random_string.instance_password.result}"
      port     = "${var.IsWinRMConfiguredForSSL == "true" ? 5986 : 5985}"
      https    = "${var.IsWinRMConfiguredForSSL}"
      insecure = "true"                                                               #self-signed certificate
    }

    content     = "${data.template_file.setup_ps1.rendered}"
    destination = "c:/setup.ps1"
  }

  provisioner "remote-exec" {
    connection {
      type     = "winrm"
      agent    = false
      timeout  = "1m"
      host     = "${oci_core_instance.TFInstance.public_ip}"
      user     = "${data.oci_core_instance_credentials.InstanceCredentials.username}"
      password = "${random_string.instance_password.result}"
      port     = "${var.IsWinRMConfiguredForSSL == "true" ? 5986 : 5985}"
      https    = "${var.IsWinRMConfiguredForSSL}"
      insecure = "true"                                                               #self-signed certificate
    }

    inline = [
      "powershell.exe -file c:/setup.ps1",
    ]
  }
}
