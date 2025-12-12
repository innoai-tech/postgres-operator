import { component$, rx, subscribeOnMountedUntilUnmount, useRequest } from "@nodepkg/runtime";
import {
  cancelRestoreRequest,
  currentRestoreRequest,
  restart,
} from "@webapp/console/client/postgresOperator.ts";
import {
  Box,
  FilledButton,
  Icon,
  IconButton,
  mdiCancel,
  mdiClose,
  mdiRestartAlert,
  SheetDialogContainer,
  SheetDialogContent,
  SheetDialogFooter,
  SheetDialogHeading,
  SheetDialogHeadingTitle,
  styled,
  Tooltip,
  useTopSheetDialog,
} from "@nodepkg/dashboard";
import { interval, map, merge, of, tap } from "@nodepkg/runtime/rxjs";

export const SystemControlView = component$<{}>(({}, { render }) => {
  const restart$ = useRequest(restart);
  const currentRestoreRequest$ = useRequest(currentRestoreRequest);
  const cancelRestoreRequest$ = useRequest(cancelRestoreRequest);

  rx(
    merge(of(0), interval(3000)),
    tap(() => {
      currentRestoreRequest$.next();
    }),
    subscribeOnMountedUntilUnmount(),
  );

  rx(
    cancelRestoreRequest$,
    tap(() => {
      currentRestoreRequest$.next();
    }),
    subscribeOnMountedUntilUnmount(),
  );

  const $currentRestoreRequest = rx(
    merge(
      rx(
        currentRestoreRequest$,
        map((resp) => {
          return resp.body.code;
        }),
      ),
      rx(
        currentRestoreRequest$.error$,
        map(() => {
          return "";
        }),
      ),
    ),
    render((code) => {
      return code ? (
        <AlertPanel>
          <span>重启后将恢复备份 {code}</span>
          <Tooltip $title={"取消恢复"}>
            <IconButton
              onClick={() => {
                cancelRestoreRequest$.next();
              }}
            >
              <Icon path={mdiCancel} />
            </IconButton>
          </Tooltip>
        </AlertPanel>
      ) : null;
    }),
  );

  const $dialog = useTopSheetDialog(() => {
    return () => {
      return (
        <SheetDialogContainer sx={{ h: "auto" }}>
          <SheetDialogHeading>
            <SheetDialogHeadingTitle>是否重启数据库?</SheetDialogHeadingTitle>
            <IconButton
              onClick={() => {
                $dialog.hide();
              }}
            >
              <Icon path={mdiClose} />
            </IconButton>
          </SheetDialogHeading>
          <SheetDialogContent>谨慎操作</SheetDialogContent>
          <SheetDialogFooter>
            <FilledButton
              onClick={() => {
                restart$.next();
              }}
            >
              确定重启
            </FilledButton>
          </SheetDialogFooter>
        </SheetDialogContainer>
      );
    };
  });

  rx(
    restart$,
    tap(() => {
      $dialog.hide();
    }),
    subscribeOnMountedUntilUnmount(),
  );

  return () => {
    return (
      <Box sx={{ display: "flex", alignItems: "center", gap: 8 }}>
        <Tooltip $title={"重启数据库"}>
          <IconButton
            onClick={() => {
              $dialog.show();
            }}
          >
            <Icon path={mdiRestartAlert} />
            {$dialog.$elem}
          </IconButton>
        </Tooltip>
        {$currentRestoreRequest}
      </Box>
    );
  };
});

const AlertPanel = styled("div")({
  display: "flex",
  alignItems: "center",
  containerStyle: "sys.error-container",
  pl: 8,
  gap: 4,
  rounded: "xs",
});
