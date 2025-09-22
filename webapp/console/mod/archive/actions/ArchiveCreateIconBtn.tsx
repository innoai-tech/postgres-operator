import { component$, rx, subscribeUntilUnmount } from "@nodepkg/runtime";
import { useRequest } from "@nodepkg/runtime";
import { createArchive } from "@webapp/console/client/postgresOperator.ts";
import {
  FilledButton,
  Icon,
  IconButton,
  mdiPlus,
  mdiClose,
  SheetDialogContainer,
  SheetDialogFooter,
  SheetDialogHeading,
  SheetDialogHeadingTitle,
  Tooltip,
  useTopSheetDialog,
} from "@nodepkg/dashboard";
import { tap } from "@nodepkg/runtime/rxjs";

export const ArchiveAddIconBtn = component$<{}>(() => {
  const create$ = useRequest(createArchive);

  const $dialog = useTopSheetDialog(() => {
    return () => {
      return (
        <SheetDialogContainer sx={{ h: "auto" }}>
          <SheetDialogHeading>
            <SheetDialogHeadingTitle>是否创建备份?</SheetDialogHeadingTitle>
            <IconButton
              onClick={() => {
                $dialog.hide();
              }}
            >
              <Icon path={mdiClose} />
            </IconButton>
          </SheetDialogHeading>
          <SheetDialogFooter>
            <FilledButton
              onClick={() => {
                create$.next();
              }}
            >
              确定创建
            </FilledButton>
          </SheetDialogFooter>
        </SheetDialogContainer>
      );
    };
  });

  rx(
    create$,
    tap(() => {
      $dialog.hide();
    }),
    subscribeUntilUnmount(),
  );

  return () => {
    return (
      <Tooltip $title={"创建备份"}>
        <IconButton
          onClick={() => {
            $dialog.show();
          }}
        >
          <Icon path={mdiPlus} />
          {$dialog.$elem}
        </IconButton>
      </Tooltip>
    );
  };
});
