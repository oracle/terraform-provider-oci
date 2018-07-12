#cloud-config

write_files:
  # Create file to be used when enabling ip forwarding
  - path: /etc/sysctl.d/98-ip-forward.conf
    content: |
      net.ipv4.ip_forward = 1

runcmd:
  # Run firewall commands to enable masquerading and port forwarding
  # Enable ip forwarding by setting sysctl kernel parameter
  - firewall-offline-cmd --direct --add-rule ipv4 filter FORWARD 0 -i ens3 -j ACCEPT
  - firewall-offline-cmd --direct --add-rule ipv4 filter FORWARD 0 -i ens4 -j ACCEPT
  - /bin/systemctl restart firewalld
  - sysctl -p /etc/sysctl.d/98-ip-forward.conf

