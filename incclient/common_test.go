package incclient

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MaxAttempts = 30
)

// waitingCheckTxInBlock waits and checks until a transaction has been included in a block.
//
// In case the transaction is invalid, it stops.
func waitingCheckTxInBlock(txHash string) error {
	for {
		isInBlock, err := ic.CheckTxInBlock(txHash)
		if err != nil {
			if !strings.Contains(err.Error(), "-m") {
				log.Printf("CheckTxInBlock of %v error: %v\n", txHash, err)
				return err
			} else {
				time.Sleep(10 * time.Second)
				continue
			}
		}
		if isInBlock {
			log.Printf("Tx %v is in block\n", txHash)
			return nil
		} else {
			log.Printf("Tx %v is in mempool\n", txHash)
			time.Sleep(10 * time.Second)
		}
	}
}

// waitingCheckTxInBlock waits and checks until a transaction has been included in a block.
//
// In case the transaction is invalid, it stops.
func waitingCheckBalanceUpdated(privateKey, tokenID string, oldAmount, expectedNewAmount uint64) error {
	for {
		balance, err := ic.GetBalance(privateKey, tokenID)
		if err != nil {
			return err
		}

		if balance == oldAmount {
			log.Printf("balance not updated\n")
			time.Sleep(10 * time.Second)
			continue
		}

		if balance != expectedNewAmount {
			return fmt.Errorf("expect balance to be %v, got %v", expectedNewAmount, balance)
		} else {
			log.Printf("balance updated correctly\n")
			return nil
		}
	}
}
