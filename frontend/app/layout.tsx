import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import 'prismjs/themes/prism-dark.min.css';
import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
	title: 'AtomScript',
	description: 'AtomScript - Tiny Code, Big Reactions!',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
	return (
		<html lang="en">
			<body className="dark bg-background font-sans">{children}</body>
		</html>
	);
}
