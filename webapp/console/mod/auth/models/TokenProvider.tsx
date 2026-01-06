import { createProvider, ext, ManifestProvider, persist } from "@nodepkg/runtime";
import { BehaviorSubject } from "@nodepkg/runtime/rxjs";
import type { OpenidV1Token } from "./OpenidConnect.tsx";

// return as ms
const expiresIn = (tokenStr?: string) => {
  if (tokenStr) {
    try {
      const exp = JSON.parse(atob(tokenStr.split(".")[1]!)).exp;

      return Math.round((exp - Date.now() / 1000 - 10) * 1000);
    } catch (_) {}
  }
  return 0;
};

const validateToken = (tokenStr?: string) => {
  return expiresIn(tokenStr) > 0;
};

export const TokenProvider = createProvider(
  () => {
    const manifest = ManifestProvider.use();

    const token$ = persist(
      `${manifest.name}/token:${manifest.baseHref ?? "/"}`,
      (v: OpenidV1Token | null) => new BehaviorSubject<OpenidV1Token | null>(v),
    );

    return ext(token$, {
      validateToken,
      expiresIn,
      logout: () => token$.next(null),
    });
  },
  {
    name: "Token",
  },
);
