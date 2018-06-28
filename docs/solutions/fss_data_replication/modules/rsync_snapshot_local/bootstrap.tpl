#cloud-config

packages:
  - rsync
  - nfs-utils  
runcmd:
  - sudo yum -y install rsync nfs-utils
  - sudo mkdir -p ${src_mount_path}
  - sudo mkdir -p ${dst_mount_path}
  - sudo mount ${src_mount_target_private_ip}:${src_export_path} ${src_mount_path}
  - sudo mount ${dst_mount_target_private_ip}:${dst_export_path} ${dst_mount_path}
  - sudo crontab /etc/cron.d/fss-sync-up-snapshot

write_files:     
  - path: /etc/cron.d/fss-sync-up-snapshot
    owner: root:root    
    permissions: 0644
    content: |
      ${snapshot_frequency} mkdir ${src_mount_path}.snapshot/tf-fss-snapshot-`date -u +\%Y-\%m-\%dT\%H` 
      ${data_sync_frequency} /usr/bin/flock -n /var/run/fss-sync-up-snapshot.lck rsync -aHAXxv --numeric-ids --delete ${src_mount_path}.snapshot/tf-fss-snapshot-`date -u +\%Y-\%m-\%dT\%H`/ ${dst_mount_path} 
      