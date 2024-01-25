// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsLogAnalyticsEntitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsLogAnalyticsEntities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_type_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_management_agent_id_null": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_details_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata_equals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_analytics_entity_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LogAnalyticsLogAnalyticsEntityResource()),
						},
					},
				},
			},
		},
	}
}

func readLogAnalyticsLogAnalyticsEntities(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsEntitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListLogAnalyticsEntitiesResponse
}

func (s *LogAnalyticsLogAnalyticsEntitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsEntitiesDataSourceCrud) Get() error {
	request := oci_log_analytics.ListLogAnalyticsEntitiesRequest{}

	if cloudResourceId, ok := s.D.GetOkExists("cloud_resource_id"); ok {
		tmp := cloudResourceId.(string)
		request.CloudResourceId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if entityTypeName, ok := s.D.GetOkExists("entity_type_name"); ok {
		interfaces := entityTypeName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("states") {
			request.EntityTypeName = tmp
		}
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if hostnameContains, ok := s.D.GetOkExists("hostname_contains"); ok {
		tmp := hostnameContains.(string)
		request.HostnameContains = &tmp
	}

	if isManagementAgentIdNull, ok := s.D.GetOkExists("is_management_agent_id_null"); ok {
		request.IsManagementAgentIdNull = oci_log_analytics.ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum(isManagementAgentIdNull.(string))
	}

	if lifecycleDetailsContains, ok := s.D.GetOkExists("lifecycle_details_contains"); ok {
		tmp := lifecycleDetailsContains.(string)
		request.LifecycleDetailsContains = &tmp
	}

	if metadataEquals, ok := s.D.GetOkExists("metadata_equals"); ok {
		interfaces := metadataEquals.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("metadata_equals") {
			request.MetadataEquals = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if sourceId, ok := s.D.GetOkExists("source_id"); ok {
		tmp := sourceId.(string)
		request.SourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_log_analytics.ListLogAnalyticsEntitiesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListLogAnalyticsEntities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogAnalyticsEntities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsEntitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsEntitiesDataSource-", LogAnalyticsLogAnalyticsEntitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	logAnalyticsEntity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsEntitySummaryToMap(item))
	}
	logAnalyticsEntity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsLogAnalyticsEntitiesDataSource().Schema["log_analytics_entity_collection"].Elem.(*schema.Resource).Schema)
		logAnalyticsEntity["items"] = items
	}

	resources = append(resources, logAnalyticsEntity)
	if err := s.D.Set("log_analytics_entity_collection", resources); err != nil {
		return err
	}

	return nil
}
