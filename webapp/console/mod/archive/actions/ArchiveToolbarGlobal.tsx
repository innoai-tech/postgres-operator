import { component$ } from "@nodepkg/runtime";
import { ArchiveImportIconBtn } from "./ArchiveImportIconBtn.tsx";
import { ArchiveAddIconBtn } from "./ArchiveCreateIconBtn.tsx";

export const ArchiveToolbarGlobal = component$<{}>(() => {
  return () => {
    return (
      <>
        <ArchiveAddIconBtn />
        <ArchiveImportIconBtn />
      </>
    );
  };
});
