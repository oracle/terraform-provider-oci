// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiHostedApplicationStoragesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGenerativeAiHostedApplicationStoragesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosted_application_storage_type": {
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
			"hosted_application_storage_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiHostedApplicationStorageResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiHostedApplicationStoragesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationStoragesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiHostedApplicationStoragesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.ListHostedApplicationStoragesResponse
}

func (s *GenerativeAiHostedApplicationStoragesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiHostedApplicationStoragesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.ListHostedApplicationStoragesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if hostedApplicationStorageType, ok := s.D.GetOkExists("hosted_application_storage_type"); ok {
		request.HostedApplicationStorageType = oci_generative_ai.HostedApplicationStorageStorageTypeEnum(hostedApplicationStorageType.(string))
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_generative_ai.HostedApplicationStorageLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.ListHostedApplicationStorages(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHostedApplicationStorages(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiHostedApplicationStoragesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiHostedApplicationStoragesDataSource-", GenerativeAiHostedApplicationStoragesDataSource(), s.D))
	resources := []map[string]interface{}{}
	hostedApplicationStorage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, HostedApplicationStorageSummaryToMap(item))
	}
	hostedApplicationStorage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiHostedApplicationStoragesDataSource().Schema["hosted_application_storage_collection"].Elem.(*schema.Resource).Schema)
		hostedApplicationStorage["items"] = items
	}

	resources = append(resources, hostedApplicationStorage)
	if err := s.D.Set("hosted_application_storage_collection", resources); err != nil {
		return err
	}

	return nil
}
