// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v58/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlAccessRequestHistoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOperatorAccessControlAccessRequestHistory,
		Schema: map[string]*schema.Schema{
			"access_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"actions_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"duration": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_auto_approved": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularOperatorAccessControlAccessRequestHistory(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlAccessRequestHistoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AccessRequestsClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlAccessRequestHistoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.AccessRequestsClient
	Res    *oci_operator_access_control.ListAccessRequestHistoriesResponse
}

func (s *OperatorAccessControlAccessRequestHistoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlAccessRequestHistoryDataSourceCrud) Get() error {
	request := oci_operator_access_control.ListAccessRequestHistoriesRequest{}

	if accessRequestId, ok := s.D.GetOkExists("access_request_id"); ok {
		tmp := accessRequestId.(string)
		request.AccessRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.ListAccessRequestHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OperatorAccessControlAccessRequestHistoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OperatorAccessControlAccessRequestHistoryDataSource-", OperatorAccessControlAccessRequestHistoryDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AccessRequestHistorySummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AccessRequestHistorySummaryToMap(obj oci_operator_access_control.AccessRequestHistorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["actions_list"] = obj.ActionsList

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Duration != nil {
		result["duration"] = int(*obj.Duration)
	}

	if obj.IsAutoApproved != nil {
		result["is_auto_approved"] = bool(*obj.IsAutoApproved)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeOfAction != nil {
		result["time_of_action"] = obj.TimeOfAction.String()
	}

	if obj.UserId != nil {
		result["user_id"] = string(*obj.UserId)
	}

	return result
}
