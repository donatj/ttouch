(() => {
	const parts = SplitPath(VM.AbsFilename)[0].split(/[\/\\]/g);
	const name = parts[parts.length - 2];
console.log(name)
	return `# ${name}\n\n`;
})();