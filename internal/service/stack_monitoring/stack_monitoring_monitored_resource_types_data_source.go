// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourceTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringMonitoredResourceTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"exclude_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_exclude_system_types": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"metric_namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitored_resource_types_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringMonitoredResourceTypeResource()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringMonitoredResourceTypes(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoredResourceTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListMonitoredResourceTypesResponse
}

func (s *StackMonitoringMonitoredResourceTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoredResourceTypesDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListMonitoredResourceTypesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if excludeFields, ok := s.D.GetOkExists("exclude_fields"); ok {
		interfaces := excludeFields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("exclude_fields") {
			request.ExcludeFields = tmp
		}
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if isExcludeSystemTypes, ok := s.D.GetOkExists("is_exclude_system_types"); ok {
		tmp := isExcludeSystemTypes.(bool)
		request.IsExcludeSystemTypes = &tmp
	}

	if metricNamespace, ok := s.D.GetOkExists("metric_namespace"); ok {
		tmp := metricNamespace.(string)
		request.MetricNamespace = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceCategory, ok := s.D.GetOkExists("resource_category"); ok {
		request.ResourceCategory = oci_stack_monitoring.ListMonitoredResourceTypesResourceCategoryEnum(resourceCategory.(string))
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		request.SourceType = oci_stack_monitoring.ListMonitoredResourceTypesSourceTypeEnum(sourceType.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_stack_monitoring.ListMonitoredResourceTypesStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListMonitoredResourceTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitoredResourceTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringMonitoredResourceTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringMonitoredResourceTypesDataSource-", StackMonitoringMonitoredResourceTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoredResourceType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredResourceTypeSummaryToMap(item))
	}
	monitoredResourceType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringMonitoredResourceTypesDataSource().Schema["monitored_resource_types_collection"].Elem.(*schema.Resource).Schema)
		monitoredResourceType["items"] = items
	}

	resources = append(resources, monitoredResourceType)
	if err := s.D.Set("monitored_resource_types_collection", resources); err != nil {
		return err
	}

	return nil
}
