import { getHighlightedCodeHtmlString } from '@/lib/highlight-code';
import React, { ChangeEvent, KeyboardEvent, PropsWithChildren, ReactNode, useEffect, useRef, useState } from 'react';
import './style.css';

export enum ColorMode {
	Light,
	Dark,
}

export interface Props {
	name?: string;
	prompt?: string;
	height?: string;
	colorMode?: ColorMode;
	children?: ReactNode;
	onInput?: ((input: string) => void) | null | undefined;
	startingInputValue?: string;
	redBtnCallback?: () => void;
	yellowBtnCallback?: () => void;
	greenBtnCallback?: () => void;
}

const Terminal = ({
	name,
	prompt,
	height = '600px',
	colorMode,
	onInput,
	children,
	startingInputValue = '',
	redBtnCallback,
	yellowBtnCallback,
	greenBtnCallback,
}: Props) => {
	const [currentLineInput, setCurrentLineInput] = useState('');
	const [cursorPos, setCursorPos] = useState(0);

	const scrollIntoViewRef = useRef<HTMLDivElement>(null);

	const updateCurrentLineInput = (event: ChangeEvent<HTMLInputElement>) => {
		setCurrentLineInput(event.target.value);
	};

	// Calculates the total width in pixels of the characters to the right of the cursor.
	// Create a temporary span element to measure the width of the characters.
	const calculateInputWidth = (inputElement: HTMLInputElement, chars: string) => {
		const span = document.createElement('span');
		span.style.visibility = 'hidden';
		span.style.position = 'absolute';
		span.style.fontSize = window.getComputedStyle(inputElement).fontSize;
		span.style.fontFamily = window.getComputedStyle(inputElement).fontFamily;
		span.innerText = chars;
		document.body.appendChild(span);
		const width = span.getBoundingClientRect().width;
		document.body.removeChild(span);
		// Return the negative width, since the cursor position is to the left of the input suffix
		return -width;
	};

	const clamp = (value: number, min: number, max: number) => {
		if (value > max) return max;
		if (value < min) return min;
		return value;
	};

	const handleInputKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
		if (!onInput) {
			return;
		}
		if (event.key === 'Enter') {
			onInput(currentLineInput);
			setCursorPos(0);
			setCurrentLineInput('');
		} else if (['ArrowLeft', 'ArrowRight', 'ArrowDown', 'ArrowUp', 'Delete'].includes(event.key)) {
			const inputElement = event.currentTarget;
			let charsToRightOfCursor = '';
			let cursorIndex = currentLineInput.length - (inputElement.selectionStart || 0);
			cursorIndex = clamp(cursorIndex, 0, currentLineInput.length);

			if (event.key === 'ArrowLeft') {
				if (cursorIndex > currentLineInput.length - 1) cursorIndex--;
				charsToRightOfCursor = currentLineInput.slice(currentLineInput.length - 1 - cursorIndex);
			} else if (event.key === 'ArrowRight' || event.key === 'Delete') {
				charsToRightOfCursor = currentLineInput.slice(currentLineInput.length - cursorIndex + 1);
			} else if (event.key === 'ArrowUp') {
				charsToRightOfCursor = currentLineInput.slice(0);
			}

			const inputWidth = calculateInputWidth(inputElement, charsToRightOfCursor);
			setCursorPos(inputWidth);
		}
	};

	useEffect(() => {
		setCurrentLineInput(startingInputValue.trim());
	}, [startingInputValue]);

	// An effect that handles scrolling into view the last line of terminal input or output
	const performScrolldown = useRef(false);
	useEffect(() => {
		if (performScrolldown.current) {
			// skip scrolldown when the component first loads
			setTimeout(() => scrollIntoViewRef?.current?.scrollIntoView({ behavior: 'auto', block: 'nearest' }), 500);
		}
		performScrolldown.current = true;
	}, [children]);

	// We use a hidden input to capture terminal input; make sure the hidden input is focused when clicking anywhere on the terminal
	useEffect(() => {
		if (onInput == null) {
			return;
		}
		// keep reference to listeners so we can perform cleanup
		const elListeners: { terminalEl: Element; listener: EventListenerOrEventListenerObject }[] = [];

		let allTerminalEls = Array.from(document.getElementsByClassName('react-terminal-wrapper'));

		for (const terminalEl of allTerminalEls) {
			const listener = () => (terminalEl?.querySelector('.terminal-hidden-input') as HTMLElement)?.focus();
			terminalEl?.addEventListener('click', listener);
			elListeners.push({ terminalEl, listener });
		}

		return function cleanup() {
			elListeners.forEach(elListener => {
				elListener.terminalEl.removeEventListener('click', elListener.listener);
			});
		};
	}, [onInput]);

	const classes = ['react-terminal-wrapper', 'font-mono'];
	if (colorMode === ColorMode.Light) {
		classes.push('react-terminal-light');
	}
	return (
		<div className={classes.join(' ')} data-terminal-name={name}>
			<div className="react-terminal-window-buttons">
				<button
					className={`${yellowBtnCallback ? 'clickable' : ''} red-btn`}
					disabled={!redBtnCallback}
					onClick={redBtnCallback}
				/>
				<button
					className={`${yellowBtnCallback ? 'clickable' : ''} yellow-btn`}
					disabled={!yellowBtnCallback}
					onClick={yellowBtnCallback}
				/>
				<button
					className={`${greenBtnCallback ? 'clickable' : ''} green-btn`}
					disabled={!greenBtnCallback}
					onClick={greenBtnCallback}
				/>
			</div>
			<div className="react-terminal" style={{ height }}>
				{children}
				{onInput && (
					<div
						className="react-terminal-line react-terminal-input react-terminal-active-input flex"
						data-terminal-prompt={prompt || '$'}
						key="terminal-line-prompt"
					>
						<pre
							dangerouslySetInnerHTML={{
								__html: getHighlightedCodeHtmlString(currentLineInput),
							}}
						></pre>
						<span className="cursor" style={{ left: `${cursorPos + 1}px` }}></span>
					</div>
				)}
				<div ref={scrollIntoViewRef}></div>
			</div>
			<input
				className="terminal-hidden-input"
				placeholder="Terminal Hidden Input"
				value={currentLineInput}
				autoFocus={onInput != null}
				onChange={updateCurrentLineInput}
				onKeyDown={handleInputKeyDown}
			/>
		</div>
	);
};

export { TerminalInput, TerminalOutput };
export default Terminal;

const TerminalOutput = ({ children }: { children?: React.ReactChild }) => {
	return <div className="react-terminal-line">{children}</div>;
};

type TerminalInputProps = PropsWithChildren<{
	prompt?: string;
}>;

const TerminalInput = ({ children, prompt }: TerminalInputProps) => {
	return (
		<div className="react-terminal-line react-terminal-input" data-terminal-prompt={prompt || '$'}>
			{children}
		</div>
	);
};
