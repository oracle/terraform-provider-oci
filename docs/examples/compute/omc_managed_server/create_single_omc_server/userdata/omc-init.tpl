#cloud-config
# vim: syntax=yaml
#

groups:
  - oinstall

# Add users to the system. Users are added after groups are added.
users:
  - default
  - name: oracle
    gecos: Oracle Installer
    sudo: ALL=(ALL) NOPASSWD:ALL
    primary-group: oinstall
    groups: users
    expiredate: 2012-09-01
    lock_passwd: true
    ssh-authorized-keys:
        - ${ssh_public_key}

package_upgrade: true

packages:
 - gcc
 - gcc-c++
 - wget
 - git
 - perl
 - unzip
 - bind-utils
 - bc
 - rng-tools
 - libffi-devel
 - python-devel
 - openssl-devel
 - xfsprogs
 - xfsdump
 - mdadm
 - jq

runcmd:
  - [ sh, -xc, "systemctl stop firewalld && systemctl disable firewalld" ]
  - [ sh, -xc, "echo ulimit -S -n 10000 >> ~oracle/.bash_profile" ]
  - [ sh, -xc, "echo umask 022 >> ~oracle/.bash_profile" ]
  - [ sh, -xc, "echo oracle soft nofile 65536 >> /etc/security/limits.conf" ]
  - [ sh, -xc, "echo oracle hard nofile 65536 >> /etc/security/limits.conf" ]
  - [ sh, -xc, "mkdir -p /omc"]
  - [ sh, -xc, "mkdir -p /omc/install"]
  - [ sh, -xc, "mkdir -p /omc/stage"]
  - [ sh, -xc, "mkdir -p /omc/app"]
  - [ sh, -xc, "mkdir -p /omc/apm"]
  - [ sh, -xc, "chown -R oracle:oinstall /omc"]
  - [ sh, -xc, "rngd -r /dev/urandom -o /dev/random"]
  - [ sh, -xc, "touch /tmp/signal"]

