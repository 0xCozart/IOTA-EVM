// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableTestCoreState struct {
	id int32
}

func (s ImmutableTestCoreState) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxStateCounter])
}

func (s ImmutableTestCoreState) HnameEP() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxStateHnameEP])
}

func (s ImmutableTestCoreState) Ints() MapStringToImmutableInt64 {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateInts], wasmlib.TYPE_MAP)
	return MapStringToImmutableInt64{objID: mapID}
}

func (s ImmutableTestCoreState) MintedColor() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.id, idxMap[IdxStateMintedColor])
}

func (s ImmutableTestCoreState) MintedSupply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxStateMintedSupply])
}

type MutableTestCoreState struct {
	id int32
}

func (s MutableTestCoreState) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxStateCounter])
}

func (s MutableTestCoreState) HnameEP() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxStateHnameEP])
}

func (s MutableTestCoreState) Ints() MapStringToMutableInt64 {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateInts], wasmlib.TYPE_MAP)
	return MapStringToMutableInt64{objID: mapID}
}

func (s MutableTestCoreState) MintedColor() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.id, idxMap[IdxStateMintedColor])
}

func (s MutableTestCoreState) MintedSupply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxStateMintedSupply])
}
