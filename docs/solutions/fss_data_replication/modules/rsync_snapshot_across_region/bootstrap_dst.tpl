#cloud-config

packages:
  - rsync
  - nfs-utils  
runcmd:
  - sudo yum -y install rsync nfs-utils
  - sudo mkdir -p ${dst_mount_path}
  - sudo mount ${dst_mount_target_private_ip}:${dst_export_path} ${dst_mount_path}
  - sudo crontab /etc/cron.d/fss-sync-up-snapshot

write_files:     
  - path: /etc/cron.d/fss-sync-up-snapshot
    owner: root:root    
    permissions: 0644
    content: |      
      ${data_sync_frequency} /usr/bin/flock -n /var/run/fss-sync-up-snapshot.lck rsync -aHAXxve --numeric-ids --delete -e "ssh -i /home/opc/.ssh/id_rsa -o StrictHostKeyChecking=no" opc@${src_host}:${src_mount_path}.snapshot/tf-fss-snapshot-`date -u +\%Y-\%m-\%dT\%H`/ ${dst_mount_path} 

