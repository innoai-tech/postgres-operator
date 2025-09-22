import { styled } from "@innoai-tech/vueuikit";

export const Spacer = styled("div")({
  flex: 1,
});

export const FullView = styled("div")({
  flex: 1,
  height: "100%",
  width: "100%",
  overflow: "hidden",
});

export const FullColumnView = styled("div")({
  containerStyle: "sys.surface",
  height: "100%",
  width: "100%",
  display: "flex",
  flexDirection: "column",
  justifyContent: "stretch",
  overflow: "hidden",
});

export const FullRowView = styled("div")({
  height: "100%",
  width: "100%",
  display: "flex",
  flexDirection: "flex",
  justifyContent: "stretch",
  overflow: "hidden",
});

export const Heading = styled("div")({
  display: "flex",
  alignItems: "center",
  px: 24,
  gap: 16,
  py: 12,
});

export const Title = styled("h1")({
  textStyle: "sys.headline-small",
  textTransform: "uppercase",
  padding: 0,
  margin: 0,
});

export const Content = styled(FullView)({
  px: 24,
  py: 24,
});

export const Container = styled("div")({
  containerStyle: "sys.surface-container-lowest",
  minHeight: 0,
  height: "100%",
  width: "100%",
  rounded: "xs",
  overflowY: "auto",
});

export const Action = styled("div")({
  display: "flex",
  alignItems: "center",
  gap: 0,
});

export const Scaffold = styled("div")({
  containerStyle: "sys.surface",
  width: "100vw",
  height: "100vh",
  display: "flex",
  flexDirection: "column",
  justifyContent: "stretch",
  overflow: "hidden",
});
