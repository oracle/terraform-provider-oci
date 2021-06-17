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
	RegisterResource("oci_mysql_heat_wave_cluster", MysqlHeatWaveClusterResource())
}

func MysqlHeatWaveClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("1h"),
			Update: getTimeoutDuration("1h"),
			Delete: getTimeoutDuration("1h"),
		},
		Create: createMysqlHeatWaveCluster,
		Read:   readMysqlHeatWaveCluster,
		Update: updateMysqlHeatWaveCluster,
		Delete: deleteMysqlHeatWaveCluster,
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
					string(oci_mysql.HeatWaveClusterLifecycleStateInactive),
					string(oci_mysql.HeatWaveClusterLifecycleStateActive),
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

func createMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.WorkRequestClient = m.(*OracleClients).mysqlWorkRequestsClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_mysql.HeatWaveClusterLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_mysql.HeatWaveClusterLifecycleStateInactive {
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
		sync.D.Set("state", oci_mysql.HeatWaveClusterLifecycleStateInactive)
	}

	return nil
}

func readMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	return ReadResource(sync)
}

func updateMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.WorkRequestClient = m.(*OracleClients).mysqlWorkRequestsClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_mysql.HeatWaveClusterLifecycleStateActive == oci_mysql.HeatWaveClusterLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_mysql.HeatWaveClusterLifecycleStateInactive == oci_mysql.HeatWaveClusterLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	// switch to power on
	if powerOn {
		if err := sync.Start(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.HeatWaveClusterLifecycleStateActive)
	}

	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.Stop(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.HeatWaveClusterLifecycleStateInactive)
	}

	return nil
}

func deleteMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*OracleClients).mysqlWorkRequestsClient()

	return DeleteResource(d, sync)
}

type MysqlHeatWaveClusterResourceCrud struct {
	BaseCrud
	Client                 *oci_mysql.DbSystemClient
	Res                    *oci_mysql.HeatWaveCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_mysql.WorkRequestsClient
}

func (s *MysqlHeatWaveClusterResourceCrud) ID() string {
	return getHeatWaveClusterCompositeId(s.D.Get("db_system_id").(string))
}

func (s *MysqlHeatWaveClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateCreating),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateActive),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateDeleting),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateDeleted),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateUpdating),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateActive),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) Create() error {
	request := oci_mysql.AddHeatWaveClusterRequest{}

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

	response, err := s.Client.AddHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHeatWaveClusterFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "mysql"), oci_mysql.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *MysqlHeatWaveClusterResourceCrud) getHeatWaveClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_mysql.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	heatWaveClusterId, err := heatWaveClusterWaitForWorkRequest(workId, "mysql",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*heatWaveClusterId)

	return s.Get()
}

func heatWaveClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func heatWaveClusterWaitForWorkRequest(wId *string, entityType string, action oci_mysql.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_mysql.WorkRequestsClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "mysql")
	retryPolicy.ShouldRetryOperation = heatWaveClusterWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromMysqlHeatWaveClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromMysqlHeatWaveClusterWorkRequest(client *oci_mysql.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_mysql.WorkRequestResourceActionTypeEnum) error {
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

func (s *MysqlHeatWaveClusterResourceCrud) Get() error {
	request := oci_mysql.GetHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	dbSystemId, err := parseHeatWaveClusterCompositeId(s.D.Id())
	if err == nil {
		request.DbSystemId = &dbSystemId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HeatWaveCluster
	return nil
}

func (s *MysqlHeatWaveClusterResourceCrud) Update() error {
	request := oci_mysql.UpdateHeatWaveClusterRequest{}

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

	response, err := s.Client.UpdateHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHeatWaveClusterFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "mysql"), oci_mysql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlHeatWaveClusterResourceCrud) Delete() error {
	request := oci_mysql.DeleteHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.DeleteHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := heatWaveClusterWaitForWorkRequest(workId, "mysql",
		oci_mysql.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *MysqlHeatWaveClusterResourceCrud) SetData() error {

	dbSystemId, err := parseHeatWaveClusterCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("db_system_id", &dbSystemId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	clusterNodes := []interface{}{}
	for _, item := range s.Res.ClusterNodes {
		clusterNodes = append(clusterNodes, HeatWaveNodeToMap(item))
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

func getHeatWaveClusterCompositeId(dbSystemId string) string {
	dbSystemId = url.PathEscape(dbSystemId)
	compositeId := "dbSystem/" + dbSystemId + "/heatWaveCluster"
	return compositeId
}

func parseHeatWaveClusterCompositeId(compositeId string) (dbSystemId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dbSystem/.*/heatWaveCluster", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dbSystemId, _ = url.PathUnescape(parts[1])

	return
}

func HeatWaveNodeToMap(obj oci_mysql.HeatWaveNode) map[string]interface{} {
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

func (s *MysqlHeatWaveClusterResourceCrud) Stop() error {
	request := oci_mysql.StopHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StopHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.HeatWaveClusterLifecycleStateInactive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlHeatWaveClusterResourceCrud) Start() error {
	request := oci_mysql.StartHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StartHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.HeatWaveClusterLifecycleStateActive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
