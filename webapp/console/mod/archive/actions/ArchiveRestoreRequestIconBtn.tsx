import { component$, rx, subscribeUntilUnmount } from "@nodepkg/runtime";
import type { Archive } from "@webapp/console/mod/archive/models";
import {
  FilledButton,
  Icon,
  IconButton,
  mdiBackupRestore,
  mdiClose,
  SheetDialogContainer,
  SheetDialogContent,
  SheetDialogFooter,
  SheetDialogHeading,
  SheetDialogHeadingTitle,
  Tooltip,
  useTopSheetDialog,
} from "@nodepkg/dashboard";
import { useRequest } from "@nodepkg/runtime";
import { requestRestoreArchive } from "@webapp/console/client/postgresOperator.ts";
import { tap } from "@nodepkg/runtime/rxjs";

export const ArchiveRestoreRequestIconBtn = component$<{
  archive: Archive;
}>((props) => {
  const restore$ = useRequest(requestRestoreArchive);

  const $dialog = useTopSheetDialog(() => {
    return () => {
      return (
        <SheetDialogContainer sx={{ h: "auto" }}>
          <SheetDialogHeading>
            <SheetDialogHeadingTitle>是否从此备份恢复数据库?</SheetDialogHeadingTitle>
            <IconButton
              onClick={() => {
                $dialog.hide();
              }}
            >
              <Icon path={mdiClose} />
            </IconButton>
          </SheetDialogHeading>
          <SheetDialogContent>谨慎操作，重启后生效</SheetDialogContent>
          <SheetDialogFooter>
            <FilledButton
              onClick={() => {
                restore$.next({
                  archiveCode: props.archive.code,
                });
              }}
            >
              确定恢复
            </FilledButton>
          </SheetDialogFooter>
        </SheetDialogContainer>
      );
    };
  });

  rx(
    restore$,
    tap(() => {
      $dialog.hide();
    }),
    subscribeUntilUnmount(),
  );

  return () => {
    return (
      <Tooltip $title={"恢复备份"}>
        <IconButton
          onClick={() => {
            $dialog.show();
          }}
        >
          <Icon path={mdiBackupRestore} />
          {$dialog.$elem}
        </IconButton>
      </Tooltip>
    );
  };
});
