// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMediaServicesMediaWorkflowConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_workflow_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MediaServicesMediaWorkflowConfigurationResource()),
						},
					},
				},
			},
		},
	}
}

func readMediaServicesMediaWorkflowConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaWorkflowConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.ListMediaWorkflowConfigurationsResponse
}

func (s *MediaServicesMediaWorkflowConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaWorkflowConfigurationsDataSourceCrud) Get() error {
	request := oci_media_services.ListMediaWorkflowConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_media_services.MediaWorkflowLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.ListMediaWorkflowConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMediaWorkflowConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MediaServicesMediaWorkflowConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesMediaWorkflowConfigurationsDataSource-", MediaServicesMediaWorkflowConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	mediaWorkflowConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MediaWorkflowConfigurationSummaryToMap(item))
	}
	mediaWorkflowConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MediaServicesMediaWorkflowConfigurationsDataSource().Schema["media_workflow_configuration_collection"].Elem.(*schema.Resource).Schema)
		mediaWorkflowConfiguration["items"] = items
	}

	resources = append(resources, mediaWorkflowConfiguration)
	if err := s.D.Set("media_workflow_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}
