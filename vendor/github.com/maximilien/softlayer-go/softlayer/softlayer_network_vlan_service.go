package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Network_Vlan_Service interface {
	Service

	GetObject(id int) (datatypes.SoftLayer_Network_Vlan, error)
}
