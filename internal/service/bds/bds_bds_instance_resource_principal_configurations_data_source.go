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

func BdsBdsInstanceResourcePrincipalConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceResourcePrincipalConfigurations,
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
			"resource_principal_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceResourcePrincipalConfigurationResource()),
			},
		},
	}
}

func readBdsBdsInstanceResourcePrincipalConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourcePrincipalConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceResourcePrincipalConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListResourcePrincipalConfigurationsResponse
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationsDataSourceCrud) Get() error {
	request := oci_bds.ListResourcePrincipalConfigurationsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.ResourcePrincipalConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListResourcePrincipalConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourcePrincipalConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceResourcePrincipalConfigurationsDataSource-", BdsBdsInstanceResourcePrincipalConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceResourcePrincipalConfiguration := map[string]interface{}{
			"bds_instance_id": *r.BdsInstanceId,
		}

		if r.DisplayName != nil {
			bdsInstanceResourcePrincipalConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bdsInstanceResourcePrincipalConfiguration["id"] = *r.Id
		}

		bdsInstanceResourcePrincipalConfiguration["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceResourcePrincipalConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeTokenExpiry != nil {
			bdsInstanceResourcePrincipalConfiguration["time_token_expiry"] = r.TimeTokenExpiry.String()
		}

		if r.TimeTokenRefreshed != nil {
			bdsInstanceResourcePrincipalConfiguration["time_token_refreshed"] = r.TimeTokenRefreshed.String()
		}

		if r.TimeUpdated != nil {
			bdsInstanceResourcePrincipalConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, bdsInstanceResourcePrincipalConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceResourcePrincipalConfigurationsDataSource().Schema["resource_principal_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("resource_principal_configurations", resources); err != nil {
		return err
	}

	return nil
}
