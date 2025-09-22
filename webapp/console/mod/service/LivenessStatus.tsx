import {
  component$,
  rx,
  subscribeOnMountedUntilUnmount,
} from "@nodepkg/runtime";
import { styled } from "@nodepkg/dashboard";
import { useRequest } from "@nodepkg/runtime";
import { liveness } from "@webapp/console/client/postgresOperator.ts";
import { interval, map, merge, of, tap } from "@nodepkg/runtime/rxjs";

export const LivenessStatus = component$<{}>(({}, { render }) => {
  const liveness$ = useRequest(liveness);

  rx(
    merge(of(0), interval(1000)),
    tap(() => {
      liveness$.next();
    }),
    subscribeOnMountedUntilUnmount(),
  );

  return rx(
    merge(
      rx(
        liveness$,
        map((resp) => {
          return resp.body["ready"];
        }),
      ),
      rx(
        liveness$.error$,
        map(() => {
          return false;
        }),
      ),
    ),
    render((ready) => {
      return <Status data-ready={ready} />;
    }),
  );
});

const Status = styled("div")({
  boxSize: 24,
  rounded: 12,
  bgColor: "sys.error",

  _ready: {
    bgColor: "sys.success",
  },
});
