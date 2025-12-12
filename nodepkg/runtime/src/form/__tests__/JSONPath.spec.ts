import { describe, expect, it } from "bun:test";
import { JSONPath } from "../JSONPath.ts";

describe("JSONPath", () => {
  it("parse", () => {
    expect(JSONPath.parse("cards[0].value")).toEqual(["cards", "0", "value"]);
    expect(JSONPath.parse("cards['x'].value")).toEqual(["cards", "x", "value"]);
    expect(JSONPath.parse("cards.'x'.value")).toEqual(["cards", "x", "value"]);
    expect(JSONPath.parse("cards.'k\"l'.value")).toEqual(["cards", 'k"l', "value"]);
  });
});
