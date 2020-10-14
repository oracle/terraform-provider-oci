// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v27/blockchain"
)

func init() {
	RegisterDatasource("oci_blockchain_peer", BlockchainPeerDataSource())
}

func BlockchainPeerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["blockchain_platform_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["peer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(BlockchainPeerResource(), fieldMap, readSingularBlockchainPeer)
}

func readSingularBlockchainPeer(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()

	return ReadResource(sync)
}

type BlockchainPeerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.GetPeerResponse
}

func (s *BlockchainPeerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainPeerDataSourceCrud) Get() error {
	request := oci_blockchain.GetPeerRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	if peerId, ok := s.D.GetOkExists("peer_id"); ok {
		tmp := peerId.(string)
		request.PeerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "blockchain")

	response, err := s.Client.GetPeer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BlockchainPeerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("ad", s.Res.Ad)

	if s.Res.Alias != nil {
		s.D.Set("alias", *s.Res.Alias)
	}

	if s.Res.Host != nil {
		s.D.Set("host", *s.Res.Host)
	}

	if s.Res.OcpuAllocationParam != nil {
		s.D.Set("ocpu_allocation_param", []interface{}{OcpuAllocationNumberParamToMap(s.Res.OcpuAllocationParam)})
	} else {
		s.D.Set("ocpu_allocation_param", nil)
	}

	if s.Res.PeerKey != nil {
		s.D.Set("peer_key", *s.Res.PeerKey)
	}

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
