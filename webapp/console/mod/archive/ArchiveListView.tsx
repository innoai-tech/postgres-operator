import { component, component$, rx, subscribeUntilUnmount, useRequest } from "@nodepkg/runtime";
import { listArchive } from "@webapp/console/client/postgresOperator.ts";
import { of, tap } from "@nodepkg/runtime/rxjs";
import type { Archive } from "./models";
import {
  Box,
  Detail,
  DetailList,
  DetailRow,
  FullColumnView,
  FullView,
  ObjectView,
  Spacer,
  styled,
  TimestampView,
} from "@nodepkg/dashboard";
import { filesize } from "filesize";
import { ArchiveToolbar, ArchiveToolbarGlobal } from "./actions";

const FileList = styled("div")({
  width: "100%",
  containerStyle: "sys.surface",
  px: 18,
  py: 12,

  "& > *": {
    display: "grid",
    gridTemplateColumns: "2fr 1fr 1fr",
    textAlign: "right",

    "& > [data-role=filename]": {
      textAlign: "left",
    },
  },
});

const ArchiveListItem = component<{ archive: Archive }>((props) => {
  return () => {
    return (
      <Detail
        $title={<ObjectView object={props.archive} />}
        $action={<ArchiveToolbar archive={props.archive} />}
        $aside={
          <DetailRow component={"small"}>
            创建于
            <TimestampView
              timestamp={props.archive.creationTimestamp}
              asDistance={{ addSuffix: true }}
            />
          </DetailRow>
        }
      >
        <FileList>
          {props.archive.files?.map((file) => (
            <div>
              <span data-role={"filename"}>{file.name}</span>
              <span>{filesize(file.size ?? 0, { standard: "iec" })}</span>
              <TimestampView timestamp={file.lastModifiedAt} />
            </div>
          ))}
        </FileList>
      </Detail>
    );
  };
});

export const ArchiveListView = component$(({}, { render }) => {
  const list$ = useRequest(listArchive);

  rx(
    of(1),
    tap(() => {
      list$.next(undefined);
    }),
    subscribeUntilUnmount(),
  );

  const $list = rx(
    list$,
    render((list) => {
      return list.body.items?.map((archive) => <ArchiveListItem archive={archive} />);
    }),
  );

  return () => {
    return (
      <FullColumnView>
        <Box sx={{ minH: 40, display: "flex" }}>
          <Spacer />
          <ArchiveToolbarGlobal />
        </Box>
        <FullView>
          <DetailList>{$list}</DetailList>
        </FullView>
      </FullColumnView>
    );
  };
});
