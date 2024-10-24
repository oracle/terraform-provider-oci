// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"strings"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"
)

func FileStorageMountTargetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageMountTarget,
		Read:     readFileStorageMountTarget,
		Update:   updateFileStorageMountTarget,
		Delete:   deleteFileStorageMountTarget,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"hostname_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"idmap_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"kerberos": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kerberos_realm": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"backup_key_tab_secret_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"current_key_tab_secret_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"is_kerberos_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key_tab_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"ldap_idmap": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cache_lifetime_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cache_refresh_interval_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"group_search_base": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"negative_cache_lifetime_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"outbound_connector1id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"outbound_connector2id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"schema_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_search_base": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"requested_throughput": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"export_set_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"observed_throughput": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"reserved_storage_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_billing_cycle_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFileStorageMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageMountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageMountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageMountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageMountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageMountTargetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	Res                    *oci_file_storage.MountTarget
	DisableNotFoundRetries bool
}

func (s *FileStorageMountTargetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageMountTargetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateCreating),
	}
}

func (s *FileStorageMountTargetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateActive),
		string(oci_file_storage.MountTargetLifecycleStateFailed),
	}
}

func (s *FileStorageMountTargetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateDeleting),
	}
}

func (s *FileStorageMountTargetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateDeleted),
		string(oci_file_storage.MountTargetLifecycleStateFailed),
	}
}

func (s *FileStorageMountTargetResourceCrud) Create() error {
	request := oci_file_storage.CreateMountTargetRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if idmapType, ok := s.D.GetOkExists("idmap_type"); ok {
		request.IdmapType = oci_file_storage.MountTargetIdmapTypeEnum(idmapType.(string))
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	if kerberos, ok := s.D.GetOkExists("kerberos"); ok {
		if tmpList := kerberos.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kerberos", 0)
			tmp, err := s.mapToCreateKerberosDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Kerberos = &tmp
		}
	}

	if ldapIdmap, ok := s.D.GetOkExists("ldap_idmap"); ok {
		if tmpList := ldapIdmap.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ldap_idmap", 0)
			tmp, err := s.mapToCreateLdapIdmapDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LdapIdmap = &tmp
		}
	}

	if locks, ok := s.D.GetOkExists("locks"); ok {
		interfaces := locks.([]interface{})
		tmp := make([]oci_file_storage.ResourceLock, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
			converted, err := s.mapToResourceLock(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("locks") {
			request.Locks = tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if requestedThroughput, ok := s.D.GetOkExists("requested_throughput"); ok {
		tmp := requestedThroughput.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert requestedThroughput string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.RequestedThroughput = &tmpInt64
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MountTarget
	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageMountTargetResourceCrud) Get() error {
	request := oci_file_storage.GetMountTargetRequest{}

	tmp := s.D.Id()
	request.MountTargetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MountTarget
	return nil
}

func (s *FileStorageMountTargetResourceCrud) Update() error {
	if _, ok := s.D.GetOkExists("state"); ok &&
		s.D.Get("state") == string(oci_file_storage.MountTargetLifecycleStateUpdating) &&
		!s.D.HasChange("requested_throughput") && !s.D.HasChange("compartment_id") &&
		!s.D.HasChange("display_name") && !s.D.HasChange("freeform_tags") &&
		!s.D.HasChange("idmap_type") && !s.D.HasChange("kerberos") &&
		!s.D.HasChange("nsg_ids") && !s.D.HasChange("ldap_idmap") {
		/* Since there is no trigger handle for CancelDowngradeShapeMountTarget,
		so we will check if the UPDATE operation is triggered by other fields changes
		IMPORTANT: Need to update the condition here if there is a new field in Update operation.
		Otherwise, it will trigger CancelDowngrade in HPMT MT test case */
		err := s.CancelDowngradeShapeMountTarget()
		if err != nil {
			return err
		}
	}

	if _, ok := s.D.GetOkExists("requested_throughput"); ok && s.D.HasChange("requested_throughput") {
		oldRaw, newRaw := s.D.GetChange("requested_throughput")
		if oldRaw != "" && newRaw != "" {
			oldV := oldRaw.(string)
			newV := newRaw.(string)
			if strings.Compare(oldV, newV) < 0 {
				err := s.UpgradeShapeMountTarget()
				if err != nil {
					return err
				}
			} else {
				err := s.ScheduleDowngradeShapeMountTarget()
				if err != nil {
					return err
				}
			}
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
	request := oci_file_storage.UpdateMountTargetRequest{}

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

	if idmapType, ok := s.D.GetOkExists("idmap_type"); ok {
		request.IdmapType = oci_file_storage.MountTargetIdmapTypeEnum(idmapType.(string))
	}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	if kerberos, ok := s.D.GetOkExists("kerberos"); ok {
		if tmpList := kerberos.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kerberos", 0)
			tmp, err := s.mapToUpdateKerberosDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Kerberos = &tmp
		}
	}

	if ldapIdmap, ok := s.D.GetOkExists("ldap_idmap"); ok {
		if tmpList := ldapIdmap.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ldap_idmap", 0)
			tmp, err := s.mapToUpdateLdapIdmapDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LdapIdmap = &tmp
		}
	}

	tmp := s.D.Id()
	request.MountTargetId = &tmp

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MountTarget
	return nil
}

func (s *FileStorageMountTargetResourceCrud) Delete() error {
	request := oci_file_storage.DeleteMountTargetRequest{}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.MountTargetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteMountTarget(context.Background(), request)
	return err
}

func (s *FileStorageMountTargetResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExportSetId != nil {
		s.D.Set("export_set_id", *s.Res.ExportSetId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("idmap_type", s.Res.IdmapType)

	if s.Res.Kerberos != nil {
		s.D.Set("kerberos", []interface{}{KerberosToMap(s.Res.Kerberos)})
	} else {
		s.D.Set("kerberos", nil)
	}

	if s.Res.LdapIdmap != nil {
		s.D.Set("ldap_idmap", []interface{}{LdapIdmapToMap(s.Res.LdapIdmap)})
	} else {
		s.D.Set("ldap_idmap", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, MountTargetResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.ObservedThroughput != nil {
		s.D.Set("observed_throughput", strconv.FormatInt(*s.Res.ObservedThroughput, 10))
	}

	s.D.Set("private_ip_ids", s.Res.PrivateIpIds)

	if s.Res.RequestedThroughput != nil {
		s.D.Set("requested_throughput", strconv.FormatInt(*s.Res.RequestedThroughput, 10))
	}

	if s.Res.ReservedStorageCapacity != nil {
		s.D.Set("reserved_storage_capacity", strconv.FormatInt(*s.Res.ReservedStorageCapacity, 10))
	}

	// Service returns only 1 item in this field
	if len(s.Res.PrivateIpIds) > 0 {
		err := s.setPrivateIpDetails(s.Res.PrivateIpIds[0])
		if err != nil {
			return err
		}
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeBillingCycleEnd != nil {
		s.D.Set("time_billing_cycle_end", s.Res.TimeBillingCycleEnd.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *FileStorageMountTargetResourceCrud) CancelDowngradeShapeMountTarget() error {
	request := oci_file_storage.CancelDowngradeShapeMountTargetRequest{}

	idTmp := s.D.Id()
	request.MountTargetId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.CancelDowngradeShapeMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageMountTargetResourceCrud) ScheduleDowngradeShapeMountTarget() error {
	request := oci_file_storage.ScheduleDowngradeShapeMountTargetRequest{}

	idTmp := s.D.Id()
	request.MountTargetId = &idTmp

	if requestedThroughput, ok := s.D.GetOkExists("requested_throughput"); ok {
		tmp := requestedThroughput.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert requestedThroughput string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.RequestedThroughput = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ScheduleDowngradeShapeMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageMountTargetResourceCrud) UpgradeShapeMountTarget() error {
	request := oci_file_storage.UpgradeShapeMountTargetRequest{}

	idTmp := s.D.Id()
	request.MountTargetId = &idTmp

	if requestedThroughput, ok := s.D.GetOkExists("requested_throughput"); ok {
		tmp := requestedThroughput.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert requestedThroughput string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.RequestedThroughput = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.UpgradeShapeMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageMountTargetResourceCrud) mapToCreateKerberosDetails(fieldKeyFormat string) (oci_file_storage.CreateKerberosDetails, error) {
	result := oci_file_storage.CreateKerberosDetails{}

	if backupKeyTabSecretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_key_tab_secret_version")); ok {
		tmp := backupKeyTabSecretVersion.(int)
		result.BackupKeyTabSecretVersion = &tmp
	}

	if currentKeyTabSecretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "current_key_tab_secret_version")); ok {
		tmp := currentKeyTabSecretVersion.(int)
		result.CurrentKeyTabSecretVersion = &tmp
	}

	if isKerberosEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_kerberos_enabled")); ok {
		tmp := isKerberosEnabled.(bool)
		result.IsKerberosEnabled = &tmp
	}

	if kerberosRealm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kerberos_realm")); ok {
		tmp := kerberosRealm.(string)
		result.KerberosRealm = &tmp
	}

	if keyTabSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_tab_secret_id")); ok {
		tmp := keyTabSecretId.(string)
		result.KeyTabSecretId = &tmp
	}

	return result, nil
}

func (s *FileStorageMountTargetResourceCrud) mapToUpdateKerberosDetails(fieldKeyFormat string) (oci_file_storage.UpdateKerberosDetails, error) {
	result := oci_file_storage.UpdateKerberosDetails{}

	if backupKeyTabSecretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_key_tab_secret_version")); ok {
		tmp := backupKeyTabSecretVersion.(int)
		result.BackupKeyTabSecretVersion = &tmp
	}

	if currentKeyTabSecretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "current_key_tab_secret_version")); ok {
		tmp := currentKeyTabSecretVersion.(int)
		result.CurrentKeyTabSecretVersion = &tmp
	}

	if isKerberosEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_kerberos_enabled")); ok {
		tmp := isKerberosEnabled.(bool)
		result.IsKerberosEnabled = &tmp
	}

	if kerberosRealm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kerberos_realm")); ok {
		tmp := kerberosRealm.(string)
		result.KerberosRealm = &tmp
	}

	if keyTabSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_tab_secret_id")); ok {
		tmp := keyTabSecretId.(string)
		result.KeyTabSecretId = &tmp
	}

	return result, nil
}

func KerberosToMap(obj *oci_file_storage.Kerberos) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupKeyTabSecretVersion != nil {
		result["backup_key_tab_secret_version"] = int(*obj.BackupKeyTabSecretVersion)
	}

	if obj.CurrentKeyTabSecretVersion != nil {
		result["current_key_tab_secret_version"] = int(*obj.CurrentKeyTabSecretVersion)
	}

	if obj.IsKerberosEnabled != nil {
		result["is_kerberos_enabled"] = bool(*obj.IsKerberosEnabled)
	}

	if obj.KerberosRealm != nil {
		result["kerberos_realm"] = string(*obj.KerberosRealm)
	}

	if obj.KeyTabSecretId != nil {
		result["key_tab_secret_id"] = string(*obj.KeyTabSecretId)
	}

	return result
}

func (s *FileStorageMountTargetResourceCrud) mapToCreateLdapIdmapDetails(fieldKeyFormat string) (oci_file_storage.CreateLdapIdmapDetails, error) {
	result := oci_file_storage.CreateLdapIdmapDetails{}

	if cacheLifetimeSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cache_lifetime_seconds")); ok {
		tmp := cacheLifetimeSeconds.(int)
		result.CacheLifetimeSeconds = &tmp
	}

	if cacheRefreshIntervalSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cache_refresh_interval_seconds")); ok {
		tmp := cacheRefreshIntervalSeconds.(int)
		result.CacheRefreshIntervalSeconds = &tmp
	}

	if groupSearchBase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_search_base")); ok {
		tmp := groupSearchBase.(string)
		result.GroupSearchBase = &tmp
	}

	if negativeCacheLifetimeSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "negative_cache_lifetime_seconds")); ok {
		tmp := negativeCacheLifetimeSeconds.(int)
		result.NegativeCacheLifetimeSeconds = &tmp
	}

	if outboundConnector1Id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_connector1id")); ok {
		tmp := outboundConnector1Id.(string)
		result.OutboundConnector1Id = &tmp
	}

	if outboundConnector2Id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_connector2id")); ok {
		tmp := outboundConnector2Id.(string)
		result.OutboundConnector2Id = &tmp
	}

	if schemaType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_type")); ok {
		result.SchemaType = oci_file_storage.CreateLdapIdmapDetailsSchemaTypeEnum(schemaType.(string))
	}

	if userSearchBase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_search_base")); ok {
		tmp := userSearchBase.(string)
		result.UserSearchBase = &tmp
	}

	return result, nil
}

func (s *FileStorageMountTargetResourceCrud) mapToUpdateLdapIdmapDetails(fieldKeyFormat string) (oci_file_storage.UpdateLdapIdmapDetails, error) {
	result := oci_file_storage.UpdateLdapIdmapDetails{}

	if cacheLifetimeSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cache_lifetime_seconds")); ok {
		tmp := cacheLifetimeSeconds.(int)
		result.CacheLifetimeSeconds = &tmp
	}

	if cacheRefreshIntervalSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cache_refresh_interval_seconds")); ok {
		tmp := cacheRefreshIntervalSeconds.(int)
		result.CacheRefreshIntervalSeconds = &tmp
	}

	if groupSearchBase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_search_base")); ok {
		tmp := groupSearchBase.(string)
		result.GroupSearchBase = &tmp
	}

	if negativeCacheLifetimeSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "negative_cache_lifetime_seconds")); ok {
		tmp := negativeCacheLifetimeSeconds.(int)
		result.NegativeCacheLifetimeSeconds = &tmp
	}

	if outboundConnector1Id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_connector1id")); ok {
		tmp := outboundConnector1Id.(string)
		result.OutboundConnector1Id = &tmp
	}

	if outboundConnector2Id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_connector2id")); ok {
		tmp := outboundConnector2Id.(string)
		result.OutboundConnector2Id = &tmp
	}

	if schemaType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_type")); ok {
		result.SchemaType = oci_file_storage.UpdateLdapIdmapDetailsSchemaTypeEnum(schemaType.(string))
	}

	if userSearchBase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_search_base")); ok {
		tmp := userSearchBase.(string)
		result.UserSearchBase = &tmp
	}

	return result, nil
}

func LdapIdmapToMap(obj *oci_file_storage.LdapIdmap) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CacheLifetimeSeconds != nil {
		result["cache_lifetime_seconds"] = int(*obj.CacheLifetimeSeconds)
	}

	if obj.CacheRefreshIntervalSeconds != nil {
		result["cache_refresh_interval_seconds"] = int(*obj.CacheRefreshIntervalSeconds)
	}

	if obj.GroupSearchBase != nil {
		result["group_search_base"] = string(*obj.GroupSearchBase)
	}

	if obj.NegativeCacheLifetimeSeconds != nil {
		result["negative_cache_lifetime_seconds"] = int(*obj.NegativeCacheLifetimeSeconds)
	}

	if obj.OutboundConnector1Id != nil {
		result["outbound_connector1id"] = string(*obj.OutboundConnector1Id)
	}

	if obj.OutboundConnector2Id != nil {
		result["outbound_connector2id"] = string(*obj.OutboundConnector2Id)
	}

	result["schema_type"] = string(obj.SchemaType)

	if obj.UserSearchBase != nil {
		result["user_search_base"] = string(*obj.UserSearchBase)
	}

	return result
}

func (s *FileStorageMountTargetResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_file_storage.ResourceLock, error) {
	result := oci_file_storage.ResourceLock{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_file_storage.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func MountTargetResourceLockToMap(obj oci_file_storage.ResourceLock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *FileStorageMountTargetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_file_storage.ChangeMountTargetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.MountTargetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ChangeMountTargetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageMountTargetResourceCrud) setPrivateIpDetails(privateIpOcid string) error {
	request := oci_core.GetPrivateIpRequest{}

	request.PrivateIpId = &privateIpOcid

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

	response, err := s.VirtualNetworkClient.GetPrivateIp(context.Background(), request)
	if err != nil {
		return err
	}
	if response.HostnameLabel != nil {
		s.D.Set("hostname_label", *response.HostnameLabel)
	}

	if response.IpAddress != nil {
		s.D.Set("ip_address", *response.IpAddress)
	}
	return nil
}
