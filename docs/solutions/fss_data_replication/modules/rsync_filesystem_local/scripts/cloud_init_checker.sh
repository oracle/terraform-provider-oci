#!/bin/bash -e
echo "waiting up to 600 seconds for cloud-init to complete"
timeout 600 /bin/bash -c \
  'until stat /var/lib/cloud/instance/boot-finished 2>/dev/null; do echo waiting cloud-init to complete...; sleep 10; done'
echo "done"
