// @ts-ignore
import normalizeCss from "normalize.css/normalize.css?raw";
import conf, { CONFIG } from "@webapp/console/config";
import {
  currentUserInfo,
  exchangeToken,
  jwKs,
} from "@webapp/console/client/postgresOperator.ts";
import { CustomFetcherProvider } from "./CustomFetcherProvider.tsx";
import { component, ManifestProvider, RouterView } from "@nodepkg/runtime";
import { CSSReset, GlobalStyle, ThemeProvider } from "@nodepkg/dashboard";
import { OpenidConnectProvider, TokenProvider } from "@webapp/console/mod/auth";
import { theming } from "./theming.ts";

export const c = conf();

export const App = component(() => {
  return () => (
    <ThemeProvider value={theming}>
      <ManifestProvider
        value={{
          name: c.name,
          description: CONFIG.manifest["description"],
        }}
      >
        <CSSReset />
        <GlobalStyle styles={normalizeCss} />
        <GlobalStyle
          styles={{
            body: {
              width: "100vw",
              height: "100vh",
              overflow: "hidden",
              textStyle: "sys.body-small",

              a: {
                textDecoration: "none",
              },
            },
          }}
        />
        <CustomFetcherProvider
          baseUrl={c.API_POSTGRES_OPERATOR}
          baseHref={
            c.ALL_API_PREFIX_WITH_BASE_HREF == "enabled" ? c.baseHref : "/"
          }
          jwks={jwKs}
        >
          <OpenidConnectProvider
            value={{
              token: exchangeToken,
              userinfo: currentUserInfo,
            }}
          >
            <TokenProvider>
              <RouterView />
            </TokenProvider>
          </OpenidConnectProvider>
        </CustomFetcherProvider>
      </ManifestProvider>
    </ThemeProvider>
  );
});
