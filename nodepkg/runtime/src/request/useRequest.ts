import {
  createRequestSubject,
  type Fetcher,
  type RequestConfigCreator,
  type RequestSubject,
  type UploadProgress,
} from "@innoai-tech/fetcher";
import { FetcherProvider } from "./FetcherProvider";
import { createRequest } from "./client.ts";

export { type Fetcher };

export const useRequest = <TReq, TRespData>(
  createConfig: RequestConfigCreator<TReq, TRespData>,
): RequestSubject<TReq, TRespData, RespError> => {
  const fetcher = FetcherProvider.use();

  return createRequestSubject<TReq, TRespData, RespError>(createConfig, fetcher);
};

export const useRequestWithUploadProgress = <TReq, TRespData>(
  createConfig: RequestConfigCreator<TReq, TRespData>,
  onUploadProgress: (p: UploadProgress) => void,
): RequestSubject<TReq, TRespData, RespError> => {
  return useRequest(
    createRequest(createConfig.operationID, (inputs) => {
      const c = createConfig(inputs);
      c.onUploadProgress = onUploadProgress;
      return c as any;
    }),
  );
};

// common error
export interface RespError {
  code?: number;
  msg?: string;
  errors?: Array<ErrorDescriptor>;
}

export interface ErrorDescriptor {
  code?: string;
  message?: string;
  location?: string;
  pointer?: string;
  source?: string;
}
