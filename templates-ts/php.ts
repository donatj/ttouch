(() => {
	var result = ""

	if (VM.Flags.Executable) {
		result = "#!/usr/bin/env php\n";
	}

	return `<?php

${result}`;
})();