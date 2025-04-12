'use client';
import React, { useEffect, useState } from 'react';

import Terminal, { ColorMode, TerminalOutput } from '@/components/terminal';
import { getHighlightedCodeHtmlString } from '@/lib/highlight-code';
import services from '@/services';

export type CodeBlock = {
  code: string;
  isExecuted: boolean;
};

const REPL = ({ height = '300px' }) => {
  const [terminalLineData, setTerminalLineData] = useState<React.ReactNode[]>([]);
  const [loading, setLoading] = useState(false);

  const [code, setCode] = useState<CodeBlock[]>([]);

  const executeCommand = async (input: string) => {
    setLoading(true);

    setTerminalLineData(prev => [
      ...prev,
      <TerminalOutput key={Math.random()}>
        <span dangerouslySetInnerHTML={{ __html: '>> ' + getHighlightedCodeHtmlString(input) }}></span>
      </TerminalOutput>,
    ]);

    if (!input.endsWith('}') && !input.endsWith(';')) {
      input = input + ';';
    }

    let payload = [...code, { code: input, isExecuted: false }];

    try {
      const { data } = await services.evaluateRepl(payload);
      setTerminalLineData(prev => [
        ...prev,
        ...data.map((line: string) => <TerminalOutput key={Math.random()}>{line}</TerminalOutput>),
      ]);

      payload = payload.map((block, i) => {
        if (i === payload.length - 1) {
          return { ...block, isExecuted: true };
        } else {
          return block;
        }
      });

      setCode(payload);
    } catch (error: any) {
      if (error?.response?.data?.errors) {
        setTerminalLineData(prev => [
          ...prev,
          <TerminalOutput key={Math.random()}>{error.response.data.errors.join('\n')}</TerminalOutput>,
        ]);

        return;
      }
      setTerminalLineData(prev => [...prev, <TerminalOutput key={Math.random()}>{error.message}</TerminalOutput>]);
    } finally {
      setLoading(false);
    }
  };

  let output = terminalLineData;

  if (loading) {
    output = [
      ...output,
      <TerminalOutput key={Math.random()}>
        <TerminalLoader />
      </TerminalOutput>,
    ];
  }

  return (
    <div className="container">
      <Terminal
        height={height}
        prompt=">>"
        name="AtomScript REPL"
        colorMode={ColorMode.Dark}
        onInput={executeCommand}
        greenBtnCallback={() => setTerminalLineData([])}
        redBtnCallback={() => setTerminalLineData([])}
        yellowBtnCallback={() => setTerminalLineData([])}
        startingInputValue='puts("Hello world")'
      >
        {output}
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
