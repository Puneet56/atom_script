import Terminal, { TerminalInput, TerminalOutput } from '@/components/terminal';
import { HomeRepl, TryNow } from '@/components/terminal/home-repl';
import { Button } from '@/components/ui/button';
import { getHighlightedCodeHtmlString } from '@/lib/highlight-code';
import { LucideGithub } from 'lucide-react';
import Link from 'next/link';

export default function Home() {
	return (
		<main
			className="mx-auto min-h-screen bg-repeat p-8"
			style={{ backgroundImage: 'linear-gradient(176deg, #471010 300px, hsla(20 14.3% 4.1%) 300px)' }}
		>
			<div className="container">
				<div className="mb-12 flex flex-col items-center justify-center gap-4 pt-44 text-center text-white md:gap-8">
					<h1 className="text-shadow text-4xl font-semibold sm:text-6xl md:text-8xl">AtomScript</h1>
					<h2 className="text-shadow text-xl font-bold sm:text-2xl md:text-4xl">Tiny Code, Big Reactions!</h2>

					<div className="flex items-center justify-center gap-8">
						<Button className="gap-2">
							<a
								href="https://github.com/Puneet56/atom_script"
								target="_blank"
								className="flex items-center justify-center gap-2"
							>
								Code <LucideGithub />
							</a>
						</Button>

						<TryNow />

						<Button className="gap-2" variant={'ghost'}>
							<Link href={'/code-editor'} className="flex items-center justify-center gap-2">
								Take a deep dive
							</Link>
						</Button>
					</div>
				</div>

				<div className="mx-auto mt-8 flex flex-col items-center justify-between gap-x-8 gap-y-4 md:flex-row">
					<div className="w-full">
						<h2 className="text-center text-lg md:text-2xl">
							Define your atoms. <br />
							React them with a reaction and produce a new molecule.
						</h2>
					</div>
					<div className="w-full">
						<TerminalCard
							code={[
								'atom sodium = "Na";',
								'atom chlorine = "Cl";',
								`reaction getSalt(metal, halogen) {\n  produce metal + halogen;\n}`,
								`molecule salt = getSalt(sodium, chlorine);`,
								`salt`,
							]}
						/>
					</div>
				</div>

				<div className="mx-auto mt-12 flex flex-col items-center justify-between gap-x-8 gap-y-4 md:flex-row-reverse">
					<div className="w-full">
						<h2 className="text-center text-lg md:text-2xl">Perfrom arithmetic operations.</h2>
					</div>

					<div className="w-full">
						<Terminal height="200px">
							<TerminalInput prompt=">> ">
								<span
									dangerouslySetInnerHTML={{
										__html: getHighlightedCodeHtmlString('1 + 1'),
									}}
								></span>
							</TerminalInput>
							<TerminalOutput>
								<span
									dangerouslySetInnerHTML={{
										__html: getHighlightedCodeHtmlString('2'),
									}}
								></span>
							</TerminalOutput>
							<TerminalInput prompt=">> ">
								<span
									dangerouslySetInnerHTML={{
										__html: getHighlightedCodeHtmlString('100-31'),
									}}
								></span>
							</TerminalInput>
							<TerminalOutput>
								<span
									dangerouslySetInnerHTML={{
										__html: getHighlightedCodeHtmlString('69'),
									}}
								></span>
							</TerminalOutput>
							<TerminalInput prompt=">> ">
								<span
									dangerouslySetInnerHTML={{
										__html: getHighlightedCodeHtmlString('21 * 20'),
									}}
								></span>
							</TerminalInput>
							<TerminalOutput>
								<span
									dangerouslySetInnerHTML={{
										__html: getHighlightedCodeHtmlString('420'),
									}}
								></span>
							</TerminalOutput>
						</Terminal>
					</div>
				</div>

				<div className="mx-auto mt-8 flex flex-col items-center justify-between gap-x-8 gap-y-4 md:flex-row">
					<div className="w-full">
						<h2 className="text-center text-lg md:text-2xl">
							Complex molecules (Data Structures) <br />
							Arrays, Objects
						</h2>
					</div>
					<div className="w-full">
						<TerminalCard
							code={[
								'molecule elements = [1, 2, 3, 4, 5, 6, 7];',
								'molecule result = { \n  "temperature" : "300", \n  "pressure" : "1atm" \n}',
							]}
						/>
					</div>
				</div>

				<div className="mx-auto mt-12 flex flex-col items-center justify-between gap-x-8 gap-y-4 md:flex-row-reverse">
					<div className="w-full">
						<h2 className="text-center text-lg md:text-2xl">Closures!</h2>
					</div>

					<div className="w-full">
						<TerminalCard
							code={['atom baseWeight = 30;', `reaction getWeight(weight) {\n  produce weight + baseWeight;\n}`]}
						/>
					</div>
				</div>

				<HomeRepl />
			</div>
		</main>
	);
}

const TerminalCard = ({ code }: { code?: string[] }) => {
	return (
		<div className="react-terminal-wrapper">
			<div className="react-terminal-window-buttons">
				<button className="red-btn" />
				<button className="yellow-btn" />
				<button className="green-btn" />
			</div>
			<div className="react-terminal">
				{code?.map((line, i) => (
					<TerminalOutput key={i}>
						<span
							dangerouslySetInnerHTML={{
								__html: getHighlightedCodeHtmlString(line),
							}}
						></span>
					</TerminalOutput>
				))}
			</div>
		</div>
	);
};
