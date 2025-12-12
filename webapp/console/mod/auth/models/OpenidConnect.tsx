import type { RequestConfigCreator } from "@innoai-tech/fetcher";
import { createProvider } from "@nodepkg/runtime";

export type Permissions = Record<string, any[]>;

export interface OpenidConnect {
  token: RequestConfigCreator<{ body: OpenidV1GrantPayload }, OpenidV1Token>;
  userinfo: RequestConfigCreator<void, OpenidV1UserInfo>;
}

export type OpenidV1ProviderMeta = {
  code: string;
  name: string;
  authorizationEndpoint?: string;
};

export type OpenidV1ConfigurationWithExternalProviders = {
  authorization_endpoint?: string;
  claims_supported: Array<string>;
  externalProviders?: Array<OpenidV1ProviderMeta>;
  grant_types_supported: Array<string>;
  id_token_signing_alg_values_supported: Array<string>;
  issuer: string;
  jwks_uri: string;
  response_types_supported: Array<string>;
  scopes_supported: Array<string>;
  subject_types_supported: Array<string>;
  token_endpoint: string;
  token_endpoint_auth_methods_supported: Array<string>;
  userinfo_endpoint: string;
};

export type OpenidV1UserInfo = {
  sub: string;
  email?: string;
  name?: string;
  phone_number?: string;
};

export type OpenidV1GrantPayload =
  | OpenidV1GantClientCredentialsGrant
  | OpenidV1PasswordGrant
  | OpenidV1AuthorizationCodeGrant
  | OpenidV1RefreshTokenGrant;

export type OpenidV1GantClientCredentialsGrant = {
  grant_type: "client_credentials";
  client_id?: string;
  client_secret?: string;
  scope?: string;
};

export type OpenidV1AuthorizationCodeGrant = {
  client_id?: string;
  client_secret?: string;
  code: string;
  code_verifier?: string;
  grant_type: "authorization_code";
  redirect_uri?: string;
};

export type OpenidV1PasswordGrant = {
  client_id?: string;
  client_secret?: string;
  grant_type: "password";
  password: string;
  scope?: string;
  username: string;
};

export type OpenidV1RefreshTokenGrant = {
  client_id?: /* @type:string */ string;
  client_secret?: string;
  grant_type: "refresh_token";
  refresh_token: string;
  scope?: string;
};

export type OpenidV1Token = {
  access_token: string;
  expires_in?: number;
  refresh_token?: string;
  scope?: string;
  token_type: string;
};

export const OpenidConnectProvider = createProvider<OpenidConnect>(() => ({}) as any);
