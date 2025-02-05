// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseScheduledActionParamsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseScheduledActionParams,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action_param_values_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"default_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"parameter_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter_values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseScheduledActionParams(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionParamsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseScheduledActionParamsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListParamsForActionTypeResponse
}

func (s *DatabaseScheduledActionParamsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseScheduledActionParamsDataSourceCrud) Get() error {
	request := oci_database.ListParamsForActionTypeRequest{}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_database.RecommendedScheduledActionSummaryActionTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListParamsForActionType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListParamsForActionType(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseScheduledActionParamsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseScheduledActionParamsDataSource-", DatabaseScheduledActionParamsDataSource(), s.D))
	resources := []map[string]interface{}{}
	scheduledActionParam := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ActionParamValuesSummaryToMap(item))
	}
	scheduledActionParam["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseScheduledActionParamsDataSource().Schema["action_param_values_collection"].Elem.(*schema.Resource).Schema)
		scheduledActionParam["items"] = items
	}

	resources = append(resources, scheduledActionParam)
	if err := s.D.Set("action_param_values_collection", resources); err != nil {
		return err
	}

	return nil
}

func ActionParamValuesSummaryToMap(obj oci_database.ActionParamValuesSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.ParameterName != nil {
		result["parameter_name"] = string(*obj.ParameterName)
	}

	result["parameter_type"] = string(obj.ParameterType)

	result["parameter_values"] = obj.ParameterValues

	return result
}
