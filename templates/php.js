"use strict";
function getPsrMap() {
    let nss = {};
    let items = ScanUp("composer.json");
    if (items.length > 0) {
        let composerPath = items[0];
        let composerDir = SplitPath(composerPath)[0];
        let content = ReadFile(composerPath);
        if (content !== null) {
            let cjson = JSON.parse(content);
            if (cjson.autoload && cjson.autoload["psr-0"]) {
                for (let x in cjson.autoload["psr-0"]) {
                    nss[composerDir + cjson.autoload["psr-0"][x]] = "";
                }
            }
            if (cjson.autoload && cjson.autoload["psr-4"]) {
                for (let x in cjson.autoload["psr-4"]) {
                    nss[composerDir + cjson.autoload["psr-4"][x]] = x;
                }
            }
        }
    }
    return nss;
}
function trimSlashes(s) {
    return s.replace(/(^\\+)|(\\+$)/g, "");
}
(() => {
    let map = getPsrMap();
    let ns = "";
    for (var m in map) {
        if (VM.AbsFilename.indexOf(m) === 0) {
            let dir = SplitPath(VM.AbsFilename)[0];
            let suffix = dir.substr(m.length).replace(/(.*?)[\/]*$/g, "$1").replace(/[\/\\]+/g, "\\");
            let prefix = trimSlashes(map[m]);
            ns = trimSlashes(`${prefix}\\${suffix}`);
            break;
        }
    }
    let result = "";
    if (VM.Flags.Executable) {
        result += "#!/usr/bin/env php\n";
    }
    result += "<?php\n\n";
    if (ns !== "") {
        result += `namespace ${ns};\n\n`;
    }
    return result;
})();
