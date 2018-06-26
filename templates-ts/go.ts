(() => {
	const existing = Glob("*.go")
	let pkg = "main";

	if (existing && existing.length) {
		const content = ReadFile(existing[0]);

		if (content !== null) {
			const match = content.match(/^package (\w+)$/m);
			if (match !== null && match.length) {
				pkg = match[1];
			}
		}
	}

	let contents = "";
	if(pkg == "main" && VM.Filename == "main.go") {
		contents = `func main() {\n\n}\n`;
	}

	return `package ${pkg}

${contents}`;
})();