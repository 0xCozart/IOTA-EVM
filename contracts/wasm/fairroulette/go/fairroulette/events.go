// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:gocritic
package fairroulette

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type FairRouletteEvents struct{}

func (e FairRouletteEvents) Bet(address wasmlib.ScAddress, amount int64, number int64) {
	wasmlib.NewEventEncoder("fairroulette.bet").
		Address(address).
		Int64(amount).
		Int64(number).
		Emit()
}

func (e FairRouletteEvents) Payout(address wasmlib.ScAddress, amount int64) {
	wasmlib.NewEventEncoder("fairroulette.payout").
		Address(address).
		Int64(amount).
		Emit()
}

func (e FairRouletteEvents) Round(number int64) {
	wasmlib.NewEventEncoder("fairroulette.round").
		Int64(number).
		Emit()
}

func (e FairRouletteEvents) Start() {
	wasmlib.NewEventEncoder("fairroulette.start").
		Emit()
}

func (e FairRouletteEvents) Stop() {
	wasmlib.NewEventEncoder("fairroulette.stop").
		Emit()
}

func (e FairRouletteEvents) Winner(number int64) {
	wasmlib.NewEventEncoder("fairroulette.winner").
		Int64(number).
		Emit()
}
