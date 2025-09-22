import {
  component$,
  f,
  type Field,
  FormData,
  rx,
  subscribeUntilUnmount,
  t,
  useRequestWithUploadProgress,
} from "@nodepkg/runtime";
import { importArchiveFromTar } from "@webapp/console/client/postgresOperator.ts";
import {
  FilledButton,
  FormControls,
  Icon,
  IconButton,
  mdiClose,
  mdiUploadOutline,
  Progress,
  SheetDialogContainer,
  SheetDialogContent,
  SheetDialogFooter,
  SheetDialogHeading,
  SheetDialogHeadingTitle,
  Tooltip,
  useTopSheetDialog,
} from "@nodepkg/dashboard";
import { FileSelectInput } from "@webapp/console/mod/archive/inputs/OciTarFileSelectInput.tsx";
import { tap } from "@nodepkg/runtime/rxjs";
import { ImmerBehaviorSubject } from "@innoai-tech/vuekit";
import type { UploadProgress } from "@innoai-tech/fetcher";

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

export const ArchiveImportIconBtn = component$<{}>(({}, { render }) => {
  const $dialog = useTopSheetDialog(() => {
    const progress$ = new ImmerBehaviorSubject<UploadProgress | null>(null);

    const import$ = useRequestWithUploadProgress(importArchiveFromTar, (p) => {
      progress$.next({
        loaded: p.loaded,
        total: p.total,
      });
    });

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
        progress$.next(null);
        $dialog.hide();
      }),
      subscribeUntilUnmount(),
    );

    const $progress = rx(
      progress$,
      render((progress) => {
        if (progress && progress.total > 0) {
          return (
            <SheetDialogContent>
              <Progress progress={progress.loaded / progress.total} />
            </SheetDialogContent>
          );
        }

        return null;
      }),
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
            {$progress}
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
