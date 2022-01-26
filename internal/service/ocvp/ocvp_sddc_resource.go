// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v56/ocvp"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func OcvpSddcResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("6h"),
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
			"hcx_action": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					UpgradeHcxAction,
					DowngradeHcxAction,
					CancelDowngradeHcxAction,
				}, true),
			},
			"initial_sku": {
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
			"is_hcx_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"provisioning_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refresh_hcx_license_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"replication_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserving_hcx_on_premise_license_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"hcx_on_prem_licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"activation_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hcx_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_hcx_enterprise_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_hcx_pending_downgrade": {
				Type:     schema.TypeBool,
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
			"time_hcx_billing_cycle_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_hcx_license_status_updated": {
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

const (
	UpgradeHcxAction         string = "UPGRADE"
	DowngradeHcxAction       string = "DOWNGRADE"
	CancelDowngradeHcxAction string = "CANCEL_DOWNGRADE"
)

func createOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpSddcResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.SddcClient
	Res                    *oci_ocvp.Sddc
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hcxVlanId, ok := s.D.GetOkExists("hcx_vlan_id"); ok {
		tmp := hcxVlanId.(string)
		request.HcxVlanId = &tmp
	}

	if initialSku, ok := s.D.GetOkExists("initial_sku"); ok {
		request.InitialSku = oci_ocvp.SkuEnum(initialSku.(string))
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

	if provisioningVlanId, ok := s.D.GetOkExists("provisioning_vlan_id"); ok {
		tmp := provisioningVlanId.(string)
		request.ProvisioningVlanId = &tmp
	}

	if replicationVlanId, ok := s.D.GetOkExists("replication_vlan_id"); ok {
		tmp := replicationVlanId.(string)
		request.ReplicationVlanId = &tmp
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

	if _, ok := s.D.GetOk("reserving_hcx_on_premise_license_keys"); ok {
		return fmt.Errorf("reserving_hcx_on_premise_license_keys should not be provided during SDDC creation.")
	}

	if hcxAction, ok := s.D.GetOk("hcx_action"); ok {
		hcxAction = strings.ToUpper(hcxAction.(string))
		if hcxAction == UpgradeHcxAction {
			_tmp := true
			request.IsHcxEnterpriseEnabled = &_tmp
		} else {
			return fmt.Errorf("hcx_action '%s' is not supported during SDDC creation. ", hcxAction)
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateSddc(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	creationError := s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))

	if creationError != nil {
		return creationError
	} else if hcxAction, ok := s.D.GetOk("hcx_action"); ok {
		s.D.Set("hcx_action", hcxAction)
	}

	if refresh, ok := s.D.GetOk("refresh_hcx_license_status"); ok {
		tmp := s.D.Id()
		return s.refreshHcxLicenseStatus(&tmp, refresh)
	}

	return nil
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

func sddcWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = sddcWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOcvpSddcWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpSddcWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
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

func (s *OcvpSddcResourceCrud) Get() error {
	request := oci_ocvp.GetSddcRequest{}

	tmp := s.D.Id()
	request.SddcId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if provisioningVlanId, ok := s.D.GetOkExists("provisioning_vlan_id"); ok {
		tmp := provisioningVlanId.(string)
		request.ProvisioningVlanId = &tmp
	}

	if replicationVlanId, ok := s.D.GetOkExists("replication_vlan_id"); ok {
		tmp := replicationVlanId.(string)
		request.ReplicationVlanId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateSddc(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.Sddc

	if _, exists := s.D.GetOk("reserving_hcx_on_premise_license_keys"); exists {
		if hcxAction, ok := s.D.GetOk("hcx_action"); !ok || strings.ToUpper(hcxAction.(string)) != DowngradeHcxAction {
			return fmt.Errorf("reserving_hcx_on_premise_license_keys can only be set when hcx_action is DOWNGRADE")
		}
	}

	if action, exists := s.D.GetOk("hcx_action"); exists && strings.ToUpper(action.(string)) == DowngradeHcxAction {
		if _, ok := s.D.GetOk("reserving_hcx_on_premise_license_keys"); !ok {
			return fmt.Errorf("reserving_hcx_on_premise_license_keys must exist when hcx_action is DOWNGRADE")
		} else if s.D.HasChange("reserving_hcx_on_premise_license_keys") && !s.D.HasChange("hcx_action") {
			return fmt.Errorf("reserving_hcx_on_premise_license_keys cannot be changed when hcx_action is already DOWNGRADE")
		}
	}

	var updateHcxError error

	if hcxAction, ok := s.D.GetOk("hcx_action"); ok && s.D.HasChange("hcx_action") {
		action := strings.ToUpper(hcxAction.(string))
		sddcId := s.D.Id()

		if action == UpgradeHcxAction {
			hcxRequest := oci_ocvp.UpgradeHcxRequest{}
			hcxRequest.SddcId = &sddcId
			hcxRes, hcxErr := s.Client.UpgradeHcx(context.Background(), hcxRequest)
			if hcxErr != nil {
				return hcxErr
			}
			workId := hcxRes.OpcWorkRequestId
			updateHcxError = s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if updateHcxError == nil {
				s.D.Set("hcx_action", hcxAction)
			}
		} else if action == DowngradeHcxAction {
			hcxRequest := oci_ocvp.DowngradeHcxRequest{}
			hcxRequest.SddcId = &sddcId
			if reservingKeys, ok := s.D.GetOk("reserving_hcx_on_premise_license_keys"); ok {
				var keys []string
				for _, key := range reservingKeys.([]interface{}) {
					keys = append(keys, key.(string))
				}
				hcxRequest.ReservingHcxOnPremiseLicenseKeys = keys
			}
			hcxRes, hcxErr := s.Client.DowngradeHcx(context.Background(), hcxRequest)
			if hcxErr != nil {
				return hcxErr
			}
			workId := hcxRes.OpcWorkRequestId
			updateHcxError = s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if updateHcxError == nil {
				s.D.Set("hcx_action", hcxAction)
			}
		} else if action == CancelDowngradeHcxAction {
			hcxRequest := oci_ocvp.CancelDowngradeHcxRequest{}
			hcxRequest.SddcId = &sddcId
			hcxRes, hcxErr := s.Client.CancelDowngradeHcx(context.Background(), hcxRequest)
			if hcxErr != nil {
				return hcxErr
			}
			workId := hcxRes.OpcWorkRequestId
			updateHcxError = s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if updateHcxError == nil {
				s.D.Set("hcx_action", hcxAction)
			}
		} else {
			return fmt.Errorf("hcx_action '%s' is not supported. ", hcxAction)
		}
	}

	if updateHcxError != nil {
		return updateHcxError
	}

	if refresh, ok := s.D.GetOk("refresh_hcx_license_status"); ok && s.D.HasChange("refresh_hcx_license_status") {
		tmp := s.D.Id()
		return s.refreshHcxLicenseStatus(&tmp, refresh)
	}

	return nil
}

func (s *OcvpSddcResourceCrud) refreshHcxLicenseStatus(sddcId *string, refresh interface{}) error {
	hcxRequest := oci_ocvp.RefreshHcxLicenseStatusRequest{}
	hcxRequest.SddcId = sddcId
	hcxRes, hcxErr := s.Client.RefreshHcxLicenseStatus(context.Background(), hcxRequest)
	if hcxErr != nil {
		return hcxErr
	}
	workId := hcxRes.OpcWorkRequestId
	err := s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err == nil {
		s.D.Set("refresh_hcx_license_status", refresh)
	}
	return err
}

func (s *OcvpSddcResourceCrud) Delete() error {
	request := oci_ocvp.DeleteSddcRequest{}

	tmp := s.D.Id()
	request.SddcId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

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

	hcxOnPremLicenses := []interface{}{}
	for _, item := range s.Res.HcxOnPremLicenses {
		hcxOnPremLicenses = append(hcxOnPremLicenses, HcxLicenseSummaryToMap(item))
	}
	s.D.Set("hcx_on_prem_licenses", hcxOnPremLicenses)

	if s.Res.HcxPrivateIpId != nil {
		s.D.Set("hcx_private_ip_id", *s.Res.HcxPrivateIpId)
	}

	if s.Res.HcxVlanId != nil {
		s.D.Set("hcx_vlan_id", *s.Res.HcxVlanId)
	}

	s.D.Set("initial_sku", s.Res.InitialSku)

	if s.Res.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", *s.Res.InstanceDisplayNamePrefix)
	}

	if s.Res.IsHcxEnabled != nil {
		s.D.Set("is_hcx_enabled", *s.Res.IsHcxEnabled)
	}

	if s.Res.IsHcxEnterpriseEnabled != nil {
		s.D.Set("is_hcx_enterprise_enabled", *s.Res.IsHcxEnterpriseEnabled)
	}

	if s.Res.IsHcxPendingDowngrade != nil {
		s.D.Set("is_hcx_pending_downgrade", *s.Res.IsHcxPendingDowngrade)
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

	if s.Res.ProvisioningVlanId != nil {
		s.D.Set("provisioning_vlan_id", *s.Res.ProvisioningVlanId)
	}

	if s.Res.ReplicationVlanId != nil {
		s.D.Set("replication_vlan_id", *s.Res.ReplicationVlanId)
	}

	if s.Res.SshAuthorizedKeys != nil {
		s.D.Set("ssh_authorized_keys", *s.Res.SshAuthorizedKeys)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeHcxBillingCycleEnd != nil {
		s.D.Set("time_hcx_billing_cycle_end", s.Res.TimeHcxBillingCycleEnd.String())
	}

	if s.Res.TimeHcxLicenseStatusUpdated != nil {
		s.D.Set("time_hcx_license_status_updated", s.Res.TimeHcxLicenseStatusUpdated.String())
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

func HcxLicenseSummaryToMap(obj oci_ocvp.HcxLicenseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActivationKey != nil {
		result["activation_key"] = string(*obj.ActivationKey)
	}

	result["status"] = string(obj.Status)

	if obj.SystemName != nil {
		result["system_name"] = string(*obj.SystemName)
	}

	return result
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
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
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

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	_, err := s.Client.ChangeSddcCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
