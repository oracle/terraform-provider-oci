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

func BdsBdsInstanceIdentityConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceIdentityConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"identity_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceIdentityConfigurationResource()),
			},
		},
	}
}

func readBdsBdsInstanceIdentityConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceIdentityConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceIdentityConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListIdentityConfigurationsResponse
}

func (s *BdsBdsInstanceIdentityConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceIdentityConfigurationsDataSourceCrud) Get() error {
	request := oci_bds.ListIdentityConfigurationsRequest{}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.IdentityConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListIdentityConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIdentityConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceIdentityConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceIdentityConfigurationsDataSource-", BdsBdsInstanceIdentityConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceIdentityConfiguration := map[string]interface{}{}

		if r.DisplayName != nil {
			bdsInstanceIdentityConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bdsInstanceIdentityConfiguration["id"] = *r.Id
		}

		bdsInstanceIdentityConfiguration["state"] = r.LifecycleState

		resources = append(resources, bdsInstanceIdentityConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceIdentityConfigurationsDataSource().Schema["identity_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("identity_configurations", resources); err != nil {
		return err
	}

	return nil
}
