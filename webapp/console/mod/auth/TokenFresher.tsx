import {
  component$,
  createRequest,
  rx,
  subscribeOnMountedUntilUnmount,
  useRequest,
} from "@nodepkg/runtime";
import {
  OpenidConnectProvider,
  type OpenidV1GrantPayload,
  type OpenidV1Token,
  TokenProvider,
} from "./models";
import { EMPTY, switchMap, tap, timer } from "@nodepkg/runtime/rxjs";

export const TokenFresher = component$(({}, {}) => {
  const openid = OpenidConnectProvider.use();

  const exchangeToken = createRequest<
    {
      body: OpenidV1GrantPayload;
    },
    OpenidV1Token
  >("console.ExchangeToken", ({ ...inputs }) => {
    const c = openid.token(inputs);

    return {
      ...c,
      headers: {
        ...c.headers,
        Accept: "application/x-www-form-urlencoded+encrypted",
      },
    };
  });

  const exchange$ = useRequest(exchangeToken);

  const token$ = TokenProvider.use();

  rx(
    exchange$,
    tap((resp) => {
      token$.next(resp.body);
    }),
    subscribeOnMountedUntilUnmount(),
  );

  rx(
    token$,
    switchMap((token) => {
      if (token) {
        const refreshToken = token.refresh_token;

        if (refreshToken) {
          const t = token$.expiresIn(token.access_token);

          console.log(`will refresh token in ${t / 1000}s`);

          return rx(
            timer(t),
            tap(() => {
              exchange$.next({
                body: {
                  grant_type: "refresh_token",
                  refresh_token: refreshToken,
                },
              });
            }),
          );
        }
      }
      return EMPTY;
    }),
    subscribeOnMountedUntilUnmount(),
  );

  return () => null;
});
