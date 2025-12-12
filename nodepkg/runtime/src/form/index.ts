import { has } from "es-toolkit/compat";

export * from "./FormData";

export {
  type Field,
  type FieldMeta,
  type FieldState,
  type InputComponentProps,
} from "./FormDataCore.ts";

export * as f from "./f.ts";

export const delegate = <T extends { [k: string]: any }>(target: T, options: Partial<T>): T => {
  return new Proxy(target, {
    get(target, p) {
      if (has(options, p)) {
        return (options as any)[p];
      }
      return (target as any)[p];
    },
  });
};
