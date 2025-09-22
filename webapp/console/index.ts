import { createRouter, createWebHistory } from "@nodepkg/runtime";
import { App } from "./App";
import { createApp } from "@nodepkg/runtime/vue";

// from vite plugin
// @ts-ignore
import routes from "~pages";

const base = new URL(document.querySelector("base")?.href ?? "/");

createApp(App)
  .use(
    createRouter({
      history: createWebHistory(base.pathname),
      routes: routes,
    }),
  )
  .mount("#root");
