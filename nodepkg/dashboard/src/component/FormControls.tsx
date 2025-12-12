import { component$, type Field, FormData, rx } from "@nodepkg/runtime";
import { TextField } from "@innoai-tech/vuematerial";
import { onUnmounted } from "@nodepkg/runtime/vue";
import { combineLatest } from "@nodepkg/runtime/rxjs";

export const FormControls = component$<{
  form$: FormData<any>;
}>((props, { render }) => {
  return rx(
    props.form$.inputs$,
    render(() => {
      return (
        <>
          {[...props.form$.fields(props.form$.typedef)].map((field$) => {
            return <FieldControl field$={field$} />;
          })}
        </>
      );
    }),
  );
});

const FieldControl = component$<{
  field$: Field;
}>(({ field$ }, { render }) => {
  onUnmounted(() => {
    field$.destroy();
  });

  function labelOrName(field$: Field): string {
    return field$.meta["label"] ?? field$.meta["title"] ?? field$.name;
  }

  return rx(
    combineLatest([field$]),
    render(([s]) => {
      let hidden: any = field$.meta.hidden;

      if (hidden) {
        return null;
      }

      let Input: any = field$.meta.inputBy;

      const $fieldLabel = labelOrName(field$);
      const $fieldHint = field$.meta.hint;

      const readOnly = (field$.meta.readOnlyWhenInitialExist ?? false) && !!s.initial;

      return (
        <TextField
          invalid={!!s.error}
          focus={!!s.focus}
          disabled={false}
          valued={true}
          $label={
            <span>
              {$fieldLabel}
              {field$.optional ? "" : " *"}
            </span>
          }
          $supporting={<span>{s.error?.join("; ") ?? $fieldHint}</span>}
          $trailing={(Input as any).$trailing}
        >
          <Input field$={field$} readOnly={readOnly} />
        </TextField>
      );
    }),
  );
});
