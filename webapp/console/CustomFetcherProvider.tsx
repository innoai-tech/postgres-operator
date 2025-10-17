import {
  component,
  createSecFetcher,
  FetcherProvider,
  type OpenidV1Jwks,
  type VNodeChild,
} from "@nodepkg/runtime";
import { trimEnd, trimStart } from "@nodepkg/runtime/lodash";
import type { RequestConfigCreator } from "@nodepkg/runtime/fetcher";

const fixBaseURL = (baseURL: string = "", baseHref = "/") => {
  let [protocol, url] = baseURL.split("//", 2);
  let [host, pathname] = (url ?? "/").split("/", 2);

  if (!protocol) {
    protocol = location.protocol;
  }

  if (!host) {
    host = location.host;
  } else {
    if (host.startsWith(":")) {
      host = `${location.hostname}${host}`;
    }
  }

  pathname = `${baseHref}${trimStart(pathname ?? "/", "/")}`;

  return `${protocol}//${host}${trimEnd(pathname, "/")}`;
};

export const CustomFetcherProvider = component<{
  baseUrl: string;
  baseHref?: string;
  enableBodyEncrypt?: boolean;
  jwks: RequestConfigCreator<void, OpenidV1Jwks>;
  $default: VNodeChild;
}>((props, { slots }) => {
  const fetcher = createSecFetcher(
    {
      bodyEncrypt: props.enableBodyEncrypt ? "always" : "auto",
      jwks: props.jwks,
    },
    (requestConfig) => {
      if (
        !(
          requestConfig.url.startsWith("//") ||
          requestConfig.url.startsWith("http:") ||
          requestConfig.url.startsWith("https://")
        )
      ) {
        requestConfig.url = `${fixBaseURL(props.baseUrl, props.baseHref)}${requestConfig.url}`;
      }

      return requestConfig;
    },
  );

  return () => (
    <FetcherProvider value={fetcher}>{slots.default?.()}</FetcherProvider>
  );
});
