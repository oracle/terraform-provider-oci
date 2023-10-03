// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityPolicyManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSecurityPolicyManagement,
		Read:     readDataSafeSecurityPolicyManagement,
		Update:   updateDataSafeSecurityPolicyManagement,
		Delete:   deleteDataSafeSecurityPolicyManagement,
		Schema: map[string]*schema.Schema{
			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSecurityPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.GetIdFromDbSecurityConfigWorkReq()
	err1 := sync.Get()
	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	return updateDataSafeSecurityPolicyManagement(d, m)
}

func readDataSafeSecurityPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSecurityPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSecurityPolicyManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeSecurityPolicyManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SecurityPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) GetIdFromDbSecurityConfigWorkReq() error {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	tmp := "CREATE_DB_SECURITY_CONFIG"
	listWorkRequestsRequest.OperationType = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		listWorkRequestsRequest.TargetDatabaseId = &tmp
	}

	listWorkRequestsResponse, err := s.Client.ListWorkRequests(context.Background(), listWorkRequestsRequest)
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		var tmp1 = &listWorkRequestsResponse.Items[0]
		workId = tmp1.Id
	}

	if err != nil {
		return err
	}

	if workId != nil {
		err = s.getDbSecurityConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		return s.GetIdFromSecurityPolicyDeploymentList()
	} else {
		return s.GetIdFromSecurityPolicyDeploymentList()
	}
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) GetIdFromSecurityPolicyDeploymentList() error {
	deploymentsRequest := oci_data_safe.ListSecurityPolicyDeploymentsRequest{}
	var securityPolicy = new(oci_data_safe.SecurityPolicy)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		deploymentsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		deploymentsRequest.TargetId = &tmp
	}
	deploymentsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListSecurityPolicyDeployments(context.Background(), deploymentsRequest)
	if err != nil {
		return err
	}
	if response.SecurityPolicyDeploymentCollection.Items != nil && len(response.SecurityPolicyDeploymentCollection.Items) > 0 {
		tmp1 := &response.SecurityPolicyDeploymentCollection.Items[0]
		securityPolicy.Id = tmp1.SecurityPolicyId
	}

	if securityPolicy.Id == nil {
		return nil
	}

	s.D.SetId(*securityPolicy.Id)
	return nil
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) getDbSecurityConfigFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := databaseSecurityConfigWaitForWorkRequest(workId, "databasesecurityconfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	return nil
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) getSecurityPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	securityPolicyId, err := securityPolicyWaitForWorkRequest(workId, "securitypolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*securityPolicyId)

	return s.Get()
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) Get() error {
	request := oci_data_safe.GetSecurityPolicyRequest{}

	tmp := s.D.Id()
	request.SecurityPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSecurityPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityPolicy
	return nil
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSecurityPolicyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SecurityPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSecurityPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSecurityPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeSecurityPolicyManagementResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSecurityPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecurityPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeSecurityPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSecurityPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
