// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_asset_distribution_channel_attachment_channel_id" {
  default = "//"
}

data "oci_media_services_media_asset_distribution_channel_attachment" "test_media_asset_distribution_channel_attachment" {
  #Required
  media_asset_id          = oci_media_services_media_asset.test_media_asset.id
  distribution_channel_id = var.media_asset_distribution_channel_attachment_channel_id
}

