import {
  component$,
  createProvider,
  ext,
  ImmerBehaviorSubject,
  onMounted,
  useRequest,
  type VNodeChild,
} from "@nodepkg/runtime";
import {
  OpenidConnectProvider,
  type OpenidV1UserInfo,
  type Permissions,
} from "./OpenidConnect.tsx";
import { rx, subscribeUntilUnmount } from "@nodepkg/runtime";
import { tap } from "@nodepkg/runtime/rxjs";

export class ClusterUser extends ImmerBehaviorSubject<
  | (OpenidV1UserInfo & {
      permissions: Permissions;
      resource_access?: Record<string, Record<string, { roles: string[] }>>;
    })
  | null
> {}

export const CurrentUserContext = createProvider(
  () => {
    return new ClusterUser(null);
  },
  { name: "CurrentUser" },
);

const CurrentUserProviderFactory = component$<{ $default?: VNodeChild }>(({}, { slots }) => {
  const openid = OpenidConnectProvider.use();

  const currentUser$ = useRequest(openid.userinfo);

  const currentUserAndPermissions$ = new ClusterUser(null);

  rx(
    currentUser$,
    tap((resp) => {
      currentUserAndPermissions$.next({
        ...resp.body,
        permissions: {},
      });
    }),
    subscribeUntilUnmount(),
  );

  onMounted(() => {
    currentUser$.next();
  });

  return () => {
    return (
      <CurrentUserContext value={currentUserAndPermissions$}>
        {slots.default?.()}
      </CurrentUserContext>
    );
  };
});

export const CurrentUserProvider = ext(CurrentUserProviderFactory, {
  use: () => CurrentUserContext.use(),
});
