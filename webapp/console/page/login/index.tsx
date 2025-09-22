import { component } from "@nodepkg/runtime";
import { styled } from "@nodepkg/dashboard";
import { LoginCard, MustLogout } from "@webapp/console/mod/auth";

export default component(() => () => (
  <MustLogout>
    <Container>
      <div />
      <LoginContainer>
        <LoginCard />
      </LoginContainer>
    </Container>
  </MustLogout>
));

const Container = styled("div")({
  display: "grid",
  w: "100vw",
  h: "100vh",
  gridTemplateColumns: "1fr min-content",
  containerStyle: "sys.primary",
});

const LoginContainer = styled("div")({
  gridArea: "main",
  containerStyle: "sys.surface",
  h: "100vh",

  minWidth: "20vw",
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  justifyContent: "center",

  px: "4em",

  "@media (max-width: 600px)": {
    minWidth: "100vw",
    px: 0,
  },
});
