import { component$ } from "@nodepkg/runtime";
import type { Archive } from "@webapp/console/mod/archive/models";
import { ArchiveExportIconBtn } from "./ArchiveExportIconBtn.tsx";
import { ArchiveDeleteIconBtn } from "./ArchiveDeleteIconBtn.tsx";
import { ArchiveRestoreRequestIconBtn } from "./ArchiveRestoreRequestIconBtn.tsx";

export const ArchiveToolbar = component$<{
  archive: Archive;
}>((props) => {
  return () => {
    return (
      <>
        <ArchiveRestoreRequestIconBtn archive={props.archive} />
        <ArchiveDeleteIconBtn archive={props.archive} />
        <ArchiveExportIconBtn archive={props.archive} />
      </>
    );
  };
});
