import Prism from 'prismjs';

const atomScriptGrammer: Prism.Grammar = {
	keyword: /\b(atom|molecule|reaction|true|false|if|else|produce)\b/,
	function: /\b(reaction)\b/,
	string: /(")(\\?.)*?\1/,
	operator: /(\+|-|\*|\/|==|!=|>|<|>=|<=)/,
	number: /\b(\d*\.?\d+)\b/,
	punctuation: /[{}[\];(),.:]/,
	boolean: /\b(true|false)\b/,
};

export const getHighlightedCodeHtmlString = (code: string) => {
	return Prism.highlight(code, atomScriptGrammer, 'atomscript');
};
