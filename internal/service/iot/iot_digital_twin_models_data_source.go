// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinModelsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readIotDigitalTwinModelsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"spec_uri_starts_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"digital_twin_model_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(IotDigitalTwinModelResource()),
						},
					},
				},
			},
		},
	}
}

func readIotDigitalTwinModelsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.ListDigitalTwinModelsResponse
}

func (s *IotDigitalTwinModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinModelsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.ListDigitalTwinModelsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	if specUriStartsWith, ok := s.D.GetOkExists("spec_uri_starts_with"); ok {
		tmp := specUriStartsWith.(string)
		request.SpecUriStartsWith = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_iot.ListDigitalTwinModelsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.ListDigitalTwinModels(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDigitalTwinModels(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IotDigitalTwinModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotDigitalTwinModelsDataSource-", IotDigitalTwinModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	digitalTwinModel := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DigitalTwinModelSummaryToMap(item))
	}
	digitalTwinModel["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, IotDigitalTwinModelsDataSource().Schema["digital_twin_model_collection"].Elem.(*schema.Resource).Schema)
		digitalTwinModel["items"] = items
	}

	resources = append(resources, digitalTwinModel)
	if err := s.D.Set("digital_twin_model_collection", resources); err != nil {
		return err
	}

	return nil
}
