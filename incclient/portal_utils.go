package incclient

import (
	"github.com/incognitochain/go-incognito-sdk-v2/metadata"
	"github.com/incognitochain/go-incognito-sdk-v2/rpchandler"
)

// GetPortalShieldingRequestStatus retrieves the status of a port shielding request.
func (client *IncClient) GetPortalShieldingRequestStatus(shieldID string) (*metadata.PortalShieldingRequestStatus, error) {
	responseInBytes, err := client.rpcServer.GetPortalShieldingRequestStatus(shieldID)
	if err != nil {
		return nil, err
	}

	var res *metadata.PortalShieldingRequestStatus
	err = rpchandler.ParseResponse(responseInBytes, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GeneratePortalShieldingAddressFromRPC returns a multi-sig shielding address via an RPC when using the Portal.
func (client *IncClient) GeneratePortalShieldingAddressFromRPC(paymentAddressStr, tokenIDStr string) (string, error) {
	responseInBytes, err := client.rpcServer.GenerateShieldingMultiSigAddress(paymentAddressStr, tokenIDStr)
	if err != nil {
		return "", err
	}

	var res string
	err = rpchandler.ParseResponse(responseInBytes, &res)
	if err != nil {
		return "", err
	}

	return res, nil
}

// GetPortalUnShieldingRequestStatus retrieves the status of a port un-shielding request.
func (client *IncClient) GetPortalUnShieldingRequestStatus(unShieldID string) (*metadata.PortalUnshieldRequestStatus, error) {
	responseInBytes, err := client.rpcServer.GetPortalUnShieldingRequestStatus(unShieldID)
	if err != nil {
		return nil, err
	}

	var res *metadata.PortalUnshieldRequestStatus
	err = rpchandler.ParseResponse(responseInBytes, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
