import { component } from "@nodepkg/runtime";
import { MustLogon } from "@webapp/console/mod/auth";
import { Scaffold, ScaffoldContent } from "@nodepkg/dashboard";
import { PageHeading } from "@webapp/console/common";
import { Explorer } from "@webapp/console/mod/explorer";

export default component(
  () => () => (
    <MustLogon>
      <Scaffold>
        <PageHeading />
        <ScaffoldContent>
          <Explorer />
        </ScaffoldContent>
      </Scaffold>
    </MustLogon>
  ),
  { inheritAttrs: false },
);
