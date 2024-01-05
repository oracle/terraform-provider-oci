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

func BdsBdsInstanceMetastoreConfigsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceMetastoreConfigs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_api_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metastore_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metastore_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bds_metastore_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceMetastoreConfigResource()),
			},
		},
	}
}

func readBdsBdsInstanceMetastoreConfigs(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceMetastoreConfigsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceMetastoreConfigsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsMetastoreConfigurationsResponse
}

func (s *BdsBdsInstanceMetastoreConfigsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceMetastoreConfigsDataSourceCrud) Get() error {
	request := oci_bds.ListBdsMetastoreConfigurationsRequest{}

	if bdsApiKeyId, ok := s.D.GetOkExists("bds_api_key_id"); ok {
		tmp := bdsApiKeyId.(string)
		request.BdsApiKeyId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if metastoreId, ok := s.D.GetOkExists("metastore_id"); ok {
		tmp := metastoreId.(string)
		request.MetastoreId = &tmp
	}

	if metastoreType, ok := s.D.GetOkExists("metastore_type"); ok {
		request.MetastoreType = oci_bds.BdsMetastoreConfigurationMetastoreTypeEnum(metastoreType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.BdsMetastoreConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsMetastoreConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsMetastoreConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceMetastoreConfigsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceMetastoreConfigsDataSource-", BdsBdsInstanceMetastoreConfigsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceMetastoreConfig := map[string]interface{}{}

		if r.BdsApiKeyId != nil {
			bdsInstanceMetastoreConfig["bds_api_key_id"] = *r.BdsApiKeyId
		}

		if r.DisplayName != nil {
			bdsInstanceMetastoreConfig["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bdsInstanceMetastoreConfig["id"] = *r.Id
		}

		if r.MetastoreId != nil {
			bdsInstanceMetastoreConfig["metastore_id"] = *r.MetastoreId
		}

		bdsInstanceMetastoreConfig["metastore_type"] = r.MetastoreType

		bdsInstanceMetastoreConfig["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceMetastoreConfig["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			bdsInstanceMetastoreConfig["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, bdsInstanceMetastoreConfig)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceMetastoreConfigsDataSource().Schema["bds_metastore_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bds_metastore_configurations", resources); err != nil {
		return err
	}

	return nil
}
