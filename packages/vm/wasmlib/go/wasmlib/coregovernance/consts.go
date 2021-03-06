// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coregovernance

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "governance"
	ScDescription = "Core governance contract"
	HScName       = wasmlib.ScHname(0x17cf909f)
)

const (
	ParamChainOwner             = "oi"
	ParamFeeColor               = "fc"
	ParamHname                  = "hn"
	ParamMaxBlobSize            = "bs"
	ParamMaxEventSize           = "es"
	ParamMaxEventsPerReq        = "ne"
	ParamOwnerFee               = "of"
	ParamStateControllerAddress = "S"
	ParamValidatorFee           = "vf"
)

const (
	ResultAllowedStateControllerAddresses = "a"
	ResultChainID                         = "c"
	ResultChainOwnerID                    = "o"
	ResultDefaultOwnerFee                 = "do"
	ResultDefaultValidatorFee             = "dv"
	ResultDescription                     = "d"
	ResultFeeColor                        = "f"
	ResultMaxBlobSize                     = "mb"
	ResultMaxEventSize                    = "me"
	ResultMaxEventsPerReq                 = "mr"
	ResultOwnerFee                        = "of"
	ResultValidatorFee                    = "vf"
)

const (
	FuncAddAllowedStateControllerAddress    = "addAllowedStateControllerAddress"
	FuncClaimChainOwnership                 = "claimChainOwnership"
	FuncDelegateChainOwnership              = "delegateChainOwnership"
	FuncRemoveAllowedStateControllerAddress = "removeAllowedStateControllerAddress"
	FuncRotateStateController               = "rotateStateController"
	FuncSetChainInfo                        = "setChainInfo"
	FuncSetContractFee                      = "setContractFee"
	FuncSetDefaultFee                       = "setDefaultFee"
	ViewGetAllowedStateControllerAddresses  = "getAllowedStateControllerAddresses"
	ViewGetChainInfo                        = "getChainInfo"
	ViewGetFeeInfo                          = "getFeeInfo"
	ViewGetMaxBlobSize                      = "getMaxBlobSize"
)

const (
	HFuncAddAllowedStateControllerAddress    = wasmlib.ScHname(0x9469d567)
	HFuncClaimChainOwnership                 = wasmlib.ScHname(0x03ff0fc0)
	HFuncDelegateChainOwnership              = wasmlib.ScHname(0x93ecb6ad)
	HFuncRemoveAllowedStateControllerAddress = wasmlib.ScHname(0x31f69447)
	HFuncRotateStateController               = wasmlib.ScHname(0x244d1038)
	HFuncSetChainInfo                        = wasmlib.ScHname(0x702f5d2b)
	HFuncSetContractFee                      = wasmlib.ScHname(0x8421a42b)
	HFuncSetDefaultFee                       = wasmlib.ScHname(0x3310ecd0)
	HViewGetAllowedStateControllerAddresses  = wasmlib.ScHname(0xf3505183)
	HViewGetChainInfo                        = wasmlib.ScHname(0x434477e2)
	HViewGetFeeInfo                          = wasmlib.ScHname(0x9fe54b48)
	HViewGetMaxBlobSize                      = wasmlib.ScHname(0xe1db3d28)
)
