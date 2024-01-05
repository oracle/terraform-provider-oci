// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOcvpCluster,
		Read:     readOcvpCluster,
		Update:   updateOcvpCluster,
		Delete:   deleteOcvpCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compute_availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"esxi_hosts_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"network_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"nsx_edge_vtep_vlan_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"nsx_vtep_vlan_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"provisioning_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vmotion_vlan_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vsan_vlan_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"hcx_vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nsx_edge_uplink1vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nsx_edge_uplink2vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"provisioning_vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"replication_vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vsphere_vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"sddc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"datastores": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"block_volume_ids": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"datastore_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
						"capacity": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esxi_software_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"initial_commitment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"initial_host_ocpu_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"initial_host_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_display_name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_shielded_instance_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vmware_software_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workload_network_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"actual_esxi_hosts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compartment_id": {
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
			"upgrade_licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"license_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"vsphere_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsphere_upgrade_objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"download_link": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"link_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createOcvpCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.ClusterClient
	Res                    *oci_ocvp.Cluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpClusterResourceCrud) Create() error {
	request := oci_ocvp.CreateClusterRequest{}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if computeAvailabilityDomain, ok := s.D.GetOkExists("compute_availability_domain"); ok {
		tmp := computeAvailabilityDomain.(string)
		request.ComputeAvailabilityDomain = &tmp
	}

	if datastores, ok := s.D.GetOkExists("datastores"); ok {
		interfaces := datastores.([]interface{})
		tmp := make([]oci_ocvp.DatastoreInfo, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "datastores", stateDataIndex)
			converted, err := s.mapToDatastoreInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("datastores") {
			request.Datastores = tmp
		}
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

	if esxiHostsCount, ok := s.D.GetOkExists("esxi_hosts_count"); ok {
		tmp := esxiHostsCount.(int)
		request.EsxiHostsCount = &tmp
	}

	if esxiSoftwareVersion, ok := s.D.GetOkExists("esxi_software_version"); ok {
		tmp := esxiSoftwareVersion.(string)
		request.EsxiSoftwareVersion = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if initialCommitment, ok := s.D.GetOkExists("initial_commitment"); ok {
		request.InitialCommitment = oci_ocvp.CommitmentEnum(initialCommitment.(string))
	}

	if initialHostOcpuCount, ok := s.D.GetOkExists("initial_host_ocpu_count"); ok {
		tmp := float32(initialHostOcpuCount.(float64))
		request.InitialHostOcpuCount = &tmp
	}

	if initialHostShapeName, ok := s.D.GetOkExists("initial_host_shape_name"); ok {
		tmp := initialHostShapeName.(string)
		request.InitialHostShapeName = &tmp
	}

	if instanceDisplayNamePrefix, ok := s.D.GetOkExists("instance_display_name_prefix"); ok {
		tmp := instanceDisplayNamePrefix.(string)
		request.InstanceDisplayNamePrefix = &tmp
	}

	if isShieldedInstanceEnabled, ok := s.D.GetOkExists("is_shielded_instance_enabled"); ok {
		tmp := isShieldedInstanceEnabled.(bool)
		request.IsShieldedInstanceEnabled = &tmp
	}

	if networkConfiguration, ok := s.D.GetOkExists("network_configuration"); ok {
		if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_configuration", 0)
			tmp, err := s.mapToNetworkConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfiguration = &tmp
		}
	}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
		tmp := vmwareSoftwareVersion.(string)
		request.VmwareSoftwareVersion = &tmp
	}

	if workloadNetworkCidr, ok := s.D.GetOkExists("workload_network_cidr"); ok {
		tmp := workloadNetworkCidr.(string)
		request.WorkloadNetworkCidr = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_ocvp.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_ocvp.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "cluster") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpClusterResourceCrud) getClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	clusterId, err := clusterWaitForWorkRequest(workId, "cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

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
		if tfresource.ShouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func clusterWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = clusterWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ocvp.OperationStatusInProgress),
			string(oci_ocvp.OperationStatusAccepted),
			string(oci_ocvp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ocvp.OperationStatusSucceeded),
			string(oci_ocvp.OperationStatusFailed),
			string(oci_ocvp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ocvp.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpClusterWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ocvp.ListWorkRequestErrorsRequest{
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

func (s *OcvpClusterResourceCrud) Get() error {
	request := oci_ocvp.GetClusterRequest{}

	tmp := s.D.Id()
	request.ClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Cluster
	return nil
}

func (s *OcvpClusterResourceCrud) Update() error {
	request := oci_ocvp.UpdateClusterRequest{}

	tmp := s.D.Id()
	request.ClusterId = &tmp

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

	if esxiSoftwareVersion, ok := s.D.GetOkExists("esxi_software_version"); ok {
		tmp := esxiSoftwareVersion.(string)
		request.EsxiSoftwareVersion = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if networkConfiguration, ok := s.D.GetOkExists("network_configuration"); ok {
		if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_configuration", 0)
			tmp, err := s.mapToNetworkConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfiguration = &tmp
		}
	}

	if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
		tmp := vmwareSoftwareVersion.(string)
		request.VmwareSoftwareVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Cluster
	return nil
}

func (s *OcvpClusterResourceCrud) Delete() error {
	request := oci_ocvp.DeleteClusterRequest{}

	tmp := s.D.Id()
	request.ClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := clusterWaitForWorkRequest(workId, "cluster",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpClusterResourceCrud) SetData() error {
	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	datastores := []interface{}{}
	for _, item := range s.Res.Datastores {
		datastores = append(datastores, DatastoreDetailsToMap(item))
	}
	s.D.Set("datastores", datastores)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	// We Update value of esxi_hosts_count in state file only if the esxi_hosts_count of the
	// SDDC is modified in the TF config by the user.
	// As there could a scenario where the SDDC esxi_hosts_count on the cloud could be different as esxi host can be attached to the SDDC
	// Then we do not Update the size but instead Update the actual_esxi_hosts_count in the state file.
	if s.Res.EsxiHostsCount != nil {
		_, ok := s.D.GetOk("esxi_hosts_count") // This checks if size is in the state or not. If not and size in response is not nil it could be that user is importing and hence we need to updated the size
		if !ok {
			log.Printf("[DEBUG] esxi_hosts_count does not exists in state, hence assuming user is importing resource")
		}
		if s.D.HasChange("esxi_hosts_count") || !ok {
			oldValue, newValue := s.D.GetChange("esxi_hosts_count")
			log.Printf("[DEBUG] esxi_hosts_count has been updated in config from %v to %v", oldValue, newValue)
			s.D.Set("esxi_hosts_count", *s.Res.EsxiHostsCount)
		}
		s.D.Set("actual_esxi_hosts_count", *s.Res.EsxiHostsCount)
	}

	if s.Res.EsxiSoftwareVersion != nil {
		s.D.Set("esxi_software_version", *s.Res.EsxiSoftwareVersion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("initial_commitment", s.Res.InitialCommitment)

	if s.Res.InitialHostOcpuCount != nil {
		s.D.Set("initial_host_ocpu_count", *s.Res.InitialHostOcpuCount)
	}

	if s.Res.InitialHostShapeName != nil {
		s.D.Set("initial_host_shape_name", *s.Res.InitialHostShapeName)
	}

	if s.Res.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", *s.Res.InstanceDisplayNamePrefix)
	}

	if s.Res.IsShieldedInstanceEnabled != nil {
		s.D.Set("is_shielded_instance_enabled", *s.Res.IsShieldedInstanceEnabled)
	}

	if s.Res.NetworkConfiguration != nil {
		s.D.Set("network_configuration", []interface{}{NetworkConfigurationToMap(s.Res.NetworkConfiguration)})
	} else {
		s.D.Set("network_configuration", nil)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	upgradeLicenses := []interface{}{}
	for _, item := range s.Res.UpgradeLicenses {
		upgradeLicenses = append(upgradeLicenses, VsphereLicenseToMap(item))
	}
	s.D.Set("upgrade_licenses", upgradeLicenses)

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	s.D.Set("vsphere_type", s.Res.VsphereType)

	vsphereUpgradeObjects := []interface{}{}
	for _, item := range s.Res.VsphereUpgradeObjects {
		vsphereUpgradeObjects = append(vsphereUpgradeObjects, VsphereUpgradeObjectToMap(item))
	}
	s.D.Set("vsphere_upgrade_objects", vsphereUpgradeObjects)

	if s.Res.WorkloadNetworkCidr != nil {
		s.D.Set("workload_network_cidr", *s.Res.WorkloadNetworkCidr)
	}

	return nil
}

func ClusterSummaryToMap(obj oci_ocvp.ClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeAvailabilityDomain != nil {
		result["compute_availability_domain"] = string(*obj.ComputeAvailabilityDomain)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EsxiHostsCount != nil {
		result["esxi_hosts_count"] = int(*obj.EsxiHostsCount)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InitialHostOcpuCount != nil {
		result["initial_host_ocpu_count"] = float32(*obj.InitialHostOcpuCount)
	}

	if obj.InitialHostShapeName != nil {
		result["initial_host_shape_name"] = string(*obj.InitialHostShapeName)
	}

	if obj.IsShieldedInstanceEnabled != nil {
		result["is_shielded_instance_enabled"] = bool(*obj.IsShieldedInstanceEnabled)
	}

	if obj.SddcId != nil {
		result["sddc_id"] = string(*obj.SddcId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VmwareSoftwareVersion != nil {
		result["vmware_software_version"] = string(*obj.VmwareSoftwareVersion)
	}

	result["vsphere_type"] = string(obj.VsphereType)

	return result
}

func (s *OcvpClusterResourceCrud) mapToDatastoreInfo(fieldKeyFormat string) (oci_ocvp.DatastoreInfo, error) {
	result := oci_ocvp.DatastoreInfo{}

	if blockVolumeIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_ids")); ok {
		interfaces := blockVolumeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volume_ids")) {
			result.BlockVolumeIds = tmp
		}
	}

	if datastoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "datastore_type")); ok {
		result.DatastoreType = oci_ocvp.DatastoreTypesEnum(datastoreType.(string))
	}

	return result, nil
}

func DatastoreDetailsToMap(obj oci_ocvp.DatastoreDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["block_volume_ids"] = obj.BlockVolumeIds
	result["block_volume_ids"] = obj.BlockVolumeIds

	if obj.Capacity != nil {
		result["capacity"] = float64(*obj.Capacity)
	}

	result["datastore_type"] = string(obj.DatastoreType)

	return result
}

func (s *OcvpClusterResourceCrud) mapToNetworkConfiguration(fieldKeyFormat string) (oci_ocvp.NetworkConfiguration, error) {
	result := oci_ocvp.NetworkConfiguration{}

	if hcxVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hcx_vlan_id")); ok {
		tmp := hcxVlanId.(string)
		result.HcxVlanId = &tmp
	}

	if nsxEdgeUplink1VlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsx_edge_uplink1vlan_id")); ok {
		tmp := nsxEdgeUplink1VlanId.(string)
		result.NsxEdgeUplink1VlanId = &tmp
	}

	if nsxEdgeUplink2VlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsx_edge_uplink2vlan_id")); ok {
		tmp := nsxEdgeUplink2VlanId.(string)
		result.NsxEdgeUplink2VlanId = &tmp
	}

	if nsxEdgeVTepVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsx_edge_vtep_vlan_id")); ok {
		tmp := nsxEdgeVTepVlanId.(string)
		result.NsxEdgeVTepVlanId = &tmp
	}

	if nsxVTepVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsx_vtep_vlan_id")); ok {
		tmp := nsxVTepVlanId.(string)
		result.NsxVTepVlanId = &tmp
	}

	if provisioningSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "provisioning_subnet_id")); ok {
		tmp := provisioningSubnetId.(string)
		result.ProvisioningSubnetId = &tmp
	}

	if provisioningVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "provisioning_vlan_id")); ok {
		tmp := provisioningVlanId.(string)
		result.ProvisioningVlanId = &tmp
	}

	if replicationVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replication_vlan_id")); ok {
		tmp := replicationVlanId.(string)
		result.ReplicationVlanId = &tmp
	}

	if vmotionVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vmotion_vlan_id")); ok {
		tmp := vmotionVlanId.(string)
		result.VmotionVlanId = &tmp
	}

	if vsanVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vsan_vlan_id")); ok {
		tmp := vsanVlanId.(string)
		result.VsanVlanId = &tmp
	}

	if vsphereVlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vsphere_vlan_id")); ok {
		tmp := vsphereVlanId.(string)
		result.VsphereVlanId = &tmp
	}

	return result, nil
}

func NetworkConfigurationToMap(obj *oci_ocvp.NetworkConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HcxVlanId != nil {
		result["hcx_vlan_id"] = string(*obj.HcxVlanId)
	}

	if obj.NsxEdgeUplink1VlanId != nil {
		result["nsx_edge_uplink1vlan_id"] = string(*obj.NsxEdgeUplink1VlanId)
	}

	if obj.NsxEdgeUplink2VlanId != nil {
		result["nsx_edge_uplink2vlan_id"] = string(*obj.NsxEdgeUplink2VlanId)
	}

	if obj.NsxEdgeVTepVlanId != nil {
		result["nsx_edge_vtep_vlan_id"] = string(*obj.NsxEdgeVTepVlanId)
	}

	if obj.NsxVTepVlanId != nil {
		result["nsx_vtep_vlan_id"] = string(*obj.NsxVTepVlanId)
	}

	if obj.ProvisioningSubnetId != nil {
		result["provisioning_subnet_id"] = string(*obj.ProvisioningSubnetId)
	}

	if obj.ProvisioningVlanId != nil {
		result["provisioning_vlan_id"] = string(*obj.ProvisioningVlanId)
	}

	if obj.ReplicationVlanId != nil {
		result["replication_vlan_id"] = string(*obj.ReplicationVlanId)
	}

	if obj.VmotionVlanId != nil {
		result["vmotion_vlan_id"] = string(*obj.VmotionVlanId)
	}

	if obj.VsanVlanId != nil {
		result["vsan_vlan_id"] = string(*obj.VsanVlanId)
	}

	if obj.VsphereVlanId != nil {
		result["vsphere_vlan_id"] = string(*obj.VsphereVlanId)
	}

	return result
}

func VsphereLicenseToMap(obj oci_ocvp.VsphereLicense) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LicenseKey != nil {
		result["license_key"] = string(*obj.LicenseKey)
	}

	if obj.LicenseType != nil {
		result["license_type"] = string(*obj.LicenseType)
	}

	return result
}

func VsphereUpgradeObjectToMap(obj oci_ocvp.VsphereUpgradeObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DownloadLink != nil {
		result["download_link"] = string(*obj.DownloadLink)
	}

	if obj.LinkDescription != nil {
		result["link_description"] = string(*obj.LinkDescription)
	}

	return result
}
