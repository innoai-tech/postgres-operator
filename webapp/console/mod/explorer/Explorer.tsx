import {
  component$,
  ImmerBehaviorSubject,
  observableRef,
  rx,
  subscribeOnMountedUntilUnmount,
  useRequest,
  useRoute,
  useRouter,
} from "@nodepkg/runtime";
import {
  combineLatest,
  EMPTY,
  fromEvent,
  map,
  of,
  skip,
  switchMap,
  tap,
} from "@nodepkg/runtime/rxjs";
import {
  ExplorerContainer,
  ExplorerInput,
  ExplorerNav,
  ExplorerTable,
} from "@webapp/console/mod/explorer/views";
import {
  listDatabase,
  queryDatabase,
} from "@webapp/console/client/postgresOperator.ts";
import { DatabaseItem } from "@webapp/console/mod/explorer/DatabaseItem.tsx";
import {
  alpha,
  Box,
  Icon,
  IconButton,
  mdiPlay,
  styled,
  variant,
} from "@nodepkg/dashboard";

export const Explorer = component$(({}, { render }) => {
  const listDatabase$ = useRequest(listDatabase);
  const queryDatabase$ = useRequest(queryDatabase);
  const router = useRouter();
  const r = useRoute();

  rx(
    of(1),
    tap(() => {
      listDatabase$.next(undefined);
    }),
    subscribeOnMountedUntilUnmount(),
  );

  const focus$ = ImmerBehaviorSubject.seed<string | null>(
    r.query["focus"] as string,
  );

  rx(
    focus$,
    skip(1),
    tap((focus) => {
      router.replace({
        query: focus
          ? {
              focus: focus,
            }
          : {},
      });
    }),
    subscribeOnMountedUntilUnmount(),
  );

  const $databases = rx(
    combineLatest([
      rx(
        listDatabase$,
        map((resp) => resp.body.items ?? []),
      ),
      focus$,
    ]),
    render(([databases, focus]) => {
      return (
        <DatabaseListView>
          {databases.map((database) => {
            return (
              <DatabaseItem
                database={database}
                focused={database.code == focus}
                onFocus={(d) => {
                  focus$.next(d.code);
                }}
              />
            );
          })}
        </DatabaseListView>
      );
    }),
  );

  const error$ = ImmerBehaviorSubject.seed("");

  const $textArea = observableRef<HTMLTextAreaElement | null>(null);

  rx(
    $textArea,
    switchMap((el) => {
      if (!el) {
        return EMPTY;
      }
      return rx(
        fromEvent<KeyboardEvent>(el, "keyup"),
        tap((evt) => {
          console.log(evt);
          if (evt.shiftKey && evt.key == "Enter") {
            submit();
          }
        }),
      );
    }),
    subscribeOnMountedUntilUnmount(),
  );

  const submit = () => {
    error$.next("");

    if (focus$.value) {
      if ($textArea.value && $textArea.value.value) {
        queryDatabase$.next({
          databaseCode: focus$.value,
          body: {
            sql: $textArea.value.value,
          },
        });
      }
    } else {
      error$.next("未指定数据库");
    }
  };

  rx(
    queryDatabase$.error$,
    tap((resp) => {
      error$.next(resp.body.msg);
    }),
    subscribeOnMountedUntilUnmount(),
  );

  const $error = rx(
    error$,
    render((err) => {
      return (
        err && (
          <Box
            sx={{
              px: 12,
              py: 8,
              color: "sys.error",
            }}
          >
            {err}
          </Box>
        )
      );
    }),
  );

  const $result = rx(
    queryDatabase$,
    render((resp) => {
      return (
        <ExplorerQueryResultView>
          <thead>
            <tr>
              {resp.body.columns?.map((c) => {
                return <th>{c.code}</th>;
              })}
              <th data-expand>&nbsp;</th>
            </tr>
          </thead>
          <tbody>
            {resp.body.data.map((data) => {
              return (
                <tr>
                  {data.map((d: any) => {
                    return <td>{`${d}`}</td>;
                  })}
                  <td></td>
                </tr>
              );
            })}
          </tbody>
        </ExplorerQueryResultView>
      );
    }),
  );

  return () => {
    return (
      <ExplorerContainer>
        <ExplorerNav>{$databases}</ExplorerNav>
        <ExplorerInput>
          <ExplorerTextArea
            ref={$textArea}
            name="sql"
            rows="10"
            value={`select count(1) from t_account`}
          />
          <IconButton
            sx={{
              pos: "absolute",
              bottom: 16,
              right: 16,
            }}
            onClick={() => {
              submit();
            }}
          >
            <Icon path={mdiPlay} />
          </IconButton>
        </ExplorerInput>
        <ExplorerTable>
          {$result}
          {$error}
        </ExplorerTable>
      </ExplorerContainer>
    );
  };
});

const ExplorerTextArea = styled("textarea")({
  width: "100%",
  bgColor: "inherit",
  px: 12,
  py: 12,
  mx: 12,
  my: 8,
  border: "none",
  outline: "none",
});

const ExplorerQueryResultView = styled("table")({
  width: "100%",
  height: "100%",
  tableLayout: "auto",
  rounded: "xs",
  overflow: "auto",
  borderCollapse: "collapse",
  borderSpacing: 0,
  bgColor: "inherit",

  "tr,th,td": {
    bgColor: "inherit",
    border: "1px solid",
    borderColor: variant("sys.outline-variant", alpha(0.38)),
  },

  "th,td": {
    px: 12,
    py: 4,

    "&[data-expand]": {
      width: "100%",
    },
  },

  thead: {
    bgColor: "inherit",
    textAlign: "left",

    th: {
      position: "sticky",
      top: 0,
      zIndex: 1,
      bgColor: "inherit",
    },
  },
  tbody: {
    textAlign: "right",
  },
});

const DatabaseListView = styled("div")({
  py: 12,
  display: "flex",
  flexDirection: "column",
  gap: 8,
  width: "100%",
  height: "100%",
  overflow: "auto",
});
