// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

output "desktop_pool_id" {
  value = [data.oci_desktops_desktop_pools.test_desktop_pools_datasource.desktop_pool_collection.0.items.*.id]
}

# Output the desktop IDs of the desktop pool
output "desktop_pool_desktop_ids" {
  value = [data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops_datasource.desktop_pool_desktop_collection.0.items.*.desktop_id]
}

# Output the volume IDs of the desktop pool
output "desktop_pool_volume_ids" {
  value = [data.oci_desktops_desktop_pool_volumes.test_desktop_pool_volumes_datasource.desktop_pool_volume_collection.0.items.*.id]
}

# Output the desktop IDs of the desktop pool
output "desktop_ids" {
  value = [data.oci_desktops_desktops.test_desktops_datasource.desktop_collection.0.items.*.id]
}
