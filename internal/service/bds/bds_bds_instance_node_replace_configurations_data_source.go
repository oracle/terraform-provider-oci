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

func BdsBdsInstanceNodeReplaceConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceNodeReplaceConfigurations,
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
			"node_replace_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceNodeReplaceConfigurationResource()),
			},
		},
	}
}

func readBdsBdsInstanceNodeReplaceConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeReplaceConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceNodeReplaceConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListNodeReplaceConfigurationsResponse
}

func (s *BdsBdsInstanceNodeReplaceConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceNodeReplaceConfigurationsDataSourceCrud) Get() error {
	request := oci_bds.ListNodeReplaceConfigurationsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.NodeReplaceConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListNodeReplaceConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNodeReplaceConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceNodeReplaceConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceNodeReplaceConfigurationsDataSource-", BdsBdsInstanceNodeReplaceConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceNodeReplaceConfiguration := map[string]interface{}{
			"bds_instance_id": *r.BdsInstanceId,
		}

		if r.DisplayName != nil {
			bdsInstanceNodeReplaceConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bdsInstanceNodeReplaceConfiguration["id"] = *r.Id
		}

		if r.LevelTypeDetails != nil {
			levelTypeDetailsArray := []interface{}{}
			if levelTypeDetailsMap := LevelTypeDetailsToMap(&r.LevelTypeDetails); levelTypeDetailsMap != nil {
				levelTypeDetailsArray = append(levelTypeDetailsArray, levelTypeDetailsMap)
			}
			bdsInstanceNodeReplaceConfiguration["level_type_details"] = levelTypeDetailsArray
		} else {
			bdsInstanceNodeReplaceConfiguration["level_type_details"] = nil
		}

		bdsInstanceNodeReplaceConfiguration["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceNodeReplaceConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			bdsInstanceNodeReplaceConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, bdsInstanceNodeReplaceConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceNodeReplaceConfigurationsDataSource().Schema["node_replace_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("node_replace_configurations", resources); err != nil {
		return err
	}

	return nil
}
