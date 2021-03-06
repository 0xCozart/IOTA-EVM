// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// Provide host with details about funcs and views in this smart contract

import {ROOT, ScFuncContext, ScViewContext} from "./context";
import {KEY_EXPORTS, KEY_PARAMS, KEY_RESULTS, KEY_STATE, KEY_ZZZZZZZ} from "./keys";
import {ScMutableStringArray} from "./mutable";
import {getObjectID, OBJ_ID_PARAMS, OBJ_ID_RESULTS, OBJ_ID_ROOT, OBJ_ID_STATE, TYPE_MAP} from "./host";

// Note that we do not use the Wasm export symbol table on purpose
// because Wasm does not allow us to determine whether the symbols
// are meant as view or func, or meant as extra public callbacks
// generated by the compilation of the the Wasm code.
// There are only 2 symbols the ISCP host will actually look for
// in the export table:
// on_load (which must be defined by the SC code) and
// on_call (which is defined here as part of WasmLib)

type ScFuncContextFunc = (f: ScFuncContext) => void;
type ScViewContextFunc = (v: ScViewContext) => void;

let funcs: Array<ScFuncContextFunc> = [];
let views: Array<ScViewContextFunc> = [];

// general entrypoint for the host to call any SC function
// the host will pass the index of one of the entry points
// that was provided by onLoad during SC initialization
export function onCall(index: i32): void {
    let ctx = new ScFuncContext();
    ctx.require(getObjectID(OBJ_ID_ROOT, KEY_STATE, TYPE_MAP) == OBJ_ID_STATE, "object id mismatch");
    ctx.require(getObjectID(OBJ_ID_ROOT, KEY_PARAMS, TYPE_MAP) == OBJ_ID_PARAMS, "object id mismatch");
    ctx.require(getObjectID(OBJ_ID_ROOT, KEY_RESULTS, TYPE_MAP) == OBJ_ID_RESULTS, "object id mismatch");

    if ((index & 0x8000) != 0) {
        // immutable view function, invoke with a view context
        let view = views[index & 0x7fff];
        view(new ScViewContext());
        return;
    }

    // mutable full function, invoke with a func context
    let func = funcs[index];
    func(ctx);
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// context for onLoad function to be able to tell host which
// funcs and views are available as entry points to the SC
export class ScExports {
    exports: ScMutableStringArray;

    // constructs the symbol export context for the onLoad function
    constructor() {
        this.exports = ROOT.getStringArray(KEY_EXPORTS);
        // tell host what value our special predefined key is
        // this helps detect versioning problems between host
        // and client versions of WasmLib
        this.exports.getString(KEY_ZZZZZZZ.keyID).setValue("TypeScript:KEY_ZZZZZZZ");
    }

    // defines the external name of a smart contract func
    // and the entry point function associated with it
    addFunc(name: string, f: ScFuncContextFunc): void {
        let index = funcs.length;
        funcs.push(f);
        this.exports.getString(index).setValue(name);
    }

    // defines the external name of a smart contract view
    // and the entry point function associated with it
    addView(name: string, v: ScViewContextFunc): void {
        let index = views.length as i32;
        views.push(v);
        this.exports.getString(index | 0x8000).setValue(name);
    }
}

