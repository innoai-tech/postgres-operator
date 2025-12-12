import { TokenProvider } from "./models";
import { component$, rx, useRoute, useRouter, type VNodeChild } from "@nodepkg/runtime";
import { tap } from "@nodepkg/runtime/rxjs";

export const MustLogout = component$<{
  $default?: VNodeChild;
}>((_, { slots, render }) => {
  const token$ = TokenProvider.use();
  const router = useRouter();
  const r = useRoute();

  return rx(
    token$,
    tap((token) => {
      if (token$.validateToken(token?.access_token)) {
        const query = r.query as { redirect_uri?: string };

        void router.replace(query.redirect_uri ? query.redirect_uri : "/");
      }
    }),
    render((token) => {
      if (!token$.validateToken(token?.access_token)) {
        return slots.default?.();
      }
      return null;
    }),
  );
});
