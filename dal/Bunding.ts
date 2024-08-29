/**
 * This file is responsible for binding the C library to the Bun runtime.
 */
import { dlopen, FFIType, suffix, ptr, toBuffer } from "bun:ffi";

const libname = `lib/clib.${suffix}`;
const libpath = libname;

const {
  symbols: { InitSQLite, CreateRowIterator, NextRow, GetLen, Free, Cleanup },
} = dlopen(libpath, {
  InitSQLite: {
    args: [FFIType.cstring],
    returns: FFIType.void,
  },
  CreateRowIterator: {
    args: [FFIType.cstring, FFIType.i32],
    returns: FFIType.i32,
  },
  NextRow: {
    args: [FFIType.i32],
    returns: FFIType.ptr,
  },
  GetLen: {
    args: [FFIType.i32],
    returns: FFIType.i32,
  },
  Free: {
    args: [FFIType.ptr],
    returns: FFIType.void,
  },
  Cleanup: {
    args: [FFIType.i32],
    returns: FFIType.void,
  },
});

function initSQLite(pragmas: string) {
  const buf = Buffer.from(pragmas);
  InitSQLite(ptr(buf));
}

function rowIterator(buf: Buffer) {
  const iter = CreateRowIterator(ptr(buf), buf.length);
  const next = () => {
    const pointer = NextRow(iter);
    if (pointer === null) {
      return null;
    }
    const buf = toBuffer(pointer, 0, GetLen(iter));
    //Free(pointer) //should be resolved by GC;
    return buf;
  };

  const cleanup = () => {
    Cleanup(iter);
  };

  return {
    next,
    cleanup,
  };
}

export default {
  initSQLite,
  rowIterator,
};
