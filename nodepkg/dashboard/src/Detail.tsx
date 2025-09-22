import { alpha, styled, variant } from "@innoai-tech/vueuikit";
import type { VNodeChild } from "@nodepkg/runtime";

export const DetailList = styled("div")({
  "& > *": {
    borderBottom: "1px solid",
    borderColor: variant("sys.outline-variant", alpha(0.38)),

    "&:last-child": {
      borderBottom: "none",
    },
  },
});

export const Detail = styled<
  {
    $leading?: VNodeChild;
    $title?: VNodeChild;
    $summary?: VNodeChild;
    $action?: VNodeChild;
    $main?: VNodeChild;
    $aside?: VNodeChild;
    $default?: VNodeChild;
  },
  "div"
>("div", ({}, { slots }) => {
  return (Wrap) => {
    return (
      <Wrap>
        <DetailLeading>{slots.leading?.()}</DetailLeading>
        <DetailTitle>{slots.title?.()}</DetailTitle>
        <DetailSummary>{slots.summary?.()}</DetailSummary>
        <DetailAction>{slots.action?.()}</DetailAction>
        <DetailAside>{slots.aside?.()}</DetailAside>
        <DetailMain>{slots.default?.()}</DetailMain>
      </Wrap>
    );
  };
})({
  py: 16,
  px: 24,
  gap: 8,
  display: "grid",
  gridTemplate: `
"leading title summary action"
"holder aside main  main"
/ min-content minmax(min-content, 20vw) 1fr min-content
`,

  "@media (max-width: 1200px)": {
    gap: 16,

    gridTemplate: `
"leading title summary action"
"holder main  main    main"
"holder aside aside   aside"
/ min-content min-content 1fr min-content
  `,
  },

  "@media (max-width: 600px)": {
    gap: 16,
    gridTemplate: `
"leading title   action"
"holder summary summary"
"holder main    main"
"holder aside   aside"
/ min-content 1fr min-content
  `,
  },
});

const DetailLeading = styled("div")({
  gridArea: "leading",
  display: "flex",
  alignItems: "center",
});

const DetailTitle = styled("div")({
  gridArea: "title",
  display: "flex",
  alignItems: "center",
  gap: 4,
});

const DetailSummary = styled("div")({
  gridArea: "summary",
  display: "flex",
  alignItems: "center",
  gap: 8,
});

const DetailAction = styled("div")({
  gridArea: "action",
  display: "flex",
  alignItems: "center",
});

const DetailMain = styled("div")({
  gridArea: "main",
  display: "flex",
  flexDirection: "column",
  gap: 4,
  textStyle: "sys.body-small",
});

const DetailAside = styled("div")({
  gridArea: "aside",
  textStyle: "sys.body-small",
  wordBreak: "keep-all",
  whiteSpace: "nowrap",

  display: "flex",
  flexDirection: "column",
  justifyContent: "flex-end",

  "@media (max-width: 600px)": {
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "flex-start",
  },
});

export const DetailRow = styled("div")({
  display: "flex",
  alignItems: "center",
});
