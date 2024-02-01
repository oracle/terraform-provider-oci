// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentDataSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagementAgentManagementAgentDataSource,
		Read:     readManagementAgentManagementAgentDataSource,
		Update:   updateManagementAgentManagementAgentDataSource,
		Delete:   deleteManagementAgentManagementAgentDataSource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"KUBERNETES_CLUSTER",
					"PROMETHEUS_EMITTER",
				}, true),
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"allow_metrics": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metric_dimensions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"proxy_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"read_data_limit_in_kilobytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"read_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_mins": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"is_daemon_set": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"data_source_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"read_data_limit": {
				Type:     schema.TypeInt,
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
	}
}

func createManagementAgentManagementAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readManagementAgentManagementAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

func updateManagementAgentManagementAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagementAgentManagementAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagementAgentManagementAgentDataSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_management_agent.ManagementAgentClient
	Res                    *oci_management_agent.DataSource
	DisableNotFoundRetries bool
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) ID() string {
	managementAgentDataSource := *s.Res
	return GetManagementAgentsDataSourcesCompositeId(s.D.Get("management_agent_id").(string), *managementAgentDataSource.GetKey())
}
func GetManagementAgentsDataSourcesCompositeId(managementAgentId string, dataSourceKey string) string {
	managementAgentId = url.PathEscape(managementAgentId)
	dataSourceKey = url.PathEscape(dataSourceKey)
	compositeId := "managementAgents/" + managementAgentId + "/dataSources/" + dataSourceKey
	log.Printf("[DEBUG] GetManagementAgentsDataSourcesCompositeId: %v\n", compositeId)
	return compositeId
}
func (s *ManagementAgentManagementAgentDataSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesCreating),
	}
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesActive),
	}
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesDeleting),
	}
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesTerminated),
		string(oci_management_agent.LifecycleStatesDeleted),
	}
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) Create() error {
	request := oci_management_agent.CreateDataSourceRequest{}
	err := s.populateTopLevelPolymorphicCreateDataSourceRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.CreateDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getManagementAgentDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent"), oci_management_agent.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) getManagementAgentDataSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_management_agent.ActionTypesEnum, timeout time.Duration) error {

	if workId == nil {
		return fmt.Errorf("work request is empty")
	}
	// Wait until it finishes
	managementAgentDataSourceId, err := managementAgentDataSourceWaitForWorkRequest(workId, "managementagent",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, managementAgentDataSourceId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_management_agent.DeleteWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*managementAgentDataSourceId)

	return s.Get()
}

func managementAgentDataSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "management_agent", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_management_agent.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func managementAgentDataSourceWaitForWorkRequest(wId *string, entityType string, action oci_management_agent.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_management_agent.ManagementAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "management_agent")
	retryPolicy.ShouldRetryOperation = managementAgentDataSourceWorkRequestShouldRetryFunc(timeout)

	response := oci_management_agent.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_management_agent.OperationStatusInProgress),
			string(oci_management_agent.OperationStatusAccepted),
			string(oci_management_agent.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_management_agent.OperationStatusSucceeded),
			string(oci_management_agent.OperationStatusFailed),
			string(oci_management_agent.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_management_agent.GetWorkRequestRequest{
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
	var identifierStr string

	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			var pattern = `/\d*/`
			var regex = regexp.MustCompile(pattern)
			identifierStr = regex.ReplaceAllString(*res.EntityUri, "")
			identifier = &identifierStr
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_management_agent.OperationStatusFailed || response.Status == oci_management_agent.OperationStatusCanceled {
		return nil, getErrorFromManagementAgentManagementAgentDataSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromManagementAgentManagementAgentDataSourceWorkRequest(client *oci_management_agent.ManagementAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_management_agent.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_management_agent.ListWorkRequestErrorsRequest{
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

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) Get() error {
	request := oci_management_agent.GetDataSourceRequest{}

	tmp := s.D.Id()
	request.DataSourceKey = &tmp

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	dataSourceKey, managementAgentId, err := parseManagementAgentDataSourceCompositeId(s.D.Id())
	if err == nil {
		request.DataSourceKey = &dataSourceKey
		request.ManagementAgentId = &managementAgentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.GetDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataSource
	return nil
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) Update() error {
	request := oci_management_agent.UpdateDataSourceRequest{}
	err := s.populateTopLevelPolymorphicUpdateDataSourceRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.UpdateDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getManagementAgentDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent"), oci_management_agent.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) Delete() error {
	request := oci_management_agent.DeleteDataSourceRequest{}

	tmp := s.D.Id()
	request.DataSourceKey = &tmp

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.DeleteDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := managementAgentDataSourceWaitForWorkRequest(workId, "managementagent",
		oci_management_agent.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) SetData() error {

	dataSourceKey, managementAgentId, err := parseManagementAgentDataSourceCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(dataSourceKey)
		s.D.Set("management_agent_id", &managementAgentId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_management_agent.KubernetesClusterDataSource:
		s.D.Set("type", "KUBERNETES_CLUSTER")

		if v.IsDaemonSet != nil {
			s.D.Set("is_daemon_set", *v.IsDaemonSet)
		}

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Key != nil {
			s.D.Set("data_source_key", *v.Key)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		s.D.Set("state", v.State)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_management_agent.PrometheusEmitterDataSource:
		s.D.Set("type", "PROMETHEUS_EMITTER")

		if v.AllowMetrics != nil {
			s.D.Set("allow_metrics", *v.AllowMetrics)
		}

		if v.ConnectionTimeout != nil {
			s.D.Set("connection_timeout", *v.ConnectionTimeout)
		}

		metricDimensions := []interface{}{}
		for _, item := range v.MetricDimensions {
			metricDimensions = append(metricDimensions, MetricDimensionToMap(item))
		}
		s.D.Set("metric_dimensions", metricDimensions)

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.ProxyUrl != nil {
			s.D.Set("proxy_url", *v.ProxyUrl)
		}

		if v.ReadDataLimit != nil {
			s.D.Set("read_data_limit", *v.ReadDataLimit)
		}

		if v.ReadTimeout != nil {
			s.D.Set("read_timeout", *v.ReadTimeout)
		}

		if v.ResourceGroup != nil {
			s.D.Set("resource_group", *v.ResourceGroup)
		}

		if v.ScheduleMins != nil {
			s.D.Set("schedule_mins", *v.ScheduleMins)
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Key != nil {
			s.D.Set("data_source_key", *v.Key)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		s.D.Set("state", v.State)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetManagementAgentDataSourceCompositeId(dataSourceKey string, managementAgentId string) string {
	//If the Key is already from an ID, it will contain the child information already,  if so, just return this as the Id
	match, _ := regexp.MatchString("managementAgents/.*/dataSources/.*", dataSourceKey)
	if match {
		return dataSourceKey
	}
	dataSourceKey = url.PathEscape(dataSourceKey)
	managementAgentId = url.PathEscape(managementAgentId)
	compositeId := "managementAgents/" + managementAgentId + "/dataSources/" + dataSourceKey
	return compositeId
}

func parseManagementAgentDataSourceCompositeId(compositeId string) (dataSourceKey string, managementAgentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("managementAgents/.*/dataSources/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	managementAgentId, _ = url.PathUnescape(parts[1])
	dataSourceKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) mapToMetricDimension(fieldKeyFormat string) (oci_management_agent.MetricDimension, error) {
	result := oci_management_agent.MetricDimension{}

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

func MetricDimensionToMap(obj oci_management_agent.MetricDimension) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) populateTopLevelPolymorphicCreateDataSourceRequest(request *oci_management_agent.CreateDataSourceRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("PROMETHEUS_EMITTER"):
		details := oci_management_agent.CreatePrometheusEmitterDataSourceDetails{}
		if allowMetrics, ok := s.D.GetOkExists("allow_metrics"); ok {
			tmp := allowMetrics.(string)
			details.AllowMetrics = &tmp
		}
		if connectionTimeout, ok := s.D.GetOkExists("connection_timeout"); ok {
			tmp := connectionTimeout.(int)
			details.ConnectionTimeout = &tmp
		}
		if metricDimensions, ok := s.D.GetOkExists("metric_dimensions"); ok {
			interfaces := metricDimensions.([]interface{})
			tmp := make([]oci_management_agent.MetricDimension, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metric_dimensions", stateDataIndex)
				converted, err := s.mapToMetricDimension(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("metric_dimensions") {
				details.MetricDimensions = tmp
			}
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if proxyUrl, ok := s.D.GetOkExists("proxy_url"); ok {
			tmp := proxyUrl.(string)
			details.ProxyUrl = &tmp
		}
		if readDataLimitInKilobytes, ok := s.D.GetOkExists("read_data_limit_in_kilobytes"); ok {
			tmp := readDataLimitInKilobytes.(int)
			details.ReadDataLimitInKilobytes = &tmp
		}
		if readTimeout, ok := s.D.GetOkExists("read_timeout"); ok {
			tmp := readTimeout.(int)
			details.ReadTimeout = &tmp
		}
		if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
			tmp := resourceGroup.(string)
			details.ResourceGroup = &tmp
		}
		if scheduleMins, ok := s.D.GetOkExists("schedule_mins"); ok {
			tmp := scheduleMins.(int)
			details.ScheduleMins = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
			tmp := managementAgentId.(string)
			request.ManagementAgentId = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateDataSourceDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *ManagementAgentManagementAgentDataSourceResourceCrud) populateTopLevelPolymorphicUpdateDataSourceRequest(request *oci_management_agent.UpdateDataSourceRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("PROMETHEUS_EMITTER"):
		details := oci_management_agent.UpdatePrometheusEmitterDataSourceDetails{}
		if allowMetrics, ok := s.D.GetOkExists("allow_metrics"); ok {
			tmp := allowMetrics.(string)
			details.AllowMetrics = &tmp
		}
		if connectionTimeout, ok := s.D.GetOkExists("connection_timeout"); ok {
			tmp := connectionTimeout.(int)
			details.ConnectionTimeout = &tmp
		}
		if metricDimensions, ok := s.D.GetOkExists("metric_dimensions"); ok {
			interfaces := metricDimensions.([]interface{})
			tmp := make([]oci_management_agent.MetricDimension, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metric_dimensions", stateDataIndex)
				converted, err := s.mapToMetricDimension(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("metric_dimensions") {
				details.MetricDimensions = tmp
			}
		}
		if proxyUrl, ok := s.D.GetOkExists("proxy_url"); ok {
			tmp := proxyUrl.(string)
			details.ProxyUrl = &tmp
		}
		if readDataLimitInKilobytes, ok := s.D.GetOkExists("read_data_limit_in_kilobytes"); ok {
			tmp := readDataLimitInKilobytes.(int)
			details.ReadDataLimitInKilobytes = &tmp
		}
		if readTimeout, ok := s.D.GetOkExists("read_timeout"); ok {
			tmp := readTimeout.(int)
			details.ReadTimeout = &tmp
		}
		if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
			tmp := resourceGroup.(string)
			details.ResourceGroup = &tmp
		}
		if scheduleMins, ok := s.D.GetOkExists("schedule_mins"); ok {
			tmp := scheduleMins.(int)
			details.ScheduleMins = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		tmp := s.D.Id()
		request.DataSourceKey = &tmp

		if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
			tmp := managementAgentId.(string)
			request.ManagementAgentId = &tmp
		}

		dataSourceKey, managementAgentId, err := parseManagementAgentDataSourceCompositeId(s.D.Id())
		if err == nil {
			request.DataSourceKey = &dataSourceKey
			request.ManagementAgentId = &managementAgentId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
		request.UpdateDataSourceDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
