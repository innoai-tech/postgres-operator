import { styled } from "@nodepkg/dashboard";

export const ExplorerContainer = styled("div")({
  display: "grid",
  gridTemplate: `
 "nav input" 200px
 "nav table" 1fr
 / minmax(300px, max-content) 1fr 
  `,
  gap: 16,
  height: "100%",
  width: "100%",
});

export const ExplorerNav = styled("div")({
  gridArea: "nav",
  containerStyle: "sys.surface-container-lowest",
  rounded: "xs",
  overflow: "hidden",
});

export const ExplorerInput = styled("div")({
  gridArea: "input",
  containerStyle: "sys.surface-container-lowest",
  rounded: "xs",
  overflow: "hidden",
  position: "relative",
});

export const ExplorerTable = styled("div")({
  gridArea: "table",
  containerStyle: "sys.surface-container-lowest",
  rounded: "xs",
  overflow: "auto",
});
