// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v42/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v42/mysql"
)

func init() {
	RegisterResource("oci_mysql_analytics_cluster", MysqlAnalyticsClusterResource())
}

func MysqlAnalyticsClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("1h"),
			Update: getTimeoutDuration("1h"),
			Delete: getTimeoutDuration("1h"),
		},
		Create: createMysqlAnalyticsCluster,
		Read:   readMysqlAnalyticsCluster,
		Update: updateMysqlAnalyticsCluster,
		Delete: deleteMysqlAnalyticsCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.AnalyticsClusterLifecycleStateInactive),
					string(oci_mysql.AnalyticsClusterLifecycleStateActive),
				}, true),
			},

			// Computed
			"cluster_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"node_id": {
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
					},
				},
			},
			"lifecycle_details": {
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

func createMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.WorkRequestClient = m.(*OracleClients).mysqlWorkRequestsClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_mysql.AnalyticsClusterLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_mysql.AnalyticsClusterLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}

	// switch to power off
	if powerOff {
		if err := sync.Stop(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.AnalyticsClusterLifecycleStateInactive)
	}

	return nil
}

func readMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	return ReadResource(sync)
}

func updateMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.WorkRequestClient = m.(*OracleClients).mysqlWorkRequestsClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_mysql.AnalyticsClusterLifecycleStateActive == oci_mysql.AnalyticsClusterLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_mysql.AnalyticsClusterLifecycleStateInactive == oci_mysql.AnalyticsClusterLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	// switch to power on
	if powerOn {
		if err := sync.Start(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.AnalyticsClusterLifecycleStateActive)
	}

	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.Stop(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.AnalyticsClusterLifecycleStateInactive)
	}

	return nil
}

func deleteMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*OracleClients).mysqlWorkRequestsClient()

	return DeleteResource(d, sync)
}

type MysqlAnalyticsClusterResourceCrud struct {
	BaseCrud
	Client                 *oci_mysql.DbSystemClient
	Res                    *oci_mysql.AnalyticsCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_mysql.WorkRequestsClient
}

func (s *MysqlAnalyticsClusterResourceCrud) ID() string {
	return getAnalyticsClusterCompositeId(s.D.Get("db_system_id").(string))
}

func (s *MysqlAnalyticsClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateCreating),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateActive),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateDeleting),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateDeleted),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateUpdating),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateActive),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) Create() error {
	request := oci_mysql.AddAnalyticsClusterRequest{}

	if clusterSize, ok := s.D.GetOkExists("cluster_size"); ok {
		tmp := clusterSize.(int)
		request.ClusterSize = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.AddAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsClusterFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "mysql"), oci_mysql.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *MysqlAnalyticsClusterResourceCrud) getAnalyticsClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_mysql.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	analyticsClusterId, err := analyticsClusterWaitForWorkRequest(workId, "mysql",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*analyticsClusterId)

	return s.Get()
}

func analyticsClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "mysql", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_mysql.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func analyticsClusterWaitForWorkRequest(wId *string, entityType string, action oci_mysql.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_mysql.WorkRequestsClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "mysql")
	retryPolicy.ShouldRetryOperation = analyticsClusterWorkRequestShouldRetryFunc(timeout)

	response := oci_mysql.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_mysql.WorkRequestOperationStatusInProgress),
			string(oci_mysql.WorkRequestOperationStatusAccepted),
			string(oci_mysql.WorkRequestOperationStatusCanceling),
		},
		Target: []string{
			string(oci_mysql.WorkRequestOperationStatusSucceeded),
			string(oci_mysql.WorkRequestOperationStatusFailed),
			string(oci_mysql.WorkRequestOperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_mysql.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_mysql.WorkRequestOperationStatusFailed || response.Status == oci_mysql.WorkRequestOperationStatusCanceled {
		return nil, getErrorFromMysqlAnalyticsClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromMysqlAnalyticsClusterWorkRequest(client *oci_mysql.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_mysql.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_mysql.ListWorkRequestErrorsRequest{
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

func (s *MysqlAnalyticsClusterResourceCrud) Get() error {
	request := oci_mysql.GetAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	dbSystemId, err := parseAnalyticsClusterCompositeId(s.D.Id())
	if err == nil {
		request.DbSystemId = &dbSystemId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyticsCluster
	return nil
}

func (s *MysqlAnalyticsClusterResourceCrud) Update() error {
	request := oci_mysql.UpdateAnalyticsClusterRequest{}

	if clusterSize, ok := s.D.GetOkExists("cluster_size"); ok {
		tmp := clusterSize.(int)
		request.ClusterSize = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.UpdateAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsClusterFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "mysql"), oci_mysql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlAnalyticsClusterResourceCrud) Delete() error {
	request := oci_mysql.DeleteAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.DeleteAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := analyticsClusterWaitForWorkRequest(workId, "mysql",
		oci_mysql.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *MysqlAnalyticsClusterResourceCrud) SetData() error {

	dbSystemId, err := parseAnalyticsClusterCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("db_system_id", &dbSystemId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	clusterNodes := []interface{}{}
	for _, item := range s.Res.ClusterNodes {
		clusterNodes = append(clusterNodes, AnalyticsClusterNodeToMap(item))
	}
	s.D.Set("cluster_nodes", clusterNodes)

	if s.Res.ClusterSize != nil {
		s.D.Set("cluster_size", *s.Res.ClusterSize)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func getAnalyticsClusterCompositeId(dbSystemId string) string {
	dbSystemId = url.PathEscape(dbSystemId)
	compositeId := "dbSystem/" + dbSystemId + "/analyticsCluster"
	return compositeId
}

func parseAnalyticsClusterCompositeId(compositeId string) (dbSystemId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dbSystem/.*/analyticsCluster", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dbSystemId, _ = url.PathUnescape(parts[1])

	return
}

func AnalyticsClusterNodeToMap(obj oci_mysql.AnalyticsClusterNode) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NodeId != nil {
		result["node_id"] = string(*obj.NodeId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *MysqlAnalyticsClusterResourceCrud) Stop() error {
	request := oci_mysql.StopAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StopAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.AnalyticsClusterLifecycleStateInactive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlAnalyticsClusterResourceCrud) Start() error {
	request := oci_mysql.StartAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StartAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.AnalyticsClusterLifecycleStateActive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
