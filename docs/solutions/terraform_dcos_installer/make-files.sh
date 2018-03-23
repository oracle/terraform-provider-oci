#!/usr/bin/env bash
. ./ips.txt
# Make some config files
cat > config.yaml << FIN
bootstrap_url: http://$BOOTSTRAP:4040
cluster_name: $CLUSTER_NAME
exhibitor_storage_backend: zookeeper
exhibitor_zk_hosts: $BOOTSTRAP:2181
exhibitor_zk_path: /$CLUSTER_NAME
log_directory: /genconf/logs
master_discovery: static
master_list:
- $MASTER_00
- $MASTER_01
- $MASTER_02
- $MASTER_03
- $MASTER_04
resolvers: 
- 8.8.4.4
- 8.8.8.8
FIN


# Make a script

cat > do-install.sh << FIN
#!/usr/bin/env bash
mkdir /tmp/dcos && cd /tmp/dcos
printf "Waiting for installer to appear at Bootstrap URL"
until \$(curl -m 2 --connect-timeout 2 --output /dev/null --silent --head --fail http://$BOOTSTRAP:4040/dcos_install.sh); do
    sleep 1
done           
curl -O http://$BOOTSTRAP:4040/dcos_install.sh
sudo bash dcos_install.sh \$1
FIN
rm -rf ./ips.txt
