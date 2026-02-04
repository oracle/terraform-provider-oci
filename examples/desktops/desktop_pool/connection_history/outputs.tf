// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Output - The desktops of the desktop pool
output "desktop_pool_desktops" {
  value = [data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops_datasource.desktop_pool_desktop_collection.0.items]
}

# Output - The desktops in a compartment
output "oci_desktops_desktops" {
  value = [data.oci_desktops_desktops.test_desktops_datasource.desktop_collection.0.items]
}

# Output - A specific desktop
output "oci_desktops_desktop" {
  value = [data.oci_desktops_desktop.test_desktop_datasource]
}
