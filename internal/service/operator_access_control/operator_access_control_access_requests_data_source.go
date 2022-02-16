// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v58/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlAccessRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOperatorAccessControlAccessRequests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
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
			"access_request_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_reason_summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"action_requests_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"approver_comment": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audit_type": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"closure_comment": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"duration": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"extend_duration": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_auto_approved": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"opctl_additional_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"opctl_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"opctl_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operator_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reason": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"request_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_creation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_modification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_user_creation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"workflow_id": {
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

func readOperatorAccessControlAccessRequests(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlAccessRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AccessRequestsClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlAccessRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.AccessRequestsClient
	Res    *oci_operator_access_control.ListAccessRequestsResponse
}

func (s *OperatorAccessControlAccessRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlAccessRequestsDataSourceCrud) Get() error {
	request := oci_operator_access_control.ListAccessRequestsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := string(oci_operator_access_control.ResourceTypesEnum(resourceType.(string)))
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_operator_access_control.ListAccessRequestsLifecycleStateEnum(state.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.ListAccessRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OperatorAccessControlAccessRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OperatorAccessControlAccessRequestsDataSource-", OperatorAccessControlAccessRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	accessRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AccessRequestsSummaryToMap(item))
	}
	accessRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OperatorAccessControlAccessRequestsDataSource().Schema["access_request_collection"].Elem.(*schema.Resource).Schema)
		accessRequest["items"] = items
	}

	resources = append(resources, accessRequest)
	if err := s.D.Set("access_request_collection", resources); err != nil {
		return err
	}

	return nil
}

func AccessRequestsSummaryToMap(obj oci_operator_access_control.AccessRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessReasonSummary != nil {
		result["access_reason_summary"] = string(*obj.AccessReasonSummary)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Duration != nil {
		result["duration"] = int(*obj.Duration)
	}

	if obj.ExtendDuration != nil {
		result["extend_duration"] = int(*obj.ExtendDuration)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAutoApproved != nil {
		result["is_auto_approved"] = bool(*obj.IsAutoApproved)
	}

	if obj.RequestId != nil {
		result["request_id"] = string(*obj.RequestId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["severity"] = string(obj.Severity)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeOfCreation != nil {
		result["time_of_creation"] = obj.TimeOfCreation.String()
	}

	if obj.TimeOfModification != nil {
		result["time_of_modification"] = obj.TimeOfModification.String()
	}

	if obj.TimeOfUserCreation != nil {
		result["time_of_user_creation"] = obj.TimeOfUserCreation.String()
	}

	return result
}
