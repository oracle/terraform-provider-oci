// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageOutboundConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageOutboundConnector,
		Read:     readFileStorageOutboundConnector,
		Update:   updateFileStorageOutboundConnector,
		Delete:   deleteFileStorageOutboundConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"bind_distinguished_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connector_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"LDAPBIND",
				}, true),
			},
			"endpoints": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"port": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
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
			"password_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"password_secret_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
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

func createFileStorageOutboundConnector(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageOutboundConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageOutboundConnector(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageOutboundConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageOutboundConnector(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageOutboundConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageOutboundConnector(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageOutboundConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageOutboundConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.OutboundConnector
	DisableNotFoundRetries bool
}

func (s *FileStorageOutboundConnectorResourceCrud) ID() string {
	outboundConnector := *s.Res
	return *outboundConnector.GetId()
}

func (s *FileStorageOutboundConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.OutboundConnectorLifecycleStateCreating),
	}
}

func (s *FileStorageOutboundConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.OutboundConnectorLifecycleStateActive),
	}
}

func (s *FileStorageOutboundConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.OutboundConnectorLifecycleStateDeleting),
	}
}

func (s *FileStorageOutboundConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.OutboundConnectorLifecycleStateDeleted),
	}
}

func (s *FileStorageOutboundConnectorResourceCrud) Create() error {
	request := oci_file_storage.CreateOutboundConnectorRequest{}
	err := s.populateTopLevelPolymorphicCreateOutboundConnectorRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateOutboundConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OutboundConnector
	return nil
}

func (s *FileStorageOutboundConnectorResourceCrud) Get() error {
	request := oci_file_storage.GetOutboundConnectorRequest{}

	tmp := s.D.Id()
	request.OutboundConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetOutboundConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OutboundConnector
	return nil
}

func (s *FileStorageOutboundConnectorResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_file_storage.UpdateOutboundConnectorRequest{}

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

	tmp := s.D.Id()
	request.OutboundConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateOutboundConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OutboundConnector
	return nil
}

func (s *FileStorageOutboundConnectorResourceCrud) Delete() error {
	request := oci_file_storage.DeleteOutboundConnectorRequest{}

	tmp := s.D.Id()
	request.OutboundConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteOutboundConnector(context.Background(), request)
	return err
}

func (s *FileStorageOutboundConnectorResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_file_storage.LdapBindAccount:
		s.D.Set("connector_type", "LDAPBIND")

		if v.BindDistinguishedName != nil {
			s.D.Set("bind_distinguished_name", *v.BindDistinguishedName)
		}

		endpoints := []interface{}{}
		for _, item := range v.Endpoints {
			endpoints = append(endpoints, EndpointToMap(item))
		}
		s.D.Set("endpoints", endpoints)

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PasswordSecretVersion != nil {
			s.D.Set("password_secret_version", *v.PasswordSecretVersion)
		}

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *FileStorageOutboundConnectorResourceCrud) mapToEndpoint(fieldKeyFormat string) (oci_file_storage.Endpoint, error) {
	result := oci_file_storage.Endpoint{}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert port string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.Port = &tmpInt64
	}

	return result, nil
}

func EndpointToMap(obj oci_file_storage.Endpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Port != nil {
		result["port"] = strconv.FormatInt(*obj.Port, 10)
	}

	return result
}

func (s *FileStorageOutboundConnectorResourceCrud) populateTopLevelPolymorphicCreateOutboundConnectorRequest(request *oci_file_storage.CreateOutboundConnectorRequest) error {
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists("connector_type")
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("LDAPBIND"):
		details := oci_file_storage.CreateLdapBindAccountDetails{}
		if bindDistinguishedName, ok := s.D.GetOkExists("bind_distinguished_name"); ok {
			tmp := bindDistinguishedName.(string)
			details.BindDistinguishedName = &tmp
		}
		if endpoints, ok := s.D.GetOkExists("endpoints"); ok {
			interfaces := endpoints.([]interface{})
			tmp := make([]oci_file_storage.Endpoint, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "endpoints", stateDataIndex)
				converted, err := s.mapToEndpoint(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("endpoints") {
				details.Endpoints = tmp
			}
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if passwordSecretVersion, ok := s.D.GetOkExists("password_secret_version"); ok {
			tmp := passwordSecretVersion.(int)
			details.PasswordSecretVersion = &tmp
		}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateOutboundConnectorDetails = details
	default:
		return fmt.Errorf("unknown connector_type '%v' was specified", connectorType)
	}
	return nil
}

func (s *FileStorageOutboundConnectorResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_file_storage.ChangeOutboundConnectorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OutboundConnectorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ChangeOutboundConnectorCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
