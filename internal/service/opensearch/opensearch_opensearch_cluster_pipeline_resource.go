// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpensearchOpensearchClusterPipelineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("45m"),
			Update: tfresource.GetTimeoutDuration("45m"),
			Delete: tfresource.GetTimeoutDuration("45m"),
		},
		Create: createOpensearchOpensearchClusterPipeline,
		Read:   readOpensearchOpensearchClusterPipeline,
		Update: updateOpensearchOpensearchClusterPipeline,
		Delete: deleteOpensearchOpensearchClusterPipeline,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_prepper_configuration_body": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ocpu_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pipeline_configuration_body": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"node_shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsg_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"opc_dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"reverse_connection_endpoints": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"customer_fqdn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"customer_ip": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"subnet_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"opensearch_pipeline_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opensearch_pipeline_private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pipeline_mode": {
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

func createOpensearchOpensearchClusterPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterPipelineClient()
	sync.ClusterClient = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readOpensearchOpensearchClusterPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterPipelineClient()
	sync.ClusterClient = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.ReadResource(sync)
}

func updateOpensearchOpensearchClusterPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterPipelineClient()
	sync.ClusterClient = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpensearchOpensearchClusterPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterPipelineClient()
	sync.ClusterClient = m.(*client.OracleClients).OpensearchClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpensearchOpensearchClusterPipelineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opensearch.OpensearchClusterPipelineClient
	ClusterClient          *oci_opensearch.OpensearchClusterClient
	Res                    *oci_opensearch.OpensearchClusterPipeline
	DisableNotFoundRetries bool
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterPipelineLifecycleStateCreating),
	}
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterPipelineLifecycleStateActive),
	}
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterPipelineLifecycleStateDeleting),
	}
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterPipelineLifecycleStateDeleted),
	}
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) Create() error {
	request := oci_opensearch.CreateOpensearchClusterPipelineRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataPrepperConfigurationBody, ok := s.D.GetOkExists("data_prepper_configuration_body"); ok {
		tmp := dataPrepperConfigurationBody.(string)
		request.DataPrepperConfigurationBody = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if memoryGB, ok := s.D.GetOkExists("memory_gb"); ok {
		tmp := memoryGB.(int)
		request.MemoryGB = &tmp
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
	}

	if nodeShape, ok := s.D.GetOkExists("node_shape"); ok {
		tmp := nodeShape.(string)
		request.NodeShape = &tmp
	}

	if nsgId, ok := s.D.GetOkExists("nsg_id"); ok {
		tmp := nsgId.(string)
		request.NsgId = &tmp
	}

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
		tmp := ocpuCount.(int)
		request.OcpuCount = &tmp
	}

	if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
		tmp := opcDryRun.(bool)
		request.OpcDryRun = &tmp
	}

	if pipelineConfigurationBody, ok := s.D.GetOkExists("pipeline_configuration_body"); ok {
		tmp := pipelineConfigurationBody.(string)
		request.PipelineConfigurationBody = &tmp
	}

	if reverseConnectionEndpoints, ok := s.D.GetOkExists("reverse_connection_endpoints"); ok {
		interfaces := reverseConnectionEndpoints.([]interface{})
		tmp := make([]oci_opensearch.OpensearchPipelineReverseConnectionEndpoint, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reverse_connection_endpoints", stateDataIndex)
			converted, err := s.mapToOpensearchPipelineReverseConnectionEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("reverse_connection_endpoints") {
			request.ReverseConnectionEndpoints = tmp
		}
	}

	if subnetCompartmentId, ok := s.D.GetOkExists("subnet_compartment_id"); ok {
		tmp := subnetCompartmentId.(string)
		request.SubnetCompartmentId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnCompartmentId, ok := s.D.GetOkExists("vcn_compartment_id"); ok {
		tmp := vcnCompartmentId.(string)
		request.VcnCompartmentId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.CreateOpensearchClusterPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_opensearch.GetWorkRequestResponse{}
	workRequestResponse, err = s.ClusterClient.GetWorkRequest(context.Background(),
		oci_opensearch.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "opensearchpipeline") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getOpensearchClusterPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch"), oci_opensearch.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) getOpensearchClusterPipelineFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opensearch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	opensearchClusterPipelineId, err := opensearchClusterPipelineWaitForWorkRequest(workId, "opensearchpipeline",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.ClusterClient)

	if err != nil {
		return err
	}
	s.D.SetId(*opensearchClusterPipelineId)

	return s.Get()
}

func opensearchClusterPipelineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opensearch", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opensearch.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func opensearchClusterPipelineWaitForWorkRequest(wId *string, entityType string, action oci_opensearch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opensearch.OpensearchClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opensearch")
	retryPolicy.ShouldRetryOperation = opensearchClusterPipelineWorkRequestShouldRetryFunc(timeout)

	response := oci_opensearch.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_opensearch.OperationStatusInProgress),
			string(oci_opensearch.OperationStatusAccepted),
			string(oci_opensearch.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opensearch.OperationStatusSucceeded),
			string(oci_opensearch.OperationStatusFailed),
			string(oci_opensearch.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opensearch.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opensearch.OperationStatusFailed || response.Status == oci_opensearch.OperationStatusCanceled {
		return nil, getErrorFromOpensearchOpensearchClusterPipelineWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpensearchOpensearchClusterPipelineWorkRequest(client *oci_opensearch.OpensearchClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opensearch.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opensearch.ListWorkRequestErrorsRequest{
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

func (s *OpensearchOpensearchClusterPipelineResourceCrud) Get() error {
	request := oci_opensearch.GetOpensearchClusterPipelineRequest{}

	tmp := s.D.Id()
	request.OpensearchClusterPipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.GetOpensearchClusterPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OpensearchClusterPipeline
	return nil
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) Update() error {
	request := oci_opensearch.UpdateOpensearchClusterPipelineRequest{}

	if dataPrepperConfigurationBody, ok := s.D.GetOkExists("data_prepper_configuration_body"); ok {
		tmp := dataPrepperConfigurationBody.(string)
		request.DataPrepperConfigurationBody = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if memoryGB, ok := s.D.GetOkExists("memory_gb"); ok {
		tmp := memoryGB.(int)
		request.MemoryGB = &tmp
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
	}

	if nodeShape, ok := s.D.GetOkExists("node_shape"); ok {
		tmp := nodeShape.(string)
		request.NodeShape = &tmp
	}

	if nsgId, ok := s.D.GetOkExists("nsg_id"); ok {
		tmp := nsgId.(string)
		request.NsgId = &tmp
	}

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
		tmp := ocpuCount.(int)
		request.OcpuCount = &tmp
	}

	if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
		tmp := opcDryRun.(bool)
		request.OpcDryRun = &tmp
	}

	tmp := s.D.Id()
	request.OpensearchClusterPipelineId = &tmp

	if pipelineConfigurationBody, ok := s.D.GetOkExists("pipeline_configuration_body"); ok {
		tmp := pipelineConfigurationBody.(string)
		request.PipelineConfigurationBody = &tmp
	}

	if reverseConnectionEndpoints, ok := s.D.GetOkExists("reverse_connection_endpoints"); ok {
		interfaces := reverseConnectionEndpoints.([]interface{})
		tmp := make([]oci_opensearch.OpensearchPipelineReverseConnectionEndpoint, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reverse_connection_endpoints", stateDataIndex)
			converted, err := s.mapToOpensearchPipelineReverseConnectionEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("reverse_connection_endpoints") {
			request.ReverseConnectionEndpoints = tmp
		}
	}

	if subnetCompartmentId, ok := s.D.GetOkExists("subnet_compartment_id"); ok {
		tmp := subnetCompartmentId.(string)
		request.SubnetCompartmentId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnCompartmentId, ok := s.D.GetOkExists("vcn_compartment_id"); ok {
		tmp := vcnCompartmentId.(string)
		request.VcnCompartmentId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.UpdateOpensearchClusterPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOpensearchClusterPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch"), oci_opensearch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) Delete() error {
	request := oci_opensearch.DeleteOpensearchClusterPipelineRequest{}

	tmp := s.D.Id()
	request.OpensearchClusterPipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.DeleteOpensearchClusterPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := opensearchClusterPipelineWaitForWorkRequest(workId, "opensearchpipeline",
		oci_opensearch.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.ClusterClient)
	return delWorkRequestErr
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataPrepperConfigurationBody != nil {
		s.D.Set("data_prepper_configuration_body", *s.Res.DataPrepperConfigurationBody)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MemoryGB != nil {
		s.D.Set("memory_gb", *s.Res.MemoryGB)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.NodeShape != nil {
		s.D.Set("node_shape", *s.Res.NodeShape)
	}

	if s.Res.NsgId != nil {
		s.D.Set("nsg_id", *s.Res.NsgId)
	}

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.OpensearchPipelineFqdn != nil {
		s.D.Set("opensearch_pipeline_fqdn", *s.Res.OpensearchPipelineFqdn)
	}

	if s.Res.OpensearchPipelinePrivateIp != nil {
		s.D.Set("opensearch_pipeline_private_ip", *s.Res.OpensearchPipelinePrivateIp)
	}

	if s.Res.PipelineConfigurationBody != nil {
		s.D.Set("pipeline_configuration_body", *s.Res.PipelineConfigurationBody)
	}

	s.D.Set("pipeline_mode", s.Res.PipelineMode)

	reverseConnectionEndpoints := []interface{}{}
	for _, item := range s.Res.ReverseConnectionEndpoints {
		reverseConnectionEndpoints = append(reverseConnectionEndpoints, OpensearchPipelineReverseConnectionEndpointToMap(item))
	}
	s.D.Set("reverse_connection_endpoints", reverseConnectionEndpoints)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetCompartmentId != nil {
		s.D.Set("subnet_compartment_id", *s.Res.SubnetCompartmentId)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
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

	if s.Res.VcnCompartmentId != nil {
		s.D.Set("vcn_compartment_id", *s.Res.VcnCompartmentId)
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func OpensearchClusterPipelineSummaryToMap(obj oci_opensearch.OpensearchClusterPipelineSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DataPrepperConfigurationBody != nil {
		result["data_prepper_configuration_body"] = string(*obj.DataPrepperConfigurationBody)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MemoryGB != nil {
		result["memory_gb"] = int(*obj.MemoryGB)
	}

	if obj.NodeCount != nil {
		result["node_count"] = int(*obj.NodeCount)
	}

	if obj.NodeShape != nil {
		result["node_shape"] = string(*obj.NodeShape)
	}

	if obj.OcpuCount != nil {
		result["ocpu_count"] = int(*obj.OcpuCount)
	}

	if obj.PipelineConfigurationBody != nil {
		result["pipeline_configuration_body"] = string(*obj.PipelineConfigurationBody)
	}

	result["pipeline_mode"] = string(obj.PipelineMode)

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
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

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *OpensearchOpensearchClusterPipelineResourceCrud) mapToOpensearchPipelineReverseConnectionEndpoint(fieldKeyFormat string) (oci_opensearch.OpensearchPipelineReverseConnectionEndpoint, error) {
	result := oci_opensearch.OpensearchPipelineReverseConnectionEndpoint{}

	if customerFqdn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_fqdn")); ok {
		tmp := customerFqdn.(string)
		result.CustomerFqdn = &tmp
	}

	if customerIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_ip")); ok {
		tmp := customerIp.(string)
		result.CustomerIp = &tmp
	}

	return result, nil
}

func OpensearchPipelineReverseConnectionEndpointToMap(obj oci_opensearch.OpensearchPipelineReverseConnectionEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomerFqdn != nil {
		result["customer_fqdn"] = string(*obj.CustomerFqdn)
	}

	if obj.CustomerIp != nil {
		result["customer_ip"] = string(*obj.CustomerIp)
	}

	return result
}
