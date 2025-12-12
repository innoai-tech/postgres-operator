import { component$ } from "@nodepkg/runtime";
import { format, formatDistance, parseISO } from "@nodepkg/runtime/date-fns";
import { zhCN } from "date-fns/locale";
import { isBoolean } from "@nodepkg/runtime/lodash";
import { styled } from "@innoai-tech/vueuikit";

export const TimestampView = component$<{
  timestamp?: string;
  format?: string;
  asDistance?:
    | true
    | {
        addSuffix?: true;
        includeSeconds?: true;
      };
}>((props) => {
  return () => {
    if (!props.timestamp) {
      return null;
    }

    if (props.asDistance) {
      let v = formatDistance(props.timestamp, Date.now(), {
        ...(isBoolean(props.asDistance) ? {} : props.asDistance),
        locale: zhCN,
      });

      if (/[0-9A-Za-z]/.test(v[0] ?? "")) {
        return (
          <TimestampValueContainer data-timestamp={props.timestamp}>
            <span>&nbsp;</span>
            <span>{v}</span>
          </TimestampValueContainer>
        );
      }

      return (
        <TimestampValueContainer data-timestamp={props.timestamp}>{v}</TimestampValueContainer>
      );
    }

    return (
      <TimestampValueContainer data-timestamp={props.timestamp}>
        {format(parseISO(props.timestamp), props.format ?? "yyyy-MM-dd HH:mm:ss")}
      </TimestampValueContainer>
    );
  };
});

const TimestampValueContainer = styled("div")({});
