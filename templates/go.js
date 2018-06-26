"use strict";
(function () {
    var existing = Glob("*.go");
    var pkg = "main";
    if (existing && existing.length) {
        var content = ReadFile(existing[0]);
        if (content !== null) {
            var match = content.match(/^package (\w+)$/m);
            if (match !== null && match.length) {
                pkg = match[1];
            }
        }
    }
    var contents = "";
    if (pkg == "main" && VM.Filename == "main.go") {
        contents = "func main(){\n\n}\n";
    }
    return "package " + pkg + "\n\n" + contents;
})();
