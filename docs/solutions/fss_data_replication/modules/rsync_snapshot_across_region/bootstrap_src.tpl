#cloud-config

packages:
  - rsync
  - nfs-utils  
runcmd:
  - sudo yum -y install rsync nfs-utils
  - sudo mkdir -p ${src_mount_path}
  - sudo mount ${src_mount_target_private_ip}:${src_export_path} ${src_mount_path}
  - sudo crontab /etc/cron.d/fss-sync-up-snapshot

write_files:     
  - path: /etc/cron.d/fss-sync-up-snapshot
    owner: root:root    
    permissions: 0644
    content: |
      ${snapshot_frequency} mkdir ${src_mount_path}.snapshot/tf-fss-snapshot-`date -u +\%Y-\%m-\%dT\%H` 
