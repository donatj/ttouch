// abc = 2 + twoPlus(2);
// console.log("The value of abc is " + abc); // 4
// console.log(VM.Flags.Executable);

var result = "// " + VM.Filename + "\n\n";

var existing = Glob("*.go")


if( existing && existing.length ) {
	var content = ReadFile(existing[0]);

	var match = content.match(/^package (\w+)$/m);
	if( match.length ) {
		result += "package " + match[1] + "\n\n";
	}
}

result;