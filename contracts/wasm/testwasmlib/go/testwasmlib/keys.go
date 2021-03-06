// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamAddress     = 0
	IdxParamAgentID     = 1
	IdxParamBlockIndex  = 2
	IdxParamBool        = 3
	IdxParamBytes       = 4
	IdxParamChainID     = 5
	IdxParamColor       = 6
	IdxParamHash        = 7
	IdxParamHname       = 8
	IdxParamIndex       = 9
	IdxParamInt16       = 10
	IdxParamInt32       = 11
	IdxParamInt64       = 12
	IdxParamInt8        = 13
	IdxParamName        = 14
	IdxParamParam       = 15
	IdxParamRecordIndex = 16
	IdxParamRequestID   = 17
	IdxParamString      = 18
	IdxParamUint16      = 19
	IdxParamUint32      = 20
	IdxParamUint64      = 21
	IdxParamUint8       = 22
	IdxParamValue       = 23

	IdxResultCount  = 24
	IdxResultIotas  = 25
	IdxResultLength = 26
	IdxResultRandom = 27
	IdxResultRecord = 28
	IdxResultValue  = 29

	IdxStateArrays = 30
	IdxStateRandom = 31
)

const keyMapLen = 32

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAddress,
	ParamAgentID,
	ParamBlockIndex,
	ParamBool,
	ParamBytes,
	ParamChainID,
	ParamColor,
	ParamHash,
	ParamHname,
	ParamIndex,
	ParamInt16,
	ParamInt32,
	ParamInt64,
	ParamInt8,
	ParamName,
	ParamParam,
	ParamRecordIndex,
	ParamRequestID,
	ParamString,
	ParamUint16,
	ParamUint32,
	ParamUint64,
	ParamUint8,
	ParamValue,
	ResultCount,
	ResultIotas,
	ResultLength,
	ResultRandom,
	ResultRecord,
	ResultValue,
	StateArrays,
	StateRandom,
}

var idxMap [keyMapLen]wasmlib.Key32
