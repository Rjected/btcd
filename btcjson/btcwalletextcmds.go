// Copyright (c) 2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a wallet server with btcwallet extensions.

package btcjson

// CreateNewAccountCmd defines the createnewaccount JSON-RPC command.
type CreateNewAccountCmd struct {
	Account string
}

// NewCreateNewAccountCmd returns a new instance which can be used to issue a
// createnewaccount JSON-RPC command.
func NewCreateNewAccountCmd(account string) *CreateNewAccountCmd {
	return &CreateNewAccountCmd{
		Account: account,
	}
}

// DumpWalletCmd defines the dumpwallet JSON-RPC command.
type DumpWalletCmd struct {
	Filename string
}

// NewDumpWalletCmd returns a new instance which can be used to issue a
// dumpwallet JSON-RPC command.
func NewDumpWalletCmd(filename string) *DumpWalletCmd {
	return &DumpWalletCmd{
		Filename: filename,
	}
}

// ImportAddressCmd defines the importaddress JSON-RPC command.
type ImportAddressCmd struct {
	Address string
	Account string
	Rescan  *bool `jsonrpcdefault:"true"`
}

// NewImportAddressCmd returns a new instance which can be used to issue an
// importaddress JSON-RPC command.
func NewImportAddressCmd(address string, account string, rescan *bool) *ImportAddressCmd {
	return &ImportAddressCmd{
		Address: address,
		Account: account,
		Rescan:  rescan,
	}
}

// ImportPubKeyCmd defines the importpubkey JSON-RPC command.
type ImportPubKeyCmd struct {
	PubKey string
	Rescan *bool `jsonrpcdefault:"true"`
}

// NewImportPubKeyCmd returns a new instance which can be used to issue an
// importpubkey JSON-RPC command.
func NewImportPubKeyCmd(pubKey string, rescan *bool) *ImportPubKeyCmd {
	return &ImportPubKeyCmd{
		PubKey: pubKey,
		Rescan: rescan,
	}
}

// ImportWalletCmd defines the importwallet JSON-RPC command.
type ImportWalletCmd struct {
	Filename string
}

// NewImportWalletCmd returns a new instance which can be used to issue a
// importwallet JSON-RPC command.
func NewImportWalletCmd(filename string) *ImportWalletCmd {
	return &ImportWalletCmd{
		Filename: filename,
	}
}

// RenameAccountCmd defines the renameaccount JSON-RPC command.
type RenameAccountCmd struct {
	OldAccount string
	NewAccount string
}

// NewRenameAccountCmd returns a new instance which can be used to issue a
// renameaccount JSON-RPC command.
func NewRenameAccountCmd(oldAccount, newAccount string) *RenameAccountCmd {
	return &RenameAccountCmd{
		OldAccount: oldAccount,
		NewAccount: newAccount,
	}
}

// PsbtInput represents an input to include in the PSBT created by the
// WalletCreateFundedPsbtCmd command.
type PsbtInput struct {
	Txid     string `json:"txid"`
	Vout     int64  `json:"vout"`
	Sequence int64  `json:"sequence"`
}

// PsbtOutput represents an output to include in the PSBT created by the
// WalletCreateFundedPsbtCmd command.
type PsbtOutput map[string]float64

// NewPsbtOutput returns a new instance of a PSBT output to use with the
// WalletCreateFundedPsbtCmd command.
func NewPsbtOutput(address string, amount float64) PsbtOutput {
	out := make(map[string]float64)
	out[address] = amount
	return out
}

// WalletCreateFundedPsbtOpts represents the optional options struct provided
// with a WalletCreateFundedPsbtCmd command.
type WalletCreateFundedPsbtOpts struct {
	ChangeAddress          *string  `json:"changeAddress,omitempty"`
	ChangePosition         *int64   `json:"changePosition,omitempty"`
	ChangeType             *string  `json:"change_type,omitempty"`
	IncludeWatching        *bool    `json:"includeWatching,omitempty"`
	LockUnspents           *bool    `json:"lockUnspents,omitempty"`
	FeeRate                *int64   `json:"feeRate,omitempty"`
	SubtractFeeFromOutputs *[]int64 `json:"subtractFeeFromOutputs,omitempty"`
	Replaceable            *bool    `json:"replaceable,omitempty"`
	ConfTarget             *int64   `json:"conf_target,omitempty"`
	EstimateMode           *string  `json:"estimate_mode,omitempty"`
}

// WalletCreateFundedPsbtCmd defines the walletcreatefundedpsbt JSON-RPC command.
type WalletCreateFundedPsbtCmd struct {
	Inputs      []PsbtInput
	Outputs     []PsbtOutput
	Locktime    *int64
	Options     *WalletCreateFundedPsbtOpts
	Bip32Derivs *bool
}

// NewWalletCreateFundedPsbtCmd returns a new instance which can be used to issue a
// walletcreatefundedpsbt JSON-RPC command.
func NewWalletCreateFundedPsbtCmd(
	inputs []PsbtInput, outputs []PsbtOutput, locktime *int64,
	options *WalletCreateFundedPsbtOpts, bip32Derivs *bool,
) *WalletCreateFundedPsbtCmd {
	return &WalletCreateFundedPsbtCmd{
		Inputs:      inputs,
		Outputs:     outputs,
		Locktime:    locktime,
		Options:     options,
		Bip32Derivs: bip32Derivs,
	}
}

// WalletProcessPsbtCmd defines the walletprocesspsbt JSON-RPC command.
type WalletProcessPsbtCmd struct {
	Psbt        string
	Sign        *bool `jsonrpcdefault:"true"`
	SighashType *string
	Bip32Derivs *bool
}

// NewWalletProcessPsbtCmd returns a new instance which can be used to issue a
// walletprocesspsbt JSON-RPC command.
func NewWalletProcessPsbtCmd(psbt string, sign *bool, sighashType *string, bip32Derivs *bool) *WalletProcessPsbtCmd {
	return &WalletProcessPsbtCmd{
		Psbt:        psbt,
		Sign:        sign,
		SighashType: sighashType,
		Bip32Derivs: bip32Derivs,
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	flags := UFWalletOnly

	MustRegisterCmd("createnewaccount", (*CreateNewAccountCmd)(nil), flags)
	MustRegisterCmd("dumpwallet", (*DumpWalletCmd)(nil), flags)
	MustRegisterCmd("importaddress", (*ImportAddressCmd)(nil), flags)
	MustRegisterCmd("importpubkey", (*ImportPubKeyCmd)(nil), flags)
	MustRegisterCmd("importwallet", (*ImportWalletCmd)(nil), flags)
	MustRegisterCmd("renameaccount", (*RenameAccountCmd)(nil), flags)
	MustRegisterCmd("walletcreatefundedpsbt", (*WalletCreateFundedPsbtCmd)(nil), flags)
	MustRegisterCmd("walletprocesspsbt", (*WalletProcessPsbtCmd)(nil), flags)
}
