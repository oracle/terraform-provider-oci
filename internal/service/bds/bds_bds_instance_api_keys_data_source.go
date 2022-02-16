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

func BdsBdsInstanceApiKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceApiKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bds_api_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceApiKeyResource()),
			},
		},
	}
}

func readBdsBdsInstanceApiKeys(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceApiKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceApiKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsApiKeysResponse
}

func (s *BdsBdsInstanceApiKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceApiKeysDataSourceCrud) Get() error {
	request := oci_bds.ListBdsApiKeysRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.BdsApiKeyLifecycleStateEnum(state.(string))
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsApiKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsApiKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceApiKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceApiKeysDataSource-", BdsBdsInstanceApiKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceApiKey := map[string]interface{}{}

		if r.DefaultRegion != nil {
			bdsInstanceApiKey["default_region"] = *r.DefaultRegion
		}

		if r.Id != nil {
			bdsInstanceApiKey["id"] = *r.Id
		}

		if r.KeyAlias != nil {
			bdsInstanceApiKey["key_alias"] = *r.KeyAlias
		}

		bdsInstanceApiKey["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceApiKey["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, bdsInstanceApiKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceApiKeysDataSource().Schema["bds_api_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bds_api_keys", resources); err != nil {
		return err
	}

	return nil
}
