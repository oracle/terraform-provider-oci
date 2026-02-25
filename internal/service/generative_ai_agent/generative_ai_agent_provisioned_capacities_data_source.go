// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentProvisionedCapacitiesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGenerativeAiAgentProvisionedCapacitiesWithContext,
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
			"provisioned_capacity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provisioned_capacity_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiAgentProvisionedCapacityResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiAgentProvisionedCapacitiesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiAgentProvisionedCapacitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiAgentProvisionedCapacitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.ListProvisionedCapacitiesResponse
}

func (s *GenerativeAiAgentProvisionedCapacitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentProvisionedCapacitiesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai_agent.ListProvisionedCapacitiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if provisionedCapacityId, ok := s.D.GetOkExists("id"); ok {
		tmp := provisionedCapacityId.(string)
		request.ProvisionedCapacityId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_generative_ai_agent.ProvisionedCapacityLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.ListProvisionedCapacities(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProvisionedCapacities(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiAgentProvisionedCapacitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiAgentProvisionedCapacitiesDataSource-", GenerativeAiAgentProvisionedCapacitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	provisionedCapacity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProvisionedCapacitySummaryToMap(item))
	}
	provisionedCapacity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiAgentProvisionedCapacitiesDataSource().Schema["provisioned_capacity_collection"].Elem.(*schema.Resource).Schema)
		provisionedCapacity["items"] = items
	}

	resources = append(resources, provisionedCapacity)
	if err := s.D.Set("provisioned_capacity_collection", resources); err != nil {
		return err
	}

	return nil
}
