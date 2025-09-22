import type { FetcherResponse } from "@innoai-tech/fetcher";
import {
  type Fetcher,
  type FetcherErrorResponse,
  type RequestConfigCreator,
  type RequestSubject,
} from "@innoai-tech/fetcher";
import { FetcherProvider } from "./FetcherProvider";
import {
  BehaviorSubject,
  catchError,
  from,
  ignoreElements,
  mergeMap,
  Observable,
  of,
  Subject,
  tap,
} from "rxjs";
import { isFunction } from "es-toolkit/compat";
import { fetchEventSource } from "@microsoft/fetch-event-source";
import type { RespError } from "./useRequest.ts";

export const useEventSource = <TReq, TRespData>(
  createConfig: RequestConfigCreator<TReq, TRespData>,
): EventSourceSubject<TReq, TRespData, RespError> => {
  const fetcher = FetcherProvider.use();

  return new EventSourceSubject(createConfig, fetcher);
};

export class EventSourceSubject<TInputs, TBody, TError>
  extends Observable<FetcherResponse<TInputs, TBody>>
  implements RequestSubject<TInputs, TBody, TError>
{
  error$ = new Subject<FetcherErrorResponse<TInputs, TError>>();

  private _message$ = new Subject<FetcherResponse<TInputs, TBody>>();
  private _input$ = new Subject<TInputs>();

  constructor(
    private createConfig: RequestConfigCreator<TInputs, TBody>,
    private fetcher: Fetcher,
  ) {
    super((subscriber) => {
      return this._message$.subscribe(subscriber);
    });
  }

  requesting$ = new BehaviorSubject<boolean>(false);

  unsubscribe = this._input$
    .pipe(
      mergeMap((input) => {
        const config = this.fetcher.build(this.createConfig(input));

        const ctrl = new AbortController();

        const fetch = fetchEventSource(this.fetcher.toHref(config), {
          method: config.method,
          headers: config.headers,
          body: this.fetcher.toRequestBody(config),
          signal: ctrl.signal,
          onerror(err) {
            ctrl.abort(err);
          },
          onmessage: (evt) => {
            if (evt.data) {
              this._message$.next({
                status: 200,
                headers: {},
                config,
                body: JSON.parse(evt.data),
              });
            }
          },
        });

        return from(fetch).pipe(
          catchError((errorResp) => {
            this.error$.next(errorResp);
            return of(errorResp);
          }),
          tap((resp) => {
            this._message$.next(resp);
          }),
        );
      }),
      ignoreElements(),
    )
    .subscribe();

  private _prevInputs?: TInputs;

  next = (inputs: TInputs | ((prevInputs?: TInputs) => TInputs)) => {
    const next = isFunction(inputs) ? inputs(this._prevInputs) : inputs;
    this._prevInputs = next;
    this._input$.next(next);
  };

  get operationID() {
    return this.createConfig.operationID;
  }

  toHref(inputs: TInputs): string {
    return this.fetcher.toHref(this.createConfig(inputs));
  }
}
