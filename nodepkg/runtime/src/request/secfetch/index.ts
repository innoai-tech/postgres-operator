import {
  chacha20_decrypt,
  chacha20_encrypt,
  generate_chacha20_key,
  type JWK,
  rsa_oaep_encrypt,
} from "@innoai-tech/crypto";
import {
  applyRequestInterceptors,
  createFetcher as _createFetcher,
  type Fetcher,
  type FetcherErrorResponse,
  type FetcherResponse,
  paramsSerializer,
  type RequestConfig,
  type RequestInterceptor,
  transformRequestBody,
} from "@nodepkg/runtime/fetcher";
import type { RequestConfigCreator } from "@innoai-tech/fetcher";

export type OpenidV1Jwks = {
  keys: OpenidV1Jwk[];
};

export type OpenidV1Jwk = {
  alg: string;
  e: string;
  kid: string;
  kty: string;
  n: string;
  use: string;
};

export const createSecFetcher = (
  {
    jwks,
    bodyEncrypt = "auto",
  }: {
    bodyEncrypt?: "auto" | "always";
    jwks?: RequestConfigCreator<void, OpenidV1Jwks>;
  },
  ...requestInterceptors: RequestInterceptor[]
): Fetcher => {
  const fetcher = applyRequestInterceptors(...requestInterceptors)(
    _createFetcher({
      paramsSerializer,
      transformRequestBody,
    }),
  );

  const fetchJWK = async (): Promise<OpenidV1Jwk> => {
    const resp = await fetcher.request(jwks!());

    const key = (((resp.body as any)?.["keys"] as any[]) ?? []).find(
      (k) => k.use == "enc",
    );

    if (!key) {
      throw Error("no key for enc");
    }

    return key;
  };

  let encKey: Promise<JWK>;

  const resolveJWKForEnc = (): Promise<JWK> => {
    return (encKey ??= fetchJWK());
  };

  const patchEncryptedIfNeed = (accept: string) => {
    if (accept.endsWith("+encrypted")) {
      return accept;
    }
    return accept + "+encrypted";
  };

  return applyRequestInterceptors(...requestInterceptors)({
    ...fetcher,

    async request<TInputs, TRespData>(
      requestConfig: RequestConfig<TInputs>,
    ): Promise<FetcherResponse<TInputs, TRespData>> {
      if (!jwks) {
        return fetcher.request(requestConfig);
      }

      let accept = requestConfig.headers?.["Accept"] ?? "application/json";

      if (bodyEncrypt !== "always" && !accept.endsWith("+encrypted")) {
        return fetcher.request(requestConfig);
      }

      let cipherKey: string | null = null;

      if (accept != "application/octet-stream") {
        cipherKey = await generate_chacha20_key();

        const protectedKey = await rsa_oaep_encrypt(
          new TextEncoder().encode(cipherKey),
          await resolveJWKForEnc(),
        );

        requestConfig.headers = {
          ...requestConfig.headers,
          ["Accept"]: `${patchEncryptedIfNeed(accept)}; protected="${protectedKey}",*/*`,
        };

        const requestBodyContentType = requestConfig.headers?.["Content-Type"];

        switch (requestBodyContentType) {
          case "application/x-www-form-urlencoded":
          case "application/json":
            const buf = new TextEncoder().encode(
              transformRequestBody(requestConfig.body, requestConfig.headers),
            );

            requestConfig.headers = {
              ...requestConfig.headers,
              ["Content-Type"]: `${requestBodyContentType}+encrypted; protected="${protectedKey}"`,
            };

            requestConfig.body = await chacha20_encrypt(buf, cipherKey);
            break;
        }
      }

      const reqInit: RequestInit = {
        method: requestConfig.method,
        headers: requestConfig.headers || {},
        body: requestConfig.body,
      };

      const res = await fetch(
        `${requestConfig.url}?${paramsSerializer(requestConfig.params)}`,
        reqInit,
      );

      let respContentType: string =
        res.headers.get("Content-Type")?.split(";")[0] ?? "";

      let body: any;

      if (cipherKey && respContentType.endsWith("+encrypted")) {
        try {
          const buf = await res.arrayBuffer();
          const arr = new Uint8Array(buf);
          const decrypted = await chacha20_decrypt(arr, cipherKey);
          body = new TextDecoder("utf-8").decode(decrypted);

          respContentType = respContentType.slice(
            0,
            respContentType.length - "+encrypted".length,
          );
        } catch (err) {
          throw err;
        }
      } else {
        if (respContentType != "application/octet-stream") {
          body = await res.text();
        }
      }

      switch (respContentType) {
        case "application/octet-stream":
          body = await res.blob();
          break;
        case "application/json":
          body = JSON.parse(body);
          break;
      }

      const resp: any = {
        config: requestConfig,
        status: res.status,
        headers: {},
      };

      for (const [key, value] of res.headers) {
        resp.headers[key] = value;
      }

      resp.body = body as TRespData;

      if (resp.status >= 400) {
        (resp as FetcherErrorResponse<TInputs, any>).error = resp.body;
        throw resp;
      }

      return resp as FetcherResponse<TInputs, TRespData>;
    },
  });
};
