// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneSubscriptionMappingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneSubscriptionMappings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_mapping_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_mapping_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(TenantmanagercontrolplaneSubscriptionMappingResource()),
						},
					},
				},
			},
		},
	}
}

func readTenantmanagercontrolplaneSubscriptionMappings(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSubscriptionMappingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneSubscriptionMappingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SubscriptionClient
	Res    *oci_tenantmanagercontrolplane.ListSubscriptionMappingsResponse
}

func (s *TenantmanagercontrolplaneSubscriptionMappingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneSubscriptionMappingsDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListSubscriptionMappingsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_tenantmanagercontrolplane.SubscriptionMappingLifecycleStateEnum(state.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionMappingId, ok := s.D.GetOkExists("id"); ok {
		tmp := subscriptionMappingId.(string)
		request.SubscriptionMappingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListSubscriptionMappings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubscriptionMappings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneSubscriptionMappingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneSubscriptionMappingsDataSource-", TenantmanagercontrolplaneSubscriptionMappingsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subscriptionMapping := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubscriptionMappingSummaryToMap(item))
	}
	subscriptionMapping["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneSubscriptionMappingsDataSource().Schema["subscription_mapping_collection"].Elem.(*schema.Resource).Schema)
		subscriptionMapping["items"] = items
	}

	resources = append(resources, subscriptionMapping)
	if err := s.D.Set("subscription_mapping_collection", resources); err != nil {
		return err
	}

	return nil
}
