// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"strings"

	"time"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

var (
	clusterOperationMaxTime = 60 * time.Minute
)

func ClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCluster,
		Read:     readCluster,
		Update:   updateCluster,
		Delete:   deleteCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"options": {
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
						"add_ons": {
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
									"is_kubernetes_dashboard_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_tiller_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"kubernetes_network_config": {
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
									"pods_cidr": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"services_cidr": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"service_lb_subnet_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"available_kubernetes_upgrades": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"kubernetes": {
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
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"created_by_user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_work_request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deleted_by_user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deleted_by_work_request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_deleted": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by_user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by_work_request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return CreateResource(d, sync)
}

func readCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

func updateCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return UpdateResource(d, sync)
}

func deleteCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ClusterResourceCrud struct {
	BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.Cluster
	DisableNotFoundRetries bool
}

func (s *ClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateCreating),
	}
}

func (s *ClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateActive),
	}
}

func (s *ClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateDeleting),
	}
}

func (s *ClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateDeleted),
	}
}

//containerEngineWorkRequestShouldRetryFunc Custom retry function for containerengine service
func containerEngineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	stopTime := time.Now().Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		//Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		//Make sure we stop on default rules
		if shouldRetry(response, false, "containerengine") {
			return true
		}

		// Only stop if the time Finished is set
		if okeRes, ok := response.Response.(oci_containerengine.GetWorkRequestResponse); ok {
			return okeRes.TimeFinished == nil
		}
		return false
	}
}

//containerEngineWaitForWorkRequest custom logic to extract an identifier from a workRequest
func containerEngineWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = containerEngineWorkRequestShouldRetryFunc(timeout)

	response, err := client.GetWorkRequest(context.Background(),
		oci_containerengine.GetWorkRequestRequest{
			WorkRequestId: wId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return nil, err
	}

	var identifier *string
	//The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			identifier = res.Identifier
			if res.ActionType == action {
				return res.Identifier, nil
			}
		}
	}

	//Otherwise the operation ended unsucessfully
	errorMessage, _ := getErrorFromWorkRequest(wId, response.CompartmentId, client, disableFoundRetries)
	return identifier, fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
}

func (s *ClusterResourceCrud) Create() error {
	request := oci_containerengine.CreateClusterRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if options, ok := s.D.GetOkExists("options"); ok {
		if tmpList := options.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", 0)
			tmp, err := s.mapToClusterCreateOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Options = &tmp
		}
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}
	//Trigger a create request
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	response, err := s.Client.CreateCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	//Wait until it finishes
	clusterID, err := containerEngineWaitForWorkRequest(workId, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeCreated, clusterOperationMaxTime, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		if clusterID != nil {
			//Try to clean up
			delReq := oci_containerengine.DeleteClusterRequest{}
			delReq.ClusterId = clusterID
			delReq.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

			//Issue the delete delReq
			delRes, delErr := s.Client.DeleteCluster(context.Background(), delReq)
			if delErr != nil {
				return err
			}
			delWorkRequest := delRes.OpcWorkRequestId

			//Wait until request finishes
			_, _ = containerEngineWaitForWorkRequest(delWorkRequest, "cluster",
				oci_containerengine.WorkRequestResourceActionTypeDeleted, clusterOperationMaxTime, s.DisableNotFoundRetries, s.Client)
		}
		return err
	}

	//Fetch the cluster object
	requestGet := oci_containerengine.GetClusterRequest{}
	requestGet.ClusterId = clusterID
	requestGet.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetCluster(context.Background(), requestGet)
	if err != nil {
		return err
	}
	s.Res = &responseGet.Cluster

	return nil
}

func (s *ClusterResourceCrud) Get() error {
	id := s.D.Id()
	request := oci_containerengine.GetClusterRequest{}
	request.ClusterId = &id
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Cluster
	return nil
}

func (s *ClusterResourceCrud) Update() error {
	request := oci_containerengine.UpdateClusterRequest{}

	tmp := s.D.Id()
	request.ClusterId = &tmp

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	//Issue update request
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	response, err := s.Client.UpdateCluster(context.Background(), request)
	if err != nil {
		return err
	}
	workRequest := response.OpcWorkRequestId

	//Wait until request finishes
	clusterID, err := containerEngineWaitForWorkRequest(workRequest, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeUpdated,
		clusterOperationMaxTime, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}

	//Refresh data
	requestGet := oci_containerengine.GetClusterRequest{}
	requestGet.ClusterId = clusterID
	requestGet.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetCluster(context.Background(), requestGet)
	if err != nil {
		return err
	}

	s.Res = &responseGet.Cluster
	return nil
}

func (s *ClusterResourceCrud) Delete() error {
	request := oci_containerengine.DeleteClusterRequest{}
	tmp := s.D.Id()
	request.ClusterId = &tmp
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	//Issue the delete request
	response, err := s.Client.DeleteCluster(context.Background(), request)
	if err != nil {
		return err
	}
	workRequest := response.OpcWorkRequestId

	//Wait until request finishes
	_, err = containerEngineWaitForWorkRequest(workRequest, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeDeleted, clusterOperationMaxTime, s.DisableNotFoundRetries, s.Client)

	return err
}

func (s *ClusterResourceCrud) SetData() error {
	s.D.Set("available_kubernetes_upgrades", s.Res.AvailableKubernetesUpgrades)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Endpoints != nil {
		s.D.Set("endpoints", []interface{}{ClusterEndpointsToMap(s.Res.Endpoints)})
	} else {
		s.D.Set("endpoints", nil)
	}

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ClusterMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Options != nil {
		s.D.Set("options", []interface{}{ClusterCreateOptionsToMap(s.Res.Options)})
	} else {
		s.D.Set("options", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *ClusterResourceCrud) mapToAddOnOptions(fieldKeyFormat string) (oci_containerengine.AddOnOptions, error) {
	result := oci_containerengine.AddOnOptions{}

	if isKubernetesDashboardEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_kubernetes_dashboard_enabled")); ok {
		tmp := isKubernetesDashboardEnabled.(bool)
		result.IsKubernetesDashboardEnabled = &tmp
	}

	if isTillerEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_tiller_enabled")); ok {
		tmp := isTillerEnabled.(bool)
		result.IsTillerEnabled = &tmp
	}

	return result, nil
}

func AddOnOptionsToMap(obj *oci_containerengine.AddOnOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsKubernetesDashboardEnabled != nil {
		result["is_kubernetes_dashboard_enabled"] = bool(*obj.IsKubernetesDashboardEnabled)
	}

	if obj.IsTillerEnabled != nil {
		result["is_tiller_enabled"] = bool(*obj.IsTillerEnabled)
	}

	return result
}

func (s *ClusterResourceCrud) mapToClusterCreateOptions(fieldKeyFormat string) (oci_containerengine.ClusterCreateOptions, error) {
	result := oci_containerengine.ClusterCreateOptions{}

	if addOns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "add_ons")); ok {
		if tmpList := addOns.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "add_ons"), 0)
			tmp, err := s.mapToAddOnOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			result.AddOns = &tmp
		}
	}

	if kubernetesNetworkConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kubernetes_network_config")); ok {
		if tmpList := kubernetesNetworkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "kubernetes_network_config"), 0)
			tmp, err := s.mapToKubernetesNetworkConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			result.KubernetesNetworkConfig = &tmp
		}
	}

	if serviceLbSubnetIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_lb_subnet_ids")); ok {
		interfaces := serviceLbSubnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			tmp[i] = interfaces[i].(string)
		}
		result.ServiceLbSubnetIds = tmp
	}

	return result, nil
}

func ClusterCreateOptionsToMap(obj *oci_containerengine.ClusterCreateOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddOns != nil {
		result["add_ons"] = []interface{}{AddOnOptionsToMap(obj.AddOns)}
	}

	if obj.KubernetesNetworkConfig != nil {
		result["kubernetes_network_config"] = []interface{}{KubernetesNetworkConfigToMap(obj.KubernetesNetworkConfig)}
	}

	result["service_lb_subnet_ids"] = obj.ServiceLbSubnetIds

	return result
}

func ClusterEndpointsToMap(obj *oci_containerengine.ClusterEndpoints) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Kubernetes != nil {
		result["kubernetes"] = string(*obj.Kubernetes)
	}

	return result
}

func ClusterMetadataToMap(obj *oci_containerengine.ClusterMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreatedByUserId != nil {
		result["created_by_user_id"] = string(*obj.CreatedByUserId)
	}

	if obj.CreatedByWorkRequestId != nil {
		result["created_by_work_request_id"] = string(*obj.CreatedByWorkRequestId)
	}

	if obj.DeletedByUserId != nil {
		result["deleted_by_user_id"] = string(*obj.DeletedByUserId)
	}

	if obj.DeletedByWorkRequestId != nil {
		result["deleted_by_work_request_id"] = string(*obj.DeletedByWorkRequestId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeDeleted != nil {
		result["time_deleted"] = obj.TimeDeleted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UpdatedByUserId != nil {
		result["updated_by_user_id"] = string(*obj.UpdatedByUserId)
	}

	if obj.UpdatedByWorkRequestId != nil {
		result["updated_by_work_request_id"] = string(*obj.UpdatedByWorkRequestId)
	}

	return result
}

func (s *ClusterResourceCrud) mapToKubernetesNetworkConfig(fieldKeyFormat string) (oci_containerengine.KubernetesNetworkConfig, error) {
	result := oci_containerengine.KubernetesNetworkConfig{}

	if podsCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pods_cidr")); ok {
		tmp := podsCidr.(string)
		result.PodsCidr = &tmp
	}

	if servicesCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "services_cidr")); ok {
		tmp := servicesCidr.(string)
		result.ServicesCidr = &tmp
	}

	return result, nil
}

func KubernetesNetworkConfigToMap(obj *oci_containerengine.KubernetesNetworkConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PodsCidr != nil {
		result["pods_cidr"] = string(*obj.PodsCidr)
	}

	if obj.ServicesCidr != nil {
		result["services_cidr"] = string(*obj.ServicesCidr)
	}

	return result
}

// getErrorFromWorkRequest retuns a concatened string of all errors for a given work request, if there is a reading the error it returns an empty string an error
func getErrorFromWorkRequest(workRequestId *string, compartmentId *string, client *oci_containerengine.ContainerEngineClient, disableFoundAutoRetries bool) (string, error) {
	req := oci_containerengine.ListWorkRequestErrorsRequest{}
	req.WorkRequestId = workRequestId
	req.CompartmentId = compartmentId
	req.RequestMetadata.RetryPolicy = getRetryPolicy(disableFoundAutoRetries, "containerengine")
	res, err := client.ListWorkRequestErrors(context.Background(), req)

	if err != nil {
		return "", err
	}

	allErrs := make([]string, 0)
	for _, errs := range res.Items {
		allErrs = append(allErrs, *errs.Message)
	}

	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage, nil
}
