import { component } from "@nodepkg/runtime";
import { Scaffold, ScaffoldContent } from "@nodepkg/dashboard";
import { PageHeading } from "@webapp/console/common";
import { MustLogon } from "@webapp/console/mod/auth";
import { ArchiveListView } from "@webapp/console/mod/archive";

export default component(() => {
  return () => {
    return (
      <MustLogon>
        <Scaffold>
          <PageHeading />
          <ScaffoldContent>
            <ArchiveListView />
          </ScaffoldContent>
        </Scaffold>
      </MustLogon>
    );
  };
});
