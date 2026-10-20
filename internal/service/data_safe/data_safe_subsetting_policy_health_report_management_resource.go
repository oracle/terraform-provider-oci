// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicyHealthReportManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeSubsettingPolicyHealthReportManagement,
		ReadContext:   readDataSafeSubsettingPolicyHealthReportManagement,
		DeleteContext: deleteDataSafeSubsettingPolicyHealthReportManagement,
		Schema: map[string]*schema.Schema{
			"target_credentials": {
				Type: schema.TypeList, Required: true, ForceNew: true, MinItems: 1, MaxItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"user_name": {Type: schema.TypeString, Required: true},
					"password":  {Type: schema.TypeString, Required: true, Sensitive: true},
				}},
			},
			"target_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"check_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL", "TABLESPACE_CHECK"}, false),
			},
			"tablespace": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 30),
			},

			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Optional and computed because the service returns the tags generated for
			// the health report when they are not supplied by the configuration.
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"error_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"warning_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSubsettingPolicyHealthReportManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyHealthReportManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.getSubsettingPolicyHealthReportIdFromFilter(ctx)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	err1 := sync.GetWithContext(ctx)
	if err1 != nil {
		return tfresource.HandleDiagError(m, err1)
	}

	err = sync.SetData()
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func readDataSafeSubsettingPolicyHealthReportManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyHealthReportManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteDataSafeSubsettingPolicyHealthReportManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyHealthReportManagementResourceCrud{}
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.D = d

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeSubsettingPolicyHealthReportManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SubsettingPolicyHealthReport
	DisableNotFoundRetries bool
}

func (s *DataSafeSubsettingPolicyHealthReportManagementResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteSubsettingPolicyHealthReportRequest{}

	tmp := s.D.Id()
	request.SubsettingPolicyHealthReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSubsettingPolicyHealthReport(ctx, request)
	if err != nil {
		return err
	}
	if response.OpcWorkRequestId == nil {
		return fmt.Errorf("delete subsetting policy health report operation did not return a work request ID")
	}

	_, err = generateSubsettingHealthReportWaitForWorkRequest(ctx, response.OpcWorkRequestId, "subsettingPolicyHealthReport",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return err
}

func (s *DataSafeSubsettingPolicyHealthReportManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func generateSubsettingHealthReportWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = subsettingPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_data_safe.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if dataSafeWorkRequestEntityTypeMatches(res.EntityType, entityType) && res.Identifier != nil {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}
	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSubsettingPolicyWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}
func (s *DataSafeSubsettingPolicyHealthReportManagementResourceCrud) getSubsettingPolicyHealthReportIdFromFilter(ctx context.Context) error {
	// Subsetting health report will be in same compartment as of policy
	subsettingPolicyID := s.D.Get("subsetting_policy_id").(string)
	compartmentID, resolvedTargetID, err := getSubsettingPolicyTargetAndCompartment(ctx, s.Client, subsettingPolicyID)
	if err != nil {
		return err
	}

	err = s.D.Set("compartment_id", compartmentID)
	if err != nil {
		return err
	}
	if _, ok := s.D.GetOk("target_id"); !ok && resolvedTargetID != nil {
		if err := s.D.Set("target_id", *resolvedTargetID); err != nil {
			return err
		}
	}

	// Get the subsetting health report for given target and subsetting policy ID
	err = s.GetSubsettingPolicyHealthReportList(ctx)
	if err != nil {
		return err
	}

	// check if subsetting health report id is set and subsetting health report already exists
	if s.D.Id() != "" {
		return nil
	}

	// generate health report
	generateHealthReportRequest := oci_data_safe.GenerateSubsettingHealthReportRequest{}
	generateHealthReportRequest.SubsettingPolicyId = &subsettingPolicyID

	if checkType, ok := s.D.GetOkExists("check_type"); ok {
		generateHealthReportRequest.CheckType = oci_data_safe.GenerateSubsettingHealthReportDetailsCheckTypeEnum(checkType.(string))
	}
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		generateHealthReportRequest.CompartmentId = &tmp
	}
	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		generateHealthReportRequest.DefinedTags = convertedDefinedTags
	}
	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		generateHealthReportRequest.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		generateHealthReportRequest.TargetId = &tmp
	}
	if tablespace, ok := s.D.GetOkExists("tablespace"); ok {
		tmp := tablespace.(string)
		generateHealthReportRequest.Tablespace = &tmp
	}
	if credentials, ok := s.D.GetOkExists("target_credentials"); ok {
		values := credentials.([]interface{})[0].(map[string]interface{})
		userName := values["user_name"].(string)
		password := values["password"].(string)
		generateHealthReportRequest.TargetCredentials = &oci_data_safe.Credentials{UserName: &userName, Password: &password}
	}

	generateHealthReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GenerateSubsettingHealthReport(ctx, generateHealthReportRequest)
	if err != nil {
		return err
	}
	if response.OpcWorkRequestId == nil {
		return fmt.Errorf("generate health report operation did not return a work request ID")
	}

	workId := response.OpcWorkRequestId
	subsettingPolicyHealthReportId, err := generateSubsettingHealthReportWaitForWorkRequest(ctx, workId, "subsettingPolicyHealthReport",
		oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err == nil {
		s.D.SetId(*subsettingPolicyHealthReportId)
		return nil
	}
	return err
}

func (s *DataSafeSubsettingPolicyHealthReportManagementResourceCrud) GetSubsettingPolicyHealthReportList(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingPolicyHealthReportsRequest{}
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
		request.SortOrder = oci_data_safe.ListSubsettingPolicyHealthReportsSortOrderDesc
		request.SortBy = oci_data_safe.ListSubsettingPolicyHealthReportsSortByTimecreated
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	var subsettingPolicyHealthReport = new(oci_data_safe.SubsettingPolicyHealthReport)
	response, err := s.Client.ListSubsettingPolicyHealthReports(ctx, request)
	if err != nil {
		return err
	}
	if response.SubsettingPolicyHealthReportCollection.Items != nil && len(response.SubsettingPolicyHealthReportCollection.Items) > 0 {
		temp1 := response.SubsettingPolicyHealthReportCollection.Items[0]
		subsettingPolicyHealthReport.Id = temp1.Id
	}

	if subsettingPolicyHealthReport.Id == nil {
		return nil
	}

	s.D.SetId(*subsettingPolicyHealthReport.Id)
	return nil
}
func (s *DataSafeSubsettingPolicyHealthReportManagementResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingPolicyHealthReportRequest{}

	tmp := s.D.Id()
	request.SubsettingPolicyHealthReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSubsettingPolicyHealthReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SubsettingPolicyHealthReport
	return nil
}

func (s *DataSafeSubsettingPolicyHealthReportManagementResourceCrud) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.SubsettingPolicyId != nil {
		s.D.Set("subsetting_policy_id", *s.Res.SubsettingPolicyId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.ErrorCount != nil {
		s.D.Set("error_count", strconv.FormatInt(*s.Res.ErrorCount, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.WarningCount != nil {
		s.D.Set("warning_count", strconv.FormatInt(*s.Res.WarningCount, 10))
	}

	return nil
}
