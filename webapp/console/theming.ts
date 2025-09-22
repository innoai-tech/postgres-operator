import { defaultTheme, Palette, Theming } from "@nodepkg/dashboard";

const seedColors = {
  primary: "#336791",

  error: "#FF4026",
  warning: "#FCFD01",
  success: "#5AC220",
};

const light = Palette.createRoleColorRuleBuilder().build();

export const theming = Theming.create(
  {
    ...defaultTheme,
    ...Palette.fromColors(seedColors).toDesignTokens(light),
  },
  { varPrefix: "vk" },
);
