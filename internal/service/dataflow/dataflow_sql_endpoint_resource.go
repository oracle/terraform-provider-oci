// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataflowSqlEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("45m"),
			Update: tfresource.GetTimeoutDuration("45m"),
			Delete: tfresource.GetTimeoutDuration("45m"),
		},
		Create: createDataflowSqlEndpoint,
		Read:   readDataflowSqlEndpoint,
		Update: updateDataflowSqlEndpoint,
		Delete: deleteDataflowSqlEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"driver_shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"executor_shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lake_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_executor_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"metastore_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"min_executor_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"network_configuration": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"network_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SECURE_ACCESS",
								"VCN",
							}, true),
						},

						// Optional
						"access_control_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ip_notation": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vcn_ips": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"host_name_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"private_endpoint_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_endpoint_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"sql_endpoint_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"driver_shape_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"memory_in_gbs": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ocpus": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"executor_shape_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"memory_in_gbs": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ocpus": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"spark_advanced_configurations": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_dataflow.SqlEndpointLifecycleStateCreating),
					string(oci_dataflow.SqlEndpointLifecycleStateActive),
					string(oci_dataflow.SqlEndpointLifecycleStateDeleting),
					string(oci_dataflow.SqlEndpointLifecycleStateDeleted),
					string(oci_dataflow.SqlEndpointLifecycleStateFailed),
					string(oci_dataflow.SqlEndpointLifecycleStateUpdating),
					string(oci_dataflow.SqlEndpointLifecycleStateNeedsAttention),
					string(oci_dataflow.SqlEndpointLifecycleStateInactive),
				}, true),
			},

			// Computed
			"warehouse_bucket_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jdbc_endpoint_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_message": {
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

func createDataflowSqlEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowSqlEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_dataflow.SqlEndpointLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_dataflow.SqlEndpointLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopSqlEndpoint(); err != nil {
			return err
		}
		sync.D.Set("state", oci_dataflow.SqlEndpointLifecycleStateInactive)
	}
	return nil

}

func readDataflowSqlEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowSqlEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

func updateDataflowSqlEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowSqlEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_dataflow.SqlEndpointLifecycleStateActive == oci_dataflow.SqlEndpointLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_dataflow.SqlEndpointLifecycleStateInactive == oci_dataflow.SqlEndpointLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartSqlEndpoint(); err != nil {
			return err
		}
		sync.D.Set("state", oci_dataflow.SqlEndpointLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopSqlEndpoint(); err != nil {
			return err
		}
		sync.D.Set("state", oci_dataflow.SqlEndpointLifecycleStateInactive)
	}

	return nil
}

func deleteDataflowSqlEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowSqlEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataflowSqlEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataflow.DataFlowClient
	Res                    *oci_dataflow.SqlEndpoint
	DisableNotFoundRetries bool
}

func (s *DataflowSqlEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataflowSqlEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dataflow.SqlEndpointLifecycleStateCreating),
	}
}

func (s *DataflowSqlEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dataflow.SqlEndpointLifecycleStateActive),
	}
}

func (s *DataflowSqlEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dataflow.SqlEndpointLifecycleStateDeleting),
	}
}

func (s *DataflowSqlEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataflow.SqlEndpointLifecycleStateDeleted),
	}
}

func (s *DataflowSqlEndpointResourceCrud) Create() error {
	request := oci_dataflow.CreateSqlEndpointRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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

	if driverShape, ok := s.D.GetOkExists("driver_shape"); ok {
		tmp := driverShape.(string)
		request.DriverShape = &tmp
	}

	if driverShapeConfig, ok := s.D.GetOkExists("driver_shape_config"); ok {
		if tmpList := driverShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "driver_shape_config", 0)
			tmp, err := s.mapToShapeConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DriverShapeConfig = &tmp
		}
	}

	if executorShape, ok := s.D.GetOkExists("executor_shape"); ok {
		tmp := executorShape.(string)
		request.ExecutorShape = &tmp
	}

	if executorShapeConfig, ok := s.D.GetOkExists("executor_shape_config"); ok {
		if tmpList := executorShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "executor_shape_config", 0)
			tmp, err := s.mapToShapeConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ExecutorShapeConfig = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if lakeId, ok := s.D.GetOkExists("lake_id"); ok {
		tmp := lakeId.(string)
		request.LakeId = &tmp
	}

	if maxExecutorCount, ok := s.D.GetOkExists("max_executor_count"); ok {
		tmp := maxExecutorCount.(int)
		request.MaxExecutorCount = &tmp
	}

	if metastoreId, ok := s.D.GetOkExists("metastore_id"); ok {
		tmp := metastoreId.(string)
		request.MetastoreId = &tmp
	}

	if minExecutorCount, ok := s.D.GetOkExists("min_executor_count"); ok {
		tmp := minExecutorCount.(int)
		request.MinExecutorCount = &tmp
	}

	if networkConfiguration, ok := s.D.GetOkExists("network_configuration"); ok {
		if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_configuration", 0)
			tmp, err := s.mapToSqlEndpointNetworkConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfiguration = tmp
		}
	}

	if sparkAdvancedConfigurations, ok := s.D.GetOkExists("spark_advanced_configurations"); ok {
		request.SparkAdvancedConfigurations = tfresource.ObjectMapToStringMap(sparkAdvancedConfigurations.(map[string]interface{}))
	}

	if sqlEndpointVersion, ok := s.D.GetOkExists("sql_endpoint_version"); ok {
		tmp := sqlEndpointVersion.(string)
		request.SqlEndpointVersion = &tmp
	}

	if warehouseBucketUri, ok := s.D.GetOkExists("warehouse_bucket_uri"); ok {
		tmp := warehouseBucketUri.(string)
		request.WarehouseBucketUri = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.CreateSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.pollForSqlEpOperationCompletion(identifier, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow"), oci_dataflow.SqlEndpointLifecycleStateActive, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataflowSqlEndpointResourceCrud) pollForSqlEpOperationCompletion(sqlEndpointId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dataflow.SqlEndpointLifecycleStateEnum, timeout time.Duration) error {

	// Wait until it finishes
	sqlEndpointId, err := waitForSqlEndpoint(sqlEndpointId, actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*sqlEndpointId)

	return s.Get()
}

func sqlEndpointShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dataflow", startTime) {
			return true
		}

		return false
	}
}

func waitForSqlEndpoint(sqlEndpointId *string, action oci_dataflow.SqlEndpointLifecycleStateEnum,
	timeout time.Duration, disableNotFoundRetries bool, client *oci_dataflow.DataFlowClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableNotFoundRetries, "dataflow")
	retryPolicy.ShouldRetryOperation = sqlEndpointShouldRetryFunc(timeout)

	response := oci_dataflow.GetSqlEndpointResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_dataflow.SqlEndpointLifecycleStateCreating),
			string(oci_dataflow.SqlEndpointLifecycleStateDeleting),
			string(oci_dataflow.SqlEndpointLifecycleStateUpdating),
		},
		Target: []string{
			string(action),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetSqlEndpoint(context.Background(),
				oci_dataflow.GetSqlEndpointRequest{
					SqlEndpointId: sqlEndpointId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			sqlEndpoint := &response.SqlEndpoint
			return sqlEndpoint, string(sqlEndpoint.LifecycleState), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	return response.SqlEndpoint.Id, nil
}

func (s *DataflowSqlEndpointResourceCrud) Get() error {
	request := oci_dataflow.GetSqlEndpointRequest{}

	tmp := s.D.Id()
	request.SqlEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.GetSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SqlEndpoint
	return nil
}

func (s *DataflowSqlEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dataflow.UpdateSqlEndpointRequest{}

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

	if maxExecutorCount, ok := s.D.GetOkExists("max_executor_count"); ok {
		tmp := maxExecutorCount.(int)
		request.MaxExecutorCount = &tmp
	}

	if minExecutorCount, ok := s.D.GetOkExists("min_executor_count"); ok {
		tmp := minExecutorCount.(int)
		request.MinExecutorCount = &tmp
	}

	if sparkAdvancedConfigurations, ok := s.D.GetOkExists("spark_advanced_configurations"); ok {
		request.SparkAdvancedConfigurations = tfresource.ObjectMapToStringMap(sparkAdvancedConfigurations.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SqlEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.UpdateSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	return s.pollForSqlEpOperationCompletion(&tmp, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow"), oci_dataflow.SqlEndpointLifecycleStateActive, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataflowSqlEndpointResourceCrud) Delete() error {
	request := oci_dataflow.DeleteSqlEndpointRequest{}

	tmp := s.D.Id()
	request.SqlEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.DeleteSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	// Wait until sql endpoint is deleted
	_, delReqError := waitForSqlEndpoint(&tmp, oci_dataflow.SqlEndpointLifecycleStateDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delReqError
}

func (s *DataflowSqlEndpointResourceCrud) SetData() error {
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

	if s.Res.DriverShape != nil {
		s.D.Set("driver_shape", *s.Res.DriverShape)
	}

	if s.Res.DriverShapeConfig != nil {
		s.D.Set("driver_shape_config", []interface{}{ShapeConfigToMap(s.Res.DriverShapeConfig)})
	} else {
		s.D.Set("driver_shape_config", nil)
	}

	if s.Res.ExecutorShape != nil {
		s.D.Set("executor_shape", *s.Res.ExecutorShape)
	}

	if s.Res.ExecutorShapeConfig != nil {
		s.D.Set("executor_shape_config", []interface{}{ShapeConfigToMap(s.Res.ExecutorShapeConfig)})
	} else {
		s.D.Set("executor_shape_config", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.JdbcEndpointUrl != nil {
		s.D.Set("jdbc_endpoint_url", *s.Res.JdbcEndpointUrl)
	}

	if s.Res.LakeId != nil {
		s.D.Set("lake_id", *s.Res.LakeId)
	}

	if s.Res.MaxExecutorCount != nil {
		s.D.Set("max_executor_count", *s.Res.MaxExecutorCount)
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	if s.Res.MinExecutorCount != nil {
		s.D.Set("min_executor_count", *s.Res.MinExecutorCount)
	}

	if s.Res.NetworkConfiguration != nil {
		networkConfigurationArray := []interface{}{}
		if networkConfigurationMap := SqlEndpointNetworkConfigurationToMap(&s.Res.NetworkConfiguration, false); networkConfigurationMap != nil {
			networkConfigurationArray = append(networkConfigurationArray, networkConfigurationMap)
		}
		s.D.Set("network_configuration", networkConfigurationArray)
	} else {
		s.D.Set("network_configuration", nil)
	}

	s.D.Set("spark_advanced_configurations", s.Res.SparkAdvancedConfigurations)

	if s.Res.SqlEndpointVersion != nil {
		s.D.Set("sql_endpoint_version", *s.Res.SqlEndpointVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WarehouseBucketUri != nil {
		s.D.Set("warehouse_bucket_uri", *s.Res.WarehouseBucketUri)
	}

	return nil
}

func (s *DataflowSqlEndpointResourceCrud) StartSqlEndpoint() error {
	request := oci_dataflow.StartSqlEndpointRequest{}

	idTmp := s.D.Id()
	request.SqlEndpointId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.StartSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_dataflow.SqlEndpointLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataflowSqlEndpointResourceCrud) StopSqlEndpoint() error {
	request := oci_dataflow.StopSqlEndpointRequest{}

	idTmp := s.D.Id()
	request.SqlEndpointId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.StopSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_dataflow.SqlEndpointLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataflowSqlEndpointResourceCrud) mapToSecureAccessControlRule(fieldKeyFormat string) (oci_dataflow.SecureAccessControlRule, error) {
	result := oci_dataflow.SecureAccessControlRule{}

	if ipNotation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_notation")); ok {
		result.IpNotation = oci_dataflow.IpNotationTypeEnum(ipNotation.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if vcnIps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_ips")); ok {
		tmp := vcnIps.(string)
		result.VcnIps = &tmp
	}

	return result, nil
}

func SecureAccessControlRuleToMap(obj oci_dataflow.SecureAccessControlRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["ip_notation"] = string(obj.IpNotation)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	if obj.VcnIps != nil {
		result["vcn_ips"] = string(*obj.VcnIps)
	}

	return result
}

func (s *DataflowSqlEndpointResourceCrud) mapToShapeConfig(fieldKeyFormat string) (oci_dataflow.ShapeConfig, error) {
	result := oci_dataflow.ShapeConfig{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func (s *DataflowSqlEndpointResourceCrud) mapToSqlEndpointNetworkConfiguration(fieldKeyFormat string) (oci_dataflow.SqlEndpointNetworkConfiguration, error) {
	var baseObject oci_dataflow.SqlEndpointNetworkConfiguration
	//discriminator
	networkTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type"))
	var networkType string
	if ok {
		networkType = networkTypeRaw.(string)
	} else {
		networkType = "SECURE_ACCESS" // default value
	}
	switch strings.ToLower(networkType) {
	case strings.ToLower("SECURE_ACCESS"):
		details := oci_dataflow.SqlEndpointSecureAccessConfig{}
		if accessControlRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_control_rules")); ok {
			interfaces := accessControlRules.([]interface{})
			tmp := make([]oci_dataflow.SecureAccessControlRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "access_control_rules"), stateDataIndex)
				converted, err := s.mapToSecureAccessControlRule(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "access_control_rules")) {
				details.AccessControlRules = tmp
			}
		} else {
			emptyAccessControlRules := make([]oci_dataflow.SecureAccessControlRule, 0)
			details.AccessControlRules = emptyAccessControlRules
		}
		if publicEndpointIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_endpoint_ip")); ok {
			tmp := publicEndpointIp.(string)
			details.PublicEndpointIp = &tmp
		}
		baseObject = details
	case strings.ToLower("VCN"):
		details := oci_dataflow.SqlEndpointVcnConfig{}
		if hostNamePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name_prefix")); ok {
			tmp := hostNamePrefix.(string)
			details.HostNamePrefix = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
				details.NsgIds = tmp
			}
		}
		if privateEndpointIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_ip")); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
			tmp := vcnId.(string)
			details.VcnId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown network_type '%v' was specified", networkType)
	}
	return baseObject, nil
}

func SqlEndpointNetworkConfigurationToMap(obj *oci_dataflow.SqlEndpointNetworkConfiguration, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_dataflow.SqlEndpointSecureAccessConfig:
		result["network_type"] = "SECURE_ACCESS"

		accessControlRules := []interface{}{}
		for _, item := range v.AccessControlRules {
			accessControlRules = append(accessControlRules, SecureAccessControlRuleToMap(item))
		}
		result["access_control_rules"] = accessControlRules

		if v.PublicEndpointIp != nil {
			result["public_endpoint_ip"] = string(*v.PublicEndpointIp)
		}
	case oci_dataflow.SqlEndpointVcnConfig:
		result["network_type"] = "VCN"

		if v.HostNamePrefix != nil {
			result["host_name_prefix"] = string(*v.HostNamePrefix)
		}

		if v.PrivateEndpointIp != nil {
			result["private_endpoint_ip"] = string(*v.PrivateEndpointIp)
		}

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}

		if v.VcnId != nil {
			result["vcn_id"] = string(*v.VcnId)
		}
	default:
		log.Printf("[WARN] Received 'network_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SqlEndpointSummaryToMap(obj oci_dataflow.SqlEndpointSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DriverShape != nil {
		result["driver_shape"] = string(*obj.DriverShape)
	}

	if obj.DriverShapeConfig != nil {
		result["driver_shape_config"] = []interface{}{ShapeConfigToMap(obj.DriverShapeConfig)}
	}

	if obj.ExecutorShape != nil {
		result["executor_shape"] = string(*obj.ExecutorShape)
	}

	if obj.ExecutorShapeConfig != nil {
		result["executor_shape_config"] = []interface{}{ShapeConfigToMap(obj.ExecutorShapeConfig)}
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.JdbcEndpointUrl != nil {
		result["jdbc_endpoint_url"] = string(*obj.JdbcEndpointUrl)
	}

	if obj.LakeId != nil {
		result["lake_id"] = string(*obj.LakeId)
	}

	if obj.MaxExecutorCount != nil {
		result["max_executor_count"] = int(*obj.MaxExecutorCount)
	}

	if obj.MetastoreId != nil {
		result["metastore_id"] = string(*obj.MetastoreId)
	}

	if obj.MinExecutorCount != nil {
		result["min_executor_count"] = int(*obj.MinExecutorCount)
	}

	if obj.NetworkConfiguration != nil {
		networkConfigurationArray := []interface{}{}
		if networkConfigurationMap := SqlEndpointNetworkConfigurationToMap(&obj.NetworkConfiguration, false); networkConfigurationMap != nil {
			networkConfigurationArray = append(networkConfigurationArray, networkConfigurationMap)
		}
		result["network_configuration"] = networkConfigurationArray
	}

	result["spark_advanced_configurations"] = obj.SparkAdvancedConfigurations

	if obj.SqlEndpointVersion != nil {
		result["sql_endpoint_version"] = string(*obj.SqlEndpointVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StateMessage != nil {
		result["state_message"] = string(*obj.StateMessage)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.WarehouseBucketUri != nil {
		result["warehouse_bucket_uri"] = string(*obj.WarehouseBucketUri)
	}

	return result
}

func (s *DataflowSqlEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dataflow.ChangeSqlEndpointCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SqlEndpointId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.ChangeSqlEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return s.pollForSqlEpOperationCompletion(&idTmp, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow"), oci_dataflow.SqlEndpointLifecycleStateActive, s.D.Timeout(schema.TimeoutUpdate))
}
