import { component$, rx, subscribeOnMountedUntilUnmount, useRequest } from "@nodepkg/runtime";
import {
  type DatabaseV1Database,
  listTableOfDatabase,
} from "@webapp/console/client/postgresOperator.ts";
import { tap } from "@nodepkg/runtime/rxjs";
import { styled } from "@nodepkg/dashboard";
import { SummaryLabel, TableView } from "./views";

export const DatabaseItem = component$<{
  focused?: boolean;
  onFocus?: (database: DatabaseV1Database) => void;
  database: DatabaseV1Database;
}>((props, { render, emit }) => {
  const listTable$ = useRequest(listTableOfDatabase);

  rx(
    props.focused$,
    tap((focused) => {
      if (focused) {
        listTable$.next({
          databaseCode: props.database.code,
        });
      }
    }),
    subscribeOnMountedUntilUnmount(),
  );

  const $tables = rx(
    listTable$,
    render((resp) => {
      return (
        <DatabaseTableListView>
          {resp.body.items?.map((table) => {
            return <TableView table={table} key={table.code} />;
          })}
        </DatabaseTableListView>
      );
    }),
  );

  return () => {
    return (
      <DatabaseItemView
        open={props.focused}
        onToggle={(evt) => {
          if (evt.oldState == "closed" && evt.newState == "open") {
            emit("focus", props.database);
          }
        }}
      >
        <summary>
          <SummaryLabel>{props.database.code}</SummaryLabel>
        </summary>
        {$tables}
      </DatabaseItemView>
    );
  };
});

const DatabaseItemView = styled("details")({
  summary: {
    textStyle: "sys.label-small",
  },

  "&[open] > summary": {
    color: "sys.primary",
    fontWeight: "bold",
  },
});

const DatabaseTableListView = styled("div")({
  display: "flex",
  flexDirection: "column",
  pl: 16,
  pr: 8,
  py: 8,
  gap: 4,
});
