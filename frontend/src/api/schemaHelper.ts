import { paths } from "./schema";
import { UnionToIntersection, Get } from "type-fest";

export type UrlPaths = keyof paths;

export type HttpMethods = Extract<
  keyof UnionToIntersection<paths[keyof paths]>,
  string
>;

export type HttpMethodsFilteredByPath<Path extends UrlPaths> = HttpMethods &
  Extract<keyof UnionToIntersection<paths[Path]>, string>;

export type RequestParameters<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.parameters.query`>;

export type RequestData<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.requestBody.content.application/json`>;

export type ResponseData<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.responses.200.content.application/json`>;
