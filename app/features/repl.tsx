'use client';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import Terminal, { ColorMode, TerminalOutput } from 'react-terminal-ui';

const REPL = (props = {}) => {
	const [terminalLineData, setTerminalLineData] = useState<React.ReactNode[]>([]);
	const [loading, setLoading] = useState(true);

	const executeCommand = async (input: string) => {
		setLoading(true);

		try {
			const { data } = await axios.post(process.env.NEXT_PUBLIC_API_URL + '/api/eval', { code: input });
			setTerminalLineData(prev => [...prev, <TerminalOutput>{data}</TerminalOutput>]);
		} catch (error: any) {
			if (error?.response?.data?.error) {
				setTerminalLineData(prev => [...prev, <TerminalOutput>{error.response.data.error}</TerminalOutput>]);
			}
			setTerminalLineData(prev => [...prev, <TerminalOutput>{error.message}</TerminalOutput>]);
		} finally {
			setLoading(false);
		}
	};

	if (loading) {
		terminalLineData.push(
			<TerminalOutput>
				<TerminalLoader />
			</TerminalOutput>
		);
	}

	return (
		<div className="container">
			<Terminal name="AtomScript REPL" colorMode={ColorMode.Dark} onInput={executeCommand}>
				{terminalLineData}
			</Terminal>
		</div>
	);
};

export default REPL;

let loaderElements = ['|', '/', '-', '\\'];

export const TerminalLoader = () => {
	const [loaderIndex, setLoaderIndex] = useState(0);

	useEffect(() => {
		const interval = setInterval(() => {
			setLoaderIndex(prev => {
				if (prev === loaderElements.length - 1) {
					return 0;
				} else {
					return prev + 1;
				}
			});
		}, 100);

		return () => clearInterval(interval);
	}, []);

	return <span>{loaderElements[loaderIndex]}</span>;
};
