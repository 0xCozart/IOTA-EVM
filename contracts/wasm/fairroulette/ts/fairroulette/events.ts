// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";

export class FairRouletteEvents {

	bet(address: wasmlib.ScAddress, amount: i64, number: i64): void {
		new wasmlib.EventEncoder("fairroulette.bet").
		address(address).
		int64(amount).
		int64(number).
		emit();
	}

	payout(address: wasmlib.ScAddress, amount: i64): void {
		new wasmlib.EventEncoder("fairroulette.payout").
		address(address).
		int64(amount).
		emit();
	}

	round(number: i64): void {
		new wasmlib.EventEncoder("fairroulette.round").
		int64(number).
		emit();
	}

	start(): void {
		new wasmlib.EventEncoder("fairroulette.start").
		emit();
	}

	stop(): void {
		new wasmlib.EventEncoder("fairroulette.stop").
		emit();
	}

	winner(number: i64): void {
		new wasmlib.EventEncoder("fairroulette.winner").
		int64(number).
		emit();
	}
}
