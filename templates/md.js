"use strict";
(function () {
    var parts = SplitPath(VM.AbsFilename)[0].split(/[\/\\]/g);
    var name = parts[parts.length - 2];
    console.log(name);
    return "# " + name + "\n\n";
})();
