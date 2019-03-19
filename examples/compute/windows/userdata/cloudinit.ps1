#ps1_sysnative

# Template variables
$user='${instance_user}'
$password='${instance_password}'
$computerName='${instance_name}'

Write-Output "Changing $user password"
net user $user $password
Write-Output "Changed $user password"

Write-Output "Configuring WinRM"
# Allow unencrypted if you wish to use http 5985 endpoint
winrm set winrm/config/service '@{AllowUnencrypted="true"}'

# Create a self-signed certificate to configure WinRM for HTTPS
$cert = New-SelfSignedCertificate -CertStoreLocation 'Cert:\LocalMachine\My' -DnsName $computerName
Write-Output "Self-signed SSL certificate generated with details: $cert"

$valueSet = @{
    Hostname = $computerName
    CertificateThumbprint = $cert.Thumbprint
}

$selectorSet = @{
    Transport = "HTTPS"
    Address = "*"
}

# Remove any prior HTTPS listener
$listeners = Get-ChildItem WSMan:\localhost\Listener
If (!($listeners | Where {$_.Keys -like "TRANSPORT=HTTPS"}))
{
    Remove-WSManInstance -ResourceURI 'winrm/config/Listener' -SelectorSet $selectorSet
}

Write-Output "Enabling HTTPS listener"
New-WSManInstance -ResourceURI 'winrm/config/Listener' -SelectorSet $selectorSet -ValueSet $valueSet
Write-Output "Enabled HTTPS listener"

Write-Output "Configured WinRM"
