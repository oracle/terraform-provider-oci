// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v46/usageapi"
)

func init() {
	RegisterDatasource("oci_metering_computation_custom_tables", MeteringComputationCustomTablesDataSource())
}

func MeteringComputationCustomTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMeteringComputationCustomTables,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"saved_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_table_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(MeteringComputationCustomTableResource()),
						},
					},
				},
			},
		},
	}
}

func readMeteringComputationCustomTables(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCustomTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).usageapiClient()

	return ReadResource(sync)
}

type MeteringComputationCustomTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.ListCustomTablesResponse
}

func (s *MeteringComputationCustomTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationCustomTablesDataSourceCrud) Get() error {
	request := oci_metering_computation.ListCustomTablesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if savedReportId, ok := s.D.GetOkExists("saved_report_id"); ok {
		tmp := savedReportId.(string)
		request.SavedReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "metering_computation")

	response, err := s.Client.ListCustomTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCustomTables(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MeteringComputationCustomTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MeteringComputationCustomTablesDataSource-", MeteringComputationCustomTablesDataSource(), s.D))
	resources := []map[string]interface{}{}
	customTable := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CustomTableSummaryToMap(item))
	}
	customTable["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, MeteringComputationCustomTablesDataSource().Schema["custom_table_collection"].Elem.(*schema.Resource).Schema)
		customTable["items"] = items
	}

	resources = append(resources, customTable)
	if err := s.D.Set("custom_table_collection", resources); err != nil {
		return err
	}

	return nil
}
