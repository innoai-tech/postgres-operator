import { createRequest } from "./client";

import { t } from "@innoai-tech/typedef";

export const baseUrl = /*#__PURE__*/ createRequest<
  void,
  { [k: string]: string }
>("postgres-operator.BaseURL", () => ({
  method: "GET",
  url: "/api/postgres-operator/v1",
  headers: {
    Accept: "application/json",
  },
}));

export const cancelRestoreRequest = /*#__PURE__*/ createRequest<void, any>(
  "postgres-operator.CancelRestoreRequest",
  () => ({
    method: "DELETE",
    url: "/api/postgres-operator/v1/archive/request-restore",
  }),
);

export const createArchive = /*#__PURE__*/ createRequest<
  void,
  /* @type:object */ ArchiveV1Archive
>("postgres-operator.CreateArchive", () => ({
  method: "POST",
  url: "/api/postgres-operator/v1/archive/archives",
  headers: {
    Accept: "application/json",
  },
}));

export const currentRestoreRequest = /*#__PURE__*/ createRequest<
  void,
  /* @type:object */ ArchiveV1Archive
>("postgres-operator.CurrentRestoreRequest", () => ({
  method: "GET",
  url: "/api/postgres-operator/v1/archive/request-restore",
  headers: {
    Accept: "application/json",
  },
}));

export const currentUserInfo = /*#__PURE__*/ createRequest<
  void,
  /* @type:object */ OpenidV1UserInfo
>("postgres-operator.CurrentUserInfo", () => ({
  method: "GET",
  url: "/api/postgres-operator/v1/openid/user/info",
  headers: {
    Accept: "application/json",
  },
}));

export const deleteArchive = /*#__PURE__*/ createRequest<
  {
    archiveCode: /* @type:string */ ArchiveV1ArchiveCode;
  },
  any
>("postgres-operator.DeleteArchive", (x) => ({
  method: "DELETE",
  url: `/api/postgres-operator/v1/archive/archives/${x["archiveCode"]}`,
}));

export const exchangeToken = /*#__PURE__*/ createRequest<
  {
    body: /* @type:union */ OpenidV1GrantPayload;
  },
  /* @type:object */ OpenidV1Token
>("postgres-operator.ExchangeToken", (x) => ({
  method: "POST",
  url: "/api/postgres-operator/v1/openid/auth/token",
  body: x.body,
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
    Accept: "application/json",
  },
}));

export const exportArchiveAsTar = /*#__PURE__*/ createRequest<
  {
    archiveCode: /* @type:string */ ArchiveV1ArchiveCode;
  },
  any
>("postgres-operator.ExportArchiveAsTar", (x) => ({
  method: "GET",
  url: `/api/postgres-operator/v1/archive/archives/${x["archiveCode"]}/as-tar`,
}));

export const importArchiveFromTar = /*#__PURE__*/ createRequest<
  {
    archiveCode: /* @type:string */ ArchiveV1ArchiveCode;
    body: File | Blob;
  },
  any
>("postgres-operator.ImportArchiveFromTar", (x) => ({
  method: "PUT",
  url: `/api/postgres-operator/v1/archive/archives/${x["archiveCode"]}/from-tar`,
  body: x.body,
  headers: {
    "Content-Type": "application/octet-stream",
  },
}));

export const jwKs = /*#__PURE__*/ createRequest<
  void,
  /* @type:object */ OpenidV1Jwks
>("postgres-operator.JWKs", () => ({
  method: "GET",
  url: "/api/postgres-operator/v1/openid/.well-known/jwks.json",
  headers: {
    Accept: "application/json",
  },
}));

export const listArchive = /*#__PURE__*/ createRequest<
  void,
  /* @type:object */ ArchiveV1ArchiveAsList
>("postgres-operator.ListArchive", () => ({
  method: "GET",
  url: "/api/postgres-operator/v1/archive/archives",
  headers: {
    Accept: "application/json",
  },
}));

export const liveness = /*#__PURE__*/ createRequest<void, { [k: string]: any }>(
  "postgres-operator.Liveness",
  () => ({
    method: "GET",
    url: "/api/postgres-operator/v1/status/liveness",
    headers: {
      Accept: "application/json",
    },
  }),
);

export const readiness = /*#__PURE__*/ createRequest<
  void,
  { [k: string]: any }
>("postgres-operator.Readiness", () => ({
  method: "GET",
  url: "/api/postgres-operator/v1/status/readiness",
  headers: {
    Accept: "application/json",
  },
}));

export const requestRestoreArchive = /*#__PURE__*/ createRequest<
  {
    archiveCode: /* @type:string */ ArchiveV1ArchiveCode;
  },
  any
>("postgres-operator.RequestRestoreArchive", (x) => ({
  method: "PUT",
  url: `/api/postgres-operator/v1/archive/archives/${x["archiveCode"]}/restore-request`,
}));

export const restart = /*#__PURE__*/ createRequest<void, any>(
  "postgres-operator.Restart",
  () => ({
    method: "POST",
    url: "/api/postgres-operator/v1/service/restart",
  }),
);

export type ArchiveV1Archive = {
  code: /* @type:string */ ArchiveV1ArchiveCode;
  kind?: "Archive";
  apiVersion?: "archive/v1";
  name?: string;
  description?: string;
  annotations?: { [k: string]: string };
  creationTimestamp?: /* @type:string */ SqltypeTimeTimestamp;
  modificationTimestamp?: /* @type:string */ SqltypeTimeTimestamp;
  files?: Array</* @type:object */ ArchiveV1File>;
};

export type ArchiveV1ArchiveCode = string;

export type SqltypeTimeTimestamp = string;

export type ArchiveV1File = {
  name: string;
  size?: number;
  lastModifiedAt?: /* @type:string */ SqltypeTimeTimestamp;
};

export type OpenidV1UserInfo = {
  sub: string;
  name?: string;
  nickname?: string;
  preferred_username?: string;
  email?: string;
  email_verified?: boolean;
  phone_number?: string;
  phone_number_verified?: boolean;
};

export type OpenidV1GrantPayload = /* @type:object */
  | OpenidV1AuthorizationCodeGrant
  | /* @type:object */ OpenidV1ClientCredentialsGrant
  | /* @type:object */ OpenidV1PasswordGrant
  | /* @type:object */ OpenidV1RefreshTokenGrant;

export type OpenidV1AuthorizationCodeGrant = {
  grant_type: "authorization_code";
  code: string;
  redirect_uri?: string;
  code_verifier?: string;
  client_id?: string;
  client_secret?: string;
};

export type OpenidV1ClientCredentialsGrant = {
  grant_type: "client_credentials";
  scope?: string;
  client_id?: string;
  client_secret?: string;
};

export type OpenidV1PasswordGrant = {
  grant_type: "password";
  username: string;
  password: string;
  scope?: string;
  client_id?: string;
  client_secret?: string;
};

export type OpenidV1RefreshTokenGrant = {
  grant_type: "refresh_token";
  refresh_token: string;
  scope?: string;
  client_id?: string;
  client_secret?: string;
};

export type OpenidV1Token = {
  token_type: string;
  expires_in?: number;
  access_token: string;
  refresh_token?: string;
  scope?: string;
};

export type OpenidV1Jwks = {
  keys: Array</* @type:object */ OpenidV1Jwk>;
};

export type OpenidV1Jwk = {
  kty: string;
  alg: string;
  kid: string;
  use: string;
  e: string;
  n: string;
};

export type ArchiveV1ArchiveAsList = {
  items?: Array</* @type:object */ ArchiveV1Archive>;
  total?: number;
};

export class ArchiveV1FileSchema {
  @t.string()
  "name"!: string;

  @t.integer()
  @t.optional()
  "size"?: number;

  @t.string()
  @t.optional()
  "lastModifiedAt"?: /* @type:string */ SqltypeTimeTimestamp;
}

export class ArchiveV1ArchiveSchema {
  @t.annotate({ title: "编码" })
  @t.string()
  "code"!: /* @type:string */ ArchiveV1ArchiveCode;

  @t.enums(["Archive"])
  @t.optional()
  "kind"?: "Archive";

  @t.enums(["archive/v1"])
  @t.optional()
  "apiVersion"?: "archive/v1";

  @t.annotate({ title: "名称" })
  @t.string()
  @t.optional()
  "name"?: string;

  @t.annotate({ title: "描述" })
  @t.string()
  @t.optional()
  "description"?: string;

  @t.annotate({ title: "其他注解" })
  @t.record(t.string(), t.string())
  @t.optional()
  "annotations"?: { [k: string]: string };

  @t.annotate({ title: "创建时间" })
  @t.string()
  @t.optional()
  "creationTimestamp"?: /* @type:string */ SqltypeTimeTimestamp;

  @t.annotate({ title: "更新时间" })
  @t.string()
  @t.optional()
  "modificationTimestamp"?: /* @type:string */ SqltypeTimeTimestamp;

  @t.array(t.ref("ArchiveV1FileSchema", () => t.object(ArchiveV1FileSchema)))
  @t.optional()
  "files"?: Array</* @type:object */ ArchiveV1File>;
}

export class OpenidV1UserInfoSchema {
  @t.annotate({ title: "用户标识" })
  @t.string()
  "sub"!: string;

  @t.annotate({ title: "姓名" })
  @t.string()
  @t.optional()
  "name"?: string;

  @t.annotate({ title: "昵称" })
  @t.string()
  @t.optional()
  "nickname"?: string;

  @t.annotate({ title: "自定义用户名" })
  @t.string()
  @t.optional()
  "preferred_username"?: string;

  @t.annotate({ title: "邮箱" })
  @t.string()
  @t.optional()
  "email"?: string;

  @t.annotate({ title: "已验证邮箱" })
  @t.boolean()
  @t.optional()
  "email_verified"?: boolean;

  @t.annotate({ title: "手机号" })
  @t.string()
  @t.optional()
  "phone_number"?: string;

  @t.annotate({ title: "已验证手机号" })
  @t.boolean()
  @t.optional()
  "phone_number_verified"?: boolean;
}

export class OpenidV1AuthorizationCodeGrantSchema {
  @t.annotate({ title: "授权类型" })
  @t.enums(["authorization_code"])
  "grant_type"!: "authorization_code";

  @t.annotate({ title: "临时凭证 code" })
  @t.string()
  "code"!: string;

  @t.annotate({ title: "重定向地址" })
  @t.string()
  @t.optional()
  "redirect_uri"?: string;

  @t.string()
  @t.optional()
  "code_verifier"?: string;

  @t.annotate({ title: "Client ID" })
  @t.string()
  @t.optional()
  "client_id"?: string;

  @t.annotate({ title: "Client Secret" })
  @t.string()
  @t.optional()
  "client_secret"?: string;
}

export class OpenidV1ClientCredentialsGrantSchema {
  @t.annotate({ title: "授权类型" })
  @t.enums(["client_credentials"])
  "grant_type"!: "client_credentials";

  @t.annotate({ title: "授权范围" })
  @t.string()
  @t.optional()
  "scope"?: string;

  @t.annotate({ title: "Client ID" })
  @t.string()
  @t.optional()
  "client_id"?: string;

  @t.annotate({ title: "Client Secret" })
  @t.string()
  @t.optional()
  "client_secret"?: string;
}

export class OpenidV1PasswordGrantSchema {
  @t.annotate({ title: "授权类型" })
  @t.enums(["password"])
  "grant_type"!: "password";

  @t.annotate({ title: "用户标识" })
  @t.string()
  "username"!: string;

  @t.annotate({ title: "密码" })
  @t.string()
  "password"!: string;

  @t.annotate({ title: "授权范围" })
  @t.string()
  @t.optional()
  "scope"?: string;

  @t.annotate({ title: "Client ID" })
  @t.string()
  @t.optional()
  "client_id"?: string;

  @t.annotate({ title: "Client Secret" })
  @t.string()
  @t.optional()
  "client_secret"?: string;
}

export class OpenidV1RefreshTokenGrantSchema {
  @t.annotate({ title: "授权类型" })
  @t.enums(["refresh_token"])
  "grant_type"!: "refresh_token";

  @t.annotate({ title: "刷新 Token" })
  @t.string()
  "refresh_token"!: string;

  @t.annotate({ title: "授权范围" })
  @t.string()
  @t.optional()
  "scope"?: string;

  @t.annotate({ title: "Client ID" })
  @t.string()
  @t.optional()
  "client_id"?: string;

  @t.annotate({ title: "Client Secret" })
  @t.string()
  @t.optional()
  "client_secret"?: string;
}

export class OpenidV1TokenSchema {
  @t.annotate({ title: "Token 类型" })
  @t.string()
  "token_type"!: string;

  @t.annotate({ title: "过期时间（单位：秒）" })
  @t.integer()
  @t.optional()
  "expires_in"?: number;

  @t.annotate({ title: "访问凭证" })
  @t.string()
  "access_token"!: string;

  @t.annotate({ title: "刷新凭证" })
  @t.string()
  @t.optional()
  "refresh_token"?: string;

  @t.annotate({ title: "凭证范围" })
  @t.string()
  @t.optional()
  "scope"?: string;
}

export class OpenidV1JwkSchema {
  @t.annotate({ title: "密钥类型" })
  @t.string()
  "kty"!: string;

  @t.annotate({ title: "密钥算法类型" })
  @t.string()
  "alg"!: string;

  @t.annotate({ title: "密钥 ID" })
  @t.string()
  "kid"!: string;

  @t.annotate({ title: "用途" })
  @t.string()
  "use"!: string;

  @t.annotate({ title: "RSA 公钥的模数" })
  @t.string()
  "e"!: string;

  @t.annotate({ title: "RSA 公钥的指数" })
  @t.string()
  "n"!: string;
}

export class OpenidV1JwksSchema {
  @t.annotate({ title: "密钥列表" })
  @t.array(t.ref("OpenidV1JwkSchema", () => t.object(OpenidV1JwkSchema)))
  "keys"!: Array</* @type:object */ OpenidV1Jwk>;
}

export class ArchiveV1ArchiveAsListSchema {
  @t.annotate({ title: "列表" })
  @t.array(
    t.ref("ArchiveV1ArchiveSchema", () => t.object(ArchiveV1ArchiveSchema)),
  )
  @t.optional()
  "items"?: Array</* @type:object */ ArchiveV1Archive>;

  @t.annotate({ title: "总数" })
  @t.integer()
  @t.optional()
  "total"?: number;
}
