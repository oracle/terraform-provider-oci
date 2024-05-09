// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceChangeAvailabilityManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubSoftwareSourceChangeAvailabilityManagement,
		Read:     readOsManagementHubSoftwareSourceChangeAvailabilityManagement,
		Delete:   deleteOsManagementHubSoftwareSourceChangeAvailabilityManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"software_source_availabilities": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"software_source_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"availability": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"availability_at_oci": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Optional

			// Computed
		},
	}
}

func createOsManagementHubSoftwareSourceChangeAvailabilityManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubSoftwareSourceChangeAvailabilityManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubSoftwareSourceChangeAvailabilityManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.SoftwareSourceClient
	Res                    *oci_os_management_hub.ChangeAvailabilityOfSoftwareSourcesResponse
	DisableNotFoundRetries bool
}

func (s *OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceChangeAvailabilityManagementResource-", OsManagementHubSoftwareSourceChangeAvailabilityManagementResource(), s.D)
}

func (s *OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceCrud) Create() error {
	request := oci_os_management_hub.ChangeAvailabilityOfSoftwareSourcesRequest{}

	if softwareSourceAvailabilities, ok := s.D.GetOkExists("software_source_availabilities"); ok {
		interfaces := softwareSourceAvailabilities.([]interface{})
		tmp := make([]oci_os_management_hub.SoftwareSourceAvailability, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "software_source_availabilities", stateDataIndex)
			converted, err := s.mapToSoftwareSourceAvailability(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("software_source_availabilities") {
			request.SoftwareSourceAvailabilities = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeAvailabilityOfSoftwareSources(context.Background(), request)
	if err != nil {
		return err
	}

	// Bulk operation, no Get() available
	return nil
}

func (s *OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceCrud) SetData() error {
	return nil
}

func (s *OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceCrud) mapToSoftwareSourceAvailability(fieldKeyFormat string) (oci_os_management_hub.SoftwareSourceAvailability, error) {
	result := oci_os_management_hub.SoftwareSourceAvailability{}

	if availability, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability")); ok {
		result.Availability = oci_os_management_hub.AvailabilityEnum(availability.(string))
	}

	if availabilityAtOci, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_at_oci")); ok {
		result.AvailabilityAtOci = oci_os_management_hub.AvailabilityEnum(availabilityAtOci.(string))
	}

	if softwareSourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_id")); ok {
		tmp := softwareSourceId.(string)
		result.SoftwareSourceId = &tmp
	}

	return result, nil
}

func SoftwareSourceAvailabilityToMap(obj oci_os_management_hub.SoftwareSourceAvailability) map[string]interface{} {
	result := map[string]interface{}{}

	result["availability"] = string(obj.Availability)

	result["availability_at_oci"] = string(obj.AvailabilityAtOci)

	if obj.SoftwareSourceId != nil {
		result["software_source_id"] = string(*obj.SoftwareSourceId)
	}

	return result
}
