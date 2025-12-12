import { component$ } from "@nodepkg/runtime";
import { Icon, IconButton, mdiDownloadOutline, Tooltip } from "@nodepkg/dashboard";
import type { Archive } from "../models";
import { useRequest } from "@nodepkg/runtime";
import { exportArchiveAsTar } from "@webapp/console/client/postgresOperator.ts";

export const ArchiveExportIconBtn = component$<{
  archive: Archive;
}>((props) => {
  const export$ = useRequest(exportArchiveAsTar);

  return () => {
    return (
      <Tooltip $title={"导出备份"}>
        <IconButton
          onClick={() => {
            window.open(export$.toHref({ archiveCode: props.archive.code }), "blank");
          }}
        >
          <Icon path={mdiDownloadOutline} />
        </IconButton>
      </Tooltip>
    );
  };
});
