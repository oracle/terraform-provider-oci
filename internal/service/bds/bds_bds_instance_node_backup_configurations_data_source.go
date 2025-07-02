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

func BdsBdsInstanceNodeBackupConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceNodeBackupConfigurations,
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
			"node_backup_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceNodeBackupConfigurationResource()),
			},
		},
	}
}

func readBdsBdsInstanceNodeBackupConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceNodeBackupConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListNodeBackupConfigurationsResponse
}

func (s *BdsBdsInstanceNodeBackupConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceNodeBackupConfigurationsDataSourceCrud) Get() error {
	request := oci_bds.ListNodeBackupConfigurationsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.NodeBackupConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListNodeBackupConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNodeBackupConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceNodeBackupConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceNodeBackupConfigurationsDataSource-", BdsBdsInstanceNodeBackupConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceNodeBackupConfiguration := map[string]interface{}{
			"bds_instance_id": *r.BdsInstanceId,
		}

		if r.DisplayName != nil {
			bdsInstanceNodeBackupConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bdsInstanceNodeBackupConfiguration["id"] = *r.Id
		}

		if r.LevelTypeDetails != nil {
			levelTypeDetailsArray := []interface{}{}
			if levelTypeDetailsMap := LevelTypeDetailsToMap(&r.LevelTypeDetails); levelTypeDetailsMap != nil {
				levelTypeDetailsArray = append(levelTypeDetailsArray, levelTypeDetailsMap)
			}
			bdsInstanceNodeBackupConfiguration["level_type_details"] = levelTypeDetailsArray
		} else {
			bdsInstanceNodeBackupConfiguration["level_type_details"] = nil
		}

		bdsInstanceNodeBackupConfiguration["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceNodeBackupConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			bdsInstanceNodeBackupConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, bdsInstanceNodeBackupConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceNodeBackupConfigurationsDataSource().Schema["node_backup_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("node_backup_configurations", resources); err != nil {
		return err
	}

	return nil
}
