// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"strconv"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v56/dataflow"
)

func DataflowInvokeRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataflowInvokeRuns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_starts_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner_principal_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataflowInvokeRunResource()),
			},
		},
	}
}

func readDataflowInvokeRuns(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowInvokeRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowInvokeRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.ListRunsResponse
}

func (s *DataflowInvokeRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowInvokeRunsDataSourceCrud) Get() error {
	request := oci_dataflow.ListRunsRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameStartsWith, ok := s.D.GetOkExists("display_name_starts_with"); ok {
		tmp := displayNameStartsWith.(string)
		request.DisplayNameStartsWith = &tmp
	}

	if ownerPrincipalId, ok := s.D.GetOkExists("owner_principal_id"); ok {
		tmp := ownerPrincipalId.(string)
		request.OwnerPrincipalId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dataflow.ListRunsLifecycleStateEnum(state.(string))
	}

	if timeCreatedGreaterThan, ok := s.D.GetOkExists("time_created_greater_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.ListRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataflowInvokeRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataflowInvokeRunsDataSource-", DataflowInvokeRunsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		invokeRun := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ApplicationId != nil {
			invokeRun["application_id"] = *r.ApplicationId
		}

		if r.DataReadInBytes != nil {
			invokeRun["data_read_in_bytes"] = strconv.FormatInt(*r.DataReadInBytes, 10)
		}

		if r.DataWrittenInBytes != nil {
			invokeRun["data_written_in_bytes"] = strconv.FormatInt(*r.DataWrittenInBytes, 10)
		}

		if r.DefinedTags != nil {
			invokeRun["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			invokeRun["display_name"] = *r.DisplayName
		}

		invokeRun["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			invokeRun["id"] = *r.Id
		}

		invokeRun["language"] = r.Language

		if r.LifecycleDetails != nil {
			invokeRun["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.OpcRequestId != nil {
			invokeRun["opc_request_id"] = *r.OpcRequestId
		}

		if r.OwnerPrincipalId != nil {
			invokeRun["owner_principal_id"] = *r.OwnerPrincipalId
		}

		if r.OwnerUserName != nil {
			invokeRun["owner_user_name"] = *r.OwnerUserName
		}

		if r.RunDurationInMilliseconds != nil {
			invokeRun["run_duration_in_milliseconds"] = strconv.FormatInt(*r.RunDurationInMilliseconds, 10)
		}

		invokeRun["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			invokeRun["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			invokeRun["time_updated"] = r.TimeUpdated.String()
		}

		if r.TotalOCpu != nil {
			invokeRun["total_ocpu"] = *r.TotalOCpu
		}

		invokeRun["type"] = r.Type

		resources = append(resources, invokeRun)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataflowInvokeRunsDataSource().Schema["runs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("runs", resources); err != nil {
		return err
	}

	return nil
}
