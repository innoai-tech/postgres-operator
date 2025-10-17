import {
  applyRequestInterceptors,
  type RequestConfig,
} from "@nodepkg/runtime/fetcher";
import { has } from "@nodepkg/runtime/lodash";
import { component, FetcherProvider, type VNodeChild } from "@nodepkg/runtime";
import {
  OpenidConnectProvider,
  type OpenidV1Token,
  TokenProvider,
} from "./models";

export const AuthorizationInQuery = "x-param-header-Authorization";

export const TokenPatchFetcher = component<{
  $default: VNodeChild;
}>((_, { slots }) => {
  const token$ = TokenProvider.use();
  const oidc = OpenidConnectProvider.use();

  const pureFetcher = FetcherProvider.use();

  const f = applyRequestInterceptors((requestConfig) => {
    const Authorization = `Bearer ${token$.value?.access_token}`;

    return {
      ...requestConfig,
      params: {
        ...requestConfig.params,
        ...(has(requestConfig.params, AuthorizationInQuery)
          ? { [AuthorizationInQuery]: Authorization }
          : {}),
      },
      headers: {
        ...requestConfig.headers,
        Authorization,
      },
    };
  })(pureFetcher);

  const checkAccess = async () => {
    if (token$.value) {
      if (token$.value.refresh_token) {
        if (!token$.validateToken(token$.value.access_token)) {
          try {
            const resp = await pureFetcher.request(
              oidc.token({
                body: {
                  grant_type: "refresh_token",
                  refresh_token: token$.value.refresh_token,
                },
              }),
            );

            if (resp.body) {
              token$.next(resp.body as OpenidV1Token);
            }
          } catch (err) {
            if (
              (err as any)?.["status"] == 401 ||
              (err as any)?.["status"] == 403
            ) {
              token$.next(null);
            }
          }
        }
      }
    }
    return undefined;
  };

  const fetcher = {
    checkAccess: checkAccess,

    build: f.build,

    toRequestBody: f.toRequestBody,

    request: async <TInputs extends any, TRespData extends any>(
      requestConfig: RequestConfig<TInputs>,
    ) => {
      await checkAccess();

      try {
        return await f.request<TInputs, TRespData>(requestConfig);
      } catch (err) {
        if ((err as any)?.["status"] == 401) {
          token$.next(null);
        }
        if (
          (err as any)?.["body"]?.["errors"]?.some((e: any) =>
            e.code?.endsWith("ErrAccountAlreadyLogout"),
          )
        ) {
          token$.next(null);
        }
        throw err;
      }
    },

    toHref: (requestConfig: RequestConfig<any>) =>
      f.toHref({
        ...requestConfig,
        params: {
          ...requestConfig.params,
          [AuthorizationInQuery]: "", // mark AuthorizationInQuery to let Interceptor patching
        },
      }),
  };

  return () => {
    return (
      <FetcherProvider value={fetcher}>{slots.default?.()}</FetcherProvider>
    );
  };
});
