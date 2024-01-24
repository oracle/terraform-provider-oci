// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"
)

func DataflowInvokeRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataflowInvokeRun,
		Read:     readDataflowInvokeRun,
		Update:   updateDataflowInvokeRun,
		Delete:   deleteDataflowInvokeRun,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"application_log_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"log_group_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"log_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"archive_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"arguments": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"driver_shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"execute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"executor_shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"idle_timeout_in_minutes": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"logs_bucket_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"max_duration_in_minutes": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"metastore_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"num_executors": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"spark_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"warehouse_bucket_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"asynchronous": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},

			// Computed
			"class_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_read_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_written_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opc_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_dns_zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_max_host_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"private_endpoint_nsg_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_endpoint_subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"run_duration_in_milliseconds": {
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
			"total_ocpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDataflowInvokeRun(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowInvokeRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.CreateResource(d, sync)
}

func readDataflowInvokeRun(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowInvokeRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

func updateDataflowInvokeRun(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowInvokeRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataflowInvokeRun(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowInvokeRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataflowInvokeRunResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataflow.DataFlowClient
	Res                    *oci_dataflow.Run
	DisableNotFoundRetries bool
}

func (s *DataflowInvokeRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataflowInvokeRunResourceCrud) CreatedPending() []string {
	invokeAsynchronously := true
	if async, ok := s.D.GetOkExists("asynchronous"); ok {
		invokeAsynchronously = async.(bool)
	}
	if invokeAsynchronously {
		return []string{
			string(oci_dataflow.RunLifecycleStateInProgress),
		}
	}
	return []string{
		string(oci_dataflow.RunLifecycleStateInProgress),
		string(oci_dataflow.RunLifecycleStateAccepted),
	}
}

func (s *DataflowInvokeRunResourceCrud) CreatedTarget() []string {
	invokeAsynchronously := true
	if async, ok := s.D.GetOkExists("asynchronous"); ok {
		invokeAsynchronously = async.(bool)
	}
	if invokeAsynchronously {
		return []string{
			string(oci_dataflow.RunLifecycleStateAccepted),
		}
	}
	return []string{
		string(oci_dataflow.RunLifecycleStateSucceeded),
		string(oci_dataflow.RunLifecycleStateFailed),
	}
}

func (s *DataflowInvokeRunResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dataflow.RunLifecycleStateAccepted),
		string(oci_dataflow.RunLifecycleStateInProgress),
		string(oci_dataflow.RunLifecycleStateCanceling),
		string(oci_dataflow.RunLifecycleStateStopping),
		string(oci_dataflow.RunLifecycleStateCanceled),
		string(oci_dataflow.RunLifecycleStateStopped),
	}
}

func (s *DataflowInvokeRunResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataflow.RunLifecycleStateSucceeded),
		string(oci_dataflow.RunLifecycleStateFailed),
		string(oci_dataflow.RunLifecycleStateCanceled),
		string(oci_dataflow.RunLifecycleStateStopped),
	}
}

func (s *DataflowInvokeRunResourceCrud) Create() error {
	request := oci_dataflow.CreateRunRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if applicationLogConfig, ok := s.D.GetOkExists("application_log_config"); ok {
		if tmpList := applicationLogConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "application_log_config", 0)
			tmp, err := s.mapToApplicationLogConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ApplicationLogConfig = &tmp
		}
	}

	if archiveUri, ok := s.D.GetOkExists("archive_uri"); ok {
		tmp := archiveUri.(string)
		request.ArchiveUri = &tmp
	}

	if arguments, ok := s.D.GetOkExists("arguments"); ok {
		interfaces := arguments.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("arguments") {
			request.Arguments = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		request.Configuration = tfresource.ObjectMapToStringMap(configuration.(map[string]interface{}))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
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

	if execute, ok := s.D.GetOkExists("execute"); ok {
		tmp := execute.(string)
		request.Execute = &tmp
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

	if idleTimeoutInMinutes, ok := s.D.GetOkExists("idle_timeout_in_minutes"); ok {
		tmp := idleTimeoutInMinutes.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert idleTimeoutInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.IdleTimeoutInMinutes = &tmpInt64
	}

	if logsBucketUri, ok := s.D.GetOkExists("logs_bucket_uri"); ok {
		tmp := logsBucketUri.(string)
		request.LogsBucketUri = &tmp
	}

	if maxDurationInMinutes, ok := s.D.GetOkExists("max_duration_in_minutes"); ok {
		tmp := maxDurationInMinutes.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert maxDurationInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.MaxDurationInMinutes = &tmpInt64
	}

	if metastoreId, ok := s.D.GetOkExists("metastore_id"); ok {
		tmp := metastoreId.(string)
		request.MetastoreId = &tmp
	}

	if numExecutors, ok := s.D.GetOkExists("num_executors"); ok {
		tmp := numExecutors.(int)
		request.NumExecutors = &tmp
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_dataflow.ApplicationParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
			converted, err := s.mapToApplicationParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	if poolId, ok := s.D.GetOkExists("pool_id"); ok {
		tmp := poolId.(string)
		request.PoolId = &tmp
	}

	if sparkVersion, ok := s.D.GetOkExists("spark_version"); ok {
		tmp := sparkVersion.(string)
		request.SparkVersion = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_dataflow.ApplicationTypeEnum(type_.(string))
	}

	if warehouseBucketUri, ok := s.D.GetOkExists("warehouse_bucket_uri"); ok {
		tmp := warehouseBucketUri.(string)
		request.WarehouseBucketUri = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.CreateRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Run
	return nil
}

func (s *DataflowInvokeRunResourceCrud) Get() error {
	request := oci_dataflow.GetRunRequest{}

	tmp := s.D.Id()
	request.RunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.GetRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Run
	return nil
}

func (s *DataflowInvokeRunResourceCrud) Update() error {
	fmt.Println("ML: in run update here is s.D", s.D)
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dataflow.UpdateRunRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if runId, ok := s.D.GetOkExists("run_id"); ok {
		tmp := runId.(string)
		request.RunId = &tmp
	} else {
		tmp := s.D.Id()
		request.RunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.UpdateRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Run
	return nil
}

func (s *DataflowInvokeRunResourceCrud) Delete() error {
	getRequest := oci_dataflow.GetRunRequest{}

	tmp := s.D.Id()
	getRequest.RunId = &tmp

	getRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	getResponse, err := s.Client.GetRun(context.Background(), getRequest)
	if err != nil {
		return err
	}
	lifecycleState := getResponse.LifecycleState
	if lifecycleState == oci_dataflow.RunLifecycleStateSucceeded || lifecycleState == oci_dataflow.RunLifecycleStateFailed {
		return nil
	}
	request := oci_dataflow.DeleteRunRequest{}

	tmp = s.D.Id()
	request.RunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err = s.Client.DeleteRun(context.Background(), request)
	return err
}

func (s *DataflowInvokeRunResourceCrud) SetData() error {
	if s.Res.ApplicationId != nil {
		s.D.Set("application_id", *s.Res.ApplicationId)
	}

	if s.Res.ApplicationLogConfig != nil {
		s.D.Set("application_log_config", []interface{}{DataflowRunApplicationLogConfigToMap(s.Res.ApplicationLogConfig)})
	} else {
		s.D.Set("application_log_config", nil)
	}

	if s.Res.ArchiveUri != nil {
		s.D.Set("archive_uri", *s.Res.ArchiveUri)
	}

	s.D.Set("arguments", s.Res.Arguments)

	if s.Res.ClassName != nil {
		s.D.Set("class_name", *s.Res.ClassName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("configuration", s.Res.Configuration)

	if s.Res.DataReadInBytes != nil {
		s.D.Set("data_read_in_bytes", strconv.FormatInt(*s.Res.DataReadInBytes, 10))
	}

	if s.Res.DataWrittenInBytes != nil {
		s.D.Set("data_written_in_bytes", strconv.FormatInt(*s.Res.DataWrittenInBytes, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.Execute != nil {
		s.D.Set("execute", *s.Res.Execute)
	}

	if s.Res.ExecutorShape != nil {
		s.D.Set("executor_shape", *s.Res.ExecutorShape)
	}

	if s.Res.ExecutorShapeConfig != nil {
		s.D.Set("executor_shape_config", []interface{}{ShapeConfigToMap(s.Res.ExecutorShapeConfig)})
	} else {
		s.D.Set("executor_shape_config", nil)
	}

	if s.Res.FileUri != nil {
		s.D.Set("file_uri", *s.Res.FileUri)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdleTimeoutInMinutes != nil {
		s.D.Set("idle_timeout_in_minutes", strconv.FormatInt(*s.Res.IdleTimeoutInMinutes, 10))
	}

	s.D.Set("language", s.Res.Language)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogsBucketUri != nil {
		s.D.Set("logs_bucket_uri", *s.Res.LogsBucketUri)
	}

	if s.Res.MaxDurationInMinutes != nil {
		s.D.Set("max_duration_in_minutes", strconv.FormatInt(*s.Res.MaxDurationInMinutes, 10))
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	if s.Res.NumExecutors != nil {
		s.D.Set("num_executors", *s.Res.NumExecutors)
	}

	if s.Res.OpcRequestId != nil {
		s.D.Set("opc_request_id", *s.Res.OpcRequestId)
	}

	if s.Res.OwnerPrincipalId != nil {
		s.D.Set("owner_principal_id", *s.Res.OwnerPrincipalId)
	}

	if s.Res.OwnerUserName != nil {
		s.D.Set("owner_user_name", *s.Res.OwnerUserName)
	}

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, ApplicationParameterToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.PoolId != nil {
		s.D.Set("pool_id", *s.Res.PoolId)
	}

	s.D.Set("private_endpoint_dns_zones", s.Res.PrivateEndpointDnsZones)

	if s.Res.PrivateEndpointId != nil {
		s.D.Set("private_endpoint_id", *s.Res.PrivateEndpointId)
	}

	if s.Res.PrivateEndpointMaxHostCount != nil {
		s.D.Set("private_endpoint_max_host_count", *s.Res.PrivateEndpointMaxHostCount)
	}

	s.D.Set("private_endpoint_nsg_ids", s.Res.PrivateEndpointNsgIds)

	if s.Res.PrivateEndpointSubnetId != nil {
		s.D.Set("private_endpoint_subnet_id", *s.Res.PrivateEndpointSubnetId)
	}

	if s.Res.RunDurationInMilliseconds != nil {
		s.D.Set("run_duration_in_milliseconds", strconv.FormatInt(*s.Res.RunDurationInMilliseconds, 10))
	}

	if s.Res.SparkVersion != nil {
		s.D.Set("spark_version", *s.Res.SparkVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalOCpu != nil {
		s.D.Set("total_ocpu", *s.Res.TotalOCpu)
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.WarehouseBucketUri != nil {
		s.D.Set("warehouse_bucket_uri", *s.Res.WarehouseBucketUri)
	}

	return nil
}

func GetInvokeRunCompositeId(runId string) string {
	runId = url.PathEscape(runId)
	compositeId := "runs/" + runId
	return compositeId
}

func parseInvokeRunCompositeId(compositeId string) (runId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("runs/.*", compositeId)
	if !match || len(parts) != 2 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	runId, _ = url.PathUnescape(parts[1])

	return
}

func (s *DataflowInvokeRunResourceCrud) mapToApplicationLogConfig(fieldKeyFormat string) (oci_dataflow.ApplicationLogConfig, error) {
	result := oci_dataflow.ApplicationLogConfig{}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func DataflowRunApplicationLogConfigToMap(obj *oci_dataflow.ApplicationLogConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DataflowInvokeRunResourceCrud) mapToApplicationParameter(fieldKeyFormat string) (oci_dataflow.ApplicationParameter, error) {
	result := oci_dataflow.ApplicationParameter{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *DataflowInvokeRunResourceCrud) mapToShapeConfig(fieldKeyFormat string) (oci_dataflow.ShapeConfig, error) {
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

func (s *DataflowInvokeRunResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dataflow.ChangeRunCompartmentRequest{}
	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if runId := s.D.Id(); runId != "" {
		changeCompartmentRequest.RunId = &runId
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.ChangeRunCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
