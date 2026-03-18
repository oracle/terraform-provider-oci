// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabasePrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDistributedDatabaseDistributedDatabasePrivateEndpointsWithContext,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"distributed_database_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DistributedDatabaseDistributedDatabasePrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readDistributedDatabaseDistributedDatabasePrivateEndpointsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasePrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbPrivateEndpointServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedDatabasePrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedDbPrivateEndpointServiceClient
	Res    *oci_distributed_database.ListDistributedDatabasePrivateEndpointsResponse
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.ListDistributedDatabasePrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.ListDistributedDatabasePrivateEndpoints(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDistributedDatabasePrivateEndpoints(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DistributedDatabaseDistributedDatabasePrivateEndpointsDataSource-", DistributedDatabaseDistributedDatabasePrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	distributedDatabasePrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DistributedDatabasePrivateEndpointSummaryToMap(item, true))
	}
	distributedDatabasePrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DistributedDatabaseDistributedDatabasePrivateEndpointsDataSource().Schema["distributed_database_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		distributedDatabasePrivateEndpoint["items"] = items
	}

	resources = append(resources, distributedDatabasePrivateEndpoint)
	if err := s.D.Set("distributed_database_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
