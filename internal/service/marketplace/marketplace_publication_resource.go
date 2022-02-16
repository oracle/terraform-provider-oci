// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_marketplace "github.com/oracle/oci-go-sdk/v58/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func MarketplacePublicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwelveHours,
			Update: &tfresource.TwentyMinutes,
			Delete: &tfresource.TwentyMinutes,
		},
		Create: createMarketplacePublication,
		Read:   readMarketplacePublication,
		Update: updateMarketplacePublication,
		Delete: deleteMarketplacePublication,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_agreement_acknowledged": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"listing_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"eula": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Set:      eulaHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"eula_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"TEXT",
										}, true),
									},

									// Optional
									"license_text": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"operating_system": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"package_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IMAGE",
							}, true),
						},
						"package_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"image_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"short_description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"support_contacts": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      supportContactsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"email": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"phone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subject": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"long_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"icon": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"content_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_extension": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mime_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"package_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_operating_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMarketplacePublication(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.CreateResource(d, sync)
}

func readMarketplacePublication(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

func updateMarketplacePublication(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMarketplacePublication(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MarketplacePublicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_marketplace.MarketplaceClient
	Res                    *oci_marketplace.Publication
	DisableNotFoundRetries bool
}

func (s *MarketplacePublicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MarketplacePublicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_marketplace.PublicationLifecycleStateCreating),
	}
}

func (s *MarketplacePublicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_marketplace.PublicationLifecycleStateActive),
	}
}

func (s *MarketplacePublicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_marketplace.PublicationLifecycleStateDeleting),
	}
}

func (s *MarketplacePublicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_marketplace.PublicationLifecycleStateDeleted),
	}
}

func (s *MarketplacePublicationResourceCrud) Create() error {
	request := oci_marketplace.CreatePublicationRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAgreementAcknowledged, ok := s.D.GetOkExists("is_agreement_acknowledged"); ok {
		tmp := isAgreementAcknowledged.(bool)
		request.IsAgreementAcknowledged = &tmp
	}

	if listingType, ok := s.D.GetOkExists("listing_type"); ok {
		request.ListingType = oci_marketplace.ListingTypeEnum(listingType.(string))
	}

	if longDescription, ok := s.D.GetOkExists("long_description"); ok {
		tmp := longDescription.(string)
		request.LongDescription = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if packageDetails, ok := s.D.GetOkExists("package_details"); ok {
		if tmpList := packageDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "package_details", 0)
			tmp, err := s.mapToCreatePublicationPackage(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PackageDetails = tmp
		}
	}

	if shortDescription, ok := s.D.GetOkExists("short_description"); ok {
		tmp := shortDescription.(string)
		request.ShortDescription = &tmp
	}

	if supportContacts, ok := s.D.GetOkExists("support_contacts"); ok {
		set := supportContacts.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_marketplace.SupportContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := supportContactsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "support_contacts", stateDataIndex)
			converted, err := s.mapToSupportContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("support_contacts") {
			request.SupportContacts = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.CreatePublication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Publication
	return nil
}

func (s *MarketplacePublicationResourceCrud) Get() error {
	request := oci_marketplace.GetPublicationRequest{}

	tmp := s.D.Id()
	request.PublicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.GetPublication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Publication
	return nil
}

func (s *MarketplacePublicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_marketplace.UpdatePublicationRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if longDescription, ok := s.D.GetOkExists("long_description"); ok {
		tmp := longDescription.(string)
		request.LongDescription = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	tmp := s.D.Id()
	request.PublicationId = &tmp

	if shortDescription, ok := s.D.GetOkExists("short_description"); ok {
		tmp := shortDescription.(string)
		request.ShortDescription = &tmp
	}

	if supportContacts, ok := s.D.GetOkExists("support_contacts"); ok {
		set := supportContacts.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_marketplace.SupportContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := supportContactsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "support_contacts", stateDataIndex)
			converted, err := s.mapToSupportContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("support_contacts") {
			request.SupportContacts = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.UpdatePublication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Publication
	return nil
}

func (s *MarketplacePublicationResourceCrud) Delete() error {
	request := oci_marketplace.DeletePublicationRequest{}

	tmp := s.D.Id()
	request.PublicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	_, err := s.Client.DeletePublication(context.Background(), request)
	return err
}

func (s *MarketplacePublicationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Icon != nil {
		s.D.Set("icon", []interface{}{UploadDataToMap(s.Res.Icon)})
	} else {
		s.D.Set("icon", nil)
	}

	s.D.Set("listing_type", s.Res.ListingType)

	if s.Res.LongDescription != nil {
		s.D.Set("long_description", *s.Res.LongDescription)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	s.D.Set("state", s.Res.LifecycleState)

	supportContacts := []interface{}{}
	for _, item := range s.Res.SupportContacts {
		supportContacts = append(supportContacts, SupportContactToMap(item))
	}
	s.D.Set("support_contacts", schema.NewSet(supportContactsHashCodeForSets, supportContacts))

	supportedOperatingSystems := []interface{}{}
	for _, item := range s.Res.SupportedOperatingSystems {
		supportedOperatingSystems = append(supportedOperatingSystems, OperatingSystemToMap(item))
	}
	s.D.Set("supported_operating_systems", supportedOperatingSystems)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *MarketplacePublicationResourceCrud) mapToCreatePublicationPackage(fieldKeyFormat string) (oci_marketplace.CreatePublicationPackage, error) {
	var baseObject oci_marketplace.CreatePublicationPackage
	//discriminator
	packageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_type"))
	var packageType string
	if ok {
		packageType = packageTypeRaw.(string)
	} else {
		packageType = "" // default value
	}
	switch strings.ToLower(packageType) {
	case strings.ToLower("IMAGE"):
		details := oci_marketplace.CreateImagePublicationPackage{}
		if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
			tmp := imageId.(string)
			details.ImageId = &tmp
		}
		if eula, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "eula")); ok {
			set := eula.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_marketplace.Eula, len(interfaces))
			for i := range interfaces {
				stateDataIndex := eulaHashCodeForSets(interfaces[i])
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "eula"), stateDataIndex)
				converted, err := s.mapToEula(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "eula")) {
				details.Eula = tmp
			}
		}
		if operatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system")); ok {
			if tmpList := operatingSystem.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "operating_system"), 0)
				tmp, err := s.mapToOperatingSystem(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert operating_system, encountered error: %v", err)
				}
				details.OperatingSystem = &tmp
			}
		}
		if packageVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_version")); ok {
			tmp := packageVersion.(string)
			details.PackageVersion = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown package_type '%v' was specified", packageType)
	}
	return baseObject, nil
}

func CreatePublicationPackageToMap(obj *oci_marketplace.CreatePublicationPackage, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_marketplace.CreateImagePublicationPackage:
		result["package_type"] = "IMAGE"

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'package_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MarketplacePublicationResourceCrud) mapToEula(fieldKeyFormat string) (oci_marketplace.Eula, error) {
	var baseObject oci_marketplace.Eula
	//discriminator
	eulaTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "eula_type"))
	var eulaType string
	if ok {
		eulaType = eulaTypeRaw.(string)
	} else {
		eulaType = "" // default value
	}
	switch strings.ToLower(eulaType) {
	case strings.ToLower("TEXT"):
		details := oci_marketplace.TextBasedEula{}
		if licenseText, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_text")); ok {
			tmp := licenseText.(string)
			details.LicenseText = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown eula_type '%v' was specified", eulaType)
	}
	return baseObject, nil
}

func EulaToMap(obj oci_marketplace.Eula) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_marketplace.TextBasedEula:
		result["eula_type"] = "TEXT"

		if v.LicenseText != nil {
			result["license_text"] = string(*v.LicenseText)
		}
	default:
		log.Printf("[WARN] Received 'eula_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func OperatingSystemToMap(obj oci_marketplace.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *MarketplacePublicationResourceCrud) mapToSupportContact(fieldKeyFormat string) (oci_marketplace.SupportContact, error) {
	result := oci_marketplace.SupportContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if phone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "phone")); ok {
		tmp := phone.(string)
		result.Phone = &tmp
	}

	if subject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject")); ok {
		tmp := subject.(string)
		result.Subject = &tmp
	}

	return result, nil
}

func SupportContactToMap(obj oci_marketplace.SupportContact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Phone != nil {
		result["phone"] = string(*obj.Phone)
	}

	if obj.Subject != nil {
		result["subject"] = string(*obj.Subject)
	}

	return result
}

func UploadDataToMap(obj *oci_marketplace.UploadData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentUrl != nil {
		result["content_url"] = string(*obj.ContentUrl)
	}

	if obj.FileExtension != nil {
		result["file_extension"] = string(*obj.FileExtension)
	}

	if obj.MimeType != nil {
		result["mime_type"] = string(*obj.MimeType)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func eulaHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if eulaType, ok := m["eula_type"]; ok && eulaType != "" {
		buf.WriteString(fmt.Sprintf("%v-", eulaType))
	}
	if licenseText, ok := m["license_text"]; ok && licenseText != "" {
		buf.WriteString(fmt.Sprintf("%v-", licenseText))
	}
	return hashcode.String(buf.String())
}

func supportContactsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if email, ok := m["email"]; ok && email != "" {
		buf.WriteString(fmt.Sprintf("%v-", email))
	}
	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}
	if phone, ok := m["phone"]; ok && phone != "" {
		buf.WriteString(fmt.Sprintf("%v-", phone))
	}
	if subject, ok := m["subject"]; ok && subject != "" {
		buf.WriteString(fmt.Sprintf("%v-", subject))
	}
	return hashcode.String(buf.String())
}
func (s *MarketplacePublicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_marketplace.ChangePublicationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PublicationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	_, err := s.Client.ChangePublicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *MarketplacePublicationResourceCrud) mapToOperatingSystem(fieldKeyFormat string) (oci_marketplace.OperatingSystem, error) {
	result := oci_marketplace.OperatingSystem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}
