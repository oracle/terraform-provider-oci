// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v58/bds"
)

func BdsBdsInstanceApiKeyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["api_key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceApiKeyResource(), fieldMap, readSingularBdsBdsInstanceApiKey)
}

func readSingularBdsBdsInstanceApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceApiKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceApiKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetBdsApiKeyResponse
}

func (s *BdsBdsInstanceApiKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceApiKeyDataSourceCrud) Get() error {
	request := oci_bds.GetBdsApiKeyRequest{}

	if apiKeyId, ok := s.D.GetOkExists("api_key_id"); ok {
		tmp := apiKeyId.(string)
		request.ApiKeyId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetBdsApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceApiKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefaultRegion != nil {
		s.D.Set("default_region", *s.Res.DefaultRegion)
	}

	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	if s.Res.KeyAlias != nil {
		s.D.Set("key_alias", *s.Res.KeyAlias)
	}

	if s.Res.Pemfilepath != nil {
		s.D.Set("pemfilepath", *s.Res.Pemfilepath)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
