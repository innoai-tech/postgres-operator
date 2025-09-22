import {
  createRequestSubject,
  type Fetcher,
  type RequestConfigCreator,
  type RequestSubject,
} from "@innoai-tech/fetcher";
import { FetcherProvider } from "./FetcherProvider";

export { type Fetcher };

export const useRequest = <TReq, TRespData>(
  createConfig: RequestConfigCreator<TReq, TRespData>,
): RequestSubject<TReq, TRespData, RespError> => {
  const fetcher = FetcherProvider.use();

  return createRequestSubject<TReq, TRespData, RespError>(
    createConfig,
    fetcher,
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
