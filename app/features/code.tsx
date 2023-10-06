'use client';

import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@radix-ui/react-label';
import axios from 'axios';
import { useEffect, useState } from 'react';

const CodeEditor = () => {
	const [code, setCode] = useState('');
	const [output, setOutput] = useState('');

	useEffect(() => {
		axios.get(process.env.NEXT_PUBLIC_API_URL!).then(res => {
			console.log(res.data);
		});
	}, []);

	const tokenizeCode = () => {
		setOutput('Loading...');

		axios
			.post(process.env.NEXT_PUBLIC_API_URL + '/api/tokenize', code)
			.then(res => {
				setOutput(JSON.stringify(res.data, null, 2));
			})
			.catch(err => {
				setOutput(JSON.stringify(err.response.data, null, 2));
			});
	};

	return (
		<div className="flex w-full flex-col items-center gap-8">
			<div className="flex w-full items-start justify-evenly gap-8 px-8">
				<div className="w-full">
					<Label className="mb-2">Code</Label>
					<Textarea
						value={code}
						onChange={e => setCode(e.target.value)}
						rows={20}
						className="w-full font-mono text-lg"
					/>

					<Button onClick={tokenizeCode} className="mt-8" disabled={code.length === 0}>
						Tokenize
					</Button>
				</div>

				<div className="w-full">
					<Label className="mb-2">Output</Label>
					<Textarea value={output} onChange={() => {}} rows={20} className="w-full font-mono text-lg" />
				</div>
			</div>
		</div>
	);
};

export default CodeEditor;
