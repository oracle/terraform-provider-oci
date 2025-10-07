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

func IotDigitalTwinAdaptersDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readIotDigitalTwinAdaptersWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"digital_twin_model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"digital_twin_model_spec_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"digital_twin_adapter_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(IotDigitalTwinAdapterResource()),
						},
					},
				},
			},
		},
	}
}

func readIotDigitalTwinAdaptersWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinAdaptersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinAdaptersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.ListDigitalTwinAdaptersResponse
}

func (s *IotDigitalTwinAdaptersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinAdaptersDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.ListDigitalTwinAdaptersRequest{}

	if digitalTwinModelId, ok := s.D.GetOkExists("digital_twin_model_id"); ok {
		tmp := digitalTwinModelId.(string)
		request.DigitalTwinModelId = &tmp
	}

	if digitalTwinModelSpecUri, ok := s.D.GetOkExists("digital_twin_model_spec_uri"); ok {
		tmp := digitalTwinModelSpecUri.(string)
		request.DigitalTwinModelSpecUri = &tmp
	}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_iot.ListDigitalTwinAdaptersLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.ListDigitalTwinAdapters(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDigitalTwinAdapters(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IotDigitalTwinAdaptersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotDigitalTwinAdaptersDataSource-", IotDigitalTwinAdaptersDataSource(), s.D))
	resources := []map[string]interface{}{}
	digitalTwinAdapter := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DigitalTwinAdapterSummaryToMap(item))
	}
	digitalTwinAdapter["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, IotDigitalTwinAdaptersDataSource().Schema["digital_twin_adapter_collection"].Elem.(*schema.Resource).Schema)
		digitalTwinAdapter["items"] = items
	}

	resources = append(resources, digitalTwinAdapter)
	if err := s.D.Set("digital_twin_adapter_collection", resources); err != nil {
		return err
	}

	return nil
}
