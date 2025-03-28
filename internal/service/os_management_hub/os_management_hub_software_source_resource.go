// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubSoftwareSource,
		Read:     readOsManagementHubSoftwareSource,
		Update:   updateOsManagementHubSoftwareSource,
		Delete:   deleteOsManagementHubSoftwareSource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_source_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"CUSTOM",
					"PRIVATE",
					"THIRD_PARTY",
					"VENDOR",
					"VERSIONED",
				}, true),
			},

			// Optional
			"advanced_repo_options": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arch_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"custom_software_source_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"module_stream_profile_filters": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"filter_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"module_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"profile_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"stream_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"package_filters": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"filter_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"package_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"package_name_pattern": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"package_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"package_group_filters": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"filter_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"package_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},

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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"gpg_key_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_auto_resolve_dependencies": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_automatically_updated": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_created_from_package_list": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_gpg_check_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_latest_content_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_mirror_sync_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_ssl_verify_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"origin_software_source_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"packages": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_source_sub_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"software_source_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendor_software_sources": {
				Type:             schema.TypeList,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: vendorSoftwareSourcesDiffSuppressFunc,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"availability": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_at_oci": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpg_key_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpg_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_mandatory_for_autonomous_linux": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"package_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repo_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeFloat,
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
			"time_metadata_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOsManagementHubSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubSoftwareSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.SoftwareSourceClient
	Res                    *oci_os_management_hub.SoftwareSource
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubSoftwareSourceResourceCrud) ID() string {
	softwareSource := *s.Res
	return *softwareSource.GetId()
}

func (s *OsManagementHubSoftwareSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.SoftwareSourceLifecycleStateCreating),
	}
}

func (s *OsManagementHubSoftwareSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.SoftwareSourceLifecycleStateActive),
		string(oci_os_management_hub.SoftwareSourceLifecycleStateNeedsAttention),
	}
}

func (s *OsManagementHubSoftwareSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.SoftwareSourceLifecycleStateDeleting),
	}
}

func (s *OsManagementHubSoftwareSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.SoftwareSourceLifecycleStateDeleted),
	}
}

func (s *OsManagementHubSoftwareSourceResourceCrud) Create() error {
	request := oci_os_management_hub.CreateSoftwareSourceRequest{}
	err := s.populateTopLevelPolymorphicCreateSoftwareSourceRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	// CreateSoftwareSource doesn't return the opc-work-request-id header for now (will be implemented post-GA)
	s.Res = &response.SoftwareSource
	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) getSoftwareSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	softwareSourceId, err := softwareSourceWaitForWorkRequest(workId, "os_management_hub",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*softwareSourceId)

	return s.Get()
}

func softwareSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "os_management_hub", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_os_management_hub.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func softwareSourceWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = softwareSourceWorkRequestShouldRetryFunc(timeout)

	response := oci_os_management_hub.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_os_management_hub.OperationStatusInProgress),
			string(oci_os_management_hub.OperationStatusAccepted),
			string(oci_os_management_hub.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_os_management_hub.OperationStatusSucceeded),
			string(oci_os_management_hub.OperationStatusFailed),
			string(oci_os_management_hub.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_os_management_hub.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(string(res.EntityType)), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_os_management_hub.OperationStatusFailed || response.Status == oci_os_management_hub.OperationStatusCanceled {
		return nil, getErrorFromOsManagementHubSoftwareSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubSoftwareSourceWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_os_management_hub.ListWorkRequestErrorsRequest{
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

func (s *OsManagementHubSoftwareSourceResourceCrud) Get() error {
	request := oci_os_management_hub.GetSoftwareSourceRequest{}

	tmp := s.D.Id()
	request.SoftwareSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SoftwareSource
	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeSoftwareSourceCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_os_management_hub.UpdateSoftwareSourceRequest{}
	err := s.populateTopLevelPolymorphicUpdateSoftwareSourceRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err = s.Client.UpdateSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	// UpdateSoftwareSource doesn't return the opc-work-request-id header for now (will be implemented post-GA)
	err = s.Get()
	if err != nil {
		return err
	}

	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteSoftwareSourceRequest{}

	tmp := s.D.Id()
	request.SoftwareSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteSoftwareSource(context.Background(), request)
	return err
}

func (s *OsManagementHubSoftwareSourceResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_os_management_hub.CustomSoftwareSource:
		s.D.Set("software_source_type", "CUSTOM")

		if v.CustomSoftwareSourceFilter != nil {
			s.D.Set("custom_software_source_filter", []interface{}{CustomSoftwareSourceFilterToMap(v.CustomSoftwareSourceFilter)})
		} else {
			s.D.Set("custom_software_source_filter", nil)
		}

		if v.IsAutoResolveDependencies != nil {
			s.D.Set("is_auto_resolve_dependencies", *v.IsAutoResolveDependencies)
		}

		if v.IsAutomaticallyUpdated != nil {
			s.D.Set("is_automatically_updated", *v.IsAutomaticallyUpdated)
		}

		if v.IsCreatedFromPackageList != nil {
			s.D.Set("is_created_from_package_list", *v.IsCreatedFromPackageList)
		}

		if v.IsLatestContentOnly != nil {
			s.D.Set("is_latest_content_only", *v.IsLatestContentOnly)
		}

		s.D.Set("packages", v.Packages)

		s.D.Set("software_source_sub_type", v.SoftwareSourceSubType)

		if v.TimeMetadataUpdated != nil {
			s.D.Set("time_metadata_updated", v.TimeMetadataUpdated.Format(time.RFC3339Nano))
		}

		vendorSoftwareSources := []interface{}{}
		for _, item := range v.VendorSoftwareSources {
			vendorSoftwareSources = append(vendorSoftwareSources, IdToMap(&item))
		}
		s.D.Set("vendor_software_sources", vendorSoftwareSources)

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}
	case oci_os_management_hub.PrivateSoftwareSource:
		s.D.Set("software_source_type", "PRIVATE")

		if v.AdvancedRepoOptions != nil {
			s.D.Set("advanced_repo_options", *v.AdvancedRepoOptions)
		}

		if v.IsGpgCheckEnabled != nil {
			s.D.Set("is_gpg_check_enabled", *v.IsGpgCheckEnabled)
		}

		if v.IsMirrorSyncAllowed != nil {
			s.D.Set("is_mirror_sync_allowed", *v.IsMirrorSyncAllowed)
		}

		if v.IsSslVerifyEnabled != nil {
			s.D.Set("is_ssl_verify_enabled", *v.IsSslVerifyEnabled)
		}

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		s.D.Set("packages", nil)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		if v.IsMirrorSyncAllowed != nil {
			s.D.Set("is_mirror_sync_allowed", *v.IsMirrorSyncAllowed)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}
	case oci_os_management_hub.ThirdPartySoftwareSource:
		s.D.Set("software_source_type", "THIRD_PARTY")

		if v.AdvancedRepoOptions != nil {
			s.D.Set("advanced_repo_options", *v.AdvancedRepoOptions)
		}

		if v.IsGpgCheckEnabled != nil {
			s.D.Set("is_gpg_check_enabled", *v.IsGpgCheckEnabled)
		}

		if v.IsMirrorSyncAllowed != nil {
			s.D.Set("is_mirror_sync_allowed", *v.IsMirrorSyncAllowed)
		}

		if v.IsSslVerifyEnabled != nil {
			s.D.Set("is_ssl_verify_enabled", *v.IsSslVerifyEnabled)
		}

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		s.D.Set("packages", nil)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		if v.IsMirrorSyncAllowed != nil {
			s.D.Set("is_mirror_sync_allowed", *v.IsMirrorSyncAllowed)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}
	case oci_os_management_hub.VendorSoftwareSource:
		s.D.Set("software_source_type", "VENDOR")

		if v.IsMandatoryForAutonomousLinux != nil {
			s.D.Set("is_mandatory_for_autonomous_linux", *v.IsMandatoryForAutonomousLinux)
		}

		if v.OriginSoftwareSourceId != nil {
			s.D.Set("origin_software_source_id", *v.OriginSoftwareSourceId)
		}

		s.D.Set("vendor_name", v.VendorName)

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		s.D.Set("packages", nil)

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.VersionedCustomSoftwareSource:
		s.D.Set("software_source_type", "VERSIONED")

		if v.CustomSoftwareSourceFilter != nil {
			s.D.Set("custom_software_source_filter", []interface{}{CustomSoftwareSourceFilterToMap(v.CustomSoftwareSourceFilter)})
		} else {
			s.D.Set("custom_software_source_filter", nil)
		}

		if v.IsAutoResolveDependencies != nil {
			s.D.Set("is_auto_resolve_dependencies", *v.IsAutoResolveDependencies)
		}

		if v.IsCreatedFromPackageList != nil {
			s.D.Set("is_created_from_package_list", *v.IsCreatedFromPackageList)
		}

		if v.IsLatestContentOnly != nil {
			s.D.Set("is_latest_content_only", *v.IsLatestContentOnly)
		}

		s.D.Set("packages", v.Packages)

		s.D.Set("software_source_sub_type", v.SoftwareSourceSubType)

		if v.SoftwareSourceVersion != nil {
			s.D.Set("software_source_version", *v.SoftwareSourceVersion)
		}

		if v.TimeMetadataUpdated != nil {
			s.D.Set("time_metadata_updated", v.TimeMetadataUpdated.Format(time.RFC3339Nano))
		}

		vendorSoftwareSources := []interface{}{}
		for _, item := range v.VendorSoftwareSources {
			vendorSoftwareSources = append(vendorSoftwareSources, IdToMap(&item))
		}
		s.D.Set("vendor_software_sources", vendorSoftwareSources)

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}
	default:
		log.Printf("[WARN] Received 'software_source_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) ChangeSoftwareSourceCompartment() error {
	request := oci_os_management_hub.ChangeSoftwareSourceCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.SoftwareSourceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeSoftwareSourceCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) mapToCustomSoftwareSourceFilter(fieldKeyFormat string) (oci_os_management_hub.CustomSoftwareSourceFilter, error) {
	result := oci_os_management_hub.CustomSoftwareSourceFilter{}

	if moduleStreamProfileFilters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "module_stream_profile_filters")); ok {
		interfaces := moduleStreamProfileFilters.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamProfileFilter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "module_stream_profile_filters"), stateDataIndex)
			converted, err := s.mapToModuleStreamProfileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "module_stream_profile_filters")) {
			result.ModuleStreamProfileFilters = tmp
		}
	}

	if packageFilters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_filters")); ok {
		interfaces := packageFilters.([]interface{})
		tmp := make([]oci_os_management_hub.PackageFilter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "package_filters"), stateDataIndex)
			converted, err := s.mapToPackageFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "package_filters")) {
			result.PackageFilters = tmp
		}
	}

	if packageGroupFilters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_group_filters")); ok {
		interfaces := packageGroupFilters.([]interface{})
		tmp := make([]oci_os_management_hub.PackageGroupFilter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "package_group_filters"), stateDataIndex)
			converted, err := s.mapToPackageGroupFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "package_group_filters")) {
			result.PackageGroupFilters = tmp
		}
	}

	return result, nil
}

func CustomSoftwareSourceFilterToMap(obj *oci_os_management_hub.CustomSoftwareSourceFilter) map[string]interface{} {
	result := map[string]interface{}{}

	moduleStreamProfileFilters := []interface{}{}
	for _, item := range obj.ModuleStreamProfileFilters {
		moduleStreamProfileFilters = append(moduleStreamProfileFilters, ModuleStreamProfileFilterToMap(item))
	}
	result["module_stream_profile_filters"] = moduleStreamProfileFilters

	packageFilters := []interface{}{}
	for _, item := range obj.PackageFilters {
		packageFilters = append(packageFilters, PackageFilterToMap(item))
	}
	result["package_filters"] = packageFilters

	packageGroupFilters := []interface{}{}
	for _, item := range obj.PackageGroupFilters {
		packageGroupFilters = append(packageGroupFilters, PackageGroupFilterToMap(item))
	}
	result["package_group_filters"] = packageGroupFilters

	return result
}

func (s *OsManagementHubSoftwareSourceResourceCrud) mapToId(fieldKeyFormat string) (oci_os_management_hub.Id, error) {
	result := oci_os_management_hub.Id{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) mapToModuleStreamProfileFilter(fieldKeyFormat string) (oci_os_management_hub.ModuleStreamProfileFilter, error) {
	result := oci_os_management_hub.ModuleStreamProfileFilter{}

	if filterType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_type")); ok {
		result.FilterType = oci_os_management_hub.FilterTypeEnum(filterType.(string))
	}

	if moduleName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "module_name")); ok {
		tmp := moduleName.(string)
		result.ModuleName = &tmp
	}

	if profileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "profile_name")); ok {
		tmp := profileName.(string)
		result.ProfileName = &tmp
	}

	if streamName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_name")); ok {
		tmp := streamName.(string)
		result.StreamName = &tmp
	}

	return result, nil
}

func ModuleStreamProfileFilterToMap(obj oci_os_management_hub.ModuleStreamProfileFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["filter_type"] = string(obj.FilterType)

	if obj.ModuleName != nil {
		result["module_name"] = string(*obj.ModuleName)
	}

	if obj.ProfileName != nil {
		result["profile_name"] = string(*obj.ProfileName)
	}

	if obj.StreamName != nil {
		result["stream_name"] = string(*obj.StreamName)
	}

	return result
}

func (s *OsManagementHubSoftwareSourceResourceCrud) mapToPackageFilter(fieldKeyFormat string) (oci_os_management_hub.PackageFilter, error) {
	result := oci_os_management_hub.PackageFilter{}

	if filterType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_type")); ok {
		result.FilterType = oci_os_management_hub.FilterTypeEnum(filterType.(string))
	}

	if packageName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_name")); ok {
		tmp := packageName.(string)
		if tmp != "" {
			result.PackageName = &tmp
		}
	}

	if packageNamePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_name_pattern")); ok {
		tmp := packageNamePattern.(string)
		if tmp != "" {
			result.PackageNamePattern = &tmp
		}
	}

	if packageVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_version")); ok {
		tmp := packageVersion.(string)
		if tmp != "" {
			result.PackageVersion = &tmp
		}
	}

	return result, nil
}

func PackageFilterToMap(obj oci_os_management_hub.PackageFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["filter_type"] = string(obj.FilterType)

	if obj.PackageName != nil {
		result["package_name"] = string(*obj.PackageName)
	}

	if obj.PackageNamePattern != nil {
		result["package_name_pattern"] = string(*obj.PackageNamePattern)
	}

	if obj.PackageVersion != nil {
		result["package_version"] = string(*obj.PackageVersion)
	}

	return result
}

func (s *OsManagementHubSoftwareSourceResourceCrud) mapToPackageGroupFilter(fieldKeyFormat string) (oci_os_management_hub.PackageGroupFilter, error) {
	result := oci_os_management_hub.PackageGroupFilter{}

	if filterType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_type")); ok {
		result.FilterType = oci_os_management_hub.FilterTypeEnum(filterType.(string))
	}

	if packageGroups, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_groups")); ok {
		interfaces := packageGroups.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "package_groups")) {
			result.PackageGroups = tmp
		}
	}

	return result, nil
}

func PackageGroupFilterToMap(obj oci_os_management_hub.PackageGroupFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["filter_type"] = string(obj.FilterType)

	result["package_groups"] = obj.PackageGroups

	return result
}

func (s *OsManagementHubSoftwareSourceResourceCrud) populateTopLevelPolymorphicCreateSoftwareSourceRequest(request *oci_os_management_hub.CreateSoftwareSourceRequest) error {
	//discriminator
	softwareSourceTypeRaw, ok := s.D.GetOkExists("software_source_type")
	var softwareSourceType string
	if ok {
		softwareSourceType = softwareSourceTypeRaw.(string)
	} else {
		softwareSourceType = "" // default value
	}
	switch strings.ToLower(softwareSourceType) {
	case strings.ToLower("CUSTOM"):
		details := oci_os_management_hub.CreateCustomSoftwareSourceDetails{}
		if customSoftwareSourceFilter, ok := s.D.GetOkExists("custom_software_source_filter"); ok {
			if tmpList := customSoftwareSourceFilter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_software_source_filter", 0)
				tmp, err := s.mapToCustomSoftwareSourceFilter(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CustomSoftwareSourceFilter = &tmp
			}
		}
		if isAutoResolveDependencies, ok := s.D.GetOkExists("is_auto_resolve_dependencies"); ok {
			tmp := isAutoResolveDependencies.(bool)
			details.IsAutoResolveDependencies = &tmp
		}
		if isAutomaticallyUpdated, ok := s.D.GetOkExists("is_automatically_updated"); ok {
			tmp := isAutomaticallyUpdated.(bool)
			details.IsAutomaticallyUpdated = &tmp
		}
		if isCreatedFromPackageList, ok := s.D.GetOkExists("is_created_from_package_list"); ok {
			tmp := isCreatedFromPackageList.(bool)
			details.IsCreatedFromPackageList = &tmp
		}
		if isLatestContentOnly, ok := s.D.GetOkExists("is_latest_content_only"); ok {
			tmp := isLatestContentOnly.(bool)
			details.IsLatestContentOnly = &tmp
		}
		if packages, ok := s.D.GetOkExists("packages"); ok {
			interfaces := packages.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("packages") {
				details.Packages = tmp
			}
		}
		if softwareSourceSubType, ok := s.D.GetOkExists("software_source_sub_type"); ok {
			details.SoftwareSourceSubType = oci_os_management_hub.SoftwareSourceSubTypeEnum(softwareSourceSubType.(string))
		}
		if vendorSoftwareSources, ok := s.D.GetOkExists("vendor_software_sources"); ok {
			interfaces := vendorSoftwareSources.([]interface{})
			tmp := make([]oci_os_management_hub.Id, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vendor_software_sources", stateDataIndex)
				converted, err := s.mapToId(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("vendor_software_sources") {
				details.VendorSoftwareSources = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateSoftwareSourceDetails = details
	case strings.ToLower("PRIVATE"):
		details := oci_os_management_hub.CreatePrivateSoftwareSourceDetails{}
		if advancedRepoOptions, ok := s.D.GetOkExists("advanced_repo_options"); ok {
			tmp := advancedRepoOptions.(string)
			details.AdvancedRepoOptions = &tmp
		}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isGpgCheckEnabled, ok := s.D.GetOkExists("is_gpg_check_enabled"); ok {
			tmp := isGpgCheckEnabled.(bool)
			details.IsGpgCheckEnabled = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		if isSslVerifyEnabled, ok := s.D.GetOkExists("is_ssl_verify_enabled"); ok {
			tmp := isSslVerifyEnabled.(bool)
			details.IsSslVerifyEnabled = &tmp
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		request.CreateSoftwareSourceDetails = details
	case strings.ToLower("THIRD_PARTY"):
		details := oci_os_management_hub.CreateThirdPartySoftwareSourceDetails{}
		if advancedRepoOptions, ok := s.D.GetOkExists("advanced_repo_options"); ok {
			tmp := advancedRepoOptions.(string)
			details.AdvancedRepoOptions = &tmp
		}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isGpgCheckEnabled, ok := s.D.GetOkExists("is_gpg_check_enabled"); ok {
			tmp := isGpgCheckEnabled.(bool)
			details.IsGpgCheckEnabled = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		if isSslVerifyEnabled, ok := s.D.GetOkExists("is_ssl_verify_enabled"); ok {
			tmp := isSslVerifyEnabled.(bool)
			details.IsSslVerifyEnabled = &tmp
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		request.CreateSoftwareSourceDetails = details
	case strings.ToLower("VENDOR"):
		details := oci_os_management_hub.CreateVendorSoftwareSourceDetails{}
		if originSoftwareSourceId, ok := s.D.GetOkExists("origin_software_source_id"); ok {
			tmp := originSoftwareSourceId.(string)
			details.OriginSoftwareSourceId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateSoftwareSourceDetails = details
	case strings.ToLower("VERSIONED"):
		details := oci_os_management_hub.CreateVersionedCustomSoftwareSourceDetails{}
		if customSoftwareSourceFilter, ok := s.D.GetOkExists("custom_software_source_filter"); ok {
			if tmpList := customSoftwareSourceFilter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_software_source_filter", 0)
				tmp, err := s.mapToCustomSoftwareSourceFilter(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CustomSoftwareSourceFilter = &tmp
			}
		}
		if isAutoResolveDependencies, ok := s.D.GetOkExists("is_auto_resolve_dependencies"); ok {
			tmp := isAutoResolveDependencies.(bool)
			details.IsAutoResolveDependencies = &tmp
		}
		if isCreatedFromPackageList, ok := s.D.GetOkExists("is_created_from_package_list"); ok {
			tmp := isCreatedFromPackageList.(bool)
			details.IsCreatedFromPackageList = &tmp
		}
		if isLatestContentOnly, ok := s.D.GetOkExists("is_latest_content_only"); ok {
			tmp := isLatestContentOnly.(bool)
			details.IsLatestContentOnly = &tmp
		}
		if packages, ok := s.D.GetOkExists("packages"); ok {
			interfaces := packages.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("packages") {
				details.Packages = tmp
			}
		}
		if softwareSourceSubType, ok := s.D.GetOkExists("software_source_sub_type"); ok {
			details.SoftwareSourceSubType = oci_os_management_hub.SoftwareSourceSubTypeEnum(softwareSourceSubType.(string))
		}
		if softwareSourceVersion, ok := s.D.GetOkExists("software_source_version"); ok {
			tmp := softwareSourceVersion.(string)
			details.SoftwareSourceVersion = &tmp
		}
		if vendorSoftwareSources, ok := s.D.GetOkExists("vendor_software_sources"); ok {
			interfaces := vendorSoftwareSources.([]interface{})
			tmp := make([]oci_os_management_hub.Id, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vendor_software_sources", stateDataIndex)
				converted, err := s.mapToId(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("vendor_software_sources") {
				details.VendorSoftwareSources = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateSoftwareSourceDetails = details
	default:
		return fmt.Errorf("unknown software_source_type '%v' was specified", softwareSourceType)
	}
	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) populateTopLevelPolymorphicUpdateSoftwareSourceRequest(request *oci_os_management_hub.UpdateSoftwareSourceRequest) error {
	//discriminator
	softwareSourceTypeRaw, ok := s.D.GetOkExists("software_source_type")
	var softwareSourceType string
	if ok {
		softwareSourceType = softwareSourceTypeRaw.(string)
	} else {
		softwareSourceType = "" // default value
	}
	switch strings.ToLower(softwareSourceType) {
	case strings.ToLower("CUSTOM"):
		details := oci_os_management_hub.UpdateCustomSoftwareSourceDetails{}
		if customSoftwareSourceFilter, ok := s.D.GetOkExists("custom_software_source_filter"); ok {
			if tmpList := customSoftwareSourceFilter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_software_source_filter", 0)
				tmp, err := s.mapToCustomSoftwareSourceFilter(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CustomSoftwareSourceFilter = &tmp
			}
		}
		if isAutoResolveDependencies, ok := s.D.GetOkExists("is_auto_resolve_dependencies"); ok {
			tmp := isAutoResolveDependencies.(bool)
			details.IsAutoResolveDependencies = &tmp
		}
		if isAutomaticallyUpdated, ok := s.D.GetOkExists("is_automatically_updated"); ok {
			tmp := isAutomaticallyUpdated.(bool)
			details.IsAutomaticallyUpdated = &tmp
		}
		if isLatestContentOnly, ok := s.D.GetOkExists("is_latest_content_only"); ok {
			tmp := isLatestContentOnly.(bool)
			details.IsLatestContentOnly = &tmp
		}
		if vendorSoftwareSources, ok := s.D.GetOkExists("vendor_software_sources"); ok {
			interfaces := vendorSoftwareSources.([]interface{})
			tmp := make([]oci_os_management_hub.Id, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vendor_software_sources", stateDataIndex)
				converted, err := s.mapToId(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("vendor_software_sources") {
				details.VendorSoftwareSources = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.SoftwareSourceId = &tmp
		request.UpdateSoftwareSourceDetails = details
	case strings.ToLower("PRIVATE"):
		details := oci_os_management_hub.UpdatePrivateSoftwareSourceDetails{}
		if advancedRepoOptions, ok := s.D.GetOkExists("advanced_repo_options"); ok {
			tmp := advancedRepoOptions.(string)
			details.AdvancedRepoOptions = &tmp
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isGpgCheckEnabled, ok := s.D.GetOkExists("is_gpg_check_enabled"); ok {
			tmp := isGpgCheckEnabled.(bool)
			details.IsGpgCheckEnabled = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		if isSslVerifyEnabled, ok := s.D.GetOkExists("is_ssl_verify_enabled"); ok {
			tmp := isSslVerifyEnabled.(bool)
			details.IsSslVerifyEnabled = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		tmp := s.D.Id()
		request.SoftwareSourceId = &tmp
		request.UpdateSoftwareSourceDetails = details
	case strings.ToLower("THIRD_PARTY"):
		details := oci_os_management_hub.UpdateThirdPartySoftwareSourceDetails{}
		if advancedRepoOptions, ok := s.D.GetOkExists("advanced_repo_options"); ok {
			tmp := advancedRepoOptions.(string)
			details.AdvancedRepoOptions = &tmp
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isGpgCheckEnabled, ok := s.D.GetOkExists("is_gpg_check_enabled"); ok {
			tmp := isGpgCheckEnabled.(bool)
			details.IsGpgCheckEnabled = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		if isSslVerifyEnabled, ok := s.D.GetOkExists("is_ssl_verify_enabled"); ok {
			tmp := isSslVerifyEnabled.(bool)
			details.IsSslVerifyEnabled = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if gpgKeyUrl, ok := s.D.GetOkExists("gpg_key_url"); ok {
			tmp := gpgKeyUrl.(string)
			details.GpgKeyUrl = &tmp
		}
		if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
			tmp := isMirrorSyncAllowed.(bool)
			details.IsMirrorSyncAllowed = &tmp
		}
		tmp := s.D.Id()
		request.SoftwareSourceId = &tmp
		request.UpdateSoftwareSourceDetails = details
	case strings.ToLower("VENDOR"):
		details := oci_os_management_hub.UpdateVendorSoftwareSourceDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.SoftwareSourceId = &tmp
		request.UpdateSoftwareSourceDetails = details
	case strings.ToLower("VERSIONED"):
		details := oci_os_management_hub.UpdateVersionedCustomSoftwareSourceDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.SoftwareSourceId = &tmp
		request.UpdateSoftwareSourceDetails = details
	default:
		return fmt.Errorf("unknown software_source_type '%v' was specified", softwareSourceType)
	}
	return nil
}

func (s *OsManagementHubSoftwareSourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeSoftwareSourceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SoftwareSourceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeSoftwareSourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func vendorSoftwareSourcesDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	// Because Terraform passes sub-attribute keys like "vendor_software_sources.0.id",
	// extract only the parent key (e.g. "vendor_software_sources")
	parentKey := strings.Split(k, ".")[0]

	oldRaw, newRaw := d.GetChange(parentKey)
	if oldRaw == nil || newRaw == nil {
		return false
	}

	oldSlice, okOld := oldRaw.([]interface{})
	newSlice, okNew := newRaw.([]interface{})
	if !okOld || !okNew {
		// Not a slice-of-interfaces? Just do default diff.
		return false
	}

	// If lengths differ, it's definitely a diff
	if len(oldSlice) != len(newSlice) {
		return false
	}

	// Gather "id" fields from old/new
	getIDs := func(slice []interface{}) []string {
		ids := make([]string, 0, len(slice))
		for _, elem := range slice {
			if m, ok := elem.(map[string]interface{}); ok {
				// If missing or nil, skip or handle gracefully
				if val, ok := m["id"].(string); ok {
					ids = append(ids, val)
				}
			}
		}
		return ids
	}

	oldIDs := getIDs(oldSlice)
	newIDs := getIDs(newSlice)

	// If "id" extraction failed or count differs, it's a change
	if len(oldIDs) != len(newIDs) {
		return false
	}

	// Sort for order-insensitivity
	sort.Strings(oldIDs)
	sort.Strings(newIDs)

	// If sorted lists of IDs match, suppress the diff
	for i := range oldIDs {
		if oldIDs[i] != newIDs[i] {
			return false
		}
	}

	return true
}
