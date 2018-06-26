"use strict";
function getPsr4Map() {
    var nss = {};
    var items = ScanUp("composer.json");
    if (items.length > 0) {
        var composerPath = items[0];
        var composerDir = SplitPath(composerPath)[0];
        var content = ReadFile(composerPath);
        if (content !== null) {
            var cjson = JSON.parse(content);
            if (cjson.autoload && cjson.autoload["psr-4"]) {
                for (var x in cjson.autoload["psr-4"]) {
                    nss[composerDir + cjson.autoload["psr-4"][x]] = x;
                }
            }
        }
    }
    return nss;
}
(function () {
    var map = getPsr4Map();
    var ns = "";
    for (var m in map) {
        if (VM.AbsFilename.indexOf(m) === 0) {
            var dir = SplitPath(VM.AbsFilename)[0];
            var suffix = dir.substr(m.length).replace(/(.*?)[\/]*$/g, "$1").replace(/[\/\\]+/g, "\\");
            var prefix = map[m].replace(/(^\\+)|(\\+$)/g, "");
            ns = prefix + "\\" + suffix;
            break;
        }
    }
    var result = "";
    if (VM.Flags.Executable) {
        result += "#!/usr/bin/env php\n";
    }
    result += "<?php\n\n";
    if (ns !== "") {
        result += "namespace " + ns + ";\n\n";
    }
    return result;
})();
