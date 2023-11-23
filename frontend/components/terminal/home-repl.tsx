'use client';
import REPL from '@/features/repl';
import { eventBus } from '@/lib/event-bus';
import { useEffect, useRef } from 'react';
import { Button } from '../ui/button';

export const HomeRepl = () => {
	const replRef = useRef<HTMLDivElement>(null);

	useEffect(() => {
		eventBus.subscribe('repl:trynow', () => {
			replRef?.current?.scrollIntoView({ behavior: 'smooth' });
		});
	}, []);

	return (
		<div ref={replRef} className="mt-12 flex flex-col items-center justify-center gap-4">
			<h2 className="text-shadow text-xl font-bold sm:text-2xl md:text-4xl">Try now</h2>
			<REPL height="600px" />
		</div>
	);
};

export const TryNow = () => {
	return (
		<Button className="gap-2" onClick={() => eventBus.publish('repl:trynow')}>
			Try now!
		</Button>
	);
};
