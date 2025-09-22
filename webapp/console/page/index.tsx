import { component } from "@nodepkg/runtime";
import { LogoutIconBtn, MustLogon } from "@webapp/console/mod/auth";
import {
  Action,
  Container,
  Content,
  Heading,
  Scaffold,
  Spacer,
  Title,
} from "@nodepkg/dashboard";
import { ArchiveListView } from "@webapp/console/mod/archive";
import { LivenessStatus } from "@webapp/console/mod/service";

export default component(
  () => () => (
    <MustLogon>
      <Scaffold>
        <Heading>
          <Title>Postgres Console</Title>
          <LivenessStatus />
          <Spacer />
          <Action>
            <LogoutIconBtn />
          </Action>
        </Heading>
        <Content>
          <Container>
            <ArchiveListView />
          </Container>
        </Content>
      </Scaffold>
    </MustLogon>
  ),
  { inheritAttrs: false },
);
