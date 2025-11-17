import { Icon, IconButton, mdiExitToApp, Tooltip } from "@nodepkg/dashboard";
import { component } from "@nodepkg/runtime";
import { TokenProvider } from "@webapp/console/mod/auth/models";

export const LogoutIconBtn = component(() => {
  const token$ = TokenProvider.use();

  return () => (
    <Tooltip $title={"退出登录"}>
      <IconButton
        onClick={() => {
          token$.next(null);
        }}
      >
        <Icon path={mdiExitToApp} />
      </IconButton>
    </Tooltip>
  );
});
