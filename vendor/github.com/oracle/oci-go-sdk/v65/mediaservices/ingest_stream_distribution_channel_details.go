// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngestStreamDistributionChannelDetails Ingest Payload Information.
type IngestStreamDistributionChannelDetails interface {
}

type ingeststreamdistributionchanneldetails struct {
	JsonData          []byte
	IngestPayloadType string `json:"ingestPayloadType"`
}

// UnmarshalJSON unmarshals json
func (m *ingeststreamdistributionchanneldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleringeststreamdistributionchanneldetails ingeststreamdistributionchanneldetails
	s := struct {
		Model Unmarshaleringeststreamdistributionchanneldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IngestPayloadType = s.Model.IngestPayloadType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ingeststreamdistributionchanneldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.IngestPayloadType {
	case "ASSET_METADATA_MEDIA_ASSET":
		mm := AssetMetadataEntryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for IngestStreamDistributionChannelDetails: %s.", m.IngestPayloadType)
		return *m, nil
	}
}

func (m ingeststreamdistributionchanneldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ingeststreamdistributionchanneldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum Enum with underlying type: string
type IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum string

// Set of constants representing the allowable values for IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum
const (
	IngestStreamDistributionChannelDetailsIngestPayloadTypeAssetMetadataMediaAsset IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum = "ASSET_METADATA_MEDIA_ASSET"
)

var mappingIngestStreamDistributionChannelDetailsIngestPayloadTypeEnum = map[string]IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum{
	"ASSET_METADATA_MEDIA_ASSET": IngestStreamDistributionChannelDetailsIngestPayloadTypeAssetMetadataMediaAsset,
}

var mappingIngestStreamDistributionChannelDetailsIngestPayloadTypeEnumLowerCase = map[string]IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum{
	"asset_metadata_media_asset": IngestStreamDistributionChannelDetailsIngestPayloadTypeAssetMetadataMediaAsset,
}

// GetIngestStreamDistributionChannelDetailsIngestPayloadTypeEnumValues Enumerates the set of values for IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum
func GetIngestStreamDistributionChannelDetailsIngestPayloadTypeEnumValues() []IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum {
	values := make([]IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum, 0)
	for _, v := range mappingIngestStreamDistributionChannelDetailsIngestPayloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIngestStreamDistributionChannelDetailsIngestPayloadTypeEnumStringValues Enumerates the set of values in String for IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum
func GetIngestStreamDistributionChannelDetailsIngestPayloadTypeEnumStringValues() []string {
	return []string{
		"ASSET_METADATA_MEDIA_ASSET",
	}
}

// GetMappingIngestStreamDistributionChannelDetailsIngestPayloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngestStreamDistributionChannelDetailsIngestPayloadTypeEnum(val string) (IngestStreamDistributionChannelDetailsIngestPayloadTypeEnum, bool) {
	enum, ok := mappingIngestStreamDistributionChannelDetailsIngestPayloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
