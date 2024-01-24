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

func StackMonitoringDiscoveryJobsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringDiscoveryJobs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_job_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									// Required
									"resource_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"discovery_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tenant_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									// Optional
									"defined_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
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

func readStackMonitoringDiscoveryJobs(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDiscoveryJobsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringDiscoveryJobsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListDiscoveryJobsResponse
}

func (s *StackMonitoringDiscoveryJobsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringDiscoveryJobsDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListDiscoveryJobsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListDiscoveryJobs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDiscoveryJobs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringDiscoveryJobsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringDiscoveryJobsDataSource-", StackMonitoringDiscoveryJobsDataSource(), s.D))
	resources := []map[string]interface{}{}
	discoveryJob := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiscoveryJobSummaryToMap(item))
	}
	discoveryJob["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringDiscoveryJobsDataSource().Schema["discovery_job_collection"].Elem.(*schema.Resource).Schema)
		discoveryJob["items"] = items
	}

	resources = append(resources, discoveryJob)
	if err := s.D.Set("discovery_job_collection", resources); err != nil {
		return err
	}

	return nil
}
