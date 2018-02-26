package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/maximilien/softlayer-go/common"
	"github.com/maximilien/softlayer-go/softlayer"

	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type softLayer_Network_Vlan_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Network_Vlan_Service(client softlayer.Client) *softLayer_Network_Vlan_Service {
	return &softLayer_Network_Vlan_Service{
		client: client,
	}
}

func (slhs *softLayer_Network_Vlan_Service) GetName() string {
	return "SoftLayer_Network_Vlan"
}

func (slhs *softLayer_Network_Vlan_Service) GetObject(id int) (datatypes.SoftLayer_Network_Vlan, error) {
	objectMask := []string{
		"networkSpace",
	}

	response, errorCode, err := slhs.client.GetHttpClient().DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%d/getObject.json", slhs.GetName(), id), objectMask, "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Network_Vlan{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Hardware#getObject, HTTP error code: '%d'", errorCode)
		return datatypes.SoftLayer_Network_Vlan{}, errors.New(errorMessage)
	}

	err = slhs.client.GetHttpClient().CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Network_Vlan{}, err
	}

	networkVlan := datatypes.SoftLayer_Network_Vlan{}
	err = json.Unmarshal(response, &networkVlan)
	if err != nil {
		return datatypes.SoftLayer_Network_Vlan{}, err
	}

	return networkVlan, nil
}
