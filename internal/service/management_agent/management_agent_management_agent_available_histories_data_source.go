// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentAvailableHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgentAvailableHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"management_agent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_availability_status_ended_greater_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_availability_status_started_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_histories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_agent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_availability_status_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_availability_status_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readManagementAgentManagementAgentAvailableHistories(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentAvailableHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentAvailableHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListAvailabilityHistoriesResponse
}

func (s *ManagementAgentManagementAgentAvailableHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentAvailableHistoriesDataSourceCrud) Get() error {
	request := oci_management_agent.ListAvailabilityHistoriesRequest{}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if timeAvailabilityStatusEndedGreaterThan, ok := s.D.GetOkExists("time_availability_status_ended_greater_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAvailabilityStatusEndedGreaterThan.(string))
		if err != nil {
			return err
		}
		request.TimeAvailabilityStatusEndedGreaterThan = &oci_common.SDKTime{Time: tmp}
	}

	if timeAvailabilityStatusStartedLessThan, ok := s.D.GetOkExists("time_availability_status_started_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAvailabilityStatusStartedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeAvailabilityStatusStartedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.ListAvailabilityHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAvailabilityHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentManagementAgentAvailableHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentAvailableHistoriesDataSource-", ManagementAgentManagementAgentAvailableHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managementAgentAvailableHistory := map[string]interface{}{
			"management_agent_id": *r.ManagementAgentId,
		}

		managementAgentAvailableHistory["availability_status"] = r.AvailabilityStatus

		if r.TimeAvailabilityStatusEnded != nil {
			managementAgentAvailableHistory["time_availability_status_ended"] = r.TimeAvailabilityStatusEnded.String()
		}

		if r.TimeAvailabilityStatusStarted != nil {
			managementAgentAvailableHistory["time_availability_status_started"] = r.TimeAvailabilityStatusStarted.String()
		}

		resources = append(resources, managementAgentAvailableHistory)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentAvailableHistoriesDataSource().Schema["availability_histories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("availability_histories", resources); err != nil {
		return err
	}

	return nil
}
