import { defineConfig, loadEnv } from 'vite';

import federation from '@originjs/vite-plugin-federation';
import react from '@vitejs/plugin-react-swc';

export default ({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };
  const base = process.env.DEV ? '' : '/remote';
  return defineConfig({
    plugins: [
      react(),
      federation({
        name: 'remote-app',
        filename: 'remoteEntry.js',
        // Modules to expose
        exposes: {
          './Remote': './src/app.tsx',
        },
        shared: ['react'],
      }),
    ],
    base: base,
    build: {
      target: 'esnext',
    },
  });
};
