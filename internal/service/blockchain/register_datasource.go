// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_blockchain_blockchain_platform", BlockchainBlockchainPlatformDataSource())
	tfresource.RegisterDatasource("oci_blockchain_blockchain_platform_patches", BlockchainBlockchainPlatformPatchesDataSource())
	tfresource.RegisterDatasource("oci_blockchain_blockchain_platforms", BlockchainBlockchainPlatformsDataSource())
	tfresource.RegisterDatasource("oci_blockchain_osn", BlockchainOsnDataSource())
	tfresource.RegisterDatasource("oci_blockchain_osns", BlockchainOsnsDataSource())
	tfresource.RegisterDatasource("oci_blockchain_peer", BlockchainPeerDataSource())
	tfresource.RegisterDatasource("oci_blockchain_peers", BlockchainPeersDataSource())
}
