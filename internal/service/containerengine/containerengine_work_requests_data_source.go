// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"
)

func ContainerengineWorkRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineWorkRequests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"work_requests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"operation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"action_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_accepted": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_finished": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readContainerengineWorkRequests(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineWorkRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineWorkRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListWorkRequestsResponse
}

func (s *ContainerengineWorkRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineWorkRequestsDataSourceCrud) Get() error {
	request := oci_containerengine.ListWorkRequestsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_containerengine.ListWorkRequestsResourceTypeEnum(resourceType.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.Status = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListWorkRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWorkRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineWorkRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineWorkRequestsDataSource-", ContainerengineWorkRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workRequest := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Id != nil {
			workRequest["id"] = *r.Id
		}

		workRequest["operation_type"] = r.OperationType

		_resources := []interface{}{}
		for _, item := range r.Resources {
			_resources = append(_resources, WorkRequestResourceToMap(item))
		}
		workRequest["resources"] = _resources

		workRequest["status"] = r.Status

		if r.TimeAccepted != nil {
			workRequest["time_accepted"] = r.TimeAccepted.String()
		}

		if r.TimeFinished != nil {
			workRequest["time_finished"] = r.TimeFinished.String()
		}

		if r.TimeStarted != nil {
			workRequest["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, workRequest)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineWorkRequestsDataSource().Schema["work_requests"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_requests", resources); err != nil {
		return err
	}

	return nil
}

func WorkRequestResourceToMap(obj oci_containerengine.WorkRequestResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_type"] = string(obj.ActionType)

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.EntityUri != nil {
		result["entity_uri"] = string(*obj.EntityUri)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}
