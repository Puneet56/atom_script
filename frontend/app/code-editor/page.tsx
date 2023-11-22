'use client';

import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import services from '@/services';
import { useState } from 'react';

const CodeEditor = () => {
	const [code, setCode] = useState('');
	const [output, setOutput] = useState('');

	const tokenizeCode = () => {
		setOutput('Loading...');

		services
			.tokenizeCode(code)
			.then(res => {
				setOutput(JSON.stringify(res.data, null, 2));
			})
			.catch(err => {
				setOutput(JSON.stringify(err.response.data, null, 2));
			});
	};

	const generateAst = () => {
		setOutput('Loading...');

		services
			.generateAst(code)
			.then(res => {
				setOutput(JSON.stringify(res.data, null, 2));
			})
			.catch(err => {
				setOutput(JSON.stringify(err.response.data, null, 2));
			});
	};

	const parseCode = () => {
		setOutput('Loading...');

		services
			.parseCode(code)
			.then(res => {
				setOutput(res.data.join('\n'));
			})
			.catch(err => {
				setOutput(JSON.stringify(err.response.data, null, 2));
			});
	};

	const evaluateCode = () => {
		setOutput('Loading...');

		services
			.evaluateCode(code)
			.then(res => {
				setOutput(JSON.stringify(res.data, null, 2));
			})
			.catch(err => {
				setOutput(JSON.stringify(err.response.data, null, 2));
			});
	};

	return (
		<div className="mt-12 flex w-full flex-col items-center gap-8">
			<div className="flex w-full items-start justify-evenly gap-8 px-8">
				<div className="w-full">
					<p className="mb-4">AtomScript code</p>
					<Textarea
						value={code}
						onChange={e => setCode(e.target.value)}
						rows={20}
						className="w-full font-mono text-lg"
					/>
					<div className="flex gap-4 px-4">
						<Button onClick={tokenizeCode} className="mt-8" disabled={code.length === 0}>
							Tokenize
						</Button>

						<Button onClick={generateAst} className="mt-8" disabled={code.length === 0}>
							Generate AST
						</Button>

						<Button onClick={parseCode} className="mt-8" disabled={code.length === 0}>
							Parse
						</Button>

						<Button onClick={evaluateCode} className="mt-8" disabled={code.length === 0}>
							Evaluate
						</Button>
					</div>
				</div>

				<div className="w-full">
					<p className="mb-4">Output</p>
					<Textarea value={output} onChange={() => {}} rows={20} className="w-full font-mono text-lg" />
				</div>
			</div>
		</div>
	);
};

export default CodeEditor;
