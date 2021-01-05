// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v31/bds"
)

func init() {
	RegisterDatasource("oci_bds_auto_scaling_configurations", BdsAutoScalingConfigurationsDataSource())
}

func BdsAutoScalingConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsAutoScalingConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
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
			"auto_scaling_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(BdsAutoScalingConfigurationResource()),
			},
		},
	}
}

func readBdsAutoScalingConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &BdsAutoScalingConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).bdsClient()

	return ReadResource(sync)
}

type BdsAutoScalingConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListAutoScalingConfigurationsResponse
}

func (s *BdsAutoScalingConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsAutoScalingConfigurationsDataSourceCrud) Get() error {
	request := oci_bds.ListAutoScalingConfigurationsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "bds")

	response, err := s.Client.ListAutoScalingConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutoScalingConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsAutoScalingConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("BdsAutoScalingConfigurationsDataSource-", BdsAutoScalingConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autoScalingConfiguration := map[string]interface{}{}

		if r.DisplayName != nil {
			autoScalingConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			autoScalingConfiguration["id"] = *r.Id
		}

		if r.NodeType != "" {
			autoScalingConfiguration["node_type"] = r.NodeType
		}

		autoScalingConfiguration["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autoScalingConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			autoScalingConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, autoScalingConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, BdsAutoScalingConfigurationsDataSource().Schema["auto_scaling_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("auto_scaling_configurations", resources); err != nil {
		return err
	}

	return nil
}
