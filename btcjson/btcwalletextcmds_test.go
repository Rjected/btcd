// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
)

// TestBtcWalletExtCmds tests all of the btcwallet extended commands marshal and
// unmarshal into valid results include handling of optional fields being
// omitted in the marshalled command, while optional fields with defaults have
// the default assigned on unmarshalled commands.
func TestBtcWalletExtCmds(t *testing.T) {
	t.Parallel()

	testID := int(1)
	tests := []struct {
		name         string
		newCmd       func() (interface{}, error)
		staticCmd    func() interface{}
		marshalled   string
		unmarshalled interface{}
	}{
		{
			name: "createnewaccount",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("createnewaccount", "acct")
			},
			staticCmd: func() interface{} {
				return btcjson.NewCreateNewAccountCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"createnewaccount","params":["acct"],"id":1}`,
			unmarshalled: &btcjson.CreateNewAccountCmd{
				Account: "acct",
			},
		},
		{
			name: "dumpwallet",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("dumpwallet", "filename")
			},
			staticCmd: func() interface{} {
				return btcjson.NewDumpWalletCmd("filename")
			},
			marshalled: `{"jsonrpc":"1.0","method":"dumpwallet","params":["filename"],"id":1}`,
			unmarshalled: &btcjson.DumpWalletCmd{
				Filename: "filename",
			},
		},
		{
			name: "importaddress",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("importaddress", "1Address", "")
			},
			staticCmd: func() interface{} {
				return btcjson.NewImportAddressCmd("1Address", "", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importaddress","params":["1Address",""],"id":1}`,
			unmarshalled: &btcjson.ImportAddressCmd{
				Address: "1Address",
				Rescan:  btcjson.Bool(true),
			},
		},
		{
			name: "importaddress optional",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("importaddress", "1Address", "acct", false)
			},
			staticCmd: func() interface{} {
				return btcjson.NewImportAddressCmd("1Address", "acct", btcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importaddress","params":["1Address","acct",false],"id":1}`,
			unmarshalled: &btcjson.ImportAddressCmd{
				Address: "1Address",
				Account: "acct",
				Rescan:  btcjson.Bool(false),
			},
		},
		{
			name: "importpubkey",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("importpubkey", "031234")
			},
			staticCmd: func() interface{} {
				return btcjson.NewImportPubKeyCmd("031234", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importpubkey","params":["031234"],"id":1}`,
			unmarshalled: &btcjson.ImportPubKeyCmd{
				PubKey: "031234",
				Rescan: btcjson.Bool(true),
			},
		},
		{
			name: "importpubkey optional",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("importpubkey", "031234", false)
			},
			staticCmd: func() interface{} {
				return btcjson.NewImportPubKeyCmd("031234", btcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importpubkey","params":["031234",false],"id":1}`,
			unmarshalled: &btcjson.ImportPubKeyCmd{
				PubKey: "031234",
				Rescan: btcjson.Bool(false),
			},
		},
		{
			name: "importwallet",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("importwallet", "filename")
			},
			staticCmd: func() interface{} {
				return btcjson.NewImportWalletCmd("filename")
			},
			marshalled: `{"jsonrpc":"1.0","method":"importwallet","params":["filename"],"id":1}`,
			unmarshalled: &btcjson.ImportWalletCmd{
				Filename: "filename",
			},
		},
		{
			name: "renameaccount",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("renameaccount", "oldacct", "newacct")
			},
			staticCmd: func() interface{} {
				return btcjson.NewRenameAccountCmd("oldacct", "newacct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"renameaccount","params":["oldacct","newacct"],"id":1}`,
			unmarshalled: &btcjson.RenameAccountCmd{
				OldAccount: "oldacct",
				NewAccount: "newacct",
			},
		},
		{
			name: "walletcreatefundedpsbt",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd(
					"walletcreatefundedpsbt",
					[]btcjson.PsbtInput{
						{
							Txid:     "1234",
							Vout:     0,
							Sequence: 0,
						},
					},
					[]btcjson.PsbtOutput{
						btcjson.NewPsbtOutput("1234", 1234),
					},
					btcjson.Int64(1),
					btcjson.WalletCreateFundedPsbtOpts{},
					btcjson.Bool(true),
				)
			},
			staticCmd: func() interface{} {
				return btcjson.NewWalletCreateFundedPsbtCmd(
					[]btcjson.PsbtInput{
						{
							Txid:     "1234",
							Vout:     0,
							Sequence: 0,
						},
					},
					[]btcjson.PsbtOutput{
						btcjson.NewPsbtOutput("1234", 1234),
					},
					btcjson.Int64(1),
					&btcjson.WalletCreateFundedPsbtOpts{},
					btcjson.Bool(true),
				)
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletcreatefundedpsbt","params":[[{"txid":"1234","vout":0,"sequence":0}],[{"1234":1234}],1,{},true],"id":1}`,
			unmarshalled: &btcjson.WalletCreateFundedPsbtCmd{
				Inputs: []btcjson.PsbtInput{
					{
						Txid:     "1234",
						Vout:     0,
						Sequence: 0,
					},
				},
				Outputs: []btcjson.PsbtOutput{
					btcjson.NewPsbtOutput("1234", 1234),
				},
				Locktime:    btcjson.Int64(1),
				Options:     &btcjson.WalletCreateFundedPsbtOpts{},
				Bip32Derivs: btcjson.Bool(true),
			},
		},
		{
			name: "walletprocesspsbt",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd(
					"walletprocesspsbt", "1234", btcjson.Bool(true), btcjson.String("ALL"), btcjson.Bool(true))
			},
			staticCmd: func() interface{} {
				return btcjson.NewWalletProcessPsbtCmd(
					"1234", btcjson.Bool(true), btcjson.String("ALL"), btcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletprocesspsbt","params":["1234",true,"ALL",true],"id":1}`,
			unmarshalled: &btcjson.WalletProcessPsbtCmd{
				Psbt:        "1234",
				Sign:        btcjson.Bool(true),
				SighashType: btcjson.String("ALL"),
				Bip32Derivs: btcjson.Bool(true),
			},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Marshal the command as created by the new static command
		// creation function.
		marshalled, err := btcjson.MarshalCmd(testID, test.staticCmd())
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		// Ensure the command is created without error via the generic
		// new command creation function.
		cmd, err := test.newCmd()
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected NewCmd error: %v ",
				i, test.name, err)
		}

		// Marshal the command as created by the generic new command
		// creation function.
		marshalled, err = btcjson.MarshalCmd(testID, cmd)
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		var request btcjson.Request
		if err := json.Unmarshal(marshalled, &request); err != nil {
			t.Errorf("Test #%d (%s) unexpected error while "+
				"unmarshalling JSON-RPC request: %v", i,
				test.name, err)
			continue
		}

		cmd, err = btcjson.UnmarshalCmd(&request)
		if err != nil {
			t.Errorf("UnmarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(cmd, test.unmarshalled) {
			t.Errorf("Test #%d (%s) unexpected unmarshalled command "+
				"- got %s, want %s", i, test.name,
				fmt.Sprintf("(%T) %+[1]v", cmd),
				fmt.Sprintf("(%T) %+[1]v\n", test.unmarshalled))
			continue
		}
	}
}
