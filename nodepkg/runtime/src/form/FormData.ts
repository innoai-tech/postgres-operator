import { type AnyType, JSONPointer } from "@innoai-tech/vuekit";
import { FormData as OriginFormData } from "./FormDataCore.ts";
import type { RespError } from "../request";
import { JSONPath } from "./JSONPath.ts";

export class FormData<T extends AnyType = AnyType> extends OriginFormData<T> {
  static errorFromRespError(error: RespError) {
    const errors: { [k: string]: string[] } = {};

    if ((error as any).errorFields) {
      // compacted old errorFields Array<{field: string, msg: string}>
      for (const v of (error as any).errorFields) {
        errors[JSONPointer.create(JSONPath.parse(v.field))] = [v.msg];
      }
    } else if (error.errors) {
      for (const e of error.errors) {
        if (e.pointer) {
          errors[e.pointer] = [`${e.message}`];
        }
      }
    }
    return errors;
  }
}
