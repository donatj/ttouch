"use strict";
(function () {
    var result = "";
    if (VM.Flags.Executable) {
        result = "#!/usr/bin/env php\n";
    }
    return "<?php\n\n" + result;
})();
