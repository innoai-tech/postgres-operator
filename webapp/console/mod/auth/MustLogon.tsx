import { tap } from "@nodepkg/runtime/rxjs";
import { component$, useRoute, useRouter } from "@nodepkg/runtime";
import { rx, type VNodeChild } from "@nodepkg/runtime";
import { CurrentUserProvider, TokenProvider } from "./models";
import { TokenPatchFetcher } from "./TokenPatchFetcher.tsx";
import { TokenFresher } from "./TokenFresher.tsx";

export const MustLogon = component$<{
  $default?: VNodeChild;
}>((_, { slots, render }) => {
  const token$ = TokenProvider.use();
  const router = useRouter();
  const r = useRoute();

  return rx(
    token$,
    tap((token) => {
      if (
        !token$.validateToken(token?.access_token) &&
        !token$.validateToken(token?.refresh_token)
      ) {
        void router.replace({
          path: "/login",
          query: {
            redirect_uri: r.fullPath,
          },
        });
      }
    }),
    render((token) => {
      if (token) {
        return (
          <TokenPatchFetcher>
            <TokenFresher />
            <CurrentUserProvider>{slots.default?.()}</CurrentUserProvider>
          </TokenPatchFetcher>
        );
      }
      return null;
    }),
  );
});
