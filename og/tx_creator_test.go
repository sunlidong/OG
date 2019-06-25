// Copyright © 2019 Annchain Authors <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package og

import (
	"fmt"
	"github.com/annchain/OG/common/crypto"
	"github.com/annchain/OG/common/math"
	"github.com/annchain/OG/og/miner"
	"github.com/annchain/OG/types"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

type AllOkVerifier struct{}

func (AllOkVerifier) Verify(t types.Txi) bool {
	return true
}

func (AllOkVerifier) Name() string {
	return "AllOkVerifier"
}

func (a *AllOkVerifier) String() string {
	return a.Name()
}

func Init() *TxCreator {
	crypto.Signer = &crypto.SignerEd25519{}
	txc := TxCreator{
		TipGenerator:       &dummyTxPoolRandomTx{},
		Miner:              &miner.PoWMiner{},
		MaxConnectingTries: 100,
		MaxTxHash:          types.HexToHash("0x0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		MaxMinedHash:       types.HexToHash("0x00000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		GraphVerifier:      &AllOkVerifier{},
	}
	return &txc
}

func TestTxCreator(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	txc := Init()
	tx := txc.TipGenerator.GetRandomTips(1)[0].(*types.Tx)
	_, priv := crypto.Signer.RandomKeyPair()
	time1 := time.Now()
	txSigned := txc.NewSignedTx(tx.From, tx.To, tx.Value, tx.AccountNonce, priv)
	logrus.Infof("total time for Signing: %d ns", time.Since(time1).Nanoseconds())
	ok := txc.SealTx(txSigned, &priv)
	logrus.Infof("result: %t %v", ok, txSigned)
	txdata, _ := tx.MarshalMsg(nil)
	rawtx := tx.RawTx()
	rawTxData, _ := rawtx.MarshalMsg(nil)
	msg := types.MessageNewTx{
		RawTx: rawtx,
	}
	msgData, _ := msg.MarshalMsg(nil)
	logrus.WithField("len tx ", len(txdata)).WithField("len raw tx ", len(rawTxData)).WithField(
		"len msg", len(msgData)).Debug("encode msg")
	logrus.Debug(txSigned.Dump())
}

func TestSequencerCreator(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	txc := Init()
	_, priv := crypto.Signer.RandomKeyPair()
	time1 := time.Now()

	// for copy
	randomSeq := types.RandomSequencer()

	txSigned := txc.NewSignedSequencer(types.Address{}, randomSeq.Height, randomSeq.AccountNonce, priv)
	logrus.Infof("total time for Signing: %d ns", time.Since(time1).Nanoseconds())
	ok := txc.SealTx(txSigned, &priv)
	logrus.Infof("result: %t %v", ok, txSigned)
}

func sampleTxi(selfHash string, parentsHash []string, baseType types.TxBaseType) types.Txi {

	tx := &types.Tx{TxBase: types.TxBase{
		ParentsHash: types.Hashes{},
		Type:        types.TxBaseTypeNormal,
		Hash:        types.HexToHash(selfHash),
	},
	}
	for _, h := range parentsHash {
		tx.ParentsHash = append(tx.ParentsHash, types.HexToHash(h))
	}
	return tx
}

func TestBuildDag(t *testing.T) {
	crypto.Signer = &crypto.SignerEd25519{}
	logrus.SetLevel(logrus.DebugLevel)
	pool := &DummyTxPoolMiniTx{}
	pool.Init()
	txc := TxCreator{
		TipGenerator:       pool,
		Miner:              &miner.PoWMiner{},
		MaxConnectingTries: 10,
		MaxTxHash:          types.HexToHash("0x0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		MaxMinedHash:       types.HexToHash("0x000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		GraphVerifier:      &AllOkVerifier{},
	}

	_, privateKey := crypto.Signer.RandomKeyPair()

	txs := []types.Txi{
		txc.NewSignedSequencer(types.Address{}, 0, 0, privateKey),
		txc.NewSignedTx(types.HexToAddress("0x01"), types.HexToAddress("0x02"), math.NewBigInt(10), 0, privateKey),
		txc.NewSignedSequencer(types.Address{}, 1, 1, privateKey),
		txc.NewSignedTx(types.HexToAddress("0x02"), types.HexToAddress("0x03"), math.NewBigInt(9), 0, privateKey),
		txc.NewSignedTx(types.HexToAddress("0x03"), types.HexToAddress("0x04"), math.NewBigInt(8), 0, privateKey),
		txc.NewSignedSequencer(types.Address{}, 2, 2, privateKey),
	}

	txs[0].GetBase().Hash = txs[0].CalcTxHash()
	pool.Add(txs[0])
	for i := 1; i < len(txs); i++ {
		if ok := txc.SealTx(txs[i], &privateKey); ok {
			pool.Add(txs[i])
		}
	}
}

func TestNewFIFOTIpGenerator(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	f := NewFIFOTIpGenerator(&dummyTxPoolRandomTx{}, 15)
	fmt.Println(f.fifoRing)
	f.GetRandomTips(3)
	fmt.Println(f.fifoRing)
	fmt.Println(f.fifoRingPos)
	f.fifoRing[2].SetInValid(true)
	f.validation()
	fmt.Println(f.fifoRingPos)
	fmt.Println(f.fifoRing)
	f.GetRandomTips(2)
	fmt.Println(f.fifoRing)
}

func TestSlice(t *testing.T) {
	var parents types.Txis
	parentHashes := make(types.Hashes, len(parents))
	for i, parent := range parents {
		parentHashes[i] = parent.GetTxHash()
	}
}
