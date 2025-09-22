import { Icon, mdiUploadBox } from "@nodepkg/dashboard";
import { Box } from "@nodepkg/dashboard";
import {
  component$,
  type Field,
  observableRef,
  rx,
  subscribeOnMountedUntilUnmount,
} from "@nodepkg/runtime";
import { tap } from "@nodepkg/runtime/rxjs";

export const FileSelectInput = component$<{
  field$: Field<File>;
  readOnly?: boolean;
  accept?: string;
}>((props) => {
  const file$ = observableRef<File | null>(null);

  rx(
    file$,
    tap((file) => {
      if (file) {
        props.field$.update(file);
      }
    }),
    subscribeOnMountedUntilUnmount(),
  );

  return () => {
    const { readOnly, accept } = props;

    return (
      <Box
        component={"label"}
        data-input
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",

          gap: 8,

          $data_file_input: {
            display: "none",
          },

          $data_icon: {
            width: 40,
            height: 40,
            my: 40,
          },
        }}
      >
        <input
          type={"file"}
          name={props.field$.name}
          readonly={readOnly}
          accept={accept}
          data-file-input
          onChange={(e) => {
            const file = (e.target as HTMLInputElement).files?.[0];

            if (file) {
              file$.value = file;
            }
          }}
        />
        <Icon path={mdiUploadBox} />
        <span>{file$.value?.name}</span>
      </Box>
    );
  };
});
