// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"
)

func ContainerengineClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createContainerengineCluster,
		Read:   readContainerengineCluster,
		Update: updateContainerengineCluster,
		Delete: deleteContainerengineCluster,
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
			"cluster_pod_network_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cni_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"endpoint_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Optional
						"is_public_ip_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
			"image_policy_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_policy_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"kms_key_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
						"admission_controller_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"is_pod_security_policy_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
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
						"persistent_volume_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

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

									// Computed
								},
							},
						},
						"service_lb_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"kubernetes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vcn_hostname_endpoint": {
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
						"time_credential_expiration": {
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
		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("endpoint_config.0.subnet_id", func(ctx context.Context, old, new, meta interface{}) bool {
				oldSubnetId := old.(string)
				if len(oldSubnetId) > 0 {
					// OKE do not allow customer to change endpointConfig.subnetId once it is configured. Once the oldSubnetId exist, we need to force a recreation.
					return true
				}
				return false
			}),
		),
	}
}

func createContainerengineCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.Cluster
	DisableNotFoundRetries bool
}

func (s *ContainerengineClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerengineClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateCreating),
	}
}

func (s *ContainerengineClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateActive),
	}
}

func (s *ContainerengineClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateDeleting),
	}
}

func (s *ContainerengineClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterLifecycleStateDeleted),
	}
}

func (s *ContainerengineClusterResourceCrud) Create() error {
	request := oci_containerengine.CreateClusterRequest{}

	if clusterPodNetworkOptions, ok := s.D.GetOkExists("cluster_pod_network_options"); ok {
		interfaces := clusterPodNetworkOptions.([]interface{})
		tmp := make([]oci_containerengine.ClusterPodNetworkOptionDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cluster_pod_network_options", stateDataIndex)
			converted, err := s.mapToClusterPodNetworkOptionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("cluster_pod_network_options") {
			request.ClusterPodNetworkOptions = tmp
		}
	}

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

	if endpointConfig, ok := s.D.GetOkExists("endpoint_config"); ok {
		if tmpList := endpointConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "endpoint_config", 0)
			tmp, err := s.mapToCreateClusterEndpointConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EndpointConfig = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if imagePolicyConfig, ok := s.D.GetOkExists("image_policy_config"); ok {
		if tmpList := imagePolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_policy_config", 0)
			tmp, err := s.mapToCreateImagePolicyConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImagePolicyConfig = &tmp
		}
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_containerengine.ClusterTypeEnum(type_.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CreateCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	clusterIDForGet, err := getClusterIdFromWorkRequest(workId, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeCreated, s.DisableNotFoundRetries, s.Client)

	requestGet := oci_containerengine.GetClusterRequest{}
	requestGet.ClusterId = clusterIDForGet
	requestGet.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, getClusterErr := s.Client.GetCluster(context.Background(), requestGet)
	if getClusterErr != nil {
		return getClusterErr
	}

	s.Res = &responseGet.Cluster
	s.D.SetId(s.ID())
	if setDataErr := s.SetData(); setDataErr != nil {
		log.Printf("[ERROR] error setting data before clusterWaitForWorkRequest() error: %v", setDataErr)
	}

	clusterID, err := clusterWaitForWorkRequest(workId, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		if clusterID != nil {

			log.Printf("[DEBUG] creation failed, attempting to delete the cluster: %v\n", clusterID)

			delReq := oci_containerengine.DeleteClusterRequest{}
			delReq.ClusterId = clusterID
			delReq.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

			delRes, delErr := s.Client.DeleteCluster(context.Background(), delReq)
			if delErr != nil {
				return err
			}
			delWorkRequest := delRes.OpcWorkRequestId

			_, delErr = clusterWaitForWorkRequest(delWorkRequest, "cluster",
				oci_containerengine.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
			if delErr != nil {
				log.Printf("[DEBUG] cleanup delWorkRequest failed with the error: %v\n", delErr)
			}
		}
		return err
	}

	requestGet = oci_containerengine.GetClusterRequest{}
	requestGet.ClusterId = clusterID
	requestGet.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err = s.Client.GetCluster(context.Background(), requestGet)
	if err != nil {
		return err
	}
	s.Res = &responseGet.Cluster

	return nil
}

func (s *ContainerengineClusterResourceCrud) getClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_containerengine.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {
	clusterId, err := clusterWaitForWorkRequest(workId, "cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*clusterId)

	return s.Get()
}

func clusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "containerengine", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_containerengine.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func getClusterIdFromWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	response := oci_containerengine.GetWorkRequestResponse{}
	response, _ = client.GetWorkRequest(context.Background(),
		oci_containerengine.GetWorkRequestRequest{
			WorkRequestId: wId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			identifier = res.Identifier
			if res.ActionType == action {
				return res.Identifier, nil
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_containerengine.WorkRequestStatusFailed || response.Status == oci_containerengine.WorkRequestStatusCanceled {
		return nil, getErrorFromContainerengineClusterWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func clusterWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = clusterWorkRequestShouldRetryFunc(timeout)

	response := oci_containerengine.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_containerengine.WorkRequestStatusInProgress),
			string(oci_containerengine.WorkRequestStatusAccepted),
			string(oci_containerengine.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_containerengine.WorkRequestStatusSucceeded),
			string(oci_containerengine.WorkRequestStatusFailed),
			string(oci_containerengine.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_containerengine.GetWorkRequestRequest{
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
	// Set PollInterval to 1 for replay mode.
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			identifier = res.Identifier
			if res.ActionType == action {
				return res.Identifier, nil
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_containerengine.WorkRequestStatusFailed || response.Status == oci_containerengine.WorkRequestStatusCanceled {
		return nil, getErrorFromContainerengineClusterWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerengineClusterWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_containerengine.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			CompartmentId: compartmentId,
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

func (s *ContainerengineClusterResourceCrud) Get() error {
	request := oci_containerengine.GetClusterRequest{}

	tmp := s.D.Id()
	request.ClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Cluster
	return nil
}

func (s *ContainerengineClusterResourceCrud) Update() error {
	clusterID := s.D.Id()
	if endpointConfig, ok := s.D.GetOkExists("endpoint_config"); ok && s.D.HasChange("endpoint_config") {
		oldConfig, _ := s.D.GetChange("endpoint_config")
		oldConfigList := oldConfig.([]interface{})
		if len(oldConfigList) > 0 {
			//If an endpoint config is already set on the cluster, perform and UpdateClusterEndpointConfig operation, otherwise perform a MigrateClusterToNativeVCN operation
			err := s.updateClusterEndpointConfig(clusterID, endpointConfig)
			if err != nil {
				return err
			}
		} else {
			err := s.migrateClusterToNativeVCN(clusterID, endpointConfig)
			if err != nil {
				return err
			}
		}
	}

	request := oci_containerengine.UpdateClusterRequest{}
	request.ClusterId = &clusterID

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

	if imagePolicyConfig, ok := s.D.GetOkExists("image_policy_config"); ok {
		if tmpList := imagePolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_policy_config", 0)
			tmp, err := s.mapToUpdateImagePolicyConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImagePolicyConfig = &tmp
		}
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
			tmp, err := s.mapToUpdateClusterOptionsDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Options = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_containerengine.ClusterTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
func (s *ContainerengineClusterResourceCrud) updateClusterEndpointConfig(clusterID string, endpointConfig interface{}) error {
	request := oci_containerengine.UpdateClusterEndpointConfigRequest{}
	request.ClusterId = &clusterID
	if tmpList := endpointConfig.([]interface{}); len(tmpList) > 0 {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "endpoint_config", 0)
		tmp, err := s.mapToUpdateClusterEndpointConfigDetails(fieldKeyFormat)
		if err != nil {
			return err
		}
		request.UpdateClusterEndpointConfigDetails = tmp
	}

	response, err := s.Client.UpdateClusterEndpointConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workID := response.OpcWorkRequestId
	return s.getClusterFromWorkRequest(workID, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineClusterResourceCrud) migrateClusterToNativeVCN(clusterID string, endpointConfig interface{}) error {
	request := oci_containerengine.ClusterMigrateToNativeVcnRequest{}
	request.ClusterId = &clusterID

	if tmpList := endpointConfig.([]interface{}); len(tmpList) > 0 {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "endpoint_config", 0)
		tmp, err := s.mapToMigrateClusterToNativeVCNDetails(fieldKeyFormat)
		if err != nil {
			return err
		}
		request.ClusterMigrateToNativeVcnDetails = tmp
	}

	response, err := s.Client.ClusterMigrateToNativeVcn(context.Background(), request)
	if err != nil {
		return err
	}

	workID := response.OpcWorkRequestId
	return s.getClusterFromWorkRequest(workID, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineClusterResourceCrud) Delete() error {
	request := oci_containerengine.DeleteClusterRequest{}

	tmp := s.D.Id()
	request.ClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.DeleteCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := clusterWaitForWorkRequest(workId, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerengineClusterResourceCrud) SetData() error {
	s.D.Set("available_kubernetes_upgrades", s.Res.AvailableKubernetesUpgrades)

	clusterPodNetworkOptions := []interface{}{}
	for _, item := range s.Res.ClusterPodNetworkOptions {
		clusterPodNetworkOptions = append(clusterPodNetworkOptions, ClusterPodNetworkOptionDetailsToMap(item))
	}
	s.D.Set("cluster_pod_network_options", clusterPodNetworkOptions)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.EndpointConfig != nil {
		s.D.Set("endpoint_config", []interface{}{ClusterEndpointConfigToMap(s.Res.EndpointConfig, false)})
	} else {
		s.D.Set("endpoint_config", nil)
	}

	if s.Res.Endpoints != nil {
		s.D.Set("endpoints", []interface{}{ClusterEndpointsToMap(s.Res.Endpoints)})
	} else {
		s.D.Set("endpoints", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImagePolicyConfig != nil {
		s.D.Set("image_policy_config", []interface{}{ImagePolicyConfigToMap(s.Res.ImagePolicyConfig)})
	} else {
		s.D.Set("image_policy_config", nil)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
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

	s.D.Set("type", s.Res.Type)

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *ContainerengineClusterResourceCrud) mapToAddOnOptions(fieldKeyFormat string) (oci_containerengine.AddOnOptions, error) {
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

func (s *ContainerengineClusterResourceCrud) mapToAdmissionControllerOptions(fieldKeyFormat string) (oci_containerengine.AdmissionControllerOptions, error) {
	result := oci_containerengine.AdmissionControllerOptions{}

	if isPodSecurityPolicyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pod_security_policy_enabled")); ok {
		tmp := isPodSecurityPolicyEnabled.(bool)
		result.IsPodSecurityPolicyEnabled = &tmp
	}

	return result, nil
}

func AdmissionControllerOptionsToMap(obj *oci_containerengine.AdmissionControllerOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPodSecurityPolicyEnabled != nil {
		result["is_pod_security_policy_enabled"] = bool(*obj.IsPodSecurityPolicyEnabled)
	}

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToClusterCreateOptions(fieldKeyFormat string) (oci_containerengine.ClusterCreateOptions, error) {
	result := oci_containerengine.ClusterCreateOptions{}

	if addOns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "add_ons")); ok {
		if tmpList := addOns.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "add_ons"), 0)
			tmp, err := s.mapToAddOnOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert add_ons, encountered error: %v", err)
			}
			result.AddOns = &tmp
		}
	}

	if admissionControllerOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admission_controller_options")); ok {
		if tmpList := admissionControllerOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "admission_controller_options"), 0)
			tmp, err := s.mapToAdmissionControllerOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert admission_controller_options, encountered error: %v", err)
			}
			result.AdmissionControllerOptions = &tmp
		}
	}

	if kubernetesNetworkConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kubernetes_network_config")); ok {
		if tmpList := kubernetesNetworkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "kubernetes_network_config"), 0)
			tmp, err := s.mapToKubernetesNetworkConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert kubernetes_network_config, encountered error: %v", err)
			}
			result.KubernetesNetworkConfig = &tmp
		}
	}

	if persistentVolumeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "persistent_volume_config")); ok {
		if tmpList := persistentVolumeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "persistent_volume_config"), 0)
			tmp, err := s.mapToPersistentVolumeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert persistent_volume_config, encountered error: %v", err)
			}
			result.PersistentVolumeConfig = &tmp
		}
	}

	if serviceLbConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_lb_config")); ok {
		if tmpList := serviceLbConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "service_lb_config"), 0)
			tmp, err := s.mapToServiceLbConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert service_lb_config, encountered error: %v", err)
			}
			result.ServiceLbConfig = &tmp
		}
	}

	if serviceLbSubnetIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_lb_subnet_ids")); ok {
		interfaces := serviceLbSubnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "service_lb_subnet_ids")) {
			result.ServiceLbSubnetIds = tmp
		}
	}

	return result, nil
}

func ClusterCreateOptionsToMap(obj *oci_containerengine.ClusterCreateOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddOns != nil {
		result["add_ons"] = []interface{}{AddOnOptionsToMap(obj.AddOns)}
	}

	if obj.AdmissionControllerOptions != nil {
		result["admission_controller_options"] = []interface{}{AdmissionControllerOptionsToMap(obj.AdmissionControllerOptions)}
	}

	if obj.KubernetesNetworkConfig != nil {
		result["kubernetes_network_config"] = []interface{}{KubernetesNetworkConfigToMap(obj.KubernetesNetworkConfig)}
	}

	if obj.PersistentVolumeConfig != nil {
		result["persistent_volume_config"] = []interface{}{PersistentVolumeConfigDetailsToMap(obj.PersistentVolumeConfig)}
	}

	if obj.ServiceLbConfig != nil {
		result["service_lb_config"] = []interface{}{ServiceLbConfigDetailsToMap(obj.ServiceLbConfig)}
	}

	result["service_lb_subnet_ids"] = obj.ServiceLbSubnetIds

	return result
}

func ClusterEndpointsToMap(obj *oci_containerengine.ClusterEndpoints) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Kubernetes != nil {
		result["kubernetes"] = string(*obj.Kubernetes)
	}

	if obj.PrivateEndpoint != nil {
		result["private_endpoint"] = string(*obj.PrivateEndpoint)
	}

	if obj.PublicEndpoint != nil {
		result["public_endpoint"] = string(*obj.PublicEndpoint)
	}

	if obj.VcnHostnameEndpoint != nil {
		result["vcn_hostname_endpoint"] = string(*obj.VcnHostnameEndpoint)
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

	if obj.TimeCredentialExpiration != nil {
		result["time_credential_expiration"] = obj.TimeCredentialExpiration.String()
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

func (s *ContainerengineClusterResourceCrud) mapToClusterPodNetworkOptionDetails(fieldKeyFormat string) (oci_containerengine.ClusterPodNetworkOptionDetails, error) {
	var baseObject oci_containerengine.ClusterPodNetworkOptionDetails
	//discriminator
	cniTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cni_type"))
	var cniType string
	if ok {
		cniType = cniTypeRaw.(string)
	} else {
		cniType = "" // default value
	}
	switch strings.ToLower(cniType) {
	case strings.ToLower("FLANNEL_OVERLAY"):
		details := oci_containerengine.FlannelOverlayClusterPodNetworkOptionDetails{}
		baseObject = details
	case strings.ToLower("OCI_VCN_IP_NATIVE"):
		details := oci_containerengine.OciVcnIpNativeClusterPodNetworkOptionDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown cni_type '%v' was specified", cniType)
	}
	return baseObject, nil
}

func ClusterPodNetworkOptionDetailsToMap(obj oci_containerengine.ClusterPodNetworkOptionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch (obj).(type) {
	case oci_containerengine.FlannelOverlayClusterPodNetworkOptionDetails:
		result["cni_type"] = "FLANNEL_OVERLAY"
	case oci_containerengine.OciVcnIpNativeClusterPodNetworkOptionDetails:
		result["cni_type"] = "OCI_VCN_IP_NATIVE"
	default:
		log.Printf("[WARN] Received 'cni_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToUpdateClusterEndpointConfigDetails(fieldKeyFormat string) (oci_containerengine.UpdateClusterEndpointConfigDetails, error) {
	result := oci_containerengine.UpdateClusterEndpointConfigDetails{}
	if isPublicIpEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_public_ip_enabled")); ok {
		tmp := isPublicIpEnabled.(bool)
		result.IsPublicIpEnabled = &tmp
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
			result.NsgIds = tmp
		}
	}
	return result, nil
}

func (s *ContainerengineClusterResourceCrud) mapToMigrateClusterToNativeVCNDetails(fieldKeyFormat string) (oci_containerengine.ClusterMigrateToNativeVcnDetails, error) {
	result := oci_containerengine.ClusterMigrateToNativeVcnDetails{}
	endpointConfigDetails := oci_containerengine.ClusterEndpointConfig{}

	if isPublicIpEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_public_ip_enabled")); ok {
		tmp := isPublicIpEnabled.(bool)
		endpointConfigDetails.IsPublicIpEnabled = &tmp
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
			endpointConfigDetails.NsgIds = tmp
		}
	}
	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		endpointConfigDetails.SubnetId = &tmp
	}

	result.EndpointConfig = &endpointConfigDetails

	if DecommissionDelayDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "decommission_delay_duration")); ok {
		tmp := DecommissionDelayDuration.(string)
		result.DecommissionDelayDuration = &tmp
	}

	return result, nil
}

func (s *ContainerengineClusterResourceCrud) mapToCreateClusterEndpointConfigDetails(fieldKeyFormat string) (oci_containerengine.CreateClusterEndpointConfigDetails, error) {
	result := oci_containerengine.CreateClusterEndpointConfigDetails{}

	if isPublicIpEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_public_ip_enabled")); ok {
		tmp := isPublicIpEnabled.(bool)
		result.IsPublicIpEnabled = &tmp
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
			result.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func ClusterEndpointConfigToMap(obj *oci_containerengine.ClusterEndpointConfig, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPublicIpEnabled != nil {
		result["is_public_ip_enabled"] = bool(*obj.IsPublicIpEnabled)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToCreateImagePolicyConfigDetails(fieldKeyFormat string) (oci_containerengine.CreateImagePolicyConfigDetails, error) {
	result := oci_containerengine.CreateImagePolicyConfigDetails{}

	if isPolicyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_policy_enabled")); ok {
		tmp := isPolicyEnabled.(bool)
		result.IsPolicyEnabled = &tmp
	}

	if keyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_details")); ok {
		interfaces := keyDetails.([]interface{})
		tmp := make([]oci_containerengine.KeyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_details"), stateDataIndex)
			converted, err := s.mapToKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "key_details")) {
			result.KeyDetails = tmp
		}
	}

	return result, nil
}

func (s *ContainerengineClusterResourceCrud) mapToUpdateImagePolicyConfigDetails(fieldKeyFormat string) (oci_containerengine.UpdateImagePolicyConfigDetails, error) {
	result := oci_containerengine.UpdateImagePolicyConfigDetails{}

	if isPolicyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_policy_enabled")); ok {
		tmp := isPolicyEnabled.(bool)
		result.IsPolicyEnabled = &tmp
	}

	if keyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_details")); ok {
		interfaces := keyDetails.([]interface{})
		tmp := make([]oci_containerengine.KeyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_details"), stateDataIndex)
			converted, err := s.mapToKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "key_details")) {
			result.KeyDetails = tmp
		}
	}

	return result, nil
}

func ImagePolicyConfigToMap(obj *oci_containerengine.ImagePolicyConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPolicyEnabled != nil {
		result["is_policy_enabled"] = bool(*obj.IsPolicyEnabled)
	}

	keyDetails := []interface{}{}
	for _, item := range obj.KeyDetails {
		keyDetails = append(keyDetails, KeyDetailsToMap(item))
	}
	result["key_details"] = keyDetails

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToKeyDetails(fieldKeyFormat string) (oci_containerengine.KeyDetails, error) {
	result := oci_containerengine.KeyDetails{}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	return result, nil
}

func KeyDetailsToMap(obj oci_containerengine.KeyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToKubernetesNetworkConfig(fieldKeyFormat string) (oci_containerengine.KubernetesNetworkConfig, error) {
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

func (s *ContainerengineClusterResourceCrud) mapToPersistentVolumeConfigDetails(fieldKeyFormat string) (oci_containerengine.PersistentVolumeConfigDetails, error) {
	result := oci_containerengine.PersistentVolumeConfigDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func PersistentVolumeConfigDetailsToMap(obj *oci_containerengine.PersistentVolumeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToServiceLbConfigDetails(fieldKeyFormat string) (oci_containerengine.ServiceLbConfigDetails, error) {
	result := oci_containerengine.ServiceLbConfigDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func ServiceLbConfigDetailsToMap(obj *oci_containerengine.ServiceLbConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}

func (s *ContainerengineClusterResourceCrud) mapToUpdateClusterOptionsDetails(fieldKeyFormat string) (oci_containerengine.UpdateClusterOptionsDetails, error) {
	result := oci_containerengine.UpdateClusterOptionsDetails{}

	if admissionControllerOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admission_controller_options")); ok {
		if tmpList := admissionControllerOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "admission_controller_options"), 0)
			tmp, err := s.mapToAdmissionControllerOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert admission_controller_options, encountered error: %v", err)
			}
			result.AdmissionControllerOptions = &tmp
		}
	}

	if persistentVolumeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "persistent_volume_config")); ok {
		if tmpList := persistentVolumeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "persistent_volume_config"), 0)
			tmp, err := s.mapToPersistentVolumeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert persistent_volume_config, encountered error: %v", err)
			}
			result.PersistentVolumeConfig = &tmp
		}
	}

	if serviceLbConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_lb_config")); ok {
		if tmpList := serviceLbConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "service_lb_config"), 0)
			tmp, err := s.mapToServiceLbConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert service_lb_config, encountered error: %v", err)
			}
			result.ServiceLbConfig = &tmp
		}
	}

	return result, nil
}
