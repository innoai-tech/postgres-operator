export class JSONPath {
  static parse = (path: string): any[] => {
    const segments = path.split(/[\.\[\]]/g).filter((s) => !!s.trim());

    const tokens: any[] = [];

    for (const key of segments) {
      if (key.startsWith("'") && key.endsWith("'")) {
        tokens.push(key.slice(1, key.length - 1).replace(/\\'/g, "'"));
        continue;
      }

      if (key.startsWith('"') && key.endsWith('"')) {
        tokens.push(key.slice(1, key.length - 1).replace(/\\"/g, '"'));
        continue;
      }

      tokens.push(key);
    }

    return tokens;
  };
}
