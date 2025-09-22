import { component$, rx, subscribeUntilUnmount } from "@nodepkg/runtime";
import { f, type Field, FormData, t, useRequest } from "@nodepkg/runtime";
import { importArchiveFromTar } from "@webapp/console/client/postgresOperator.ts";
import {
  FilledButton,
  Icon,
  IconButton,
  mdiClose,
  mdiUploadOutline,
  SheetDialogContainer,
  SheetDialogContent,
  SheetDialogFooter,
  SheetDialogHeading,
  SheetDialogHeadingTitle,
  Tooltip,
  useTopSheetDialog,
} from "@nodepkg/dashboard";
import { FileSelectInput } from "@webapp/console/mod/archive/inputs/OciTarFileSelectInput.tsx";
import { FormControls } from "@nodepkg/dashboard";
import { tap } from "@nodepkg/runtime/rxjs";

class ArchiveImportSchema {
  @f.label("备份包")
  @f.hint("谨慎操作，导入会覆盖同名备份")
  @f.inputBy((props) => {
    const { field$, readOnly } = props;

    return (
      <FileSelectInput
        field$={field$ as Field<File>}
        readOnly={readOnly}
        accept={"application/x-tar"}
      />
    );
  })
  @t.custom<File>()
  tarFile!: File;
}

export const ArchiveImportIconBtn = component$<{}>(({}, {}) => {
  const $dialog = useTopSheetDialog(() => {
    const import$ = useRequest(importArchiveFromTar);

    const form$ = FormData.of(t.object(ArchiveImportSchema), {});

    rx(
      form$,
      tap((values) => {
        let archiveCode = "";

        if (values.tarFile.name.endsWith(".tar")) {
          archiveCode = values.tarFile.name.slice(
            0,
            values.tarFile.name.length - ".tar".length,
          );
        }

        import$.next({
          archiveCode: archiveCode,
          body: values.tarFile,
        });
      }),
      subscribeUntilUnmount(),
    );

    rx(
      import$,
      tap(() => {
        $dialog.hide();
      }),
      subscribeUntilUnmount(),
    );

    return () => {
      return (
        <SheetDialogContainer sx={{ h: "auto" }}>
          <SheetDialogHeading>
            <SheetDialogHeadingTitle>导入会覆盖备份?</SheetDialogHeadingTitle>
            <IconButton
              onClick={() => {
                $dialog.hide();
              }}
            >
              <Icon path={mdiClose} />
            </IconButton>
          </SheetDialogHeading>

          <form
            onSubmit={(e) => {
              e.preventDefault();
              form$.submit();
            }}
          >
            <SheetDialogContent>
              <FormControls form$={form$} />
            </SheetDialogContent>
            <SheetDialogFooter>
              <FilledButton type={"submit"}>导入</FilledButton>
            </SheetDialogFooter>
          </form>
        </SheetDialogContainer>
      );
    };
  });

  return () => {
    return (
      <Tooltip $title={"导入备份"}>
        <IconButton
          onClick={() => {
            $dialog.show();
          }}
        >
          <Icon path={mdiUploadOutline} />
          {$dialog.$elem}
        </IconButton>
      </Tooltip>
    );
  };
});
