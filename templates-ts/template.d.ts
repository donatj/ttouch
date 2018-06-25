declare function Glob(patterh : string) : string[];

declare function ReadFile(filename: string) : string | null;

declare namespace VM {
	export var Filename : string;

	export namespace Flags {
		export var Executable : boolean;
	}
}