import { component, RouterLink } from "@nodepkg/runtime";
import {
  Action,
  Heading,
  Icon,
  IconButton,
  mdiArchive,
  SectionRow,
  Spacer,
  Title,
  Tooltip,
} from "@nodepkg/dashboard";
import { LivenessStatus } from "@webapp/console/mod/service";
import { SystemControlView } from "@webapp/console/mod/archive";
import { LogoutIconBtn } from "@webapp/console/mod/auth";

export const PageHeading = component(() => {
  return () => {
    return (
      <Heading>
        <RouterLink to={"/"}>
          <Title>Postgres Console</Title>
        </RouterLink>
        <SectionRow>
          <LivenessStatus />
          <SystemControlView />
        </SectionRow>
        <Spacer />
        <Action>
          <RouterLink to={"/archives"}>
            <Tooltip $title={"备份管理"}>
              <IconButton>
                <Icon path={mdiArchive} />
              </IconButton>
            </Tooltip>
          </RouterLink>
          <LogoutIconBtn />
        </Action>
      </Heading>
    );
  };
});
