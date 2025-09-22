import {
  type Component,
  defineModifier,
  type Infer,
  type InferSchema,
  t,
  type Type,
} from "@innoai-tech/vuekit";
import type { Field, InputComponentProps } from "./FormDataCore";

export const label = defineModifier(
  <T extends Type>(type: T, label: string): Type<Infer<T>, InferSchema<T>> => {
    return t.annotate({ label: label }).modify(type) as T;
  },
);

export const hidden = defineModifier(
  <T extends Type>(type: T): Type<Infer<T>, InferSchema<T>> => {
    return t.annotate({ hidden: true }).modify(type) as T;
  },
);

export const hint = defineModifier(
  <T extends Type>(type: T, hint: string): Type<Infer<T>, InferSchema<T>> => {
    return t.annotate({ hint }).modify(type) as T;
  },
);

export const valueDisplay = defineModifier(
  <V extends any, T extends Type<V, any>>(
    type: T,
    valueDisplay: (value: V, field$: Field<V>) => JSX.Element | string,
  ): Type<Infer<T>, InferSchema<T>> => {
    return t
      .annotate<
        T,
        {
          valueDisplay: typeof valueDisplay;
        }
      >({ valueDisplay })
      .modify(type);
  },
);

export const inputBy = defineModifier(
  <V extends any, T extends Type<V, any>>(
    type: T,
    inputBy: Component<InputComponentProps<V>>,
  ): Type<Infer<T>, InferSchema<T>> => {
    return t
      .annotate<
        T,
        {
          inputBy: typeof inputBy;
        }
      >({ inputBy })
      .modify(type);
  },
);

export const readOnlyWhenInitialExist = defineModifier(
  <T extends Type>(
    type: T,
    readOnlyWhenInitialExist: boolean = true,
  ): Type<Infer<T>, InferSchema<T>> => {
    return t
      .annotate<
        T,
        {
          readOnlyWhenInitialExist: typeof readOnlyWhenInitialExist;
        }
      >({ readOnlyWhenInitialExist })
      .modify(type);
  },
);
