// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceNodeReplaceConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceNodeReplaceConfiguration,
		Read:     readBdsBdsInstanceNodeReplaceConfiguration,
		Update:   updateBdsBdsInstanceNodeReplaceConfiguration,
		Delete:   deleteBdsBdsInstanceNodeReplaceConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"duration_in_minutes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"level_type_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"level_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NODE_LEVEL",
								"NODE_TYPE_LEVEL",
							}, true),
						},

						// Optional
						"node_host_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"node_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"metric_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			//	"remove_trigger": {
			//		Type:     schema.TypeInt,
			//		Optional: true,
			//	},

			// Computed
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
		},
	}
}

func createBdsBdsInstanceNodeReplaceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeReplaceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	return tfresource.CreateResource(d, sync)

	//if e := tfresource.CreateResource(d, sync); e != nil {
	//	return e
	//}

	//if _, ok := sync.D.GetOkExists("remove_trigger"); ok {
	//	err := sync.RemoveNodeReplaceConfiguration()
	//	if err != nil {
	//		return err
	//	}
	//}
	//return nil

}

func readBdsBdsInstanceNodeReplaceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeReplaceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func updateBdsBdsInstanceNodeReplaceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeReplaceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	return tfresource.UpdateResource(d, sync)

	//	if _, ok := sync.D.GetOkExists("remove_trigger"); ok && sync.D.HasChange("remove_trigger") {
	//		oldRaw, newRaw := sync.D.GetChange("remove_trigger")
	//		oldValue := oldRaw.(int)
	//		newValue := newRaw.(int)
	//		if oldValue < newValue {
	//			err := sync.RemoveNodeReplaceConfiguration()
	//
	//			if err != nil {
	//				return err
	//			}
	//		} else {
	//			sync.D.Set("remove_trigger", oldRaw)
	//			return fmt.Errorf("new value of trigger should be greater than the old value")
	//		}
	//	}

	//	if err := tfresource.UpdateResource(d, sync); err != nil {
	//		return err
	//	}
	//
	//	return nil
}

func deleteBdsBdsInstanceNodeReplaceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeReplaceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceNodeReplaceConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.NodeReplaceConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) ID() string {
	return GetBdsInstanceNodeReplaceConfigurationCompositeId(s.D.Get("bds_instance_id").(string), *s.Res.Id)
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.NodeReplaceConfigurationLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.NodeReplaceConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.NodeReplaceConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.NodeReplaceConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) Create() error {
	request := oci_bds.CreateNodeReplaceConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if durationInMinutes, ok := s.D.GetOkExists("duration_in_minutes"); ok {
		tmp := durationInMinutes.(int)
		request.DurationInMinutes = &tmp
	}

	if levelTypeDetails, ok := s.D.GetOkExists("level_type_details"); ok {
		if tmpList := levelTypeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "level_type_details", 0)
			tmp, err := s.mapToLevelTypeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LevelTypeDetails = tmp
		}
	}

	if metricType, ok := s.D.GetOkExists("metric_type"); ok {
		request.MetricType = oci_bds.NodeReplaceConfigurationMetricTypeEnum(metricType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateNodeReplaceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceNodeReplaceConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) getBdsInstanceNodeReplaceConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceNodeReplaceConfigurationId, err := bdsInstanceNodeReplaceConfigurationWaitForWorkRequest(workId, "nodeReplaceConfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceNodeReplaceConfigurationId)

	return s.Get()
}

func bdsInstanceNodeReplaceConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceNodeReplaceConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceNodeReplaceConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bds.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(*res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceNodeReplaceConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceNodeReplaceConfigurationWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bds.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) Get() error {
	request := oci_bds.GetNodeReplaceConfigurationRequest{}

	tmp := s.D.Id()
	request.NodeReplaceConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsInstanceId, nodeReplaceConfigurationId, err := parseBdsInstanceNodeReplaceConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsInstanceId = &bdsInstanceId
		request.NodeReplaceConfigurationId = &nodeReplaceConfigurationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetNodeReplaceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NodeReplaceConfiguration
	return nil
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) Update() error {
	request := oci_bds.UpdateNodeReplaceConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if durationInMinutes, ok := s.D.GetOkExists("duration_in_minutes"); ok {
		tmp := durationInMinutes.(int)
		request.DurationInMinutes = &tmp
	}

	if levelTypeDetails, ok := s.D.GetOkExists("level_type_details"); ok {
		if tmpList := levelTypeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "level_type_details", 0)
			tmp, err := s.mapToLevelTypeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LevelTypeDetails = tmp
		}
	}

	if metricType, ok := s.D.GetOkExists("metric_type"); ok {
		request.MetricType = oci_bds.NodeReplaceConfigurationMetricTypeEnum(metricType.(string))
	}

	tmp := s.D.Id()
	request.NodeReplaceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateNodeReplaceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceNodeReplaceConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) SetData() error {

	bdsInstanceId, nodeReplaceConfigurationId, err := parseBdsInstanceNodeReplaceConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
		s.D.SetId(nodeReplaceConfigurationId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BdsInstanceId != nil {
		s.D.Set("bds_instance_id", *s.Res.BdsInstanceId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DurationInMinutes != nil {
		s.D.Set("duration_in_minutes", *s.Res.DurationInMinutes)
	}

	if s.Res.LevelTypeDetails != nil {
		levelTypeDetailsArray := []interface{}{}
		if levelTypeDetailsMap := LevelTypeDetailsToMap(&s.Res.LevelTypeDetails); levelTypeDetailsMap != nil {
			levelTypeDetailsArray = append(levelTypeDetailsArray, levelTypeDetailsMap)
		}
		s.D.Set("level_type_details", levelTypeDetailsArray)
	} else {
		s.D.Set("level_type_details", nil)
	}

	s.D.Set("metric_type", s.Res.MetricType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetBdsInstanceNodeReplaceConfigurationCompositeId(bdsInstanceId string, nodeReplaceConfigurationId string) string {
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	nodeReplaceConfigurationId = url.PathEscape(nodeReplaceConfigurationId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/nodeReplaceConfigurations/" + nodeReplaceConfigurationId
	return compositeId
}

func parseBdsInstanceNodeReplaceConfigurationCompositeId(compositeId string) (bdsInstanceId string, nodeReplaceConfigurationId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/nodeReplaceConfigurations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	nodeReplaceConfigurationId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) Delete() error {
	request := oci_bds.RemoveNodeReplaceConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.NodeReplaceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RemoveNodeReplaceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceNodeReplaceConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutUpdate))

}

func (s *BdsBdsInstanceNodeReplaceConfigurationResourceCrud) mapToLevelTypeDetails(fieldKeyFormat string) (oci_bds.LevelTypeDetails, error) {
	var baseObject oci_bds.LevelTypeDetails
	//discriminator
	levelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "level_type"))
	var levelType string
	if ok {
		levelType = levelTypeRaw.(string)
	} else {
		levelType = "" // default value
	}
	switch strings.ToLower(levelType) {
	case strings.ToLower("NODE_LEVEL"):
		details := oci_bds.NodeLevelDetails{}
		if nodeHostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_host_name")); ok {
			tmp := nodeHostName.(string)
			details.NodeHostName = &tmp
		}
		baseObject = details
	case strings.ToLower("NODE_TYPE_LEVEL"):
		details := oci_bds.NodeTypeLevelDetails{}
		if nodeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_type")); ok {
			details.NodeType = oci_bds.NodeNodeTypeEnum(nodeType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown level_type '%v' was specified", levelType)
	}
	return baseObject, nil
}
