// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v65/blockchain"
)

func BlockchainOsnDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["blockchain_platform_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["osn_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BlockchainOsnResource(), fieldMap, readSingularBlockchainOsnWithContext)
}

func readSingularBlockchainOsnWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BlockchainOsnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BlockchainOsnDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.GetOsnResponse
}

func (s *BlockchainOsnDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainOsnDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_blockchain.GetOsnRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	if osnId, ok := s.D.GetOkExists("osn_id"); ok {
		tmp := osnId.(string)
		request.OsnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "blockchain")

	response, err := s.Client.GetOsn(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BlockchainOsnDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BlockchainOsnDataSource-", BlockchainOsnDataSource(), s.D))

	s.D.Set("ad", s.Res.Ad)

	if s.Res.OcpuAllocationParam != nil {
		s.D.Set("ocpu_allocation_param", []interface{}{OcpuAllocationNumberParamToMap(s.Res.OcpuAllocationParam)})
	} else {
		s.D.Set("ocpu_allocation_param", nil)
	}

	if s.Res.OsnKey != nil {
		s.D.Set("osn_key", *s.Res.OsnKey)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
