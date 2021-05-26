// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v41/usageapi"
)

func init() {
	RegisterDatasource("oci_metering_computation_queries", MeteringComputationQueriesDataSource())
}

func MeteringComputationQueriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMeteringComputationQueries,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
							Elem:     GetDataSourceItemSchema(MeteringComputationQueryResource()),
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
	sync.Client = m.(*OracleClients).usageapiClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "metering_computation")

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

	s.D.SetId(GenerateDataSourceHashID("MeteringComputationQueriesDataSource-", MeteringComputationQueriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	query := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, QuerySummaryToMap(item))
	}
	query["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, MeteringComputationQueriesDataSource().Schema["query_collection"].Elem.(*schema.Resource).Schema)
		query["items"] = items
	}

	resources = append(resources, query)
	if err := s.D.Set("query_collection", resources); err != nil {
		return err
	}

	return nil
}
