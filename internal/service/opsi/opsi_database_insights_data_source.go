// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
)

func OpsiDatabaseInsightsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiDatabaseInsights,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"database_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enterprise_manager_bridge_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exadata_insight_id": {
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
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_insights_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiDatabaseInsightResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiDatabaseInsights(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiDatabaseInsightsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListDatabaseInsightsResponse
}

func (s *OpsiDatabaseInsightsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiDatabaseInsightsDataSourceCrud) Get() error {
	request := oci_opsi.ListDatabaseInsightsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		interfaces := databaseId.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("database_id") {
			request.DatabaseId = tmp
		}
	}

	if databaseType, ok := s.D.GetOkExists("database_type"); ok {
		interfaces := databaseType.([]interface{})
		tmp := make([]oci_opsi.ListDatabaseInsightsDatabaseTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.ListDatabaseInsightsDatabaseTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("database_type") {
			request.DatabaseType = tmp
		}
	}

	if enterpriseManagerBridgeId, ok := s.D.GetOkExists("enterprise_manager_bridge_id"); ok {
		tmp := enterpriseManagerBridgeId.(string)
		request.EnterpriseManagerBridgeId = &tmp
	}

	if exadataInsightId, ok := s.D.GetOkExists("exadata_insight_id"); ok {
		tmp := exadataInsightId.(string)
		request.ExadataInsightId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]oci_opsi.ListDatabaseInsightsFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.ListDatabaseInsightsFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		request.Id = []string{id.(string)}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]oci_opsi.LifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.LifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]oci_opsi.ResourceStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.ResourceStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("status") {
			request.Status = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListDatabaseInsights(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseInsights(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiDatabaseInsightsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiDatabaseInsightsDataSource-", OpsiDatabaseInsightsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseInsight := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		result := DatabaseInsightSummaryToMap(item)
		if len(result) != 0 {
			items = append(items, result)
		}

	}
	databaseInsight["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiDatabaseInsightsDataSource().Schema["database_insights_collection"].Elem.(*schema.Resource).Schema)
		databaseInsight["items"] = items
	}

	resources = append(resources, databaseInsight)
	if err := s.D.Set("database_insights_collection", resources); err != nil {
		return err
	}

	return nil
}
