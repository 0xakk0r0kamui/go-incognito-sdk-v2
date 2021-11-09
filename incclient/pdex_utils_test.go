package incclient

import (
	"encoding/json"
	"testing"
	"time"
)

func TestIncClient_GetPoolPairStateByID(t *testing.T) {
	var err error
	ic, err = NewTestNetClient()
	if err != nil {
		panic(err)
	}

	poolID := "0000000000000000000000000000000000000000000000000000000000000004-0000000000000000000000000000000000000000000000000000000000000006-56e4e9d710a01dfe865e6d5047fabd6bb98b646465863c2726ebc56538983b5d"
	poolState, err := ic.GetPoolPairStateByID(0, poolID)
	if err != nil {
		panic(err)
	}
	jsb, _ := json.MarshalIndent(poolState, "", "\t")
	Logger.Printf("state: %v\n", string(jsb))
}

func TestIncClient_GetPoolShareAmount(t *testing.T) {
	var err error
	ic, err = NewTestNetClient()
	if err != nil {
		panic(err)
	}

	poolID := "0000000000000000000000000000000000000000000000000000000000000004-0000000000000000000000000000000000000000000000000000000000000006-56e4e9d710a01dfe865e6d5047fabd6bb98b646465863c2726ebc56538983b5d"
	nftID := "d150bd389f7f881a271e1617aba13dbc6c0dde7b8d184f0cbd637e93aa83c69f"
	share, err := ic.GetPoolShareAmount(poolID, nftID)
	if err != nil {
		panic(err)
	}
	Logger.Printf("share: %v\n", share)
}

func TestIncClient_CheckTradeStatus(t *testing.T) {
	var err error
	ic, err = NewTestNetClientWithCache()
	if err != nil {
		panic(err)
	}

	txHash := "e4c13e368eb4da34ebcd04aaf9da9a401d5f55df752f3d1c650331a19f69a53a"
	status, err := ic.CheckTradeStatus(txHash)
	if err != nil {
		panic(err)
	}

	Logger.Printf("status: %v\n", status)
}

func TestIncClient_CheckNFTMintingStatus(t *testing.T) {
	var err error
	ic, err = NewTestNetClientWithCache()
	if err != nil {
		panic(err)
	}

	privateKey := "112t8rneWAhErTC8YUFTnfcKHvB1x6uAVdehy1S8GP2psgqDxK3RHouUcd69fz88oAL9XuMyQ8mBY5FmmGJdcyrpwXjWBXRpoWwgJXjsxi4j"
	encodedTx, txHash, err := ic.CreatePdexv3MintNFT(privateKey)
	if err != nil {
		panic(err)
	}
	err = ic.SendRawTx(encodedTx)
	if err != nil {
		panic(err)
	}
	Logger.Printf("TxHash: %v\n", txHash)

	time.Sleep(100 * time.Second)

	status, ID, err := ic.CheckNFTMintingStatus(txHash)
	if err != nil {
		panic(err)
	}
	Logger.Printf("status: %v, NftID: %v\n", status, ID)
}

func TestIncClient_GetListNftIDs(t *testing.T) {
	var err error
	ic, err = NewTestNetClient()
	if err != nil {
		panic(err)
	}

	nftList, err := ic.GetListNftIDs(0)
	if err != nil {
		panic(err)
	}
	Logger.Println(nftList)
}