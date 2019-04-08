$ipv4='${volume_ipv4}'
$iqn='${volume_iqn}'

echo 'Hello from terraform' > C:\hello_from_terraform_1234567890.txt

Write-Output 'Configuring ISCSI for Block Volumes'

Set-Service -Name msiscsi -StartupType Automatic
Start-Service msiscsi

New-IscsiTargetPortal -TargetPortalAddress $ipv4
Connect-IscsiTarget -NodeAddress $iqn -TargetPortalAddress $ipv4 -IsPersistent $True

Write-Output 'Configured ISCSI for Block Volumes'

Write-Output 'Configuring the new disk for a partition and file system'
Get-Disk -Number 1 | Initialize-Disk -PartitionStyle MBR -PassThru | New-Partition -AssignDriveLetter -UseMaximumSize | Format-Volume -FileSystem NTFS -NewFileSystemLabel "tfdisk" -Confirm:$false
Write-Output 'Configured the new disk'
