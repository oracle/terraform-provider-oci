// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeReportDefinitionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeReportDefinitions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data_source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_seeded": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_definition_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeReportDefinitionResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeReportDefinitions(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDefinitionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeReportDefinitionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListReportDefinitionsResponse
}

func (s *DataSafeReportDefinitionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeReportDefinitionsDataSourceCrud) Get() error {
	request := oci_data_safe.ListReportDefinitionsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListReportDefinitionsAccessLevelEnum(accessLevel.(string))
	}

	if category, ok := s.D.GetOkExists("category"); ok {
		request.Category = oci_data_safe.ListReportDefinitionsCategoryEnum(category.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if dataSource, ok := s.D.GetOkExists("data_source"); ok {
		request.DataSource = oci_data_safe.ListReportDefinitionsDataSourceEnum(dataSource.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isSeeded, ok := s.D.GetOkExists("is_seeded"); ok {
		tmp := isSeeded.(bool)
		request.IsSeeded = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListReportDefinitionsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListReportDefinitions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReportDefinitions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeReportDefinitionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeReportDefinitionsDataSource-", DataSafeReportDefinitionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	reportDefinition := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ReportDefinitionSummaryToMap(item))
	}
	reportDefinition["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeReportDefinitionsDataSource().Schema["report_definition_collection"].Elem.(*schema.Resource).Schema)
		reportDefinition["items"] = items
	}

	resources = append(resources, reportDefinition)
	if err := s.D.Set("report_definition_collection", resources); err != nil {
		return err
	}

	return nil
}
