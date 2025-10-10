import {
  api,
  type AppConfig,
  type AppContext,
  confLoader,
} from "@innoai-tech/config";

const APP_CONFIG = {
  API_POSTGRES_OPERATOR: api({
    openapi: "/api/postgres-operator",
  })((_: AppContext) => {
    return "http://0.0.0.0:8080";
  }),
  ALL_API_PREFIX_WITH_BASE_HREF: ({}: AppContext) =>
    (process.env as any).NODE_ENV == "production" ? "enabled" : "disabled",
};

export const CONFIG: AppConfig = {
  name: "postgres-operator",
  group: "",
  manifest: {
    crossorigin: "use-credentials",
  },
  config: (process.env as any).NODE_ENV == "production" ? {} : APP_CONFIG,
};

export default confLoader<keyof typeof APP_CONFIG>();
