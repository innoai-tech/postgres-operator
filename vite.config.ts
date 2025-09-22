import { defineConfig } from "rolldown-vite";
import {
  app,
  chunkCleanup,
  viteChunkSplit,
  viteVue,
} from "@innoai-tech/vue-vite-presets";
import { injectWebAppConfig } from "@innoai-tech/config/vite-plugin-inject-config";
import { generateClients } from "@innoai-tech/gents";
import { join } from "path";

export default defineConfig({
  build: {
    assetsDir: "assets", // for go embed
  },
  plugins: [
    app("console", {
      enableBaseHref: true,
      buildWithPlaceHolder: true,
    }),
    injectWebAppConfig(async (c, appConfig) => {
      if (appConfig.env !== "$") {
        try {
          await generateClients(join(c.root!, "client"), appConfig, {
            requestCreator: {
              expose: "createRequest",
              importPath: "./client",
            },
          });

          await Bun.$`prettier -w ${c.root}/client/`;
        } catch (e) {
          console.log(e);
        }
      }
    }),
    viteVue({}),
    viteChunkSplit({
      lib: [/webapp\/([^/]+)\/mod/],
    }),
    chunkCleanup({
      minify: true,
    }),
  ],
});
