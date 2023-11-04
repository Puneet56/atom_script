import REPL from '@/features/repl';

export default function Home() {
	return (
		<main className="flex flex-col items-center justify-start pt-20">
			{/* <CodeEditor /> */}
			<REPL />
		</main>
	);
}
