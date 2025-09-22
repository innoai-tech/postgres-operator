import {
  component$,
  createRequest,
  f,
  FormData,
  rx,
  subscribeOnMountedUntilUnmount,
  t,
  useRequest,
} from "@nodepkg/runtime";
import {
  OpenidConnectProvider,
  type OpenidV1GrantPayload,
  type OpenidV1Token,
  TokenProvider,
} from "./models";
import { tap } from "@nodepkg/runtime/rxjs";
import { Box, styled } from "@nodepkg/dashboard";
import { FilledButton } from "@nodepkg/dashboard";
import { FormControls } from "@nodepkg/dashboard";

class PasswordLoginSchema {
  @f.label("用户名")
  @f.inputBy((props) => {
    const { field$, readOnly, ...others } = props;

    return (
      <input
        {...others}
        data-input
        type={"text"}
        readonly={readOnly}
        name={field$.name}
        value={field$.input}
        onChange={(e) => {
          field$.update((e.target as HTMLInputElement).value);
        }}
        onFocus={() => field$.focus()}
        onBlur={() => field$.blur()}
      />
    );
  })
  @t.string()
  username!: string;

  @f.label("密码")
  @f.inputBy((props) => {
    const { field$, readOnly, ...others } = props;

    return (
      <input
        {...others}
        data-input
        type={"password"}
        readonly={readOnly}
        name={field$.name}
        autocomplete={"current-password"}
        value={field$.input}
        onChange={(e) => {
          field$.update((e.target as HTMLInputElement).value);
        }}
        onFocus={() => field$.focus()}
        onBlur={() => field$.blur()}
      />
    );
  })
  @t.string()
  password!: string;
}

export const LoginByPasswordCard = component$(({}, {}) => {
  const openid = OpenidConnectProvider.use();

  const exchangeTokenWithAuthorization = createRequest<
    {
      Authorization?: string;
      body: OpenidV1GrantPayload;
    },
    OpenidV1Token
  >("console.ExchangeToken", ({ Authorization, ...inputs }) => {
    const c = openid.token(inputs);

    return {
      ...c,
      headers: {
        ...c.headers,
        Accept: "application/x-www-form-urlencoded+encrypted",
      },
    };
  });

  const exchange$ = useRequest(exchangeTokenWithAuthorization);
  const token$ = TokenProvider.use();

  const form$ = FormData.of(t.object(PasswordLoginSchema), {});

  rx(
    exchange$.error$,
    tap((resp) => {
      const emptyErrors: Record<string, any> = {
        "/username": null,
        "/password": null,
      };

      switch (resp.status) {
        case 429:
          form$.setErrors({
            ...emptyErrors,
            "/password": [resp.body.msg],
          });
          break;
        case 403:
          form$.setErrors({
            ...emptyErrors,
            "/username": [resp.error.msg ?? ""],
            "/password": [resp.error.msg ?? ""],
          });
          break;
        case 404:
          form$.setErrors({
            ...emptyErrors,
            "/username": ["用户名不存在"],
          });
          break;
      }
    }),
    subscribeOnMountedUntilUnmount(),
  );

  rx(
    exchange$,
    tap((resp) => {
      token$.next(resp.body);
    }),
    subscribeOnMountedUntilUnmount(),
  );

  rx(
    form$,
    tap((userPassword) => {
      exchange$.next({
        body: {
          grant_type: "password",
          username: userPassword.username,
          password: userPassword.password,
        },
      });
    }),
    subscribeOnMountedUntilUnmount(),
  );

  return () => (
    <form
      novalidate
      onSubmit={(e) => {
        e.preventDefault();

        form$.submit();
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 16,
        }}
      >
        <FormControls form$={form$} />
        <FilledButton type={"submit"} sx={{ w: "100%" }}>
          登录
        </FilledButton>
      </Box>
    </form>
  );
});

export const LoginCard = component$(({}, {}) => {
  return () => {
    return (
      <LoginCardContainer>
        <Box
          sx={{
            minWidth: 300,
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            gap: 16,
            flex: 1,
          }}
        >
          <div>&nbsp;</div>
          <LoginByPasswordCard />
        </Box>
      </LoginCardContainer>
    );
  };
});

const LoginCardContainer = styled("div")({
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  gap: 16,
  h: "100%",
  w: "100%",
  py: 24,
});
