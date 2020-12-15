// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v31/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v31/ocvp"
)

func init() {
	RegisterResource("oci_ocvp_sddc", OcvpSddcResource())
}

func OcvpSddcResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("6h"),
		},
		Create: createOcvpSddc,
		Read:   readOcvpSddc,
		Update: updateOcvpSddc,
		Delete: deleteOcvpSddc,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"nsx_edge_uplink1vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nsx_edge_uplink2vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
				ForceNew: true,
			},
			"ssh_authorized_keys": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vmotion_vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vmware_software_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vsan_vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vsphere_vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
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
			"hcx_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_display_name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_hcx_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"workload_network_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"hcx_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hcx_initial_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hcx_on_prem_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hcx_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_edge_uplink_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_initial_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_overlay_segment_name": {
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
			"vcenter_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_initial_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"actual_esxi_hosts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).sddcClient()
	sync.WorkRequestClient = m.(*OracleClients).ocvpWorkRequestClient

	return CreateResource(d, sync)
}

func readOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).sddcClient()
	sync.WorkRequestClient = m.(*OracleClients).ocvpWorkRequestClient

	return ReadResource(sync)
}

func updateOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).sddcClient()
	sync.WorkRequestClient = m.(*OracleClients).ocvpWorkRequestClient

	return UpdateResource(d, sync)
}

func deleteOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).sddcClient()
	sync.WorkRequestClient = m.(*OracleClients).ocvpWorkRequestClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type OcvpSddcResourceCrud struct {
	BaseCrud
	Client                 *oci_ocvp.SddcClient
	WorkRequestClient      *oci_ocvp.WorkRequestClient
	Res                    *oci_ocvp.Sddc
	DisableNotFoundRetries bool
}

func (s *OcvpSddcResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpSddcResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpSddcResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpSddcResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpSddcResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpSddcResourceCrud) Create() error {
	request := oci_ocvp.CreateSddcRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeAvailabilityDomain, ok := s.D.GetOkExists("compute_availability_domain"); ok {
		tmp := computeAvailabilityDomain.(string)
		request.ComputeAvailabilityDomain = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hcxVlanId, ok := s.D.GetOkExists("hcx_vlan_id"); ok {
		tmp := hcxVlanId.(string)
		request.HcxVlanId = &tmp
	}

	if instanceDisplayNamePrefix, ok := s.D.GetOkExists("instance_display_name_prefix"); ok {
		tmp := instanceDisplayNamePrefix.(string)
		request.InstanceDisplayNamePrefix = &tmp
	}

	if isHcxEnabled, ok := s.D.GetOkExists("is_hcx_enabled"); ok {
		tmp := isHcxEnabled.(bool)
		request.IsHcxEnabled = &tmp
	}

	if nsxEdgeUplink1VlanId, ok := s.D.GetOkExists("nsx_edge_uplink1vlan_id"); ok {
		tmp := nsxEdgeUplink1VlanId.(string)
		request.NsxEdgeUplink1VlanId = &tmp
	}

	if nsxEdgeUplink2VlanId, ok := s.D.GetOkExists("nsx_edge_uplink2vlan_id"); ok {
		tmp := nsxEdgeUplink2VlanId.(string)
		request.NsxEdgeUplink2VlanId = &tmp
	}

	if nsxEdgeVTepVlanId, ok := s.D.GetOkExists("nsx_edge_vtep_vlan_id"); ok {
		tmp := nsxEdgeVTepVlanId.(string)
		request.NsxEdgeVTepVlanId = &tmp
	}

	if nsxVTepVlanId, ok := s.D.GetOkExists("nsx_vtep_vlan_id"); ok {
		tmp := nsxVTepVlanId.(string)
		request.NsxVTepVlanId = &tmp
	}

	if provisioningSubnetId, ok := s.D.GetOkExists("provisioning_subnet_id"); ok {
		tmp := provisioningSubnetId.(string)
		request.ProvisioningSubnetId = &tmp
	}

	if sshAuthorizedKeys, ok := s.D.GetOkExists("ssh_authorized_keys"); ok {
		tmp := sshAuthorizedKeys.(string)
		request.SshAuthorizedKeys = &tmp
	}

	if vmotionVlanId, ok := s.D.GetOkExists("vmotion_vlan_id"); ok {
		tmp := vmotionVlanId.(string)
		request.VmotionVlanId = &tmp
	}

	if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
		tmp := vmwareSoftwareVersion.(string)
		request.VmwareSoftwareVersion = &tmp
	}

	if vsanVlanId, ok := s.D.GetOkExists("vsan_vlan_id"); ok {
		tmp := vsanVlanId.(string)
		request.VsanVlanId = &tmp
	}

	if vsphereVlanId, ok := s.D.GetOkExists("vsphere_vlan_id"); ok {
		tmp := vsphereVlanId.(string)
		request.VsphereVlanId = &tmp
	}

	if workloadNetworkCidr, ok := s.D.GetOkExists("workload_network_cidr"); ok {
		tmp := workloadNetworkCidr.(string)
		request.WorkloadNetworkCidr = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateSddc(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSddcFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpSddcResourceCrud) getSddcFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	sddcId, err := sddcWaitForWorkRequest(workId, "sddc",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*sddcId)

	return s.Get()
}

func sddcWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func sddcWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = sddcWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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

	// The API Gateway workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func (s *OcvpSddcResourceCrud) Get() error {
	request := oci_ocvp.GetSddcRequest{}

	tmp := s.D.Id()
	request.SddcId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetSddc(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Sddc
	return nil
}

func (s *OcvpSddcResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ocvp.UpdateSddcRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hcxVlanId, ok := s.D.GetOkExists("hcx_vlan_id"); ok {
		tmp := hcxVlanId.(string)
		request.HcxVlanId = &tmp
	}

	if nsxEdgeUplink1VlanId, ok := s.D.GetOkExists("nsx_edge_uplink1vlan_id"); ok {
		tmp := nsxEdgeUplink1VlanId.(string)
		request.NsxEdgeUplink1VlanId = &tmp
	}

	if nsxEdgeUplink2VlanId, ok := s.D.GetOkExists("nsx_edge_uplink2vlan_id"); ok {
		tmp := nsxEdgeUplink2VlanId.(string)
		request.NsxEdgeUplink2VlanId = &tmp
	}

	if nsxEdgeVTepVlanId, ok := s.D.GetOkExists("nsx_edge_vtep_vlan_id"); ok {
		tmp := nsxEdgeVTepVlanId.(string)
		request.NsxEdgeVTepVlanId = &tmp
	}

	if nsxVTepVlanId, ok := s.D.GetOkExists("nsx_vtep_vlan_id"); ok {
		tmp := nsxVTepVlanId.(string)
		request.NsxVTepVlanId = &tmp
	}

	tmp := s.D.Id()
	request.SddcId = &tmp

	if sshAuthorizedKeys, ok := s.D.GetOkExists("ssh_authorized_keys"); ok {
		tmp := sshAuthorizedKeys.(string)
		request.SshAuthorizedKeys = &tmp
	}

	if vmotionVlanId, ok := s.D.GetOkExists("vmotion_vlan_id"); ok {
		tmp := vmotionVlanId.(string)
		request.VmotionVlanId = &tmp
	}

	if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
		tmp := vmwareSoftwareVersion.(string)
		request.VmwareSoftwareVersion = &tmp
	}

	if vsanVlanId, ok := s.D.GetOkExists("vsan_vlan_id"); ok {
		tmp := vsanVlanId.(string)
		request.VsanVlanId = &tmp
	}

	if vsphereVlanId, ok := s.D.GetOkExists("vsphere_vlan_id"); ok {
		tmp := vsphereVlanId.(string)
		request.VsphereVlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateSddc(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Sddc
	return nil
}

func (s *OcvpSddcResourceCrud) Delete() error {
	request := oci_ocvp.DeleteSddcRequest{}

	tmp := s.D.Id()
	request.SddcId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteSddc(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := sddcWaitForWorkRequest(workId, "sddc",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpSddcResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	// We update value of esxi_hosts_count in state file only if the esxi_hosts_count of the
	// SDDC is modified in the TF config by the user.
	// As there could a scenario where the SDDC esxi_hosts_count on the cloud could be different as esxi host can be attached to the SDDC
	// Then we do not update the size but instead update the actual_esxi_hosts_count in the state file.
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HcxFqdn != nil {
		s.D.Set("hcx_fqdn", *s.Res.HcxFqdn)
	}

	if s.Res.HcxInitialPassword != nil {
		s.D.Set("hcx_initial_password", *s.Res.HcxInitialPassword)
	}

	if s.Res.HcxOnPremKey != nil {
		s.D.Set("hcx_on_prem_key", *s.Res.HcxOnPremKey)
	}

	if s.Res.HcxPrivateIpId != nil {
		s.D.Set("hcx_private_ip_id", *s.Res.HcxPrivateIpId)
	}

	if s.Res.HcxVlanId != nil {
		s.D.Set("hcx_vlan_id", *s.Res.HcxVlanId)
	}

	if s.Res.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", *s.Res.InstanceDisplayNamePrefix)
	}

	if s.Res.IsHcxEnabled != nil {
		s.D.Set("is_hcx_enabled", *s.Res.IsHcxEnabled)
	}

	if s.Res.NsxEdgeUplink1VlanId != nil {
		s.D.Set("nsx_edge_uplink1vlan_id", *s.Res.NsxEdgeUplink1VlanId)
	}

	if s.Res.NsxEdgeUplink2VlanId != nil {
		s.D.Set("nsx_edge_uplink2vlan_id", *s.Res.NsxEdgeUplink2VlanId)
	}

	if s.Res.NsxEdgeUplinkIpId != nil {
		s.D.Set("nsx_edge_uplink_ip_id", *s.Res.NsxEdgeUplinkIpId)
	}

	if s.Res.NsxEdgeVTepVlanId != nil {
		s.D.Set("nsx_edge_vtep_vlan_id", *s.Res.NsxEdgeVTepVlanId)
	}

	if s.Res.NsxManagerFqdn != nil {
		s.D.Set("nsx_manager_fqdn", *s.Res.NsxManagerFqdn)
	}

	if s.Res.NsxManagerInitialPassword != nil {
		s.D.Set("nsx_manager_initial_password", *s.Res.NsxManagerInitialPassword)
	}

	if s.Res.NsxManagerPrivateIpId != nil {
		s.D.Set("nsx_manager_private_ip_id", *s.Res.NsxManagerPrivateIpId)
	}

	if s.Res.NsxManagerUsername != nil {
		s.D.Set("nsx_manager_username", *s.Res.NsxManagerUsername)
	}

	if s.Res.NsxOverlaySegmentName != nil {
		s.D.Set("nsx_overlay_segment_name", *s.Res.NsxOverlaySegmentName)
	}

	if s.Res.NsxVTepVlanId != nil {
		s.D.Set("nsx_vtep_vlan_id", *s.Res.NsxVTepVlanId)
	}

	if s.Res.ProvisioningSubnetId != nil {
		s.D.Set("provisioning_subnet_id", *s.Res.ProvisioningSubnetId)
	}

	if s.Res.SshAuthorizedKeys != nil {
		s.D.Set("ssh_authorized_keys", *s.Res.SshAuthorizedKeys)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcenterFqdn != nil {
		s.D.Set("vcenter_fqdn", *s.Res.VcenterFqdn)
	}

	if s.Res.VcenterInitialPassword != nil {
		s.D.Set("vcenter_initial_password", *s.Res.VcenterInitialPassword)
	}

	if s.Res.VcenterPrivateIpId != nil {
		s.D.Set("vcenter_private_ip_id", *s.Res.VcenterPrivateIpId)
	}

	if s.Res.VcenterUsername != nil {
		s.D.Set("vcenter_username", *s.Res.VcenterUsername)
	}

	if s.Res.VmotionVlanId != nil {
		s.D.Set("vmotion_vlan_id", *s.Res.VmotionVlanId)
	}

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	if s.Res.VsanVlanId != nil {
		s.D.Set("vsan_vlan_id", *s.Res.VsanVlanId)
	}

	if s.Res.VsphereVlanId != nil {
		s.D.Set("vsphere_vlan_id", *s.Res.VsphereVlanId)
	}

	if s.Res.WorkloadNetworkCidr != nil {
		s.D.Set("workload_network_cidr", *s.Res.WorkloadNetworkCidr)
	}

	return nil
}

func SddcSummaryToMap(obj oci_ocvp.SddcSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeAvailabilityDomain != nil {
		result["compute_availability_domain"] = string(*obj.ComputeAvailabilityDomain)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EsxiHostsCount != nil {
		result["esxi_hosts_count"] = int(*obj.EsxiHostsCount)
		result["actual_esxi_hosts_count"] = int(*obj.EsxiHostsCount)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HcxFqdn != nil {
		result["hcx_fqdn"] = string(*obj.HcxFqdn)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsHcxEnabled != nil {
		result["is_hcx_enabled"] = bool(*obj.IsHcxEnabled)
	}

	if obj.NsxManagerFqdn != nil {
		result["nsx_manager_fqdn"] = string(*obj.NsxManagerFqdn)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VcenterFqdn != nil {
		result["vcenter_fqdn"] = string(*obj.VcenterFqdn)
	}

	if obj.VmwareSoftwareVersion != nil {
		result["vmware_software_version"] = string(*obj.VmwareSoftwareVersion)
	}

	return result
}

func (s *OcvpSddcResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ocvp.ChangeSddcCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SddcId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	_, err := s.Client.ChangeSddcCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
