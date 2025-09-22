import { styled } from "@innoai-tech/vueuikit";

export type Object = {
  id?: string;
  code?: string;
  name?: string;
  description?: string;
  annotations?: Record<string, string>;
};

export const ObjectView = styled<
  {
    object: Object;
  },
  "div"
>("div", (props, {}) => {
  return (Wrap) => {
    return (
      <Wrap>
        {props.object.name ? (
          <>
            <div>{props.object.name}</div>
            <small>{props.object.code ?? props.object.id ?? ""}</small>
          </>
        ) : (
          <div>{props.object.code ?? props.object.id ?? ""}</div>
        )}
      </Wrap>
    );
  };
})({
  wordBreak: "keep-all",
  whiteSpace: "nowrap",
  display: "flex",
  flexDirection: "column",
  alignItems: "flex-start",
  textStyle: "sys.label-medium",
  fontSize: "inherit",
  overflow: "hidden",

  "& > *": {
    width: "100%",
    overflow: "hidden",
    textOverflow: "ellipsis",
  },

  "& > small": {
    opacity: 0.48,
    fontSize: "0.68em",
    lineHeight: "118%",
    whiteSpace: "nowrap",
    wordBreak: "keep-all",
  },
});
