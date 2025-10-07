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

func IotDigitalTwinRelationshipsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readIotDigitalTwinRelationshipsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"content_path": {
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
			"source_digital_twin_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_digital_twin_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"digital_twin_relationship_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(IotDigitalTwinRelationshipResource()),
						},
					},
				},
			},
		},
	}
}

func readIotDigitalTwinRelationshipsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinRelationshipsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinRelationshipsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.ListDigitalTwinRelationshipsResponse
}

func (s *IotDigitalTwinRelationshipsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinRelationshipsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.ListDigitalTwinRelationshipsRequest{}

	if contentPath, ok := s.D.GetOkExists("content_path"); ok {
		tmp := contentPath.(string)
		request.ContentPath = &tmp
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

	if sourceDigitalTwinInstanceId, ok := s.D.GetOkExists("source_digital_twin_instance_id"); ok {
		tmp := sourceDigitalTwinInstanceId.(string)
		request.SourceDigitalTwinInstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_iot.ListDigitalTwinRelationshipsLifecycleStateEnum(state.(string))
	}

	if targetDigitalTwinInstanceId, ok := s.D.GetOkExists("target_digital_twin_instance_id"); ok {
		tmp := targetDigitalTwinInstanceId.(string)
		request.TargetDigitalTwinInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.ListDigitalTwinRelationships(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDigitalTwinRelationships(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IotDigitalTwinRelationshipsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotDigitalTwinRelationshipsDataSource-", IotDigitalTwinRelationshipsDataSource(), s.D))
	resources := []map[string]interface{}{}
	digitalTwinRelationship := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DigitalTwinRelationshipSummaryToMap(item))
	}
	digitalTwinRelationship["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, IotDigitalTwinRelationshipsDataSource().Schema["digital_twin_relationship_collection"].Elem.(*schema.Resource).Schema)
		digitalTwinRelationship["items"] = items
	}

	resources = append(resources, digitalTwinRelationship)
	if err := s.D.Set("digital_twin_relationship_collection", resources); err != nil {
		return err
	}

	return nil
}
