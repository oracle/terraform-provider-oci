# Oracle Cloud Infrastructure Terraform Provider Example for Windows VM

This example contains Terraform configuration to provision a virtual machine in Oracle Cloud Infrastructure (OCI) with Microsoft Windows.

## What this example covers

-   Deploying networking resources that creates a VCN, Subnet, Route Table, Internet Gateway and a Security List to allow RDP & WinRM traffic for the VM
    -   [networking.tf](networking.tf)
-   Deploying a Windows VM instance with one of [published images on OCI](https://docs.cloud.oracle.com/iaas/images/)
    -   [windows.tf](windows.tf)
-   Using the Windows version for [Cloud-Init](https://cloud-init.io/) -  [Cloudbase-Init](https://cloudbase.it/cloudbase-init/) available on the VM to setup and configure Windows
    -   [cloudinit.ps1](userdata/cloudinit.ps1) - #ps1_sysnative
        -   Change the initial Windows VM Password
        -   Configure WinRM for HTTPS connections
    -   [cloudinit.yml](userdata/cloudinit.yml) - #cloud-config
        -   Write custom files
-   Using Terraform remote-exec execute custom scripts through WinRM
    -   [winrm.tf](winrm.tf)
    -   [setup.ps1](userdata/setup.ps1)
        -   Mount block volumes (iscsi) attached to the VM

## CloudInit

-   The latest [Windows Images](https://docs.cloud.oracle.com/iaas/images/windows-server-2012-r2-vm/) released in and after July 2018 for OCI come with Windows version of [Cloud-Init](https://cloud-init.io/), [Cloudbase-Init](https://cloudbase.it/cloudbase-init/) and WinRM enabled by default, refer to the release notes of the [images](https://docs.cloud.oracle.com/iaas/images/) to ensure the version you choose has these enabled for this example to run
-   There are multiple ways in which CloudBase-Init can be used by providing a Base64 encoded metadata in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/LaunchInstanceDetails), this example uses the #cloud-config and #ps1_sysnative variations
-   All the plugins in #cloud-config format may not be supported yet, currently script and runcmd are not supported, but some other plugins like write_files are
    -   Alternatively, to run custom Windows shell commands or Powershell commands, one can alternately use #ps1_sysnative to have them execute as part of initial setup

## WinRM

-   While WinRM is enabled on the images, if you plan to use your own images, you need to configure it using following commands.

```powershell
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
```

-   Strongly consider the security aspects of allowing unencrypted connections (HTTP). This example shows how to create a self-signed certificate using Cloudbase-Init to configure WinRM for HTTPS communication
-   Based on what ports you have configured for RDP and WinRM you want to setup the Security List for your VCN to allow those ports, this example covers these ports:
    -   3389 - RDP
    -   5985 - WinRM HTTP
    -   5986 - WinRM HTTPS

## Terraform

-   Terraform uses the Go based winrm(<https://github.com/masterzen/winrm>) library to make remote-exec connections, this library supports BasicAuth by default and can be extended to support additional authentication models
-   The remote-exec in Terraform uses SSH by default, but the type can be changed to `winrm` to execute remote commands on Windows
-   Terraform provider supports both `template_file` and `template_cloudinit_config` data sources that can be used to set metadata for LaunchInstanceDetails

## Tips and Troubleshooting

-   CloudBase-Init is installed at `C:\Program Files\Cloudbase Solutions\Cloudbase-Init` and you can find its logs in the `log` directory under it. Refer [cloudbase-init tutorial](https://cloudbase-init.readthedocs.io/en/latest/tutorial.html) for more information.
-   Use #ps1_sysnative for more advanced VM configuration through Cloudbase-init
-   The Cloudbase-Init setup may not have completed when the VM is reported to be ready, you can either introduce a wait or have the VM write to a remote location that you can poll before launching remote-exec via Terraform, this examples waits 60 seconds
-   The VM instance Cloud-Init metadata that is passed to LaunchInstanceDetails and then read over in VM is just Base64 encoded, you may want to transfer the new password in a more secure way or change it through another remote-exec that can run post Cloudbase-Init. Further, the example also has the passwords stored in the local state file.
    -   Refer Terraform recommendations for [Sensitive Data](https://www.terraform.io/docs/state/sensitive-data.html)
-   The example covers running various Powershell commands, but for a more reliable solution, you may want to add enough retries and error reporting for setup resiliency
-   While setting up HTTP over BasicAuth is easy, it is not a recommended way to connecting to these VMs, consider using HTTPS by configuring WinRM HTTPS listener using your own certificate
-   If you are facing certificate based errors for WinRM HTTPS connection it is probably due to using the self-signed certificate using New-SelfSignedCertificate that WinRM does not find compatible in newer operating systems
    -   Ideally you should use CA based certificate for configuring WinRM
    -   Alternatively, you can use this Ansible published script  [ConfigureRemotingForAnsible.ps1](https://raw.githubusercontent.com/ansible/ansible/devel/examples/scripts/ConfigureRemotingForAnsible.ps1) to generate a legacy self-signed certificate and configure WinRM to use same.
        -   This entire script can be passed as an additional part in `template_cloudinit_config` to configure WinRM for HTTPS with a self-signed certificate. If you do so, remove the certificate based section from `cloudinit.ps1` in this example
    -   To get detailed error and test WinRM connectivity from within the VM, you can use the following commands
    ```powershell
    $cred=Get-Credential
    test-wsman -Authentication Basic -UseSSL -Credential $cred
    ```
-   To test if WinRM is correctly configured on the VM and is listening, you can use any of the following methods to troubleshoot

```powershell
  # To check if winrm is listening
  curl --header "Content-Type: application/soap+xml;charset=UTF-8" --header "WSMANIDENTIFY: unauthenticated" --insecure https://<ip-address>:5986/wsman --data '&lt;s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsmid="http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"&gt;&lt;s:Header/&gt;&lt;s:Body&gt;&lt;wsmid:Identify/&gt;&lt;/s:Body&gt;&lt;/s:Envelope&gt;'

  # To check winrm with authentication:
  curl --header "Content-Type: application/soap+xml;charset=UTF-8" --insecure https://<ip-address>:5986/wsman --basic  -u opc:password --data '&lt;s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsmid="http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"&gt;&lt;s:Header/&gt;&lt;s:Body&gt;&lt;wsmid:Identify/&gt;&lt;/s:Body&gt;&lt;/s:Envelope&gt;'

  # Using winrm-cli from https://github.com/masterzen/winrm-cli that uses same underlying library that Terraform uses: https://github.com/masterzen/winrm
  ./winrm -hostname <ip-address> -username "opc" -password "password" -https -insecure "ipconfig /all"
```

-   From the VM instance you can run the following commands to get the metadata

```powershell
  curl http://169.254.169.254/opc/v1/instance/
  curl http://169.254.169.254/opc/v1/instance/metadata/
  curl http://169.254.169.254/opc/v1/instance/metadata/<any-key-name>

  # To get user_data
  curl http://169.254.169.254/opc/v1/instance/metadata/user_data
```

-   If you are facing issues with connecting or using WinRM from Terraform through remote-exec, an alternative approach can be to use local-exec with another library like [pywinrm](https://github.com/diyan/pywinrm)
