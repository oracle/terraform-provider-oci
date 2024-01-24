// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpEsxiHostResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createOcvpEsxiHost,
		Read:   readOcvpEsxiHost,
		Update: updateOcvpEsxiHost,
		Delete: deleteOcvpEsxiHost,
		Schema: map[string]*schema.Schema{
			// Required

			// Optional
			"cluster_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"sddc_id", "current_sku", "failed_esxi_host_id", "next_sku", "non_upgraded_esxi_host_id"},
			},
			"sddc_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cluster_id", "esxi_software_version"},
				// sddc_id is being changed to compute only so need to suppress diff if sddc_id is removed from config
				DiffSuppressFunc: suppressEsxiHostDeprecatedFieldRemoval,
				Deprecated:       tfresource.FieldDeprecatedForAnother("sddc_id", "cluster_id"),
			},
			"billing_donor_host_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ConflictsWith:    []string{"cluster_id", "esxi_software_version"},
				DiffSuppressFunc: suppressEsxiHostDeprecatedFieldRemoval,
				Deprecated:       "This 'billing_donor_host_id' argument has been deprecated and will be computed only.",
			},
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compute_availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"current_sku": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				// API may update current_sku in the backend, so need to suppress the diff if any change is made to current_sku after esxi_host creation
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "" && new != "" {
						return false
					}
					return true
				},
				ConflictsWith: []string{"cluster_id", "esxi_software_version"},
				Deprecated:    tfresource.FieldDeprecated("current_sku"),
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
				ForceNew: true,
			},
			"failed_esxi_host_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cluster_id", "esxi_software_version"},
				Deprecated:    "This 'failed_esxi_host_id' argument has been deprecated and will be computed only.",
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"host_ocpu_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"host_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"next_sku": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"cluster_id", "esxi_software_version"},
				Deprecated:    tfresource.FieldDeprecated("next_sku"),
			},
			"non_upgraded_esxi_host_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cluster_id", "esxi_software_version"},
				// sddc_id is being changed to compute only so need to suppress diff if sddc_id is removed from config
				DiffSuppressFunc: suppressEsxiHostDeprecatedFieldRemoval,
				Deprecated:       "This 'non_upgraded_esxi_host_id' argument has been deprecated and will be computed only.",
			},

			// Computed
			"billing_contract_end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_commitment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grace_period_end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_billing_continuation_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_billing_swapping_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"next_commitment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replacement_esxi_host_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"swap_billing_host_id": {
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
			"upgraded_replacement_esxi_host_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmware_software_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func suppressEsxiHostDeprecatedFieldRemoval(k, old string, new string, d *schema.ResourceData) bool {
	// suppress diff when resource is using new fields and old fields are removed
	if _, ok := d.GetOkExists("cluster_id"); ok && old != "" && new == "" {
		return true
	}
	return false
}

func createOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()
	sync.ClusterClient = m.(*client.OracleClients).ClusterClient()
	sync.SddcClient = m.(*client.OracleClients).SddcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	return nil

}

func readOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpEsxiHostResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.EsxiHostClient
	Res                    *oci_ocvp.EsxiHost
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
	ClusterClient          *oci_ocvp.ClusterClient
	SddcClient             *oci_ocvp.SddcClient
}

func (s *OcvpEsxiHostResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpEsxiHostResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpEsxiHostResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpEsxiHostResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpEsxiHostResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpEsxiHostResourceCrud) Create() error {
	request := oci_ocvp.CreateEsxiHostRequest{}

	// replace failed ESXi host
	if failedEsxiHostId, ok := s.D.GetOkExists("failed_esxi_host_id"); ok {
		return s.ReplaceHost(failedEsxiHostId.(string))
	}

	if billingDonorHostId, ok := s.D.GetOkExists("billing_donor_host_id"); ok {
		tmp := billingDonorHostId.(string)
		request.BillingDonorHostId = &tmp
	}
	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	} else if _, ok := s.D.GetOkExists("sddc_id"); !ok {
		return fmt.Errorf("one of cluster_id or sddc_id must be configured")
	}

	if computeAvailabilityDomain, ok := s.D.GetOkExists("compute_availability_domain"); ok {
		tmp := computeAvailabilityDomain.(string)
		request.ComputeAvailabilityDomain = &tmp
	}

	if currentCommitment, ok := s.D.GetOkExists("current_commitment"); ok {
		request.CurrentCommitment = oci_ocvp.CommitmentEnum(currentCommitment.(string))
	}

	if currentSku, ok := s.D.GetOkExists("current_sku"); ok {
		request.CurrentCommitment = oci_ocvp.CommitmentEnum(currentSku.(string))
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

	if esxiSoftwareVersion, ok := s.D.GetOkExists("esxi_software_version"); ok {
		tmp := esxiSoftwareVersion.(string)
		request.EsxiSoftwareVersion = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostOcpuCount, ok := s.D.GetOkExists("host_ocpu_count"); ok {
		tmp := float32(hostOcpuCount.(float64))
		request.HostOcpuCount = &tmp
	}

	if hostShapeName, ok := s.D.GetOkExists("host_shape_name"); ok {
		tmp := hostShapeName.(string)
		request.HostShapeName = &tmp
	}

	if nextCommitment, ok := s.D.GetOkExists("next_commitment"); ok {
		request.NextCommitment = oci_ocvp.CommitmentEnum(nextCommitment.(string))
	}
	if nextSku, ok := s.D.GetOkExists("next_sku"); ok {
		request.NextCommitment = oci_ocvp.CommitmentEnum(nextSku.(string))
	}

	if nonUpgradedEsxiHostId, ok := s.D.GetOkExists("non_upgraded_esxi_host_id"); ok {
		return s.InplaceUpgrade(nonUpgradedEsxiHostId.(string))
	}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		getSddcRequest := oci_ocvp.GetSddcRequest{}
		tmp := sddcId.(string)
		getSddcRequest.SddcId = &tmp
		getSddcRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
		getSddcResponse, getSddcErr := s.SddcClient.GetSddc(context.Background(), getSddcRequest)
		if getSddcErr != nil {
			return fmt.Errorf("cannot get SDDC %s due to error: %s", sddcId, getSddcErr)
		}

		listClustersRequest := oci_ocvp.ListClustersRequest{}
		listClustersRequest.SddcId = &tmp
		listClustersRequest.CompartmentId = getSddcResponse.CompartmentId
		listClustersRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
		listClustersResponse, listClustersErr := s.ClusterClient.ListClusters(context.Background(), listClustersRequest)
		if listClustersErr != nil {
			return fmt.Errorf("cannot list clusters for SDDC %s due to error: %s", sddcId, listClustersErr)
		}
		// by default, add ESXi host to the management cluster of the SDDC
		for _, cluster := range listClustersResponse.Items {
			if cluster.VsphereType == oci_ocvp.VsphereTypesManagement {
				request.ClusterId = cluster.Id
				break
			}
		}
		if request.ClusterId == nil {
			return fmt.Errorf("cannot find management cluster for sddc %s", sddcId)
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setEsxiHostIdFromWorkRequest(workId)
	return s.getEsxiHostFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpEsxiHostResourceCrud) setEsxiHostIdFromWorkRequest(workId *string) {
	workRequestResponse := oci_ocvp.GetWorkRequestResponse{}
	workRequestResponse, err := s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_ocvp.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "esxihost") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
}

func (s *OcvpEsxiHostResourceCrud) getEsxiHostFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	esxiHostId, err := esxiHostWaitForWorkRequest(workId, "esxihost",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*esxiHostId)

	return s.Get()
}

func esxiHostWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func esxiHostWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = esxiHostWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOcvpEsxiHostWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpEsxiHostWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
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

func (s *OcvpEsxiHostResourceCrud) Get() error {
	request := oci_ocvp.GetEsxiHostRequest{}

	tmp := s.D.Id()
	request.EsxiHostId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EsxiHost
	return nil
}

func (s *OcvpEsxiHostResourceCrud) Update() error {
	request := oci_ocvp.UpdateEsxiHostRequest{}

	if billingDonorHostId, ok := s.D.GetOkExists("billing_donor_host_id"); ok {
		tmp := billingDonorHostId.(string)
		request.BillingDonorHostId = &tmp
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

	tmp := s.D.Id()
	request.EsxiHostId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nextSku, ok := s.D.GetOkExists("next_sku"); ok {
		request.NextCommitment = oci_ocvp.CommitmentEnum(nextSku.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EsxiHost
	return nil
}

func (s *OcvpEsxiHostResourceCrud) Delete() error {
	request := oci_ocvp.DeleteEsxiHostRequest{}

	tmp := s.D.Id()
	request.EsxiHostId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := esxiHostWaitForWorkRequest(workId, "esxihost",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpEsxiHostResourceCrud) SetData() error {
	_, clusterIdExists := s.D.GetOkExists("cluster_id")
	if _, sddcIdExists := s.D.GetOkExists("sddc_id"); sddcIdExists && !clusterIdExists {
		// We will set values for deprecated fields only when deprecated fields are used in configs.
		// We will not set deprecated fields values during resource importing
		s.D.Set("current_sku", s.Res.CurrentCommitment)
		s.D.Set("next_sku", s.Res.NextCommitment)
	} else {
		if s.Res.ClusterId != nil {
			s.D.Set("cluster_id", *s.Res.ClusterId)
		}
		if s.Res.EsxiSoftwareVersion != nil {
			s.D.Set("esxi_software_version", *s.Res.EsxiSoftwareVersion)
		}
		s.D.Set("current_commitment", s.Res.CurrentCommitment)
		s.D.Set("next_commitment", s.Res.NextCommitment)
		s.D.Set("current_sku", nil)
		s.D.Set("next_sku", nil)
	}
	if s.Res.BillingContractEndDate != nil {
		s.D.Set("billing_contract_end_date", s.Res.BillingContractEndDate.String())
	}

	if s.Res.BillingDonorHostId != nil {
		s.D.Set("billing_donor_host_id", *s.Res.BillingDonorHostId)
	}
	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	if s.Res.ComputeInstanceId != nil {
		s.D.Set("compute_instance_id", *s.Res.ComputeInstanceId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedEsxiHostId != nil {
		s.D.Set("failed_esxi_host_id", *s.Res.FailedEsxiHostId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracePeriodEndDate != nil {
		s.D.Set("grace_period_end_date", s.Res.GracePeriodEndDate.String())
	}

	if s.Res.HostOcpuCount != nil {
		s.D.Set("host_ocpu_count", *s.Res.HostOcpuCount)
	}

	if s.Res.HostShapeName != nil {
		s.D.Set("host_shape_name", *s.Res.HostShapeName)
	}

	if s.Res.IsBillingContinuationInProgress != nil {
		s.D.Set("is_billing_continuation_in_progress", *s.Res.IsBillingContinuationInProgress)
	}

	if s.Res.IsBillingSwappingInProgress != nil {
		s.D.Set("is_billing_swapping_in_progress", *s.Res.IsBillingSwappingInProgress)
	}

	if s.Res.NonUpgradedEsxiHostId != nil {
		s.D.Set("non_upgraded_esxi_host_id", *s.Res.NonUpgradedEsxiHostId)
	}

	if s.Res.ReplacementEsxiHostId != nil {
		s.D.Set("replacement_esxi_host_id", *s.Res.ReplacementEsxiHostId)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SwapBillingHostId != nil {
		s.D.Set("swap_billing_host_id", *s.Res.SwapBillingHostId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpgradedReplacementEsxiHostId != nil {
		s.D.Set("upgraded_replacement_esxi_host_id", *s.Res.UpgradedReplacementEsxiHostId)
	}

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	return nil
}

func (s *OcvpEsxiHostResourceCrud) InplaceUpgrade(nonUpgradeEsxiHostId string) error {
	getNonUpgradeEsxiHostRequest := oci_ocvp.GetEsxiHostRequest{}
	getNonUpgradeEsxiHostRequest.EsxiHostId = &nonUpgradeEsxiHostId
	getNonUpgradeEsxiHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	response, err := s.Client.GetEsxiHost(context.Background(), getNonUpgradeEsxiHostRequest)
	if err != nil {
		return fmt.Errorf("cannot get non-upgrade ESXi host due to error: %s", err)
	}
	nonUpgradeEsxiHost := response.EsxiHost
	if err = validateReplacementHostDetails(nonUpgradeEsxiHost, s, true); err != nil {
		return err
	}
	if _, ok := s.D.GetOkExists("billing_donor_host_id"); ok {
		return fmt.Errorf("cannot do in-place upgrade and select donor host at the same time")
	}
	if _, ok := s.D.GetOkExists("failed_esxi_host_id"); ok {
		return fmt.Errorf("cannot do host replacement and in-place upgrade at the same time")
	}

	upgradeHostRequest := oci_ocvp.InplaceUpgradeRequest{}
	upgradeHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	upgradeHostRequest.EsxiHostId = &nonUpgradeEsxiHostId

	upgradeResponse, err := s.Client.InplaceUpgrade(context.Background(), upgradeHostRequest)
	if err != nil {
		return err
	}

	workId := upgradeResponse.OpcWorkRequestId
	s.setEsxiHostIdFromWorkRequest(workId)
	return s.getEsxiHostFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpEsxiHostResourceCrud) ReplaceHost(failedEsxiHostId string) error {
	getFailedEsxiHostRequest := oci_ocvp.GetEsxiHostRequest{}
	getFailedEsxiHostRequest.EsxiHostId = &failedEsxiHostId
	getFailedEsxiHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	response, err := s.Client.GetEsxiHost(context.Background(), getFailedEsxiHostRequest)
	if err != nil {
		return fmt.Errorf("cannot get failed ESXi host due to error: %s", err)
	}
	failedEsxiHost := response.EsxiHost
	if err = validateReplacementHostDetails(failedEsxiHost, s, false); err != nil {
		return err
	}
	if _, ok := s.D.GetOkExists("billing_donor_host_id"); ok {
		return fmt.Errorf("cannot do host replacement and select donor host at the same time")
	}
	if _, ok := s.D.GetOkExists("non_upgraded_esxi_host_id"); ok {
		return fmt.Errorf("cannot do host replacement and in-place upgrade at the same time")
	}

	replaceHostRequest := oci_ocvp.ReplaceHostRequest{}
	replaceHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	replaceHostRequest.EsxiHostId = &failedEsxiHostId

	replaceHostResponse, replaceHostErr := s.Client.ReplaceHost(context.Background(), replaceHostRequest)
	if replaceHostErr != nil {
		return replaceHostErr
	}

	workId := replaceHostResponse.OpcWorkRequestId
	s.setEsxiHostIdFromWorkRequest(workId)
	return s.getEsxiHostFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func validateReplacementHostDetails(oldHost oci_ocvp.EsxiHost, s *OcvpEsxiHostResourceCrud, isUpgrade bool) error {
	var oldHostName string
	if isUpgrade {
		oldHostName = "non-upgrade ESXi host"
	} else {
		oldHostName = "failed ESXi host"
	}

	if currentSku, ok := s.D.GetOkExists("current_sku"); ok && currentSku != oldHost.CurrentCommitment {
		return fmt.Errorf("current_sku %s is different from the current commitment %s of %s", currentSku, oldHost.CurrentCommitment, oldHostName)
	}
	if nextSku, ok := s.D.GetOkExists("next_sku"); ok && nextSku != oldHost.NextCommitment {
		return fmt.Errorf("next_sku %s is different from the next commitment %s of %s", nextSku, oldHost.NextCommitment, oldHostName)
	}
	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok && sddcId != *oldHost.SddcId {
		return fmt.Errorf("sddc_id %s is different from the sddc id %s of %s", sddcId, *oldHost.SddcId, oldHostName)
	}
	if shape, ok := s.D.GetOkExists("host_shape_name"); ok && shape != *oldHost.HostShapeName {
		return fmt.Errorf("host_shape_name %s is different from the shape name %s of %s", shape, *oldHost.HostShapeName, oldHostName)
	}
	if ocpuCount, ok := s.D.GetOkExists("host_ocpu_count"); ok && ocpuCount != *oldHost.HostOcpuCount {
		return fmt.Errorf("host_ocpu_count %f is different from the OCPU count %f of %s", ocpuCount, *oldHost.HostOcpuCount, oldHostName)
	}
	if computeAd, ok := s.D.GetOkExists("compute_availability_domain"); ok && computeAd != *oldHost.ComputeAvailabilityDomain {
		return fmt.Errorf("compute_availability_domain %s is different from the availability domain %s of %s", computeAd, *oldHost.ComputeAvailabilityDomain, oldHostName)
	}
	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok && capacityReservationId != *oldHost.CapacityReservationId {
		return fmt.Errorf("capacity_reservation_id %s is different from the capacity reservation id %s of %s", capacityReservationId, *oldHost.CapacityReservationId, oldHostName)
	}
	if _, ok := s.D.GetOkExists("display_name"); ok {
		return fmt.Errorf("cannot set display_name during host replacement or in-place upgrade")
	}
	if _, ok := s.D.GetOkExists("freeform_tags"); ok {
		return fmt.Errorf("cannot set freeform_tags during host replacement or in-place upgrade")
	}
	if _, ok := s.D.GetOkExists("defined_tags"); ok {
		return fmt.Errorf("cannot set defined_tags during host replacement or in-place upgrade")
	}
	return nil
}

func EsxiHostSummaryToMap(obj oci_ocvp.EsxiHostSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillingContractEndDate != nil {
		result["billing_contract_end_date"] = obj.BillingContractEndDate.String()
	}

	if obj.BillingDonorHostId != nil {
		result["billing_donor_host_id"] = string(*obj.BillingDonorHostId)
	}

	if obj.ClusterId != nil {
		result["cluster_id"] = string(*obj.ClusterId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeAvailabilityDomain != nil {
		result["compute_availability_domain"] = string(*obj.ComputeAvailabilityDomain)
	}

	if obj.ComputeInstanceId != nil {
		result["compute_instance_id"] = string(*obj.ComputeInstanceId)
	}

	result["current_commitment"] = string(obj.CurrentCommitment)
	result["current_sku"] = string(obj.CurrentCommitment)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FailedEsxiHostId != nil {
		result["failed_esxi_host_id"] = string(*obj.FailedEsxiHostId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GracePeriodEndDate != nil {
		result["grace_period_end_date"] = obj.GracePeriodEndDate.String()
	}

	if obj.HostOcpuCount != nil {
		result["host_ocpu_count"] = float32(*obj.HostOcpuCount)
	}

	if obj.HostShapeName != nil {
		result["host_shape_name"] = string(*obj.HostShapeName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsBillingContinuationInProgress != nil {
		result["is_billing_continuation_in_progress"] = bool(*obj.IsBillingContinuationInProgress)
	}

	if obj.IsBillingSwappingInProgress != nil {
		result["is_billing_swapping_in_progress"] = bool(*obj.IsBillingSwappingInProgress)
	}

	result["next_commitment"] = string(obj.NextCommitment)
	result["next_sku"] = string(obj.NextCommitment)

	if obj.NonUpgradedEsxiHostId != nil {
		result["non_upgraded_esxi_host_id"] = string(*obj.NonUpgradedEsxiHostId)
	}

	if obj.ReplacementEsxiHostId != nil {
		result["replacement_esxi_host_id"] = string(*obj.ReplacementEsxiHostId)
	}

	if obj.SddcId != nil {
		result["sddc_id"] = string(*obj.SddcId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SwapBillingHostId != nil {
		result["swap_billing_host_id"] = string(*obj.SwapBillingHostId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UpgradedReplacementEsxiHostId != nil {
		result["upgraded_replacement_esxi_host_id"] = string(*obj.UpgradedReplacementEsxiHostId)
	}

	if obj.VmwareSoftwareVersion != nil {
		result["vmware_software_version"] = string(*obj.VmwareSoftwareVersion)
	}

	return result
}
