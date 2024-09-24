// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func ObjectStoragePrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createObjectStoragePrivateEndpoint,
		Read:     readObjectStoragePrivateEndpoint,
		Update:   updateObjectStoragePrivateEndpoint,
		Delete:   deleteObjectStoragePrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"private_endpoint_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 10,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nsg_ids": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 5,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"access_targets": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdns": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func createObjectStoragePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResourceUsingHybridPolling(sync)
}

func readObjectStoragePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateObjectStoragePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResourceUsingHybridPolling(d, sync)
}

func deleteObjectStoragePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResourceUsingHybridPolling(d, sync)
}

type ObjectStoragePrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.PrivateEndpoint
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *ObjectStoragePrivateEndpointResourceCrud) SetData() error {
	_, namespace, err := parsePrivateEndpointCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", s.Res.SubnetId)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", s.Res.PrivateEndpointIp)
	}

	if s.Res.Prefix != nil {
		s.D.Set("prefix", s.Res.Prefix)
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	if s.Res.AccessTargets != nil {
		var accessTargets []map[string]interface{}
		for _, target := range s.Res.AccessTargets {
			accessTarget := map[string]interface{}{
				"namespace":      target.Namespace,
				"compartment_id": target.CompartmentId,
				"bucket":         target.Bucket,
				// Add other fields as needed
			}
			accessTargets = append(accessTargets, accessTarget)
		}
		s.D.Set("access_targets", accessTargets)
	} else {
		log.Printf("[WARN] SetData() unable to parse accessTargets : %s", s.Res.AccessTargets)
	}

	if s.Res.Fqdns != nil {
		fqdnsMap := make(map[string]interface{})

		// Check and set prefixFqdns if available
		if prefixFqdns, ok := s.D.GetOk("fqdns.prefixFqdns"); ok {
			prefixFqdnsMap := make(map[string]interface{})
			if m, ok := prefixFqdns.(map[string]interface{}); ok {
				if v, ok := m["objectStorageApiFqdn"].(string); ok {
					prefixFqdnsMap["objectStorageApiFqdn"] = v
				}
				if v, ok := m["s3CompatibilityApiFqdn"].(string); ok {
					prefixFqdnsMap["s3CompatibilityApiFqdn"] = v
				}
				if v, ok := m["swiftApiFqdn"].(string); ok {
					prefixFqdnsMap["swiftApiFqdn"] = v
				}
			}
			fqdnsMap["prefixFqdns"] = prefixFqdnsMap
		}

		// Check and set additionalPrefixesFqdns if available
		if additionalPrefixesFqdns, ok := s.D.GetOk("fqdns.additionalPrefixesFqdns"); ok {
			additionalPrefixesFqdnsMap := make(map[string]interface{})
			if m, ok := additionalPrefixesFqdns.(map[string]interface{}); ok {
				for key, value := range m {
					if prefixFqdns, ok := value.(map[string]interface{}); ok {
						prefixFqdnsMap := make(map[string]interface{})
						if v, ok := prefixFqdns["objectStorageApiFqdn"].(string); ok {
							prefixFqdnsMap["objectStorageApiFqdn"] = v
						}
						if v, ok := prefixFqdns["s3CompatibilityApiFqdn"].(string); ok {
							prefixFqdnsMap["s3CompatibilityApiFqdn"] = v
						}
						if v, ok := prefixFqdns["swiftApiFqdn"].(string); ok {
							prefixFqdnsMap["swiftApiFqdn"] = v
						}
						additionalPrefixesFqdnsMap[key] = prefixFqdnsMap
					}
				}
			}
			fqdnsMap["additionalPrefixesFqdns"] = additionalPrefixesFqdnsMap
		}
		// Set fqdnsMap in ResourceData
		if err := s.D.Set("fqdns", fqdnsMap); err != nil {
			return fmt.Errorf("error setting fqdns attribute: %w", err)
		}
	} else {
		return fmt.Errorf("s.Res.Fqdns is nil")
	}

	s.D.Set("additional_prefixes", s.Res.AdditionalPrefixes)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("id", *s.Res.Id)

	return nil
}

func (s *ObjectStoragePrivateEndpointResourceCrud) ID() string {
	return GetPrivateEndpointCompositeId(s.D.Get("name").(string), s.D.Get("namespace").(string))
}

func (s *ObjectStoragePrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_object_storage.PrivateEndpointLifecycleStateCreating),
	}
}

func (s *ObjectStoragePrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_object_storage.PrivateEndpointLifecycleStateActive),
		string(oci_object_storage.PrivateEndpointLifecycleStateFailed),
	}
}

func (s *ObjectStoragePrivateEndpointResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_object_storage.PrivateEndpointLifecycleStateUpdating),
	}
}

func (s *ObjectStoragePrivateEndpointResourceCrud) UpdatedTarget() []string {

	return []string{
		string(oci_object_storage.PrivateEndpointLifecycleStateActive),
		string(oci_object_storage.PrivateEndpointLifecycleStateFailed),
	}
}

func (s *ObjectStoragePrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_object_storage.PrivateEndpointLifecycleStateDeleting),
	}
}

func (s *ObjectStoragePrivateEndpointResourceCrud) DeletedTarget() []string {

	return []string{
		string(oci_object_storage.PrivateEndpointLifecycleStateDeleted),
	}
}

func (s *ObjectStoragePrivateEndpointResourceCrud) Create() error {
	request := oci_object_storage.CreatePrivateEndpointRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if prefix, ok := s.D.GetOkExists("prefix"); ok {
		tmp := prefix.(string)
		request.Prefix = &tmp
	}

	if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
		tmp := privateEndpointIp.(string)
		request.PrivateEndpointIp = &tmp
	}

	request.AccessTargets = []oci_object_storage.AccessTargetDetails{}
	if accessTargets, ok := s.D.GetOkExists("access_targets"); ok {
		interfaces := accessTargets.([]interface{})
		tmp := make([]oci_object_storage.AccessTargetDetails, len(interfaces))
		for i, item := range interfaces {
			// Assert item to map[string]interface{}
			accessTargetMap, ok := item.(map[string]interface{})
			if !ok {
				// Handle error if assertion fails
				return fmt.Errorf("unexpected type %T in access targets slice", item)
			}
			// Initialize a new AccessTargetDetails struct
			details := oci_object_storage.AccessTargetDetails{}
			// Map values from accessTargetMap to the fields of details
			if namespace, ok := accessTargetMap["namespace"].(string); ok {
				details.Namespace = &namespace
			}
			if compartmentId, ok := accessTargetMap["compartment_id"].(string); ok {
				details.CompartmentId = &compartmentId
			}
			if bucket, ok := accessTargetMap["bucket"].(string); ok {
				details.Bucket = &bucket
			}
			// Append the populated AccessTargetDetails struct to the tmp slice
			tmp[i] = details
		}
		request.AccessTargets = tmp
	}

	if additionalPrefixes, ok := s.D.GetOkExists("additional_prefixes"); ok {
		interfaces := additionalPrefixes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("additional_prefixes") {
			request.AdditionalPrefixes = tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		interfaces := nsgIds.([]interface{})
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.CreatePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workRequestId := response.OpcWorkRequestId

	workRequestErr := tfresource.ResourceRefreshForHybridPollingPreserveStateOnFailures(s.WorkRequestClient, workRequestId, "instance", oci_work_requests.WorkRequestResourceActionTypeCreated, s.DisableNotFoundRetries, s.D, s)
	if workRequestErr != nil {
		return workRequestErr
	}

	getReq := oci_object_storage.GetPrivateEndpointRequest{}
	getReq.NamespaceName = request.NamespaceName
	getReq.PeName = request.Name
	getResp, err := s.Client.GetPrivateEndpoint(context.Background(), getReq)
	s.Res = &getResp.PrivateEndpoint

	return nil
}

func (s *ObjectStoragePrivateEndpointResourceCrud) Get() error {
	request := oci_object_storage.GetPrivateEndpointRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.PeName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	peName, namespace, err := parsePrivateEndpointCompositeId(s.D.Id())
	if err == nil {
		request.PeName = &peName
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateEndpoint

	return nil
}

func (s *ObjectStoragePrivateEndpointResourceCrud) Update() error {
	request := oci_object_storage.UpdatePrivateEndpointRequest{}

	if peName, ok := s.D.GetOkExists("name"); ok {
		tmp := peName.(string)
		request.PeName = &tmp
	}

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

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.UpdatePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workRequestId := response.OpcWorkRequestId

	workRequestErr := tfresource.ResourceRefreshForHybridPollingPreserveStateOnFailures(s.WorkRequestClient, workRequestId, "instance", oci_work_requests.WorkRequestResourceActionTypeCreated, s.DisableNotFoundRetries, s.D, s)
	if workRequestErr != nil {
		return workRequestErr
	}

	getReq := oci_object_storage.GetPrivateEndpointRequest{}
	getReq.NamespaceName = request.NamespaceName
	getReq.PeName = request.Name
	getResp, err := s.Client.GetPrivateEndpoint(context.Background(), getReq)
	s.Res = &getResp.PrivateEndpoint

	return nil
}

func (s *ObjectStoragePrivateEndpointResourceCrud) Delete() error {
	request := oci_object_storage.DeletePrivateEndpointRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.PeName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.DeletePrivateEndpoint(context.Background(), request)
	if err != nil {
		log.Printf("[ERR] DeletePrivateOpertion for privateEndpoint: %s return an erro: %s", *request.PeName, err)
		return err
	}

	workRequestId := response.OpcWorkRequestId
	workRequestErr := tfresource.ResourceRefreshForHybridPollingOnDeletePreserveStateOnFailures(s.WorkRequestClient, workRequestId, "instance", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.DisableNotFoundRetries, s.D, s)

	if workRequestErr != nil {
		return workRequestErr
	}

	return nil
}

func GetPrivateEndpointCompositeId(name string, namespace string) string {
	name = url.PathEscape(name)
	namespace = url.PathEscape(namespace)
	compositeId := "n/" + namespace + "/pe/" + name
	return compositeId
}

func parsePrivateEndpointCompositeId(compositeId string) (peName string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/pe/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	peName, _ = url.PathUnescape(parts[3])

	return
}

func (s *ObjectStoragePrivateEndpointResourceCrud) mapToAccessTargets(accessTargets map[string]interface{}) (oci_object_storage.AccessTargetDetails, error) {
	details := oci_object_storage.AccessTargetDetails{}

	if tmp, ok := accessTargets["namespace"]; ok {
		ns := tmp.(string)
		details.Namespace = &ns
	}

	if tmp, ok := accessTargets["compartmentId"]; ok {
		comp := tmp.(string)
		details.CompartmentId = &comp
	}

	if tmp, ok := accessTargets["bucket"]; ok {
		bucket := tmp.(string)
		details.Bucket = &bucket
	}

	return details, nil
}
