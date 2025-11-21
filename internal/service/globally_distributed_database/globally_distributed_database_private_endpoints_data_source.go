// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globally_distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GloballyDistributedDatabasePrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGloballyDistributedDatabasePrivateEndpoints,
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
			"private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GloballyDistributedDatabasePrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readGloballyDistributedDatabasePrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabasePrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.ReadResource(sync)
}

type GloballyDistributedDatabasePrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_globally_distributed_database.ShardedDatabaseServiceClient
	Res    *oci_globally_distributed_database.ListPrivateEndpointsResponse
}

func (s *GloballyDistributedDatabasePrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GloballyDistributedDatabasePrivateEndpointsDataSourceCrud) Get() error {
	request := oci_globally_distributed_database.ListPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_globally_distributed_database.PrivateEndpointLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "globally_distributed_database")

	response, err := s.Client.ListPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GloballyDistributedDatabasePrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GloballyDistributedDatabasePrivateEndpointsDataSource-", GloballyDistributedDatabasePrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	privateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivateEndpointSummaryToMap(item, true))
	}
	privateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GloballyDistributedDatabasePrivateEndpointsDataSource().Schema["private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		privateEndpoint["items"] = items
	}

	resources = append(resources, privateEndpoint)
	if err := s.D.Set("private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
