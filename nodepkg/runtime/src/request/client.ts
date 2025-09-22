import type { RequestConfig, RequestConfigCreator } from "@innoai-tech/fetcher";

export const createRequest = <TInputs, TRespData>(
  operationID: string,
  createConfig: (input: TInputs) => Omit<RequestConfig<TInputs>, "inputs">,
): RequestConfigCreator<TInputs, TRespData> => {
  const create = (inputs: TInputs = {} as TInputs) => ({
    ...createConfig(inputs),
    inputs,
  });
  create.operationID = operationID;
  return create as RequestConfigCreator<TInputs, TRespData>;
};

export type RequestInputs<T extends RequestConfigCreator<any, any>> =
  T extends RequestConfigCreator<infer I, any> ? I : never;

export type ResponseData<T extends RequestConfigCreator<any, any>> =
  T extends RequestConfigCreator<any, infer O> ? O : never;
