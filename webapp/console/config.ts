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
