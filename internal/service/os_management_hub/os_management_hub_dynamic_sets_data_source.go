// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubDynamicSetsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readOsManagementHubDynamicSetsWithContext,
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
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_set_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubDynamicSetResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubDynamicSetsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OsManagementHubDynamicSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.DynamicSetClient
	Res    *oci_os_management_hub.ListDynamicSetsResponse
}

func (s *OsManagementHubDynamicSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubDynamicSetsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_os_management_hub.ListDynamicSetsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if dynamicSetId, ok := s.D.GetOkExists("id"); ok {
		tmp := dynamicSetId.(string)
		request.DynamicSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListDynamicSets(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDynamicSets(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubDynamicSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubDynamicSetsDataSource-", OsManagementHubDynamicSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dynamicSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DynamicSetSummaryToMap(item))
	}
	dynamicSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubDynamicSetsDataSource().Schema["dynamic_set_collection"].Elem.(*schema.Resource).Schema)
		dynamicSet["items"] = items
	}

	resources = append(resources, dynamicSet)
	if err := s.D.Set("dynamic_set_collection", resources); err != nil {
		return err
	}

	return nil
}
