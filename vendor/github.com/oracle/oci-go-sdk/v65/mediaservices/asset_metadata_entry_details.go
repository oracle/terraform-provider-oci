// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssetMetadataEntryDetails Asset Metadata entry information.
type AssetMetadataEntryDetails struct {

	// The Media Asset ID to ingest into the Distribution Channel.
	MediaAssetId *string `mandatory:"true" json:"mediaAssetId"`

	// The compartment ID where the Ingest Workflow Job will be run.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m AssetMetadataEntryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssetMetadataEntryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AssetMetadataEntryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAssetMetadataEntryDetails AssetMetadataEntryDetails
	s := struct {
		DiscriminatorParam string `json:"ingestPayloadType"`
		MarshalTypeAssetMetadataEntryDetails
	}{
		"ASSET_METADATA_MEDIA_ASSET",
		(MarshalTypeAssetMetadataEntryDetails)(m),
	}

	return json.Marshal(&s)
}
