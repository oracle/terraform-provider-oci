// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataccInfrastructuresDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataccInfrastructuresWithContext,
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
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"infrastructure_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataccInfrastructureResource()),
						},
					},
				},
			},
		},
	}
}

func readDataccInfrastructuresWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructuresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataccInfrastructuresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacc.BaseinfraClient
	Res    *oci_datacc.ListInfrastructuresResponse
}

func (s *DataccInfrastructuresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataccInfrastructuresDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.ListInfrastructuresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		stateArr := state.([]interface{})
		lifecycleStates := make([]oci_datacc.InfrastructureLifecycleStateEnum, len(stateArr))
		for i, s := range stateArr {
			lifecycleState := s.(string)
			if lifecycleStateEnum, ok := oci_datacc.GetMappingInfrastructureLifecycleStateEnum(lifecycleState); ok {
				lifecycleStates[i] = lifecycleStateEnum
			} else {
				return fmt.Errorf("Invalid Infrastructure lifecycle state %s", lifecycleState)
			}
		}
		request.LifecycleState = lifecycleStates
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacc")

	response, err := s.Client.ListInfrastructures(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInfrastructures(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataccInfrastructuresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataccInfrastructuresDataSource-", DataccInfrastructuresDataSource(), s.D))
	resources := []map[string]interface{}{}
	infrastructure := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InfrastructureSummaryToMap(item))
	}
	infrastructure["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataccInfrastructuresDataSource().Schema["infrastructure_collection"].Elem.(*schema.Resource).Schema)
		infrastructure["items"] = items
	}

	resources = append(resources, infrastructure)
	if err := s.D.Set("infrastructure_collection", resources); err != nil {
		return err
	}

	return nil
}
