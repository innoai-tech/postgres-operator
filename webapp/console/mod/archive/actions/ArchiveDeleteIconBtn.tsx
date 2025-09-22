import { component$, rx, subscribeUntilUnmount } from "@nodepkg/runtime";
import type { Archive } from "@webapp/console/mod/archive/models";
import {
  FilledButton,
  Icon,
  IconButton,
  mdiClose,
  mdiDeleteOutline,
  SheetDialogContainer,
  SheetDialogContent,
  SheetDialogFooter,
  SheetDialogHeading,
  SheetDialogHeadingTitle,
  Tooltip,
  useTopSheetDialog,
} from "@nodepkg/dashboard";
import { useRequest } from "@nodepkg/runtime";
import { deleteArchive } from "@webapp/console/client/postgresOperator.ts";
import { tap } from "@nodepkg/runtime/rxjs";

export const ArchiveDeleteIconBtn = component$<{
  archive: Archive;
}>((props) => {
  const delete$ = useRequest(deleteArchive);

  const $dialog = useTopSheetDialog(() => {
    return () => {
      return (
        <SheetDialogContainer sx={{ h: "auto" }}>
          <SheetDialogHeading>
            <SheetDialogHeadingTitle>是否删除备份?</SheetDialogHeadingTitle>
            <IconButton
              onClick={() => {
                $dialog.hide();
              }}
            >
              <Icon path={mdiClose} />
            </IconButton>
          </SheetDialogHeading>
          <SheetDialogContent>谨慎操作，删除后无法恢复</SheetDialogContent>
          <SheetDialogFooter>
            <FilledButton
              onClick={() => {
                delete$.next({
                  archiveCode: props.archive.code,
                });
              }}
            >
              确定删除
            </FilledButton>
          </SheetDialogFooter>
        </SheetDialogContainer>
      );
    };
  });

  rx(
    delete$,
    tap(() => {
      $dialog.hide();
    }),
    subscribeUntilUnmount(),
  );

  return () => {
    return (
      <Tooltip $title={"删除备份"}>
        <IconButton
          onClick={() => {
            $dialog.show();
          }}
        >
          <Icon path={mdiDeleteOutline} />
          {$dialog.$elem}
        </IconButton>
      </Tooltip>
    );
  };
});
