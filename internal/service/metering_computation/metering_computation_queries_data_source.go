// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationQueriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMeteringComputationQueries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MeteringComputationQueryResource()),
						},
					},
				},
			},
		},
	}
}

func readMeteringComputationQueries(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationQueriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationQueriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.ListQueriesResponse
}

func (s *MeteringComputationQueriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationQueriesDataSourceCrud) Get() error {
	request := oci_metering_computation.ListQueriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.ListQueries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListQueries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MeteringComputationQueriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MeteringComputationQueriesDataSource-", MeteringComputationQueriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	query := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, QuerySummaryToMap(item))
	}
	query["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MeteringComputationQueriesDataSource().Schema["query_collection"].Elem.(*schema.Resource).Schema)
		query["items"] = items
	}

	resources = append(resources, query)
	if err := s.D.Set("query_collection", resources); err != nil {
		return err
	}

	return nil
}
