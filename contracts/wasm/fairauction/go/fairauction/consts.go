// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "fairauction"
	ScDescription = "Decentralized auction to securely sell tokens to the highest bidder"
	HScName       = wasmlib.ScHname(0x1b5c43b1)
)

const (
	ParamColor       = "color"
	ParamDescription = "description"
	ParamDuration    = "duration"
	ParamMinimumBid  = "minimumBid"
	ParamOwnerMargin = "ownerMargin"
)

const (
	ResultBidders       = "bidders"
	ResultColor         = "color"
	ResultCreator       = "creator"
	ResultDeposit       = "deposit"
	ResultDescription   = "description"
	ResultDuration      = "duration"
	ResultHighestBid    = "highestBid"
	ResultHighestBidder = "highestBidder"
	ResultMinimumBid    = "minimumBid"
	ResultNumTokens     = "numTokens"
	ResultOwnerMargin   = "ownerMargin"
	ResultWhenStarted   = "whenStarted"
)

const (
	StateAuctions    = "auctions"
	StateBidderList  = "bidderList"
	StateBids        = "bids"
	StateOwnerMargin = "ownerMargin"
)

const (
	FuncFinalizeAuction = "finalizeAuction"
	FuncPlaceBid        = "placeBid"
	FuncSetOwnerMargin  = "setOwnerMargin"
	FuncStartAuction    = "startAuction"
	ViewGetInfo         = "getInfo"
)

const (
	HFuncFinalizeAuction = wasmlib.ScHname(0x8d534ddc)
	HFuncPlaceBid        = wasmlib.ScHname(0x9bd72fa9)
	HFuncSetOwnerMargin  = wasmlib.ScHname(0x1774461a)
	HFuncStartAuction    = wasmlib.ScHname(0xd5b7bacb)
	HViewGetInfo         = wasmlib.ScHname(0xcfedba5f)
)
