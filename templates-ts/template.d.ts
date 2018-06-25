declare function Glob(patterh : string) : string[];

declare function ReadFile(filename: string) : string | null;

declare function ScanUp(filename: string) : string[];

declare function SplitPath(filename: string) : string[2];

declare namespace VM {
	export var Filename : string;

	export var AbsFilename : string;

	export namespace Flags {
		export var Executable : boolean;
	}
}