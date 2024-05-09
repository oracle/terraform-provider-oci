// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubScheduledJobsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubScheduledJobs,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_restricted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"lifecycle_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"location_not_equal_to": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"managed_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operation_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduled_job_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubScheduledJobResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubScheduledJobs(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubScheduledJobsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledJobClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubScheduledJobsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ScheduledJobClient
	Res    *oci_os_management_hub.ListScheduledJobsResponse
}

func (s *OsManagementHubScheduledJobsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubScheduledJobsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListScheduledJobsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if isManagedByAutonomousLinux, ok := s.D.GetOkExists("is_managed_by_autonomous_linux"); ok {
		tmp := isManagedByAutonomousLinux.(bool)
		request.IsManagedByAutonomousLinux = &tmp
	}

	if isRestricted, ok := s.D.GetOkExists("is_restricted"); ok {
		tmp := isRestricted.(bool)
		request.IsRestricted = &tmp
	}

	if lifecycleStageId, ok := s.D.GetOkExists("lifecycle_stage_id"); ok {
		tmp := lifecycleStageId.(string)
		request.LifecycleStageId = &tmp
	}

	if location, ok := s.D.GetOkExists("location"); ok {
		interfaces := location.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("location") {
			request.Location = tmp
		}
	}

	if locationNotEqualTo, ok := s.D.GetOkExists("location_not_equal_to"); ok {
		interfaces := locationNotEqualTo.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("location_not_equal_to") {
			request.LocationNotEqualTo = tmp
		}
	}

	if managedCompartmentId, ok := s.D.GetOkExists("managed_compartment_id"); ok {
		tmp := managedCompartmentId.(string)
		request.ManagedCompartmentId = &tmp
	}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if operationType, ok := s.D.GetOkExists("operation_type"); ok {
		request.OperationType = oci_os_management_hub.ListScheduledJobsOperationTypeEnum(operationType.(string))
	}

	if scheduleType, ok := s.D.GetOkExists("schedule_type"); ok {
		request.ScheduleType = oci_os_management_hub.ListScheduledJobsScheduleTypeEnum(scheduleType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_os_management_hub.ScheduledJobLifecycleStateEnum(state.(string))
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListScheduledJobs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledJobs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubScheduledJobsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubScheduledJobsDataSource-", OsManagementHubScheduledJobsDataSource(), s.D))
	resources := []map[string]interface{}{}
	scheduledJob := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledJobSummaryToMap(item))
	}
	scheduledJob["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubScheduledJobsDataSource().Schema["scheduled_job_collection"].Elem.(*schema.Resource).Schema)
		scheduledJob["items"] = items
	}

	resources = append(resources, scheduledJob)
	if err := s.D.Set("scheduled_job_collection", resources); err != nil {
		return err
	}

	return nil
}
