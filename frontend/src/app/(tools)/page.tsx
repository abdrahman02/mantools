"use client";

import { ThumbsUp } from "lucide-react";
import Link from "next/link";
import React, { useState } from "react";

/**
 * Mantools Home Page (Next.js, TypeScript, App Router)
 * - Single-file React component suitable for `app/page.tsx` (or adapt to a route)
 * - Uses Tailwind CSS utility classes for layout + custom "neobrutal" inspired styles
 * - Includes a small interactive demo (Text Formatter for JSON/XML/SQL/Markdown — JSON implemented)
 *
 * Notes:
 * - This file is written as a Client Component so the demo can run in the browser.
 * - Keep `any` out of the code. Types are explicit where helpful.
 * - You can split components into separate files later, but this single file is easier to drop into your project.
 */

type Module = {
  id: string;
  title: string;
  subtitle: string;
  icon: string;
};

const MODULES: Module[] = [
  {
    id: "text-formatter",
    title: "Text & Document",
    subtitle: "Formatter, Case Converter, PDF Tools",
    icon: "🔤",
  },
  {
    id: "image-compressor",
    title: "Image & Media",
    subtitle: "Compressor, Converter, QR Generator",
    icon: "🖼️",
  },
  {
    id: "api-request-tester",
    title: "Developer Tools",
    subtitle: "API Tester, JWT Decoder, Hash Generator",
    icon: "🛠️",
  },
];

export default function HomePage(): React.JSX.Element {
  return (
    <div className="min-h-screen bg-gray-50 text-gray-900 antialiased">
      <header className="max-w-6xl mx-auto px-6 py-8 flex items-center justify-between">
        <div className="flex items-center gap-4">
          <div className="rounded-md border-4 border-black p-3 bg-main shadow-neobrutal">
            <span className="text-xl font-extrabold tracking-tight">
              <ThumbsUp className="size-6" />
            </span>
          </div>
          <div>
            <h1 className="text-2xl sm:text-3xl font-extrabold">Mantools</h1>
            <p className="text-sm text-gray-700">
              Versatile online toolbox — formatters, converters, generators &
              more.
            </p>
          </div>
        </div>
        <nav className="flex items-center gap-3">
          <Link
            className="px-4 py-2 border-2 border-black rounded-md hover:translate-y-[-2px] transition-transform"
            href="#modules"
          >
            Modules
          </Link>
          <Link
            className="px-4 py-2 bg-black text-white rounded-md hover:opacity-90 transition-opacity"
            href="#try"
          >
            Try demo
          </Link>
        </nav>
      </header>

      <main className="max-w-6xl mx-auto px-6 pb-24">
        {/* HERO */}
        <section className="grid grid-cols-1 md:grid-cols-2 gap-8 items-center mb-12">
          <div>
            <h2 className="text-4xl sm:text-5xl font-extrabold mb-4 leading-tight">
              Simple. Fast. Reliable.
            </h2>
            <p className="text-lg text-gray-700 mb-6">
              Mantools is a versatile toolbox website that provides a variety of
              online tools such as formatters, converters, generators, and other
              productivity utilities for free and easy to use.
            </p>
            <div className="flex gap-3">
              <Link
                href="#modules"
                className="px-6 py-3 border-4 border-black rounded-md bg-main font-semibold"
              >
                Explore tools
              </Link>
              <Link
                href="#try"
                className="px-6 py-3 rounded-md border-2 border-black"
              >
                Live demo
              </Link>
            </div>
            <div className="mt-6 flex gap-3">
              <span className="inline-block px-3 py-1 border-2 border-black rounded-md">
                Free
              </span>
              <span className="inline-block px-3 py-1 border-2 border-black rounded-md">
                No signup
              </span>
              <span className="inline-block px-3 py-1 border-2 border-black rounded-md">
                Privacy first
              </span>
            </div>
          </div>

          <div className="relative">
            <div className="neobrutal-card p-6 rounded-lg border-4 border-black bg-white shadow-neobrutal">
              <div className="text-sm text-gray-600 mb-2">Quick preview</div>
              <div className="grid grid-cols-2 gap-4">
                {MODULES.map((m) => (
                  <div
                    key={m.id}
                    className="p-3 border-2 border-black rounded-md bg-gray-50 hover:scale-105 transition-transform"
                  >
                    <div className="text-2xl">{m.icon}</div>
                    <div className="font-semibold">{m.title}</div>
                    <div className="text-xs text-gray-600">{m.subtitle}</div>
                  </div>
                ))}
              </div>
              <div className="mt-4 text-xs text-gray-600">
                All tools are built for speed and privacy — processing happens
                in your browser when possible.
              </div>
            </div>
          </div>
        </section>

        {/* MODULES GRID */}
        <section id="modules" className="mb-12">
          <h3 className="text-2xl font-extrabold mb-4">Modules</h3>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {MODULES.map((module) => (
              <ModuleCard key={module.id} module={module} />
            ))}
          </div>
        </section>

        {/* DEMO */}
        <section id="try" className="mb-12">
          <h3 className="text-2xl font-extrabold mb-4">
            Try a live demo — Text Formatter (JSON)
          </h3>
          <div className="neobrutal-card p-6 rounded-lg border-4 border-black bg-white shadow-neobrutal">
            <TextFormatterDemo />
          </div>
        </section>

        {/* FEATURES */}
        <section className="mb-12">
          <h3 className="text-2xl font-extrabold mb-4">Why Mantools?</h3>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <Feature
              title="Fast & Local"
              desc="Most tools process data in the browser — nothing is sent to the server."
            />
            <Feature
              title="Free to use"
              desc="No paywalls, core utilities available to everyone."
            />
            <Feature
              title="Designed for Developers"
              desc="Developer tools like JWT decoder, API tester, and hash generator are included."
            />
          </div>
        </section>

        {/* FOOTER */}
        <footer className="mt-12 border-t pt-6">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
            <div>
              <div className="font-bold">Mantools</div>
              <div className="text-sm text-gray-600">
                © 2025 Mantools. All rights reserved.
              </div>
            </div>
            <div className="flex gap-4">
              <span className="text-sm">Privacy</span>
              <span className="text-sm">Terms</span>
            </div>
          </div>
        </footer>
      </main>

      {/* Small style block to emulate neobrutal look using Tailwind + extra classes */}
      <style jsx>{`
        .shadow-neobrutal {
          box-shadow: 12px 12px 0 rgba(0, 0, 0, 0.08);
        }
        .neobrutal-card {
          border-style: solid;
        }
      `}</style>
    </div>
  );
}

function ModuleCard({ module }: { module: Module }): React.JSX.Element {
  return (
    <article className="p-6 border-4 border-black rounded-md bg-white hover:translate-y-[-4px] transition-transform">
      <Link href={`/${module.id}`}>
        <div className="flex items-center gap-4">
          <div className="text-3xl">{module.icon}</div>
          <div>
            <h4 className="font-bold">{module.title}</h4>
            <p className="text-sm text-gray-600">{module.subtitle}</p>
          </div>
        </div>
      </Link>
    </article>
  );
}

function Feature({
  title,
  desc,
}: {
  title: string;
  desc: string;
}): React.JSX.Element {
  return (
    <div className="p-4 border-2 border-black rounded-md bg-gray-50">
      <div className="font-semibold">{title}</div>
      <div className="text-sm text-gray-600">{desc}</div>
    </div>
  );
}

/**
 * TextFormatterDemo
 * - Small client-side tool that formats JSON input.
 * - If invalid JSON, shows an error message.
 */
function TextFormatterDemo(): React.JSX.Element {
  const [input, setInput] = useState<string>(
    `{\n  "name": "John",\n  "age": 30,\n  "nested": { "a": 1, "b": [1,2,3] }\n}`,
  );
  const [output, setOutput] = useState<string>("");
  const [error, setError] = useState<string | null>(null);
  const [pretty, setPretty] = useState<boolean>(true);

  function formatJSON(): void {
    setError(null);
    try {
      // parse then stringify to normalise and pretty-print
      const parsed = JSON.parse(input);
      const formatted = pretty
        ? JSON.stringify(parsed, null, 2)
        : JSON.stringify(parsed);
      setOutput(formatted);
    } catch (err) {
      setError((err as Error).message);
      setOutput("");
    }
  }

  function minifyJSON(): void {
    setPretty(false);
    formatJSON();
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
        <label className="text-sm font-medium">Input</label>
        <textarea
          value={input}
          onChange={(e) => setInput(e.target.value)}
          className="w-full h-56 mt-2 p-3 border-2 border-black rounded-md font-mono text-sm"
        />
        <div className="flex items-center gap-3 mt-3">
          <button
            onClick={() => {
              setPretty(true);
              formatJSON();
            }}
            className="px-4 py-2 border-2 border-black rounded-md"
          >
            Format
          </button>
          <button
            onClick={minifyJSON}
            className="px-4 py-2 border-2 border-black rounded-md"
          >
            Minify
          </button>
          <button
            onClick={() => {
              setInput("");
              setOutput("");
              setError(null);
            }}
            className="px-4 py-2 border-2 border-black rounded-md"
          >
            Clear
          </button>
        </div>
        {error && (
          <div className="mt-3 text-sm text-red-600">Error: {error}</div>
        )}
      </div>

      <div>
        <label className="text-sm font-medium">Output</label>
        <textarea
          readOnly
          value={output}
          className="w-full h-56 mt-2 p-3 border-2 border-black rounded-md font-mono text-sm bg-gray-100"
        />
        <div className="mt-3 text-xs text-gray-600">
          Formatting performed entirely in your browser.
        </div>
      </div>
    </div>
  );
}
