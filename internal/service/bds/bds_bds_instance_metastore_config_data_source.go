// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceMetastoreConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["metastore_config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceMetastoreConfigResource(), fieldMap, readSingularBdsBdsInstanceMetastoreConfig)
}

func readSingularBdsBdsInstanceMetastoreConfig(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceMetastoreConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceMetastoreConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetBdsMetastoreConfigurationResponse
}

func (s *BdsBdsInstanceMetastoreConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceMetastoreConfigDataSourceCrud) Get() error {
	request := oci_bds.GetBdsMetastoreConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if metastoreConfigId, ok := s.D.GetOkExists("metastore_config_id"); ok {
		tmp := metastoreConfigId.(string)
		request.MetastoreConfigId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetBdsMetastoreConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceMetastoreConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BdsApiKeyId != nil {
		s.D.Set("bds_api_key_id", *s.Res.BdsApiKeyId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	s.D.Set("metastore_type", s.Res.MetastoreType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
